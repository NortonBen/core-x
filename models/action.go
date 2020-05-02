package models

import "time"


type ActionData struct {
	Durable bool
	EnableHook bool
	EndPointRule []string
	TimeOut time.Duration
	Name string
	HookRegister []string
	HandlerRegister []string
}
