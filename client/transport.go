package client

import (
	"github.com/qmsk/snmpbot/snmp"
	"io"
	"net"
)

var EOF = io.EOF

type IO struct {
	Addr    net.Addr
	Packet  snmp.Packet
	PDUType snmp.PDUType
	PDU     snmp.PDU
}

type Transport interface {
	Send(IO) error
	Recv() (IO, error)
	Close() error
}
