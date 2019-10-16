package main

import (
	"github.com/klec/demo/internal/model"
	"github.com/klec/demo/internal/po"
	"github.com/klec/demo/internal/po/db"
	"github.com/klec/demo/pkg/log"
	"github.com/klec/demo/study"
)

func main() {
	log.Info("Hello demo!")

	po.SetPo(db.NewXormPo())

	model.Demo()

	study.InterfaceDemo()
}
