package capsulenats

import (
	"github.com/nats-io/nats.go"
)

var natsconn *nats.Conn

func GetCapsuleNatsConn() *nats.Conn {
	return natsconn
}

func InitNatsConn(natssrv string) (*nats.Conn, error) {

	if natsconn == nil {
		nc, err := nats.Connect(natssrv)

		return nc, err
	} else {
		return natsconn, nil
	}

}
