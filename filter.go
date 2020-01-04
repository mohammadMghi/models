package models

type Filters map[string]interface{}

func (f *Filters) Add(key string, value interface{}) {
	(*f)[key] = value
}

func (f *Filters) Delete(key string) {
	delete(*f, key)
}

func (f *Filters) Extend(filters *Filters) {
	for key, value := range *filters {
		f.Add(key, value)
	}
}
