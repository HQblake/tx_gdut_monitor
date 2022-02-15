package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp    *viper.Viper
	Hosts *HostSetting
}

// NewSetting 使用 viper 库加载configs/config.yaml
func NewSetting(path string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigFile(path)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	hosts := &HostSetting{}
	vp.UnmarshalKey("Hosts", hosts)
	return &Setting{vp, hosts}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
