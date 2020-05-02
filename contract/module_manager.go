package contract

import (
	"core_x/models"
)

type IModuleManager interface {
	Add(module models.ModuleData) (string, error)
	Update(module models.ModuleData) error
	Get(moduleName string) (models.ModuleData,error)
	Remove(moduleName string) error
	Has(moduleName string) bool
	GetListTimeOut() []string
	CreateNewIdSection() (string, error)
	SetState(moduleName string, state models.StateModule) error
	Ping(moduleName string, idSection string) error
	Check(moduleName string, idSection string) bool
	GetModuleWithToken(token string) []models.ModuleData
	RegisterProcess(moduleName string) error
	DeregisterProcess(moduleName string) error
}
