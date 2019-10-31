package study

import (
	"math"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klec/demo/pkg/log"
)

const (
	MulticastPacketLen = 1024
)

//MulticastServer receive UDP data and send to handler
func MulticastServer(addr string) {
	log.Info("Start UDP server to receive data from %s", addr)
	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, MulticastPacketLen) },
	}

	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Error("%s", err)
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, udpAddr)
	if err != nil {
		log.Error("Try to listen addr:%s err:%s", addr, err)
		return
	}

	_ = conn.SetReadBuffer(8192)

	wg.Add(1)
	go receive(conn)
	go statistics()

	wg.Wait()

	log.Info("UDP multicast server exit!")
}

func MulticastClient(addr string, interval int) {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Error("%s", err)
	}

	conn, err := net.DialUDP("udp4", nil, udpAddr)
	if err != nil {
		log.Info("Try to connect addr:%s error:%s, exit now!", addr, err)
		return
	}
	defer conn.Close()

	wg.Add(1)
	mockSinData := make([]byte, MulticastPacketLen)
	mockCosData := make([]byte, MulticastPacketLen)
	for i := 0; i < len(mockSinData); i++ {
		mockSinData[i] = byte(math.Sin(1.8 * float64(i)))
		mockCosData[i] = byte(math.Cos(1.8 * float64(i)))
	}

	go statistics()
	atomic.StoreInt32(&doIt, 1)

	for atomic.LoadInt32(&doIt) == 1 || atomic.LoadUint64(&pps) < 1000000 {
		_, err := conn.Write(mockSinData)
		if err != nil {
			log.Error("Try to write mock data to server error:%s", err)
			log.Info("total:%d, pp5s:%d", atomic.LoadUint64(&total), atomic.LoadUint64(&pps))
			return
		}
		atomic.AddUint64(&pps, uint64(1))
		time.Sleep(time.Microsecond * time.Duration(interval))
	}

	wg.Wait()
}
