package test

import (
	"cloud-disk/core/etc/helper"
	"fmt"
	"testing"
)

func TestRandomCode(t *testing.T) {
	fmt.Println(helper.RandomCode())
}
