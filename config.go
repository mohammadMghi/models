package models

type IConfig interface {
	InitializeConfig(input interface{})
}
