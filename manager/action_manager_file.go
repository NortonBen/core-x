package manager

import (
	"context"
	"core_x/contract"
	"core_x/errors"
	"core_x/models"
	"core_x/utils"
	"github.com/urfave/cli/v2"
	"log"
	"path"
)

type ActionMangerFile struct {
	list     map[string]models.ActionData
	ctx      context.Context
	fileName string
	root     string
}

func NewActionMangerFile() contract.IModule {
	return &ActionMangerFile{
		list: make(map[string]models.ActionData),
	}
}

func (t *ActionMangerFile) Init(c cli.Context) error {
	t.fileName = c.String("file_action_manager")
	t.root = c.String("root")

	t.LoadData()

	log.Println("Run action manager...")
	return nil
}

func (t *ActionMangerFile) Load() contract.AppLoad {
	return func(appContext *contract.AppContext) {
		appContext.ActionManager = t
	}
}

func (t ActionMangerFile) Flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "file_action_manager",
			DefaultText: "actions.json",
		},
	}
}

func (t ActionMangerFile) Priority() int {
	return 1
}

func (t *ActionMangerFile) Context(ctx context.Context) {
	t.ctx = ctx
}

func (t ActionMangerFile) Get(action string) (models.ActionData, error) {
	item, exit := t.list[action]
	if exit {
		return models.ActionData{}, errors.Error{
			Message: "Action Have Exist",
			Code:    errors.EXIST,
		}
	}

	return item, nil
}

func (t ActionMangerFile) Save() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.SaveFile(fileNameFullPath, t.list)
}

func (t *ActionMangerFile) LoadData() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.LoadFile(fileNameFullPath, &t.list)
}

func (a *ActionMangerFile) HasAction(action string) bool {
	_, exit := a.list[action]
	return exit
}

func (a *ActionMangerFile) Add(module models.ActionData) error {
	_, err := a.Get(module.Name)
	if err != nil {
		return err
	}
	a.list[module.Name] = module

	return a.Save()
}

func (a *ActionMangerFile) AddHandler(action string, module string) error {
	item, err := a.Get(action)
	if err != nil {
		return err
	}
	err = a.checkPermission(item, module)
	if err != nil {
		return err
	}
	item.HandlerRegister = append(item.HandlerRegister, module)
	return a.Save()
}

func (a *ActionMangerFile) AddHook(action string, module string) error {
	item, err := a.Get(action)
	if err != nil {
		return err
	}
	err = a.checkPermission(item, module)
	if err != nil {
		return err
	}
	if !item.EnableHook {
		return errors.Error{
			Message: "Not Enable Register Hook",
			Code:    errors.PERMISSION,
		}
	}
	item.HookRegister = append(item.HookRegister, module)
	return a.Save()
}

func (a *ActionMangerFile) GetHook(action string) ([]string, error) {
	item, err := a.Get(action)
	if err == nil {
		return nil, err
	}
	return item.HookRegister, nil
}

func (a *ActionMangerFile) GetHandler(action string) ([]string, error) {
	item, err := a.Get(action)
	if err == nil {
		return nil, err
	}
	return item.HandlerRegister, nil
}

func (a *ActionMangerFile) RemoveHook(action string, module string) error {
	item, err := a.Get(action)
	if err != nil {
		return err
	}

	var index = -1
	for i, value := range item.HookRegister {
		if value == module {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.Error{
			Message: "Module Not Register Hook",
			Code:    errors.EXIST,
		}
	}
	item.HookRegister = utils.RemoveItemString(item.HookRegister, index)
	return a.Save()
}

func (a *ActionMangerFile) RemoveHandler(action string, module string) error {
	item, err := a.Get(action)
	if err != nil {
		return err
	}

	var index = -1
	for i, value := range item.HandlerRegister {
		if value == module {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.Error{
			Message: "Module Not Register Handle",
			Code:    errors.EXIST,
		}
	}
	item.HandlerRegister = utils.RemoveItemString(item.HandlerRegister, index)
	return a.Save()
}

func (a *ActionMangerFile) HasHandler(action string, module string) (bool, error) {
	item, err := a.Get(action)
	if err != nil {
		return false, err
	}

	return utils.ItemExists(item.HandlerRegister, module), nil
}

func (a *ActionMangerFile) HasHook(action string, module string) (bool, error) {
	item, err := a.Get(action)
	if err != nil {
		return false, err
	}

	return utils.ItemExists(item.HookRegister, module), nil
}

func (a *ActionMangerFile) checkPermission(action models.ActionData, module string) error {
	if !utils.MatchStringArray(action.EndPointRule, module) {
		return errors.Error{
			Message: "Not Permission Register",
			Code:    errors.PERMISSION,
		}
	}
	return nil
}
