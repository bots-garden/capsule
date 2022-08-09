package registry

import (
	"fmt"
	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// IDEA: use Couchbase as a backend
// TODO: add an authentication mode with token
// TODO: add a route to delete a module
// TODO: add a route to get the list of the module (and information)

func Serve(httpPort, filesPath, crt, key string) {

	if commons.GetEnv("DEBUG", "false") == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	//router := gin.Default()
	router := gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "😢 Page not found 🥵"})
	})

	/* 🚧 this is a work in progress
	   ==============================================================
	   Get the list of all the wasm modules
	   ==============================================================
	   http://localhost:4999/modules
	*/
	router.GET("/modules", func(c *gin.Context) {
		var modules []string
		err := filepath.Walk(filesPath,

			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				modules = append(modules, path)

				//fmt.Println(path, info.Size())
				return err
			})

		if err != nil {
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, modules)

		}
	})

	/*
	   ==============================================================
	   Get information about a wasm module
	   ==============================================================
	   http://localhost:4999/info/k33g/hello/0.0.0
	*/
	router.GET("/info/:user_org/:wasm_module/:tag", func(c *gin.Context) {
		userOrg := c.Param("user_org")
		wasmModule := c.Param("wasm_module") // without extension
		tag := c.Param("tag")
		c.File(filesPath + "/" + userOrg + "/" + wasmModule + "/" + tag + "/" + wasmModule + ".info")
	})

	/*
	   ==============================================================
	   Download a wasm module
	   ==============================================================
	   http://localhost:4999/k33g/hello/0.0.0/hello.wasm
	*/
	router.GET("/:user_org/:wasm_module/:tag/:file_name", func(c *gin.Context) {
		userOrg := c.Param("user_org")
		wasmModule := c.Param("wasm_module") // without extension
		tag := c.Param("tag")
		fileName := c.Param("file_name")
		c.File(filesPath + "/" + userOrg + "/" + wasmModule + "/" + tag + "/" + fileName)
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB

	/*
	   ==============================================================
	   Upload a wasm module
	   ==============================================================
	    curl -X POST http://localhost:4999/upload/k33g/hey/0.0.0 \
	      -F "file=@../with-proxy/capsule-hey/hey.wasm" \
	      -F "info=hello hey from @k33g" \
	      -H "Content-Type: multipart/form-data"
	*/
	router.POST("/upload/:user_org/:wasm_module/:tag", func(c *gin.Context) {
		userOrg := c.Param("user_org")
		wasmModule := c.Param("wasm_module") // without extension
		tag := c.Param("tag")

		// single file
		file, _ := c.FormFile("file")
		info, _ := c.GetPostForm("info")
		log.Println(file.Filename)

		// Upload the file to specific destination.
		errMkdir := os.MkdirAll(filesPath+"/"+userOrg+"/"+wasmModule+"/"+tag, os.ModePerm)
		if errMkdir != nil {
			log.Println(errMkdir)
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR_PATH_CREATION", "message": "😡 error when creating the path"})
		}

		err := c.SaveUploadedFile(file, filesPath+"/"+userOrg+"/"+wasmModule+"/"+tag+"/"+wasmModule+".wasm")
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR_WASM_UPLOAD", "message": "😡 error when uploading the wasm module"})

		} else {
			//TODO: add an error message if the info file is not created
			// Wasm module information
			f, errInfo := os.Create(filesPath + "/" + userOrg + "/" + wasmModule + "/" + tag + "/" + wasmModule + ".info")

			if errInfo != nil {
				log.Println(errInfo)
				//c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR_INFO_CREATION", "message": "😡 error when creating the info file"})
			}

			defer f.Close()

			_, errWriteInfo := f.WriteString(info + "\n")

			if errWriteInfo != nil {
				log.Println(errWriteInfo)
				//c.JSON(http.StatusInternalServerError, gin.H{"code": "ERROR_INFO_WRITE", "message": "😡 error when writing to the info file"})
			}

			c.JSON(http.StatusOK, gin.H{"code": "OK", "message": "🙂 " + file.Filename + " uploaded!"})

		}

	})

	if crt != "" {
		// certs/procyon-registry.local.crt
		// certs/procyon-registry.local.key
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") 📦 Wasm Registry is listening on:", httpPort, "🔐🌍")

		router.RunTLS(":"+httpPort, crt, key)
	} else {
		fmt.Println("💊 Capsule (", commons.CapsuleVersion(), ") 📦 Wasm Registry is listening on:", httpPort, "🌍")
		router.Run(":" + httpPort)
	}

}
