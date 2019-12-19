package models

type Param struct {
	Key   string
	Value string
}

type ParamArray []Param

type Params struct {
	ParamArray
}

func (ps Params) Get(name string) (string, bool) {
	for _, entry := range ps.ParamArray {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

func (ps Params) ByName(name string) (va string) {
	va, _ = ps.Get(name)
	return
}

func (ps Params) Set(param *Param) {
	for _, entry := range ps.ParamArray {
		if entry.Key == param.Key {
			entry.Value = param.Value
			return
		}
	}
	ps.ParamArray = append(ps.ParamArray, *param)
}
