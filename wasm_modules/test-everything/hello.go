package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
	"github.com/tidwall/gjson"
	"strings"

	/* create json string */
	"github.com/tidwall/sjson"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func testMemoryFunctions() {
	hf.Log("===== Testing Memory functions =====")
	// +++ MemoryGet +++
	capsuleVersion, _ := hf.MemoryGet("capsule_version")
	hf.Log("🖐 Version: " + capsuleVersion)

	memSetRes, memSetErr := hf.MemorySet("message", "hello")
	if memSetErr != nil {
		hf.Log("😡 error:" + memSetErr.Error())
		hf.Log("-> Test: MemorySet: 🔴")
	} else {
		hf.Log(memSetRes)
		hf.Log("-> Test: MemorySet: 🟢")
	}
	memGetRes, memGetErr := hf.MemoryGet("message")
	if memGetErr != nil {
		hf.Log("😡 error:" + memGetErr.Error())
		hf.Log("-> Test: MemoryGet: 🔴")
	} else {
		hf.Log("message is: " + memGetRes)
		if memGetRes == "hello" {
			hf.Log("-> Test: MemoryGet: 🟢")
		} else {
			hf.Log("😡 error: not the expected value")
			hf.Log("-> Test: MemoryGet: 🔴")
		}
	}
	memKeys, memKeysErr := hf.MemoryKeys()
	// it will return an array of strings
	if memKeysErr != nil {
		hf.Log("😡 error:" + memKeysErr.Error())
		hf.Log("-> Test: MemoryKeys: 🔴")
	} else {
		for _, value := range memKeys {
			hf.Log("key: " + value)
		}
		hf.Log("-> Test: MemoryKeys: 🟢")
	}
}

func testGetHostInformation() {
	hf.Log("===== Testing GetHostInformation =====")

	// +++ GetHostInformation +++
	hostInformation := hf.GetHostInformation()
	hf.Log(hostInformation)

	if strings.HasPrefix(hostInformation, `{"httpPort":`) && strings.Contains(hostInformation, "capsuleVersion") {
		hf.Log("-> Test: GetHostInformation: 🟢")
	} else {
		hf.Log("-> Test: GetHostInformation: 🔴")
	}
	// {"httpPort":7070,"capsuleVersion":"v0.2.9 🦜 [parrot]"}
}

func testHttp() {
	hf.Log("===== Testing Http =====")

	// +++ Http GET +++
	headers := map[string]string{"Accept": "application/json", "Content-Type": "text/html; charset=UTF-8"}

	retGet, err := hf.Http("https://httpbin.org/get", "GET", headers, "[GET]👋 hello world 🌍")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
		hf.Log("-> Test: Http GET: 🔴")
	} else {
		hf.Log("💊👋 Return value from the module: " + retGet)
		hf.Log("-> Test: Http GET: 🟢")
	}

	// +++ Http POST +++
	retPost, err := hf.Http("https://httpbin.org/post", "POST", headers, "[POST]👋 hello world 🌍")
	if err != nil {
		hf.Log("😡 error:" + err.Error())
		hf.Log("-> Test: Http POST: 🔴")
	} else {
		hf.Log("💊👋 Return value from the module: " + retPost)
		hf.Log("-> Test: Http POST: 🟢")
	}

}

func testRequestParams(request hf.Request) {
	hf.Log("===== Testing Request =====")

	// TODO make the conditional tests
	hf.Log("📝 Body: " + request.Body)
	// {"message": "Golang 💚💜 wasm", "author": "Philippe"}
	hf.Log("📝 URI: " + request.Uri)
	// http://localhost:7070/
	hf.Log("📝 Method: " + request.Method)
	// Method: POST

	author := gjson.Get(request.Body, "author")
	// Philippe
	message := gjson.Get(request.Body, "message")
	// Golang 💚💜 wasm
	hf.Log("👋 " + message.String() + " by " + author.String() + " 😄")

	hf.Log("Content-Type: " + request.Headers["Content-Type"])
	// application/json; charset=utf-8
	hf.Log("Content-Length: " + request.Headers["Content-Length"])
	// 57
	hf.Log("User-Agent: " + request.Headers["User-Agent"])
	// User-Agent: curl/7.84.0
}

func testGetEnv() {
	hf.Log("===== Testing GetEnv =====")

	envMessage, err := hf.GetEnv("MESSAGE")
	if err != nil {
		hf.Log("😡 " + err.Error())
		hf.Log("-> Test: GetEnv: 🔴")

	} else {
		hf.Log("Environment variable: " + envMessage)
		hf.Log("-> Test: GetEnv: 🟢")

	}
}

//TODO test if Redis is connected
func testRedisFunctions() {
	hf.Log("===== Testing Redis functions =====")

	// +++ RedisSet +++
	// add a key, value
	res1, redisSetErr := hf.RedisSet("greetings", "Hello World")
	if redisSetErr != nil {
		hf.Log("😡 error:" + redisSetErr.Error())
		hf.Log("-> Test: RedisSet: 🔴")
	} else {
		hf.Log("RedisSet:" + res1)
		hf.Log("-> Test: RedisSet: 🟢")
	}

	// +++ RedisGet +++
	// read the value
	res2, redisGetErr := hf.RedisGet("greetings")
	if redisGetErr != nil {
		hf.Log("😡 error:" + redisGetErr.Error())
		hf.Log("-> Test: RedisGet: 🔴")
	} else {
		hf.Log("🎉 value: " + res2)
		if res2 == "Hello World" {
			hf.Log("-> Test: RedisGet: 🟢")
		} else {
			hf.Log("😡 error: not the expected value")
			hf.Log("-> Test: RedisGet: 🔴")
		}
	}

	// +++ RedisKeys +++
	_, _ = hf.RedisSet("bob1", "Bob One")
	_, _ = hf.RedisSet("bob2", "Bob Two")

	legion, redisKeysErr := hf.RedisKeys("bob*")
	if redisKeysErr != nil {
		hf.Log("😡 error:" + redisKeysErr.Error())
		hf.Log("-> Test: RedisKeys: 🔴")

	} else {
		for _, bob := range legion {
			hf.Log(bob)
		}
		hf.Log("-> Test: RedisKeys: 🟢")
	}
	//TODO: test if only 2 records

}

func Handle(request hf.Request) (response hf.Response, errResp error) {
	/*
	   bodyReq = {"author":"Philippe","message":"Golang 💚 wasm"}
	*/

	// TODO: make an array of errors and print a test report at the end

	testRequestParams(request)

	testMemoryFunctions()
	testGetHostInformation()
	testHttp()
	testGetEnv()
	testRedisFunctions()

	headersResp := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Message":      "👋 hello world 🌍",
	}

	jsondoc := `{"message": "", "author": ""}`
	jsondoc, _ = sjson.Set(jsondoc, "message", "👋 hey! What's up?")
	jsondoc, _ = sjson.Set(jsondoc, "author", "Bob")

	return hf.Response{Body: jsondoc, Headers: headersResp}, nil
}

/*
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json' \
  -d '{"message": "Golang 💚 wasm"}'
*/
