package study

import (
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klec/demo/pkg/log"
)

const (
	UDPPacketLen = 1400
)

var (
	bufferPool sync.Pool
	wg         sync.WaitGroup
	total      uint64
	pps        uint64
)

//UDPServer receive UDP data and send to handler
func UDPServer(addr string) {
	log.Info("Start UDP server to receive data from %s", addr)
	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, UDPPacketLen) },
	}

	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Error("Try to listen addr:%s err:%s", addr, err)
		return
	}

	wg.Add(1)
	go receive(conn)
	go statistics()

	wg.Wait()

	log.Info("UDP server exit!")
}

func receive(conn net.PacketConn) {
	defer conn.Close()

	for {
		data := bufferPool.Get().([]byte)
		n, addr, err := conn.ReadFrom(data)
		if err != nil {
			log.Error("Read data from addr:%s err:%s", addr, err)
			break
		}

		//log.Debug("n:%d", n)
		payloadHandler(data[:n])
		bufferPool.Put(data)
	}
}

func statistics() {
	flushTicker := time.NewTicker(time.Duration(5) * time.Second)
	for range flushTicker.C {
		// log.Printf("Ops/s %f", float64(ops)/flushInterval.Seconds())
		atomic.AddUint64(&total, pps)
		log.Info("total:%d, pp5s:%d", total, pps)
		atomic.StoreUint64(&pps, 0)
	}
}
