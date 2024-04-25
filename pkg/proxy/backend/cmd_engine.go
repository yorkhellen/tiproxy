package backend

import (
	"github.com/go-mysql-org/go-mysql/mysql"
	pnet "github.com/pingcap/tiproxy/pkg/proxy/net"
)

type CmdEngine interface {
	executeCmd(request []byte, clientIO, backendIO *pnet.PacketIO, waitingRedirect bool) (holdRequest bool, err error)
	query(packetIO *pnet.PacketIO, sql string) (result *mysql.Resultset, response []byte, err error)
	finishedTxn() bool
	setCapability(capability pnet.Capability)
	getCapability() pnet.Capability
}
