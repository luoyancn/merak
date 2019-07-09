package merak

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var once sync.Once
var EasyJson jsoniter.API

func init() {
	once.Do(func() {
		extra.SupportPrivateFields()
		EasyJson = jsoniter.ConfigCompatibleWithStandardLibrary
	})
}
