// host functions
package hf

import "strings"

//export hostHttp
//go:linkname hostHttp
func hostHttp(urlOffset, urlByteCount, methodOffSet, methodByteCount, headersOffSet, headersMethodByteCount, bodyOffset, bodyByteCount uint32, retBuffPtrPos **byte, retBuffSize *int)

func Http(url, method string, headers []string, body string) string {
  // Get parameters from the wasm module
  // Prepare parameters for the host function call
  urlStrPos, urlStrSize := GetStringPtrPositionAndSize(url)
  methodStrPos, methodStrSize := GetStringPtrPositionAndSize(method)
  headersStrPos, headersStrSize := GetStringPtrPositionAndSize(strings.Join(headers[:], "|"))
  bodyStrPos, bodyStrSize := GetStringPtrPositionAndSize(body)

  // This will be used to retrieve the return value (result)
  var bufPtr *byte
	var bufSize int

  hostHttp(urlStrPos, urlStrSize, methodStrPos, methodStrSize, headersStrPos, headersStrSize, bodyStrPos, bodyStrSize, &bufPtr, &bufSize)

  result := GetStringResult(bufPtr, bufSize)

  return result
}
