package natsconn

import "github.com/nats-io/nats.go"

var natsconn *nats.Conn
var natsErr error
var natsSubject string
var nastServer string

func GetCapsuleNatsConn() (*nats.Conn, error) {
	return natsconn, natsErr
}

func InitNatsConn(natssrv string) (*nats.Conn, error) {
	if natsconn == nil {
		natsconn, natsErr = nats.Connect(natssrv)
	}
	return natsconn, natsErr
}

func SetCapsuleNatsSubject(natssub string) {
	natsSubject = natssub
}
func GetCapsuleNatsSubject() string {
	return natsSubject
}

func SetCapsuleNatsServer(natssrv string) {
	nastServer = natssrv
}
func GetCapsuleNatsServer() string {
	return nastServer
}
