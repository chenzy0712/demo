package serv

import (
	"testing"

	"github.com/klec/demo/db/mocks"
	"github.com/stretchr/testify/assert"
)

func newMockPo() *mocks.DB {
	return &mocks.DB{}
}

func TestDaq(t *testing.T) {
	po := newMockPo()

	prps := Prps{Period: 1, Phase: 1.8, Amplitude: 12.34}
	query := Prps{Period: 1, Phase: 0, Amplitude: 0}
	want := Prps{Period: 2, Phase: 3.8, Amplitude: 15.34}

	po.On("Add", prps).Return(nil)
	po.On("Get", &query).Return(prps, nil)

	got, err := Daq(po, prps)
	assert.Equal(t, err, nil)
	assert.Equal(t, got, want)
}
