package serv

import (
	"github.com/klec/demo/po"
)

type Prps struct {
	Period    float32
	Phase     float32
	Amplitude float32
}

//Daq daq demo
func Daq(prps Prps) (Prps, error) {
	_ = po.GetPo().Add(prps)

	var want = Prps{Period: prps.Period}
	data, err := po.GetPo().Get(&want)

	value := data.(Prps)
	value.Period += 1
	value.Phase += 2
	value.Amplitude += 3

	return value, err
}
