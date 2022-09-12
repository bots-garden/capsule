package hostfunctions

import (
	"errors"
	"github.com/bots-garden/capsule/capsulemodule/memory"
	"github.com/bots-garden/capsule/commons"
	"strconv"
	_ "unsafe"
)

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
