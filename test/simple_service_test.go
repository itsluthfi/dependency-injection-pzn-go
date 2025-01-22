package test

import (
	"fmt"
	"rest-api-pzn-go/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	fmt.Println(simpleService)
	fmt.Println(err)
}
