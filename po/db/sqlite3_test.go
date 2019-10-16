package db

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/klec/demo/pkg/log"
)

var (
	mDatabse *MDatabase
)

type MDatabase struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

//NewMockPo create PO with mock
func NewMockPo() *MDatabase {
	return &MDatabase{}
}

func (m *MDatabase) Init() {
	if db, mock, err := sqlmock.New(); err != nil {
		log.Error("Init Mock PO failed:%s", err)
		panic(err)
	} else {
		m.db = db
		m.mock = mock
	}
}
