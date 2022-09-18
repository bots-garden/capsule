package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/nats-io/nats.go"
	"github.com/tetratelabs/wazero/api"
)

//TODO: NatsGetServer

// NatsGetSubject return the NATS subject of the capsule launcher
func NatsGetSubject(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	subject := commons.GetCapsuleNatsSubject()
	memory.WriteStringToMemory(subject, ctx, module, retBuffPtrPos, retBuffSize)
}

func NatsGetServer(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
	server := commons.GetCapsuleNatsServer()
	memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)

}

// NatsConnectPublish :
// only if context is cli or http
func NatsConnectPublish(ctx context.Context, module api.Module, natsSrvOffset, natsSrvByteCount, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

	var stringResultFromHost = ""

	natsSrv := memory.ReadStringFromMemory(ctx, module, natsSrvOffset, natsSrvByteCount)

	natsconn, errConn := nats.Connect(natsSrv)
	defer natsconn.Close()

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		errPub := natsconn.Publish(subject, []byte(data))

		if errPub != nil {
			//fmt.Println("2️⃣😡", errPub.Error())
			stringResultFromHost = commons.CreateStringError(errPub.Error(), 0)
			// if code 0 don't display code in the error message
		} else {
			stringResultFromHost = "[OK](" + subject + ":" + data + ")"
		}
	}
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}

//TODO: allow to create the connection inside the module

// NatsPublish :
// only if context is nats
func NatsPublish(ctx context.Context, module api.Module, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

	nc, errConn := commons.GetCapsuleNatsConn()
	// the connection already exists (we re-used it)
	// it's closed in capsule-launcher/services/nats/listen

	var stringResultFromHost = ""

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)
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
	// Write the new string stringResultFromHost to the "shared memory"
	memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
