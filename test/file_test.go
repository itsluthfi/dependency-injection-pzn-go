package test

import (
	"rest-api-pzn-go/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("database")
	assert.NotNil(t, connection)

	cleanup()
}
