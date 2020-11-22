package mymap

type SharedMap struct {
	Table map[string]string
}

func (mapFunc *SharedMap) Register(name string) string {
	mapFunc.Table[name] = "registrado! " + name
	return mapFunc.Table[name]
}

func (mapFunc SharedMap) Lookup(name string) string {
	return mapFunc.Table[name]
}
