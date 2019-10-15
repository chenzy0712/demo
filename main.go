package main

import (
	"github.com/klec/demo/db"
	"github.com/klec/demo/model"
	"github.com/klec/demo/pkg/log"
)

func main() {
	log.Info("Hello demo!")

	po := db.NewXormPo()
	po.Init()

	me := model.Person{Name: "Allen", Phone: "18758270725"}
	if err := po.Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	me = model.Person{Name: "Chenzy", Phone: "18758270725"}
	if err := po.Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	input := &model.Person{Name: "Allen"}
	if want, err := po.Get(input); err != nil || want == nil {
		log.Error("Get person for name:%s error:%s", input.Name, err)
	} else {
		log.Info("Get %+v", want)
	}

}
