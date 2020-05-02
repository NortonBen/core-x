package contract

import "core_x/models"

type ITokenManager interface {
	Check(tokenKey string) bool
	CheckAccess(tokenKey string, moduleName string) bool
	CheckPermissionAction(tokenKey string, action string,isHook  bool) bool
	Add(token models.Token) error
	Update(token models.Token) error
	Delete(tokenKey string) error
	List() ([]models.Token, error)
	Get(tokenKey string) (models.Token, error)
}
