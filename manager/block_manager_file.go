package manager

import (
	"core_x/contract"
	"core_x/errors"
	"core_x/models"
	"core_x/utils"
	"context"
	"github.com/urfave/cli/v2"
	"path"
	"regexp"
	"time"
)

type BlockMangerFile struct {
	list map[string]models.Block
	ctx context.Context
	fileName string
	root string
}

func NewBlockMangerFile() contract.IModule {
	return &BlockMangerFile{
		list: make(map[string]models.Block),
	}
}


func (t *BlockMangerFile) Init(c cli.Context) error {

	t.fileName = c.String("file_block_manager")
	t.root = c.String("root")

	t.LoadData()

	return nil
}

func (t *BlockMangerFile) Load() contract.AppLoad {
	return func(appContext *contract.AppContext) {
		appContext.BlockManager = t
	}
}

func (t *BlockMangerFile) Flag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name: "file_block_manager",
			DefaultText: "block.json",
		},
	}
}

func (t BlockMangerFile) Priority() int {
	return 1
}

func (t *BlockMangerFile) Context(ctx context.Context) {
	t.ctx = ctx
}

func (t BlockMangerFile) CheckAccessModule(moduleName string) (bool, error) {

	for _, block := range t.list {
		match, _ := regexp.MatchString(block.Key, moduleName)
		if match {
			return true, nil
		}
	}

	return false, nil
}

func (t BlockMangerFile) Add(block models.Block) error {
	_, exist := t.list[block.Key]
	if !exist {
		return errors.New("Block exist", errors.EXIST)
	}
	block.Created_at = time.Now()
	t.list[block.Key] = block
	return t.Save()
}

func (t *BlockMangerFile) Delete(blockKey string) error {
	_, exist :=  t.list[blockKey]
	if !exist {
		return errors.New("Not found token", errors.NOT_FOUND)
	}
	delete(t.list, blockKey)
	return t.Save()
}

func (t BlockMangerFile) List() ([]models.Block, error) {
	list := make([]models.Block, 10)
	for _, token := range t.list {
		list = append(list, token)
	}
	return list, nil
}

func (t BlockMangerFile) Get(blockKey string) (models.Block, error) {
	block, exist :=  t.list[blockKey]
	if !exist {
		return models.Block{}, errors.New("Not found token", errors.NOT_FOUND)
	}
	return block, nil
}

func (t BlockMangerFile) Save() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.SaveFile(fileNameFullPath, t.list)
}

func (t *BlockMangerFile) LoadData() error {
	fileNameFullPath := path.Join(t.root, t.fileName)
	return utils.LoadFile(fileNameFullPath, &t.list)
}
