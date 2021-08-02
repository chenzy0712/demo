package cmd

import (
	"sync"
	"time"

	"git.kldmp.com/learning/demo/pkg/log"
	"github.com/urfave/cli"
)

var (
	Test = cli.Command{
		Name:        "test",
		Usage:       "demo test <option>",
		Description: "run demo test",
		Subcommands: []cli.Command{
			subCmdTest,
			subCmdLog,
		},
	}

	subCmdTest = cli.Command{
		Name:        "run",
		Usage:       "demo test run",
		Description: "run demo for test",
		Action:      runTest,
	}

	subCmdLog = cli.Command{
		Name:        "log",
		Usage:       "demo test log",
		Description: "run demo for log",
		Action:      runLog,
	}
)

func runTest(c *cli.Context) error {
	var testChan chan int

	//log.Info("Start read nil chan, which will block forever ")
	//_ := <-testChan

	//log.Info("Start send nil chan, which will block forever ")
	//testChan <- 100

	//log.Info("Start send to closed chan, which will panic no matter buffered or unbuffered chan")
	//testChan = make(chan int, 10)
	//close(testChan)
	//testChan <- 100

	//log.Info("Start read from closed unbuffered chan, which will return 0, status false")
	//testChan = make(chan int)
	//go func() {
	//	testChan <- 100 //probable panic cause chan has been closed
	//}()
	//time.Sleep(time.Second)
	//close(testChan)
	//value, ok := <-testChan
	//log.Info("Get value from chan with status:%v, value:%+v", ok, value)

	//log.Info("Start read from closed buffered chan, which will return buffered value until empty, then return 0 and status false")
	//testChan = make(chan int, 10)
	//testChan <- 100 //probable panic cause chan has been closed
	//testChan <- 100 //probable panic cause chan has been closed
	//testChan <- 100 //probable panic cause chan has been closed
	//testChan <- 100 //probable panic cause chan has been closed
	//close(testChan)
	//for i := 0; i < 10; i++ {
	//	value, ok := <-testChan
	//	log.Info("Get value from chan with status:%v, value:%+v", ok, value)
	//	time.Sleep(time.Millisecond * 100)
	//}

	wg := sync.WaitGroup{}
	wg.Add(100)
	log.Info("Start send chan with 2 receiver, the receiver is random.")
	testChan = make(chan int, 10)
	go func() {
		for {
			value, ok := <-testChan
			log.Info("Get1 value from chan with status:%v, value:%+v", ok, value)
			if ok {
				wg.Done()
			}
		}
	}()

	go func() {
		for {
			value, ok := <-testChan
			log.Info("Get2 value from chan with status:%v, value:%+v", ok, value)
			if ok {
				wg.Done()
			}
		}
	}()

	for i := 0; i < 100; i++ {
		testChan <- 100
		time.Sleep(3)
	}
	wg.Wait()

	return nil
}

func runLog(c *cli.Context) error {

	_ = log.NewKLog("console", "", 7, 0)

	//s := fmt.Sprintf("hello %s.\r\n", "world")
	//fmt.Println(s)

	log.KDebug("hello")
	log.KInfo("hello")
	log.KWarn("hello %s", "this")
	log.KError("hello %s %s", "this", "poor")
	log.KFatal("hello %s %s %s!", "this", "poor", "world")
	return nil
}
