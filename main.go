package main

import (
	"github.com/klec/demo/pkg/log"
	"github.com/klec/demo/po"
)

func main() {
	log.Info("Hello demo!")

	var db po.XDatabase
	db.Init()

	me := po.Person{Name: "Allen", Phone: "18758270725"}
	if err := me.Add(); err != nil {
		log.Error("Add me failed error:%s", err)
	}
}
