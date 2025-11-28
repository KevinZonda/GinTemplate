package moduleMgr

import (
	"encoding/json"
	"strconv"
)

type ModuleStatus string

const (
	ModuleStatusInit     ModuleStatus = "init"
	ModuleStatusLoaded   ModuleStatus = "loaded"
	ModuleStatusDegraded ModuleStatus = "degraded"
	ModuleStatusError    ModuleStatus = "error"
)

type ModuleParams map[string]string

func (p ModuleParams) With(key string, value string) ModuleParams {
	p[key] = value
	return p
}

func (p ModuleParams) WithBool(key string, value bool) ModuleParams {
	p[key] = strconv.FormatBool(value)
	return p
}

func (p ModuleParams) WithInt(key string, value int) ModuleParams {
	p[key] = strconv.Itoa(value)
	return p
}

func (p ModuleParams) WithFp64(key string, value float64) ModuleParams {
	p[key] = strconv.FormatFloat(value, 'f', -1, 64)
	return p
}

func (p ModuleParams) WithFp32(key string, value float32) ModuleParams {
	p[key] = strconv.FormatFloat(float64(value), 'f', -1, 32)
	return p
}

func (p ModuleParams) string() string {
	if p == nil {
		return "nil"
	}
	if len(p) == 0 {
		return "{}"
	}
	bs, _ := json.Marshal(p)
	return string(bs)
}

type Module struct {
	Name   string
	Status ModuleStatus
	Params ModuleParams
}

type Manager struct {
	loadedModules []Module
	logger        ILogger
}

func New() *Manager {
	return &Manager{
		loadedModules: []Module{},
		logger:        nil,
	}
}

func (m *Manager) WithLogger(logger ILogger) *Manager {
	m.logger = logger
	return m
}

func (m *Manager) Modules() []Module {
	return m.loadedModules
}

func (m *Manager) RegisterModule(name string, status ModuleStatus, params ModuleParams) {
	if m.logger != nil {
		m.logger.Info("[ModuleMgr] Register module. name:", name, ", status:", status, ", params:", params.string())
	}
	m.loadedModules = append(m.loadedModules, Module{Name: name, Status: status, Params: params})
}

func (m *Manager) UpdateModuleStatus(name string, status ModuleStatus) {
	for i, module := range m.loadedModules {
		if module.Name == name {

			if m.logger != nil {
				m.logger.Info("[ModuleMgr] Update module status. name:", name,
					", status:", m.loadedModules[i].Status,
					"->", status)
			}
			m.loadedModules[i].Status = status
			return
		}
	}
}

func NewModuleParam() ModuleParams {
	return make(ModuleParams)
}
