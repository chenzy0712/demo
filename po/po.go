package po

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/klec/demo/pkg/log"
	_ "github.com/mattn/go-sqlite3"

	"github.com/go-xorm/xorm"
)

type DB interface {
	Init()
	GetPerson()
}

var (
	xDatabase *XDatabase
	mDatabse  *MDatabase
)

type XDatabase struct {
	x *xorm.Engine
}

func (e *XDatabase) Init() {
	if engine, err := xorm.NewEngine("sqlite3", "demo.db"); err != nil {
		log.Error("Init sqlite3 PO failed:%s", err)
		panic(err)
	} else {
		e.x = engine
	}

	//init tables
	_ = e.x.Sync2(Person{})
	xDatabase = e
}

type MDatabase struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (m *MDatabase) Init() {
	if db, mock, err := sqlmock.New(); err != nil {
		log.Error("Init Mock PO failed:%s", err)
		panic(err)
	} else {
		m.db = db
		m.mock = mock
	}
	mDatabse = m

}

type Person struct {
	Name  string
	Phone string
}

func (p *Person) Add() error {
	_, err := xDatabase.x.Insert(p)
	if err != nil {
		log.Error("Insert Person failed with error:%s", err)
	}

	return err
}
