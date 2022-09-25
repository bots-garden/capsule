package capsulemqtt

import (
	"context"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions"
	capsule "github.com/bots-garden/capsule/capsule-launcher/services/wasmrt"
	"github.com/bots-garden/capsule/commons"
	"github.com/bots-garden/capsule/mqttconn"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

func Listen(mqttSrv, mqttClientId, mqttTopic string, wasmFile []byte) {
	// Store the MQTT topic, server and clientId
	mqttconn.SetCapsuleMqttTopic(mqttTopic)
	mqttconn.SetCapsuleMqttServer(mqttSrv)
	mqttconn.SetCapsuleMqttClientId(mqttClientId)

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	hostfunctions.HostInformation = `{"mqttServer":"` + mqttSrv + `","capsuleVersion":"` + commons.CapsuleVersion() + `"}`

	capsule.CallExportedOnLoad(wasmFile)

	client, err := mqttconn.InitMqttConn(mqttSrv, mqttClientId, setHandler(wasmFile))
	defer client.Disconnect(250)

	if err != nil {
		StoreExitError("initialize MQTT conn", err, 1, wasmFile)
		os.Exit(1)
	}

	// Subscribe
	token := client.Subscribe(mqttTopic, 1, nil)
	token.Wait()
	err = token.Error()
	if err != nil {
		StoreExitError("initialize MQTT subscribe", err, 1, wasmFile)
		os.Exit(1)
	}

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

func setHandler(wasmFile []byte) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		wasmRuntime, wasmModule, wasmFunction, ctx := capsule.GetNewWasmRuntimeForMqtt(wasmFile)
		defer wasmRuntime.Close(ctx)

		params := string(msg.Payload())

		paramsPos, paramsLen, free, err := capsule.ReserveMemorySpaceFor(params, wasmModule, ctx)
		defer free.Call(ctx, paramsPos)

		err = capsule.ExecHandleVoidFunction(wasmFunction, wasmModule, ctx, paramsPos, paramsLen)

		if err != nil {
			StoreExitError("call MQTT ExecHandleVoidFunction (callMQTTMessageHandle)", err, 1, wasmFile)
			os.Exit(1)
			//log.Panicf("out of range of memory size")
		}
	}
}
