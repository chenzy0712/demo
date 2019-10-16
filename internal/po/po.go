package po

type PO interface {
	Init()
	Add(interface{}) error
	Get(data interface{}) (interface{}, error)
}

var x PO

func SetPo(po PO) {
	x = po
	x.Init()
}

func GetPo() PO {
	return x
}
