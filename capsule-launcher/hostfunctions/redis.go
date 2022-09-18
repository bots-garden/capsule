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
func RedisSet(ctx context.Context, module api.Module, keyOffset, keyByteCount, valueOffSet, valueByteCount, retBuffPtrPos, retBuffSize uint32) {

	InitRedisCli()
	//=========================================================
	// Read arguments values of the function call
	// get strings from the wasm module function (from memory)
	//=========================================================
	keyStr := memory.ReadStringFromMemory(ctx, module, keyOffset, keyByteCount)
	valueStr := memory.ReadStringFromMemory(ctx, module, valueOffSet, valueByteCount)

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

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}

// RedisGet :
// The wasm module will call this function like this:
// `func RedisGet(key string) (string, error)`
func RedisGet(ctx context.Context, module api.Module, keyOffset, keyByteCount, retBuffPtrPos, retBuffSize uint32) {

	InitRedisCli()
	//=========================================================
	// Read arguments values of the function call
	// get strings from the wasm module function (from memory)
	//=========================================================

	keyStr := memory.ReadStringFromMemory(ctx, module, keyOffset, keyByteCount)

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

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}

// RedisKeys : get all the keys (with a pattern)
// The wasm module will call this function like this:
// `func RedisKeys(pattern string) (string, error)`
func RedisKeys(ctx context.Context, module api.Module, patternOffset, patternByteCount, retBuffPtrPos, retBuffSize uint32) {
	InitRedisCli()
	// get the pattern parameter
	patternStr := memory.ReadStringFromMemory(ctx, module, patternOffset, patternByteCount)

	// call the redis KEYS command
	valueStr, err := getRedisCli().Keys(ctx, patternStr).Result()

	var stringResultFromHost = ""

	if err != nil {
		stringResultFromHost = commons.CreateStringError(err.Error(), 0)
		// if code 0 don't display code in the error message
	} else {
		stringResultFromHost = commons.CreateStringFromSlice(valueStr, commons.StrSeparator)
	}

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)
}
