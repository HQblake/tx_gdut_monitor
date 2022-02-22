package setting

import "github.com/spf13/viper"

type Setting struct {
	vp    *viper.Viper
	Hosts *HostSetting
}
var (
	setting *Setting
)

// InitSetting 初始化全局配置，使用 viper 库加载configs/config.yaml
func InitSetting(path string) error {
	vp := viper.New()
	vp.SetConfigFile(path)
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	hosts := &HostSetting{}
	err = vp.UnmarshalKey("Hosts", hosts)
	if err != nil {
		return err
	}
	setting = &Setting{vp, hosts}
	return nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}

func GetHostConfig() *HostSetting {
	return setting.Hosts
}
func GetRedisConfig() (*RedisSetting, error) {
	res := &RedisSetting{}
	err := setting.vp.UnmarshalKey("Redis", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetWorkerConfig() (*WorkersSetting, error) {
	res := &WorkersSetting{}
	err := setting.vp.UnmarshalKey("Workers", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func GetAlertConfig() (*AlertConfig, error) {
	res := &AlertConfig{}
	err := setting.vp.UnmarshalKey("Alert", res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
