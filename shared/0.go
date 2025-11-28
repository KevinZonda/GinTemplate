package shared

import (
	"github.com/KevinZonda/GinTemplate/lib/moduleMgr"
	"github.com/KevinZonda/GinTemplate/logger"
	"github.com/sirupsen/logrus"
)

var _mgr *moduleMgr.Manager = moduleMgr.New().WithLogger(logger.StdLogger())

// You shouldn't modify `InitAll` and `InitNoCfg`. You can only modify `InitEssential` and `InitNonEssential`.
func InitAll(path string) {
	InitCfg(path)
	InitNoCfg()
	PrintModuleMgrStatus()
}

func InitNoCfg() {
	InitEssential()
	InitNonEssential()
}

func InitEssential() {
	InitLogger()
}

func InitNonEssential() {
	InitGin()
}

func PrintModuleMgrStatus() {
	modules := _mgr.Modules()
	var moduleNames = []string{}
	for _, module := range modules {
		moduleNames = append(moduleNames, module.Name)
	}
	logger.StdLogger().WithFields(logrus.Fields{
		"modules": moduleNames,
		"count":   len(modules),
	}).Info("[ModuleMgr] Modules Loaded Information")
	for _, module := range modules {
		logger.StdLogger().WithFields(logrus.Fields{
			"module": module.Name,
			"status": module.Status,
			"params": module.Params,
		}).Info("[ModuleMgr] Module Status")
	}
}
