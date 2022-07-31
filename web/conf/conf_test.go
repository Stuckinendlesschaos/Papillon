package conf

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLoadConfig(t *testing.T) {
	convey.Convey("Given yaml file path to Load Config func", t, func() {
		file := "../example-backend.yaml"
		LoadConfig(file)
	})
	// goconvey can't execute in following statement  because sys_call
	convey.Convey("Given error file path to Load Config func:", t, func() {
		file := "../example-backend.env"
		LoadConfig(file)
	})

}
