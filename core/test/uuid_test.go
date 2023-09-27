package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	fmt.Println(uuid.NewV4().String())
}
