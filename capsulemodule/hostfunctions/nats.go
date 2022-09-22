package hostfunctions

import (
	"errors"
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
	return getStringResult(buffPtr, buffSize)
}

//export hostNatsGetServer
//go:linkname hostNatsGetServer
func hostNatsGetServer(retBuffPtrPos **byte, retBuffSize *int)

func NatsGetServer() string {
	var buffPtr *byte
	var buffSize int

	hostNatsGetServer(&buffPtr, &buffSize)

	// return the string result of the host function calling
	return getStringResult(buffPtr, buffSize)
}

//export hostNatsConnectPublish
//go:linkname hostNatsConnectPublish
func hostNatsConnectPublish(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int)

func NatsConnectPublish(natsSrv string, subject string, data string) (string, error) {

	natsSrvPtrPos, natsSrvSize := getStringPtrPositionAndSize(natsSrv)
	subjectPtrPos, subjectSize := getStringPtrPositionAndSize(subject)
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsConnectPublish(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

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

//export hostNatsConnectRequest
//go:linkname hostNatsConnectRequest
func hostNatsConnectRequest(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize, timeoutSecondDuration uint32, retBuffPtrPos **byte, retBuffSize *int)

func NatsConnectRequest(natsSrv string, subject string, data string, timeoutSecondDuration uint32) (string, error) {

	natsSrvPtrPos, natsSrvSize := getStringPtrPositionAndSize(natsSrv)
	subjectPtrPos, subjectSize := getStringPtrPositionAndSize(subject)
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsConnectRequest(natsSrvPtrPos, natsSrvSize, subjectPtrPos, subjectSize, dataPtrPos, dataSize, timeoutSecondDuration, &buffPtr, &buffSize)

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

//export hostNatsPublish
//go:linkname hostNatsPublish
func hostNatsPublish(subjectPtrPos, subjectSize, dataPtrPos, dataSize uint32, retBuffPtrPos **byte, retBuffSize *int)

// NatsPublish :
// Publish data on nats topic

func NatsPublish(subject string, data string) (string, error) {
	// transform the parameters for the host function
	subjectPtrPos, subjectSize := getStringPtrPositionAndSize(subject)
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsPublish(subjectPtrPos, subjectSize, dataPtrPos, dataSize, &buffPtr, &buffSize)

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

//export hostNatsReply
//go:linkname hostNatsReply
func hostNatsReply(dataPtrPos, dataSize, timeoutSecondDuration uint32, retBuffPtrPos **byte, retBuffSize *int)

// NatsPublish :
// Publish data on nats topic

func NatsReply(data string, timeoutSecondDuration uint32) (string, error) {
	// transform the parameters for the host function
	dataPtrPos, dataSize := getStringPtrPositionAndSize(data)

	var buffPtr *byte
	var buffSize int

	// call the host function
	// the result will be available in memory thanks to ` &buffPtr, &buffSize`
	hostNatsReply(dataPtrPos, dataSize, timeoutSecondDuration, &buffPtr, &buffSize)

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
