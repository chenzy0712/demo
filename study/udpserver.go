package study

import (
	"math"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"

	"github.com/klec/demo/pkg/log"
)

const (
	UDPPacketLen = 1300
)

var (
	bufferPool sync.Pool
	wg         sync.WaitGroup
	total      uint64
	pps        uint64
	doIt       int32
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

		payloadHandler(data[:n])
		bufferPool.Put(data)
	}
}

func statistics() {
	go signalHandler()

	//flushTicker := time.NewTicker(time.Duration(10) * time.Second)
	//for range flushTicker.C {
	//	// log.Printf("Ops/s %f", float64(ops)/flushInterval.Seconds())
	//	atomic.AddUint64(&total, pps)
	//	log.Info("total:%d, pp5s:%d", atomic.LoadUint64(&total), atomic.LoadUint64(&pps))
	//	atomic.StoreUint64(&pps, 0)
	//}
}

func signalHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	for range c {
		atomic.StoreInt32(&doIt, 0)
		atomic.AddUint64(&total, pps)
		time.Sleep(5 * time.Second)
		log.Info("total:%d, pp5s:%d", atomic.LoadUint64(&total), atomic.LoadUint64(&pps))
		wg.Done()
	}
}

func UDPClient(addr string, interval int) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Info("Try to connect addr:%s error:%s, exit now!", addr, err)
		return
	}
	defer conn.Close()

	wg.Add(1)
	mockSinData := make([]byte, UDPPacketLen)
	mockCosData := make([]byte, UDPPacketLen)
	for i := 0; i < len(mockSinData); i++ {
		mockSinData[i] = byte(math.Sin(1.8 * float64(i)))
		mockCosData[i] = byte(math.Cos(1.8 * float64(i)))
	}

	go statistics()
	atomic.StoreInt32(&doIt, 1)

	for atomic.LoadInt32(&doIt) == 1 {
		_, err := conn.Write(mockSinData)
		if err != nil {
			log.Error("Try to write mock data to server error:%s", err)
			return
		}
		atomic.AddUint64(&pps, uint64(1))
		time.Sleep(time.Microsecond * time.Duration(interval))
	}

	wg.Wait()
}
