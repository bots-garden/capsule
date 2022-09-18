package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

//export hostNatsGetSubject
//go:linkname hostNatsGetSubject
func hostNatsGetSubject(retBuffPtrPos **byte, retBuffSize *int)

func NatsGetSubject() string {
	var buffPtr *byte
	var buffSize int

	hostNatsGetSubject(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return memory.GetStringResult(buffPtr, buffSize)
}

//export hostNatsGetServer
//go:linkname hostNatsGetServer
func hostNatsGetServer(retBuffPtrPos **byte, retBuffSize *int)

func NatsGetServer() string {
	var buffPtr *byte
	var buffSize int

	hostNatsGetServer(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return memory.GetStringResult(buffPtr, buffSize)
}

//export hostNatsConnectPublish
//go:linkname hostNatsConnectPublish
func hostNatsConnectPublish(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int)

func NatsConnectPublish(natsSrv string, subject string, data string) (string, error) {

	natsSrvPtrPos, natsSrvSize := memory.GetStringPtrPositionAndSize(natsSrv)
	subjectPtrPos, subjectSize := memory.GetStringPtrPositionAndSize(subject)
	dataPtrPos, dataSize := memory.GetStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsConnectPublish(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := memory.GetStringResult(buffPtr, buffSize)

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

//export hostNatsPublish
//go:linkname hostNatsPublish
func hostNatsPublish(subjectPtrPos, subjectSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int)

// NatsPublish :
// Publish data on nats topic

func NatsPublish(subject string, data string) (string, error) {
	// transform the parameters for the host function
	subjectPtrPos, subjectSize := memory.GetStringPtrPositionAndSize(subject)
	dataPtrPos, dataSize := memory.GetStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsPublish(subjectPtrPos, subjectSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

	// transform the result to a string
	var resultStr = ""
	var err error
	valueStr := memory.GetStringResult(buffPtr, buffSize)

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
