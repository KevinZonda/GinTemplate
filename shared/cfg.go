package shared

import (
	"encoding/json"

	"github.com/KevinZonda/GinTemplate/lib/moduleMgr"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/GoX/pkg/panicx"
	"github.com/KevinZonda/GoX/pkg/ruby"
)

type Config struct {
	Addr  string `json:"addr"`
	Debug bool   `json:"debug"`
}

var _cfg *Config

func GetConfig() *Config {
	return _cfg
}

func LoadConfig(bs []byte) error {
	return json.Unmarshal(bs, &_cfg)
}

func InitCfg(path string) {
	_mgr.RegisterModule("cfg", moduleMgr.ModuleStatusInit, moduleMgr.NewModuleParam().With("path", path))

	bs := ruby.RdrErr(iox.ReadAllByte(path))
	panicx.NotNilErr(LoadConfig(bs))

	_mgr.UpdateModuleStatus("cfg", moduleMgr.ModuleStatusLoaded)
}
