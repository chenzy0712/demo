package model

import (
	"git.kldmp.com/learning/demo/internal/po"
	"git.kldmp.com/learning/demo/pkg/log"
)

type Person struct {
	Name  string
	Email string
}

func Demo() {
	me := Person{Name: "Allen", Email: "hello@gmail.com"}
	if err := po.GetPo().Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	me = Person{Name: "Steve", Email: "hello@gmail.com"}
	if err := po.GetPo().Add(me); err != nil {
		log.Error("Add me failed error:%s", err)
	}

	input := Person{Name: "Allen"}
	if want, err := po.GetPo().Get(&input); err != nil || want == nil {
		log.Error("Get person for name:%s error:%s", input.Name, err)
	} else {
		log.Info("Get %+v", want)
	}
}
