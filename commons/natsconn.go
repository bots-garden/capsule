package commons

import (
	"github.com/nats-io/nats.go"
)

var natsconn *nats.Conn
var err error

func GetCapsuleNatsConn() (*nats.Conn, error) {
	return natsconn, err
}

func InitNatsConn(natssrv string) (*nats.Conn, error) {

	if natsconn == nil {
		natsconn, err = nats.Connect(natssrv)
	}

	return natsconn, err

}
