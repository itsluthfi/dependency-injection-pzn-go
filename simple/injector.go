//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

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

// salah
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBarStruct() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}

// habis itu eksekusi pake cmd wire di folder package yg ada injectornya, file generatenya punya nama wire_gen.go
