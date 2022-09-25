package hostfunctions

import (
	"context"
	"fmt"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/bots-garden/capsule/mqttconn"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tetratelabs/wazero/api"
)

// MqttGetTopic return the MQTT Topic of the capsule launcher
func MqttGetTopic(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	topic := mqttconn.GetCapsuleMqttTopic()
	memory.WriteStringToMemory(topic, ctx, module, retBuffPtrPos, retBuffSize)
}

func MqttGetServer(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	server := mqttconn.GetCapsuleMqttServer()
	memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)

}

func MqttGetClientId(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	server := mqttconn.GetCapsuleMqttClientId()
	memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)

}

// MqttConnectPublish :
// only if context is cli or http
func MqttConnectPublish(ctx context.Context, module api.Module, mqttSrvOffset, mqttSrvByteCount, clientIdPtrOffset, clientIdByteCount, topicOffset, topicByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

	var stringResultFromHost = ""

	mqttSrv := memory.ReadStringFromMemory(ctx, module, mqttSrvOffset, mqttSrvByteCount)
	mqttClientId := memory.ReadStringFromMemory(ctx, module, clientIdPtrOffset, clientIdByteCount)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", mqttSrv))
	opts.SetClientID(mqttClientId)
	mqttClient := mqtt.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	errConn := token.Error()
	defer mqttClient.Disconnect(250)

	if errConn != nil {
		fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		topic := memory.ReadStringFromMemory(ctx, module, topicOffset, topicByteCount)
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		token := mqttClient.Publish(topic, 0, false, data)
		token.Wait()

		errPub := token.Error()

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = "[OK](" + topic + ":" + data + ")"
		}
	}
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}

//TODO: allow to create the connection inside the module

// MqttPublish :
// only if context is mqtt
func MqttPublish(ctx context.Context, module api.Module, topicOffset, topicByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

	mqttClient, errConn := mqttconn.GetCapsuleMqttConn()
	// the connection already exists (we re-used it)
	// it's closed in capsule-launcher/services/mqtt/listen

	var stringResultFromHost = ""

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		topic := memory.ReadStringFromMemory(ctx, module, topicOffset, topicByteCount)
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		token := mqttClient.Publish(topic, 0, false, data)
		token.Wait()

		errPub := token.Error()

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = "[OK](" + topic + ":" + data + ")"
		}
	}
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
