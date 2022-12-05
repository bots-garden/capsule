package hostfunctions

import (
    "context"
    "strconv"

    "github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
    "github.com/bots-garden/capsule/commons"
    "github.com/tetratelabs/wazero/api"

    "github.com/go-redis/redis/v9"
)

var redisDb *redis.Client

// TODO: handle errors

func InitRedisCli() {
    if redisDb == nil {
        defaultDb, _ := strconv.Atoi(commons.GetEnv("REDIS_DEFAULTDB", "0"))
        //fmt.Println("📦 redisdb connection ...")
        redisDb = redis.NewClient(&redis.Options{
            Addr:     commons.GetEnv("REDIS_ADDR", "localhost:6379"),
            Password: commons.GetEnv("REDIS_PWD", ""), // no password set
            DB:       defaultDb,                       // use default DB
        })
    }
}
func getRedisCli() *redis.Client {
    return redisDb
}

// RedisSet :
// The wasm module will call this function like this:
// `func RedisSet(key string, value) (string, error)`
var RedisSet = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

    InitRedisCli()
    //=========================================================
    // Read arguments values of the function call
    // get strings from the wasm module function (from memory)
    //=========================================================
    keyPosition := uint32(stack[0])
    keyLength := uint32(stack[1])
    keyStr := memory.ReadStringFromMemory(ctx, module, keyPosition, keyLength)

    valuePosition := uint32(stack[2])
    valueLength := uint32(stack[3])
    valueStr := memory.ReadStringFromMemory(ctx, module, valuePosition, valueLength)

    //==[👋 Implementation: Start]=============================
    err := getRedisCli().Set(ctx, keyStr, valueStr, 0).Err()

    var stringResultFromHost = ""

    if err != nil {
        stringResultFromHost = commons.CreateStringError(err.Error(), 0)
        // if code 0 don't display code in the error message
    } else {
        stringResultFromHost = "[OK](" + keyStr + ":" + valueStr + ")"
    }
    //==[👋 Implementation: End]===============================

    positionReturnBuffer := uint32(stack[4])
    lengthReturnBuffer := uint32(stack[5])
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

    stack[0] = 0 // return 0
})

// RedisGet :
// The wasm module will call this function like this:
// `func RedisGet(key string) (string, error)`
var RedisGet = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

    InitRedisCli()
    //=========================================================
    // Read arguments values of the function call
    // get strings from the wasm module function (from memory)
    //=========================================================
    keyPosition := uint32(stack[0])
    keyLength := uint32(stack[1])
    keyStr := memory.ReadStringFromMemory(ctx, module, keyPosition, keyLength)

    //==[👋 Implementation: Start]=============================
    valueStr, err := getRedisCli().Get(ctx, keyStr).Result()

    var stringResultFromHost = ""

    if err != nil {
        stringResultFromHost = commons.CreateStringError(err.Error(), 0)
        // if code 0 don't display code in the error message
    } else {
        stringResultFromHost = valueStr
    }
    //==[👋 Implementation: End]===============================
    positionReturnBuffer := uint32(stack[2])
    lengthReturnBuffer := uint32(stack[3])
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

    stack[0] = 0 // return 0
})

// RedisKeys : get all the keys (with a pattern)
// The wasm module will call this function like this:
// `func RedisKeys(pattern string) (string, error)`
var RedisKeys = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
    InitRedisCli()
    // get the pattern parameter
    patternPosition := uint32(stack[0])
    patternLength := uint32(stack[1])
    patternStr := memory.ReadStringFromMemory(ctx, module, patternPosition, patternLength)

    // call the redis KEYS command
    valueStr, err := getRedisCli().Keys(ctx, patternStr).Result()

    var stringResultFromHost = ""

    if err != nil {
        stringResultFromHost = commons.CreateStringError(err.Error(), 0)
        // if code 0 don't display code in the error message
    } else {
        stringResultFromHost = commons.CreateStringFromSlice(valueStr, commons.StrSeparator)
    }

    positionReturnBuffer := uint32(stack[2])
    lengthReturnBuffer := uint32(stack[3])
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, positionReturnBuffer, lengthReturnBuffer)

    stack[0] = 0 // return 0

})
