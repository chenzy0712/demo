package main

import (
	"github.com/klec/demo/model"
	"github.com/klec/demo/pkg/log"
	"github.com/klec/demo/po"
	"github.com/klec/demo/po/db"
	"github.com/klec/demo/study"
)

func main() {
	log.Info("Hello demo!")

	po.SetPo(db.NewXormPo())

	model.Demo()

	study.InterfaceDemo()
}
