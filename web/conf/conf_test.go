package conf

import (
	"io/ioutil"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v2"
)

func TestLoadConfig(t *testing.T) {
	convey.Convey("Given file path to Load Config func", t, func() {
		file := "../example-api.env"
		LoadConfig(file)
		file = "../example-backend.yaml"
		LoadConfig(file)
		var Config CONFIG
		content, err := ioutil.ReadFile(file)
		convey.So(err, convey.ShouldBeNil)
		err = yaml.Unmarshal(content, &Config)
		// convey.So(err, convey.ShouldBeNil)
		// Config.Gitlab.AccessToken = ""

	})

}
