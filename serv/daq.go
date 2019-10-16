package serv

import "github.com/klec/demo/db"

type Prps struct {
	Period    float32
	Phase     float32
	Amplitude float32
}

//Daq daq demo
func Daq(po db.DB, prps Prps) (Prps, error) {
	_ = po.Add(prps)

	var want = Prps{Period: prps.Period}
	data, err := po.Get(&want)

	value := data.(Prps)
	value.Period += 1
	value.Phase += 2
	value.Amplitude += 3

	return value, err
}
