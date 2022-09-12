package hostfunctions

import (
	"context"
	"github.com/bots-garden/capsule/capsule-launcher/hostfunctions/memory"
	"github.com/bots-garden/capsule/commons"
	"github.com/nats-io/nats.go"
	"github.com/tetratelabs/wazero/api"
)

var natsconn *nats.Conn

func InitNatsConn() (*nats.Conn, error) {

	if natsconn == nil {
		//fmt.Println("🙂 NATS_SRV:", commons.GetEnv("NATS_SRV", "nats.devsecops.fun:4222"))

		natsconn, err := nats.Connect(commons.GetEnv("NATS_SRV", "nats.devsecops.fun:4222"))
		// nats.DefaultURL

		/*
			if err != nil {
				fmt.Println("😡", err.Error())
			} else {
				fmt.Println("🙂", "connection ok")
				fmt.Println("ConnectedServerId", nc.ConnectedServerId())
			}
		*/

		return natsconn, err
	} else {
		return natsconn, nil
	}

}

func GetNewNatsConn() (*nats.Conn, error) {
	natsconn, err := nats.Connect(commons.GetEnv("NATS_SRV", "nats.devsecops.fun:4222"))
	return natsconn, err
}

func getNatsConn() *nats.Conn {
	return natsconn
}

//TODO: allow to create the connection inside the module

// NatsPublish :
// only if context is cli or http
// if context is nats, use capsuleNatsConn
func NatsPublish(ctx context.Context, module api.Module, subjectOffset, subjectByteCount, dataOffset, dataByteCount, retBuffPtrPos, retBuffSize uint32) {
	nc, errConn := GetNewNatsConn()
	//TODO: not sure that's that work
	defer nc.Close()

	var stringResultFromHost = ""

	if errConn != nil {
		//fmt.Println("1️⃣😡", errConn.Error())
		stringResultFromHost = commons.CreateStringError(errConn.Error(), 0)

	} else {
		subject := memory.ReadStringFromMemory(ctx, module, subjectOffset, subjectByteCount)
		data := memory.ReadStringFromMemory(ctx, module, dataOffset, dataByteCount)

		/*
		   fmt.Println("subject:", subject)
		   fmt.Println("data", data)
		   fmt.Println(natsconn.ConnectedServerId())
		*/

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

	/*
	   Question: do I really need: retBuffPtrPos, retBuffSize
	*/

}
