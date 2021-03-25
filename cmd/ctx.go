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
		Action:      runContextTest,
	}
)

func runContextTest(c *cli.Context) error {
	log.InitLog("console", "", "info", 1)

	//callTimeoutCtx()

	//doCancelInDeferForMistakeUsage()
	//activeDoCtxCancel()
	passiveCallCancel()
	log.Info("runContextTest finished")
	//time.Sleep(10 * time.Second)
	//log.Info("exit")

	return nil
}

func doDealLineStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			log.Info("deadline had set")
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

func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case cc := <-ctx.Done():
			log.Info("doStuff done %v", cc)
			return
		default:
			log.Info("doStuff work")
		}
	}
}

//passiveCallCancel passive call cancel() for timeout/deadline
func passiveCallCancel() {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	//go doStuff(ctx)

	go doDealLineStuff(ctx)

	time.Sleep(10 * time.Second)
	log.Info("run passiveCallCancel defer")
}

//activeDoCtxCancel() active call cancel()
func activeDoCtxCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//cancel doStuff 10 seconds later
	time.Sleep(10 * time.Second)
	cancel()
}

//callTimeoutCtx wait timeout cancel in main routine
func callTimeoutCtx() {
	log.Info("start timeout ctx")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Info("timeout ctx done")
	}
}

/*
WaitGroup work with Context
*/
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
