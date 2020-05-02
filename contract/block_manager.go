package contract

import "core_x/models"

type IBlockManager interface {
	CheckAccessModule(moduleName string) (bool, error)
	Add(block models.Block) error
	Delete(blockKey string) error
	List() ([]models.Block, error)
	Get(blockKey string) (models.Block, error)
}
