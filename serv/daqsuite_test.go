package serv

import (
	"errors"
	"testing"

	"github.com/klec/demo/po"
	"github.com/klec/demo/po/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DaqTestSuite struct {
	suite.Suite
	m *mocks.PO
}

func (daq *DaqTestSuite) SetupSuite() {
	daq.m = newMockPo()

	daq.m.On("Init")
	po.SetPo(daq.m)
}

func (daq *DaqTestSuite) TearDownSuite() {
}

func (daq *DaqTestSuite) SetupTest() {
}

func (daq *DaqTestSuite) TestPoGetFailure() {
	daq.m.On("Get", nil).Return(nil, errors.New("no table exist"))

	e, err := po.GetPo().Get(nil)

	assert.NotEmpty(daq.T(), err)
	assert.Empty(daq.T(), e)
}

func (daq *DaqTestSuite) TestDaq() {
	prps := Prps{Period: 1, Phase: 1.8, Amplitude: 12.34}
	query := Prps{Period: 1, Phase: 0, Amplitude: 0}
	want := Prps{Period: 2, Phase: 3.8, Amplitude: 15.34}

	daq.m.On("Add", prps).Return(nil)
	daq.m.On("Get", &query).Return(prps, nil)

	got, err := Daq(prps)

	assert.Equal(daq.T(), err, nil)
	assert.Equal(daq.T(), got, want)

}

func TestDaqSuite(t *testing.T) {
	suite.Run(t, new(DaqTestSuite))
}
