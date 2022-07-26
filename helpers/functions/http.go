// host functions
package hf

import (
	"errors"
	"strings"
)

//export hostHttp
//go:linkname hostHttp
func hostHttp(urlOffset uint32, urlByteCount uint32, methodOffSet uint32, methodByteCount uint32, headersOffSet uint32, headersMethodByteCount uint32, bodyOffset uint32, bodyByteCount uint32, retBuffPtrPos **byte, retBuffSize *int)

func Http(url, method string, headers []string, body string) (string, error) {
	// Get parameters from the wasm module
	// Prepare parameters for the host function call
	urlStrPos, urlStrSize := GetStringPtrPositionAndSize(url)
	methodStrPos, methodStrSize := GetStringPtrPositionAndSize(method)
	headersStrPos, headersStrSize := GetStringPtrPositionAndSize(strings.Join(headers[:], "|"))
	bodyStrPos, bodyStrSize := GetStringPtrPositionAndSize(body)

	// This will be used to retrieve the return value (result)
	var buffPtr *byte
	var buffSize int

  // call the host function
	hostHttp(urlStrPos, urlStrSize, methodStrPos, methodStrSize, headersStrPos, headersStrSize, bodyStrPos, bodyStrSize, &buffPtr, &buffSize)

  var resultStr = ""
  var err error
	valueStr := GetStringResult(buffPtr, buffSize)
  if(strings.HasPrefix(valueStr, "[ERR]")) {

    errorMessage := strings.Split(valueStr,"[ERR]")[1]

    err = errors.New(errorMessage)
  } else {
    resultStr = valueStr
  }

	return resultStr, err
}
