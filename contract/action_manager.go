package contract

import (
	"core_x/models"
)

type IActionManager interface {
	Add(module models.ActionData) error
	AddHandler(action string, module string) error
	AddHook(action string, module string) error
	GetHook(action string) ([]string, error)
	GetHandler(action string) ([]string, error)
	RemoveHook(action string, module string) error
	RemoveHandler(action string, module string) error
	HasHandler(action string, module string) (bool, error)
	HasHook(action string, module string) (bool, error)
	HasAction(action string) bool
	Get(action string) (models.ActionData, error)
}
