package shared

import (
	"github.com/KevinZonda/GinTemplate/lib/moduleMgr"
	"github.com/KevinZonda/GinTemplate/logger"
)

func InitLogger() {
	_mgr.RegisterModule("logger", moduleMgr.ModuleStatusInit, moduleMgr.NewModuleParam())

	logger.InitLogger()

	_mgr.UpdateModuleStatus("logger", moduleMgr.ModuleStatusLoaded)
}
