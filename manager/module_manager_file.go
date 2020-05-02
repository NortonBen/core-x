package manager

import (
	"core_x/contract"
	"core_x/errors"
	"core_x/models"
	"core_x/utils"
	"context"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"log"
	"path"
	"time"
)

type ModuleManagerFile struct {
	actionManager contract.IActionManager
	list map[string]models.ModuleData
	ctx context.Context
	fileName string
	root string
	timeOut time.Duration
}


func NewModuleManagerFile() contract.IModule {
	return &ModuleManagerFile{
		list: make(map[string]models.ModuleData),
		timeOut: time.Duration(1) * time.Minute,
	}
}



func (t *ModuleManagerFile) Init(c cli.Context) error {
	t.fileName = c.String("file_module_manager")
	t.root = c.String("root")

	t.LoadData()

	return  nil
}

func (t *ModuleManagerFile) Load() contract.AppLoad {
	return func(appContext *contract.AppContext) {
		t.actionManager = appContext.ActionManager
		appContext.ModuleManager = t
	}
}

func (t ModuleManagerFile) Flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "file_module_manager",
			DefaultText: "modules.json",
		},
	}
}

func (t ModuleManagerFile) Priority() int {
	return  2
}

func (t *ModuleManagerFile) Context(ctx context.Context) {
	t.ctx = ctx
}



func (t ModuleManagerFile) Add(module models.ModuleData) (string, error) {
	if t.Has(module.Id) {
		return "",  errors.New("Module has exist", errors.EXIST)
	}
	idSection, err := t.CreateNewIdSection()
	if err != nil {
		return "", nil
	}
	module.Section = idSection
	t.list[module.Id] = module
	return idSection, t.Save()
}

func (t *ModuleManagerFile) Update(module models.ModuleData) error {
	_, err := t.Get(module.Id)
	if err != nil {
		return err
	}
	t.list[module.Id] = module

	if module.State == models.StateModule_STOP {
		t.DeregisterProcess(module.Id)
	}

	return t.Save()
}

func (t ModuleManagerFile) Get(moduleName string) (models.ModuleData,error) {
	module, exist := t.list[moduleName]
	if !exist {
		return models.ModuleData{},  errors.New("Not found module", errors.NOT_FOUND)
	}
	return module, nil
}

func (t *ModuleManagerFile) Remove(moduleName string) error {
	_, err := t.Get(moduleName)
	if err != nil {
		return err
	}
	delete(t.list, moduleName)
	return t.Save()
}

func (t ModuleManagerFile) Has(moduleName string) bool {
	_, exist := t.list[moduleName]
	return exist
}

func (t ModuleManagerFile) GetListTimeOut() []string {
	list := []string{}
	timeCheck := time.Now().Add(0 - t.timeOut)
	for _, module := range t.list {
		if module.Time.Before(timeCheck) {
			list = append(list, module.Id )
		}
	}

	return list
}

func (t ModuleManagerFile) CreateNewIdSection() (string, error) {
	section := uuid.New().String()

	for {
		exist := false
		for _, module := range t.list {
			if module.Section == section {
				exist = true
				break
			}
		}
		if !exist {
			break
		}
	}


	return section, nil
}

func (t ModuleManagerFile) SetState(moduleName string, state models.StateModule)  error {
	module, err := t.Get(moduleName)
	if err !=nil {
		return err
	}
	module.State = state
	return t.Update(module)
}

func (t ModuleManagerFile) Ping(moduleName string, idSection string) error {
	module, err := t.Get(moduleName)
	if err !=nil {
		return err
	}
	if module.Section != idSection {
		return errors.New("Section not match", errors.PERMISSION)
	}
	now := time.Now()
	module.Time = &now
	return t.Update(module)
}

func (t ModuleManagerFile) Check(moduleName string, idSection string) bool {
	module, err := t.Get(moduleName)
	if err !=nil {
		return false
	}
	if module.Section != idSection {
		return false
	}
	return true
}

func (t ModuleManagerFile) GetModuleWithToken(token string) []models.ModuleData {
	list := make([]models.ModuleData, 2)
	for _, module := range t.list {
		if module.Token == token {
			list = append(list, module)
		}
	}
	return list
}

func (t ModuleManagerFile) RegisterProcess(moduleName string) error {

	module, err := t.Get(moduleName)
	if err != nil {
		return err
	}

	for _, action := range module.Actions {
		err = t.actionManager.Add(*action)
		if err != nil {
			log.Print("Have error register action: "+err.Error())
		}
	}

	for _, register := range module.Registers {
		if register.Type == models.TypeAction_EndPoint {
			err = t.actionManager.AddHandler(register.Action, moduleName)
			if err != nil {
				log.Print("Have error register handler: "+err.Error())
			}
		} else {
			err = t.actionManager.AddHook(register.Action, moduleName)
			if err != nil {
				log.Print("Have error register hook: "+err.Error())
			}
		}

	}



	return nil
}

func (t ModuleManagerFile) DeregisterProcess(moduleName string) error {

	module, err := t.Get(moduleName)
	if err != nil {
		return err
	}


	for _, register := range module.Registers {
		if register.Type == models.TypeAction_EndPoint {
			err = t.actionManager.RemoveHandler(register.Action, moduleName)
			if err != nil {
				log.Print("Have error deregister handler: "+err.Error())
			}
		} else {
			err = t.actionManager.RemoveHook(register.Action, moduleName)
			if err != nil {
				log.Print("Have error deregister hook: "+err.Error())
			}
		}

	}

	return nil
}

func (t ModuleManagerFile) Save() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.SaveFile(fileNameFullPath, t.list)
}

func (t *ModuleManagerFile) LoadData() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.LoadFile(fileNameFullPath, &t.list)
}