package serv

import (
	"errors"
	"testing"

	"github.com/klec/demo/po"

	"github.com/klec/demo/po/mocks"
	"github.com/stretchr/testify/assert"
)

func newMockPo() *mocks.PO {
	return &mocks.PO{}
}

func TestDaq(t *testing.T) {
	m := newMockPo()

	prps := Prps{Period: 1, Phase: 1.8, Amplitude: 12.34}
	query := Prps{Period: 1, Phase: 0, Amplitude: 0}
	want := Prps{Period: 2, Phase: 3.8, Amplitude: 15.34}

	m.On("Init").Return()
	m.On("Add", prps).Return(nil)
	m.On("Get", &query).Return(prps, nil)

	po.SetPo(m)
	got, err := Daq(prps)
	assert.Equal(t, err, nil)
	assert.Equal(t, got, want)
}

func TestPoGetFailure(t *testing.T) {
	m := newMockPo()

	m.On("Init").Return()
	m.On("Get", nil).Return(nil, errors.New("error"))

	po.SetPo(m)
	e, err := po.GetPo().Get(nil)
	assert.NotEmpty(t, err)
	assert.Empty(t, e)
}
