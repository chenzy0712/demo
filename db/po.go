package db

import (
	"reflect"

	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/klec/demo/model"
	"github.com/klec/demo/pkg/log"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-xorm/xorm"
)

type DB interface {
	Init()
	Add(interface{}) error
	Get(data interface{}) (interface{}, error)
}

type XDatabase struct {
	x *xorm.Engine
}

//NewXormPo create PO with xorm
func NewXormPo() *XDatabase {
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
