package manager

import (
	"core_x/contract"
	"core_x/errors"
	"core_x/models"
	"core_x/utils"
	"context"
	"github.com/urfave/cli/v2"
	"path"
)

type TokenManagerFile struct {
	list map[string]models.Token
	ctx context.Context
	fileName string
	root string
}


func NewTokenManagerFile() contract.IModule {
	return &TokenManagerFile{
		list: make(map[string]models.Token),
	}
}



func (t *TokenManagerFile) Init(c cli.Context) error {
	t.fileName = c.String("file_token_manager")
	t.root = c.String("root")

	t.LoadData()

	return  nil
}

func (t *TokenManagerFile) Load() contract.AppLoad {
	return func(appContext *contract.AppContext) {
		appContext.TokenManger = t
	}
}

func (t TokenManagerFile) Flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "file_token_manager",
			DefaultText: "token.json",
		},
	}
}

func (t TokenManagerFile) Priority() int {
	return  1
}

func (t *TokenManagerFile) Context(ctx context.Context) {
	t.ctx = ctx
}


func (t TokenManagerFile) Save() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.SaveFile(fileNameFullPath, t.list)
}

func (t *TokenManagerFile) LoadData() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.LoadFile(fileNameFullPath, &t.list)
}

func (t TokenManagerFile) Check(tokenKey string) bool {
	_, err := t.Get(tokenKey)
	if err != nil {
		return false
	}
	return true
}

func (t TokenManagerFile) CheckAccess(tokenKey string, moduleName string) bool {
	token, err := t.Get(tokenKey)
	if err != nil {
		return false
	}

	return  utils.MatchStringArray(token.ModuleRule, moduleName)
}

func (t TokenManagerFile) CheckPermissionAction(tokenKey string, action string, isHook bool) bool {
	token, err := t.Get(tokenKey)
	if err != nil {
		return false
	}
	ruleList := token.ActionRule
	if isHook {
		ruleList = token.HookRule
	}
	return  utils.MatchStringArray(ruleList, action)
}

func (t *TokenManagerFile) Add(token models.Token) error {
	_, exist :=  t.list[token.Key]
	if exist {
		return errors.New("Token exist", errors.EXIST)
	}
	t.list[token.Key] = token
	return t.Save()
}

func (t *TokenManagerFile) Update(token models.Token) error {
	_, exist :=  t.list[token.Key]
	if !exist {
		return errors.New("Not found token", errors.NOT_FOUND)
	}
	t.list[token.Key] = token
	return t.Save()
}

func (t *TokenManagerFile) Delete(tokenKey string) error {
	_, exist :=  t.list[tokenKey]
	if !exist {
		return  errors.New("Not found token", errors.NOT_FOUND)
	}
	delete(t.list, tokenKey)
	return t.Save()
}

func (t TokenManagerFile) List() ([]models.Token, error) {
	list := make([]models.Token, 1)
	for _, token := range t.list {
		list = append(list, token)
	}
	return list, nil
}

func (t TokenManagerFile) Get(tokenKey string) (models.Token, error) {
	item, exist :=  t.list[tokenKey]
	if !exist {
		return models.Token{}, errors.New("Not found token", errors.NOT_FOUND)
	}
	return item, nil
}


