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

//TODO: NatsGetServer

// NatsGetSubject return the NATS subject of the capsule launcher
func NatsGetSubject(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
    subject := natsconn.GetCapsuleNatsSubject()
    memory.WriteStringToMemory(subject, ctx, module, retBuffPtrPos, retBuffSize)
}

func NatsGetServer(ctx context.Context, module api.Module, retBuffPtrPos, retBuffSize uint32) {
    server := natsconn.GetCapsuleNatsServer()
    memory.WriteStringToMemory(server, ctx, module, retBuffPtrPos, retBuffSize)

}

// NatsConnectPublish :
// only if context is cli or http
func NatsConnectPublish(ctx context.Context, module api.Module, natsSrvOffset, natsSrvByteCount, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

    var stringResultFromHost = ""

    natsSrv := memory.ReadStringFromMemory(ctx, module, natsSrvOffset, natsSrvByteCount)

    natscn, errConn := nats.Connect(natsSrv)
    defer natscn.Close()

    if errConn != nil {
        //fmt.Println("1️⃣😡", errConn.Error())
        stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

    } else {
        subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)
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
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}

// NatsConnectRequest :
// only if context is cli or http (???)
// ref: https://docs.nats.io/using-nats/developer/sending/request_reply
func NatsConnectRequest(ctx context.Context, module api.Module, natsSrvOffset, natsSrvByteCount, subjectOffset, subjectByteCount, dataOffset, dataByteCount, timeoutSecondDuration, retBuffPtrPos, retBuffSize uint32) {

    var stringResultFromHost = ""

    natsSrv := memory.ReadStringFromMemory(ctx, module, natsSrvOffset, natsSrvByteCount)

    natscn, errConn := nats.Connect(natsSrv)
    defer natscn.Close()

    if errConn != nil {
        //fmt.Println("1️⃣😡", errConn.Error())
        stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

    } else {
        subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)
        data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

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
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}

//TODO: allow to create the connection inside the module

// ----- when the module is a subscriber (nats mode) -----

// NatsPublish :
// only if context is nats (the module is a subscriber)
func NatsPublish(ctx context.Context, module api.Module, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {

    nc, errConn := natsconn.GetCapsuleNatsConn()
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

// NatsReply :
// only if context is nats (the module is a subscriber)
func NatsReply(ctx context.Context, module api.Module, dataOffset, dataByteCount, timeoutSecondDuration, retBuffPtrPos, retBuffSize uint32) {

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
        data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

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
    // Write the new string stringResultFromHost to the "shared memory"
    memory.WriteStringToMemory(stringResultFromHost, ctx, module, retBuffPtrPos, retBuffSize)

}
