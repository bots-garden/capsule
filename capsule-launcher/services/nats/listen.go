package capsulenats

import (
	"context"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"github.com/bots-garden/capsule/natsconn"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StoreExitError(from string, err error, exitCode int, wasmFile []byte) {
	fmt.Println("🔴 [store exit error for wasm module] from:", from)
	fmt.Println("🔴 Error:", err.Error())
	// store error information for the wasm module
	commons.SetExitError(err.Error())
	commons.SetExitCode(1)
	capsule.CallExportedOnExit(wasmFile)
}

func Listen(natssrv string, subject string, wasmFile []byte) {
	// Store the Nats subject and server
	natsconn.SetCapsuleNatsSubject(subject)
	natsconn.SetCapsuleNatsServer(natssrv)

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	hostfunctions.HostInformation = `{"natsServer":"` + natssrv + `","capsuleVersion":"` + commons.CapsuleVersion() + `"}`

	capsule.CallExportedOnLoad(wasmFile)

	nc, err := natsconn.InitNatsConn(natssrv)
	defer nc.Close()

	if err != nil {
		StoreExitError("initialize NATS conn", err, 1, wasmFile)
		os.Exit(1)
	} else {

		go func() {

			// Simple Async Subscriber
			_, err := nc.Subscribe(subject, func(m *nats.Msg) {
				// this part is triggered when we have a message

				//fmt.Printf("👋 Received a message: %s\n", string(m.Data))
				// call `callNatsMessageHandle`
				wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForNats(wasmFile)
				defer wasmRuntime.Close(ctx)

				params := string(m.Data)

				paramsPos, paramsLen, free, err := capsule.ReserveMemorySpaceFor(params, wasmModule, ctx)
				defer free.Call(ctx, paramsPos)

				err = capsule.ExecHandleVoidFunction(wasmFunction, wasmModule, ctx, paramsPos, paramsLen)

				if err != nil {
					StoreExitError("call NATS ExecHandleVoidFunction (callNatsMessageHandle)", err, 1, wasmFile)
					os.Exit(1)
					//log.Panicf("out of range of memory size")
				}
			})

			if err != nil {
				StoreExitError("subscribe NATS subject", err, 1, wasmFile)
				os.Exit(1)
			}
		}()

		// Listen for the interrupt signal.
		<-ctx.Done()

		// Restore default behavior on the interrupt signal and notify user of shutdown.
		stop()
		fmt.Println("💊 Capsule shutting down gracefully ...")

		capsule.CallExportedOnExit(wasmFile)

		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		fmt.Println("💊 Capsule exiting")
	}

}
