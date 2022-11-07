package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

//export hostMqttGetTopic
//go:linkname hostMqttGetTopic
func hostMqttGetTopic(retBuffPtrPos **byte, retBuffSize *int) uint32

func MqttGetTopic() string {
	var buffPtr *byte
	var buffSize int

	hostMqttGetTopic(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return getStringResult(buffPtr, buffSize)
}

//export hostMqttGetServer
//go:linkname hostMqttGetServer
func hostMqttGetServer(retBuffPtrPos **byte, retBuffSize *int) uint32

func MqttGetServer() string {
	var buffPtr *byte
	var buffSize int

	hostMqttGetServer(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return getStringResult(buffPtr, buffSize)
}

//export hostMqttGetClientId
//go:linkname hostMqttGetClientId
func hostMqttGetClientId(retBuffPtrPos **byte, retBuffSize *int) uint32

func MqttGetClientId() string {
	var buffPtr *byte
	var buffSize int

	hostMqttGetClientId(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return getStringResult(buffPtr, buffSize)
}

//export hostMqttConnectPublish
//go:linkname hostMqttConnectPublish
func hostMqttConnectPublish(mqttSrvPtrPos, mqttSrvSize, clientIdPtrPos, clientIdSize, topicPtrPos, topicSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int) uint32

func MqttConnectPublish(mqttSrv, clientId, topic, data string) (string, error) {

	mqttSrvPtrPos, mqttSrvSize := getStringPtrPositionAndSize(mqttSrv)
	clientIdPtrPos, clientIdSize := getStringPtrPositionAndSize(clientId)
	topicPtrPos, topicSize := getStringPtrPositionAndSize(topic)
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostMqttConnectPublish(mqttSrvPtrPos, mqttSrvSize, clientIdPtrPos, clientIdSize, topicPtrPos, topicSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		resultStr = valueStr
	}
	return resultStr, err

}

//export hostMqttPublish
//go:linkname hostMqttPublish
func hostMqttPublish(topicPtrPos, topicSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int) uint32

// MqttPublish :
// Publish data on mqtt topic

func MqttPublish(topic string, data string) (string, error) {
	// transform the parameters for the host function
	topicPtrPos, topicSize := getStringPtrPositionAndSize(topic)
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostMqttPublish(topicPtrPos, topicSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := getStringResult(buffPtr, buffSize)

	// check the return value
	if commons.IsErrorString(valueStr) {
		errorMessage, errorCode := commons.GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		resultStr = valueStr
	}
	return resultStr, err

}
