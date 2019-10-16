package main

import (
	"github.com/klec/demo/model"
	"github.com/klec/demo/pkg/log"
	"github.com/klec/demo/po"
	"github.com/klec/demo/po/db"
)

func main() {
	log.Info("Hello demo!")

	po.SetPo(db.NewXormPo())

	me := model.Person{Name: "Allen", Phone: "18758270725"}
	if err := po.GetPo().Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	me = model.Person{Name: "Chenzy", Phone: "18758270725"}
	if err := po.GetPo().Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	input := model.Person{Name: "Allen"}
	if want, err := po.GetPo().Get(&input); err != nil || want == nil {
		log.Error("Get person for name:%s error:%s", input.Name, err)
	} else {
		log.Info("Get %+v", want)
	}

}
