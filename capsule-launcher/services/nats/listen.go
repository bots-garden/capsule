package capsulenats

import (
    "context"
    "fmt"
    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
    capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
    "github.com/bots-garden/capsule/commons"
    "github.com/nats-io/nats.go"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

/*
callHandleOnMessage
callHandlePublish => host function


about hf.NatsPublish => create a hf.Connect(string) too
+ onLoad?
*/

func Listen(natssrv string, subject string, wasmFile []byte) {

    // Create context that listens for the interrupt signal from the OS.
    ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    hostfunctions.HostInformation = `{"natsServer":"` + natssrv + `","capsuleVersion":"` + commons.CapsuleVersion() + `"}`

    capsule.CallExportedOnLoad(wasmFile)

    nc, err := InitNatsConn(natssrv)
    defer nc.Close()

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    } else {

        go func() {
            // Simple Async Subscriber
            _, err := nc.Subscribe(subject, func(m *nats.Msg) {
                fmt.Printf("👋 Received a message: %s\n", string(m.Data))

                wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForNats(wasmFile)
                defer wasmRuntime.Close(ctx)

                params := string(m.Data)

                paramsPos, paramsLen, free, err := capsule.ReserveMemorySpaceFor(params, wasmModule, ctx)
                defer free.Call(ctx, paramsPos)

                err = capsule.ExecHandleVoidFunction(wasmFunction, wasmModule, ctx, paramsPos, paramsLen)

                //TODO: change the error handling
                if err != nil {
                    log.Panicf("out of range of memory size")
                }
            })
            if err != nil {
                fmt.Println(err.Error())
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
