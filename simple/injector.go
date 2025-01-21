//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService() (*SimpleService, error) { // return valuenya dipilih sesuai sama provider yang mau dipake
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

// habis itu eksekusi pake cmd wire di folder package yg ada injectornya, file generatenya punya nama wire_gen.go
