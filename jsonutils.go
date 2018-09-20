package merak

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var once *sync.Once

var EasyJson = jsoniter.ConfigCompatibleWithStandardLibrary
