package models

type IAuthorization interface {
	Authenticated() bool
	HasRole(roles ...string) bool
}
