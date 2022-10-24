package hostfunctions

import (
	"errors"
	"fmt"
	"github.com/bots-garden/capsule/commons"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"sync"
)

/*
This is used by the http mode of Capsule
*/

type WasmRequestParam struct {
	JsonData string
	Uri      string
	Method   string
	Headers  string
}

/*
The id of the map is the requestId
*/

var requestParamsMap = sync.Map{}

// TODO: this is a copy paste

// GetHeadersStringFromHeadersRequest :
func GetHeadersStringFromHeadersRequest(c *fiber.Ctx) string {

	headersSlice := commons.CreateSliceFromMap(c.GetReqHeaders())
	headersParameter := commons.CreateStringFromSlice(headersSlice, commons.StrSeparator)

	return headersParameter
}

// StoreRequestParam stores json data + headers + uri + method
// and returns the requestId
// func StoreRequestParams(c *fiber.Ctx) string {
func StoreRequestParams(c *fiber.Ctx) uint32 {

	//reqId := uuid.New().String()
	reqId := uuid.New().ID()

	wrp := WasmRequestParam{
		JsonData: string(c.Body()),
		Headers:  GetHeadersStringFromHeadersRequest(c),
		Uri:      c.Request().URI().String(),
		Method:   c.Method(),
	}

	requestParamsMap.Store(reqId, wrp)

	fmt.Println("🤖[STORE]", wrp, reqId)

	return reqId
}

//func GetRequestParams(reqId string) (WasmRequestParam, error) {
func GetRequestParams(reqId uint32) (WasmRequestParam, error) {

	wrp, ok := requestParamsMap.Load(reqId)

	if ok {

		fmt.Println("🤖[READ]", wrp, reqId)

		return wrp.(WasmRequestParam), nil

	} else {
		//return WasmRequestParam{}, errors.New("😡 unable to delete request " + reqId)
		return WasmRequestParam{}, errors.New("😡 unable to delete request ")

	}
}

//func DeleteRequestParams(reqId string) {
func DeleteRequestParams(reqId uint32) {
	requestParamsMap.Delete(reqId)
}
