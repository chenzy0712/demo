package db

import (
	"reflect"

	"git.kldmp.com/learning/demo/internal/model"
	"git.kldmp.com/learning/demo/internal/po"

	"git.kldmp.com/learning/demo/pkg/log"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type XDatabase struct {
	x *xorm.Engine
}

//NewXormPo create PO with xorm
func NewXormPo() po.PO {
	return &XDatabase{}
}

func (e *XDatabase) Init() {
	if engine, err := xorm.NewEngine("sqlite3", "demo.db"); err != nil {
		log.Error("Init sqlite3 PO failed:%s", err)
		panic(err)
	} else {
		e.x = engine
	}

	//init tables
	_ = e.x.Sync2(model.Person{})

}

func (e *XDatabase) Add(data interface{}) error {
	_, err := e.x.Insert(data)
	if err != nil {
		log.Error("Insert %s failed with error:%s", reflect.TypeOf(data), err)
	}

	return err
}

func (e *XDatabase) Get(data interface{}) (interface{}, error) {
	has, err := e.x.Get(data)
	if err != nil {
		log.Error("Query %v failed with error:%s!", reflect.TypeOf(data), err)
		return nil, err
	} else if !has {
		log.Error("No record of %v exist!", data, err)
		return nil, err
	}

	return data, nil
}
