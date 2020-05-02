package contract

import "google.golang.org/grpc"

type AppContext struct {
	ActionManager  IActionManager
	BlockManager   IBlockManager
	HookManager    IHookManager
	TokenManger    ITokenManager
	HandlerManager IHandleManager
	ModuleManager  IModuleManager
	GrpcServer    *grpc.Server
}

type AppLoad func(*AppContext)
