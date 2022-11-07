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
var MqttGetTopic = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
    topic := mqttconn.GetCapsuleMqttTopic()
    retBuffPtrPos := uint32(params[0])
    retBuffSize := uint32(params[1])
    memory.WriteStringToMemory(topic, ctx, module, retBuffPtrPos, retBuffSize)
    return []uint64{0}
})

var MqttGetServer = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
    server := mqttconn.GetCapsuleMqttServer()
    retBuffPtrPos := uint32(params[0])
    retBuffSize := uint32(params[1])
    memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)
    return []uint64{0}
})

var MqttGetClientId = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {
    server := mqttconn.GetCapsuleMqttClientId()
    retBuffPtrPos := uint32(params[0])
    retBuffSize := uint32(params[1])
    memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)
    return []uint64{0}
})

// MqttConnectPublish :
// only if context is cli or http
var MqttConnectPublish = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

    var stringResultFromHost = ""

    mqttSrvOffset := uint32(params[0])
    mqttSrvByteCount := uint32(params[1])

    mqttSrv := memory.ReadStringFromMemory(ctx, module, mqttSrvOffset, mqttSrvByteCount)

    clientIdPtrOffset := uint32(params[2])
    clientIdByteCount := uint32(params[3])

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
        topicOffset := uint32(params[4])
        topicByteCount := uint32(params[5])

        topic := memory.ReadStringFromMemory(ctx, module, topicOffset, topicByteCount)

        dataOffset := uint32(params[6])
        dataByteCount := uint32(params[7])

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
    retBuffPtrPos := uint32(params[8])
    retBuffSize := uint32(params[9])
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

    return []uint64{0}

})

//TODO: allow to create the connection inside the module

// MqttPublish :
// only if context is mqtt
var MqttPublish = api.GoModuleFunc(func(ctx context.Context, module api.Module, params []uint64) []uint64 {

    mqttClient, errConn := mqttconn.GetCapsuleMqttConn()
    // the connection already exists (we re-used it)
    // it's closed in capsule-launcher/services/mqtt/listen

    var stringResultFromHost = ""

    if errConn != nil {
        //fmt.Println("1️⃣😡", errConn.Error())
        stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

    } else {

        topicOffset := uint32(params[0])
        topicByteCount := uint32(params[1])

        topic := memory.ReadStringFromMemory(ctx, module, topicOffset, topicByteCount)

        dataOffset := uint32(params[2])
        dataByteCount := uint32(params[3])

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
    retBuffPtrPos := uint32(params[4])
    retBuffSize := uint32(params[5])
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

    return []uint64{0}
})
