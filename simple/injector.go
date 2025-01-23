//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) { // return valuenya dipilih sesuai sama provider yang mau dipake, dan ini dikasih injector parameter isError
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBar() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// habis itu eksekusi pake cmd wire di folder package yg ada injectornya, file generatenya punya nama wire_gen.go
