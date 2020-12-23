package cmd

import (
	"context"
	"sync"
	"time"

	"github.com/klec/demo/pkg/log"

	"github.com/urfave/cli"
)

var (
	Ctx = cli.Command{
		Name:        "context",
		Usage:       "demo context <option>",
		Description: "run demo of context",
		Flags:       []cli.Flag{},
		Action:      runContext,
	}
)

func runContext(c *cli.Context) error {
	log.InitLog("console", "", "info", 1)

	//doContext()

	//doCancelInDeferForMistakeUsage()
	//someHandler()
	timeoutHandler()
	log.Info("end")
	//time.Sleep(10 * time.Second)
	//log.Info("exit")

	return nil
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Info("done")
			return
		default:
			log.Info("work")
		}
	}
}

func doTimeoutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			log.Info("deadline set")
			if time.Now().After(deadline) {
				log.Info("%s", ctx.Err().Error())
				return
			}
		}

		select {
		case <-ctx.Done():
			log.Info("done")
			return
		default:
			log.Info("work")
		}

	}
}

func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	go doStuff(ctx)

	//go doTimeoutStuff(ctx)

	time.Sleep(10 * time.Second)
}

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//cancel doStuff 10 seconds later
	time.Sleep(10 * time.Second)
	cancel()
}

func doWgStuff(ctx context.Context, wg sync.WaitGroup) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			log.Info("done")
			wg.Done()
			return
		default:
			log.Info("work")
		}
	}

}
func doCancelInDeferForMistakeUsage() {
	var wg sync.WaitGroup

	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go doWgStuff(ctx, wg)
	defer cancel()

	wg.Wait()
}

func doContext() {
	log.Info("start ctx")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Info("ctx done")
	}
}
