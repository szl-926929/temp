// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package tutorial

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
	shared0 "github.com/suitable/gen-go/shared"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var _ = shared0.GoUnusedProtection__

const INT32CONSTANT = 9853

var MAPCONSTANT map[string]string

func init() {
	MAPCONSTANT = map[string]string{
		"hello":     "world",
		"goodnight": "moon",
	}

}