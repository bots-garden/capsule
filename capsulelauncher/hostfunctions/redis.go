package hostfunctions

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/bots-garden/capsule/capsulelauncher/commons"
	"github.com/bots-garden/capsule/capsulelauncher/hostfunctions/memory"
	"github.com/tetratelabs/wazero/api"

	"github.com/go-redis/redis/v9"
)

var redisDb *redis.Client

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitRedisCli() {
    if redisDb==nil {
        defaultDb, _ := strconv.Atoi(getEnv("REDIS_DEFAULTDB", "0"))
        fmt.Println("📦 redisdb connection ...")
        redisDb = redis.NewClient(&redis.Options{
            Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
            Password: getEnv("REDIS_PWD", ""), // no password set
            DB:       defaultDb,               // use default DB
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
