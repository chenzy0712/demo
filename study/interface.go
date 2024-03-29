package study

import "git.kldmp.com/learning/demo/pkg/log"

type TCP interface {
	Send()
}

type UDP interface {
	BoardCast()
}

type UDPStruct struct {
}

//func (u *UDPStruct) Send() {
//	log.Info("Hello, UDP Send()")
//}

func (u *UDPStruct) BoardCast() {
	log.Info("Hello, UDP BoardCast()")
}

type NETStruct struct {
	UDPStruct
}

func (n *NETStruct) Send() {
	log.Info("Hello, TCP Send()")
}

func InterfaceDemo() {
	var (
		tcp TCP
		net *NETStruct
	)

	net = &NETStruct{}
	net.BoardCast()
	net.Send()

	tcp = net
	tcp.Send()
	//since set net to tcp, the method of BoardCast has been discarded
	//tcp.BoardCast()
}
