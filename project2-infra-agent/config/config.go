package config

import (
	"github.com/spf13/viper"
	"strings"
)

var Config = viper.New()

var replacer = strings.NewReplacer(".", "_")

func init() {
	Config.SetEnvKeyReplacer(replacer)
	Config.SetEnvPrefix("agent")
	Config.AutomaticEnv()
}

var GetString = Config.GetString

func Sub(key string) (sub *viper.Viper) {

	sub = Config.Sub(key)
  	sub.SetEnvKeyReplacer(replacer)
	sub.SetEnvPrefix("AGENT_" + replacer.Replace(key))
	sub.AutomaticEnv()

	return
}
