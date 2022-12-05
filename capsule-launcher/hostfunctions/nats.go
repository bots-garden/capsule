package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/bots-garden/capsule/natsconn"
	"github.com/nats-io/nats.go"
	"github.com/tetratelabs/wazero/api"
	"time"
)

// NatsGetSubject return the NATS subject of the capsule launcher
var NatsGetSubject = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	subject := natsconn.GetCapsuleNatsSubject()
	retBuffPtrPos := uint32(stack[0])
	retBuffSize := uint32(stack[1])
	memory.WriteStringToMemory(subject, ctx, module, retBuffPtrPos, retBuffSize)
	stack[0] = 0 // return 0
})

var NatsGetServer = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {
	server := natsconn.GetCapsuleNatsServer()
	retBuffPtrPos := uint32(stack[0])
	retBuffSize := uint32(stack[1])
	memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)
	stack[0] = 0 // return 0
})

// NatsConnectPublish :
// only if context is cli or http
var NatsConnectPublish = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	var stringResultFromHost = ""

	natsSrvOffset := uint32(stack[0])
	natsSrvByteCount := uint32(stack[1])
	natsSrv := memory.ReadStringFromMemory(ctx, module, natsSrvOffset, natsSrvByteCount)

	natscn, errConn := nats.Connect(natsSrv)
	defer natscn.Close()

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		subjectOffset := uint32(stack[2])
		subjectByteCount := uint32(stack[3])
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)

		dataOffset := uint32(stack[4])
		dataByteCount := uint32(stack[5])
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		errPub := natscn.Publish(subject, []byte(data))

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = "[OK](" + subject + ":" + data + ")"
		}
	}

	retBuffPtrPos := uint32(stack[6])
	retBuffSize := uint32(stack[7])
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

	stack[0] = 0 // return 0

})

// NatsConnectRequest :
// only if context is cli or http (???)
// ref: https://docs.nats.io/using-nats/developer/sending/request_reply
var NatsConnectRequest = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	var stringResultFromHost = ""

	natsSrvOffset := uint32(stack[0])
	natsSrvByteCount := uint32(stack[1])
	natsSrv := memory.ReadStringFromMemory(ctx, module, natsSrvOffset, natsSrvByteCount)

	natscn, errConn := nats.Connect(natsSrv)
	defer natscn.Close()

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {

		subjectOffset := uint32(stack[2])
		subjectByteCount := uint32(stack[3])
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)

		dataOffset := uint32(stack[4])
		dataByteCount := uint32(stack[5])
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		timeoutSecondDuration := uint32(stack[6])

		replyMsg, errPub := natscn.Request(subject, []byte(data), time.Duration(timeoutSecondDuration)*time.Second) // one second timeout

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			// // Use the response
			stringResultFromHost = string(replyMsg.Data)
		}
	}

	retBuffPtrPos := uint32(stack[7])
	retBuffSize := uint32(stack[8])

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

	stack[0] = 0 // return 0

})

//TODO: allow to create the connection inside the module
// ----- when the module is a subscriber (nats mode) -----

// NatsPublish :
// only if context is nats (the module is a subscriber)
// func NatsPublish(ctx context.Context, module api.Module, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {
var NatsPublish = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	nc, errConn := natsconn.GetCapsuleNatsConn()
	// the connection already exists (we re-used it)
	// it's closed in capsule-launcher/services/nats/listen

	var stringResultFromHost = ""

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {

		subjectOffset := uint32(stack[0])
		subjectByteCount := uint32(stack[1])
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)

		dataOffset := uint32(stack[2])
		dataByteCount := uint32(stack[3])
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		errPub := nc.Publish(subject, []byte(data))

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = "[OK](" + subject + ":" + data + ")"
		}
	}

	retBuffPtrPos := uint32(stack[4])
	retBuffSize := uint32(stack[5])

	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

	stack[0] = 0 // return 0

})

// NatsReply :
// only if context is nats (the module is a subscriber)
var NatsReply = api.GoModuleFunc(func(ctx context.Context, module api.Module, stack []uint64) {

	//nc, errConn := natsconn.GetCapsuleNatsConn()
	// the connection already exists (we re-used it)
	// it's closed in capsule-launcher/services/nats/listen

	sub, errSub := natsconn.GetCapsuleNatsSubscription()

	var stringResultFromHost = ""

	if errSub != nil {
		stringResultFromHost = commons.CreateStringError(errSub.Error(), 0)

	} else {
		// subjectOffset, subjectByteCount,
		//subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)

		dataOffset := uint32(stack[0])
		dataByteCount := uint32(stack[1])
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		timeoutSecondDuration := uint32(stack[2])

		msg, errMsg := sub.NextMsg(time.Duration(timeoutSecondDuration) * time.Second)

		//errPub := nc.Publish(subject, []byte(data))

		if errMsg != nil {
			stringResultFromHost = commons.CreateStringError(errMsg.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			errResp := msg.Respond([]byte(data))
			if errResp != nil {
				stringResultFromHost = commons.CreateStringError(errResp.Error(), 0)

			} else {
				stringResultFromHost = "[OK](" + natsconn.GetCapsuleNatsSubject() + ":" + data + ")"
			}
		}
	}

	retBuffPtrPos := uint32(stack[3])
	retBuffSize := uint32(stack[4])
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

	stack[0] = 0 // return 0

})
