package models

type IAuthorization interface {
	Initialize(request IRequest, authorization IAuthorization)
	GetBase() IAuthorization
	Authenticated() bool
	GetCurrentAccount(request IRequest) interface{}
	GetCurrentAccountId(request IRequest) interface{}
	HasRole(roles ...string) bool
}
