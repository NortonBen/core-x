package models

import "time"

type Token struct {
	Key string
	ModuleRule []string
	ActionRule []string
	HookRule []string
	Created_at time.Time
	Updated_at time.Time
}
