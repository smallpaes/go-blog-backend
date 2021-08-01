package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New() // create a new viper instance
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/") // path to look for config file
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig() // find and read config file
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
