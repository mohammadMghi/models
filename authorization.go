package models

type IAuthorization interface {
	Initialize(authorization IAuthorization)
	GetBase() IAuthorization
	Authenticated() bool
	GetCurrentAccount() interface{}
	GetCurrentAccountId() interface{}
	HasRole(roles ...string) bool
}
