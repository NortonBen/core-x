package handler

import (
	"context"
	"core_x/client"
	"core_x/contract"
	"core_x/contract/proto"
	"core_x/errors"
	"core_x/models"
	"core_x/utils"
	"io"
	"log"
	"time"
)

type CoreHandler struct {
	ctx context.Context
	actionManager  contract.IActionManager
	blockManager   contract.IBlockManager
	hookManager    contract.IHookManager
	tokenManger    contract.ITokenManager
	handlerManager contract.IHandleManager
	moduleManager  contract.IModuleManager
}

func (c CoreHandler) EventStream(req *proto.EventStreamRequest,stream proto.CoreService_EventStreamServer) error {

	token, err := c.checkValidateAndAuth(req.IdModule, req.IdSection, req.MetaData)

	if err != nil {
		return err
	}

	client := client.NewClientHandler(req.IdSection, token, req.IdModule, stream)


	err = c.handlerManager.AddConnect(client)

	if err != nil {
		log.Println(err.Error())
		return errors.New("server_error", errors.SERVER_ERROR)
	}


	for {
		select {
			case <- stream.Context().Done() : {
				return  nil
			}
			case <- c.ctx.Done() : {
				return errors.New("server_stop", errors.SERVER_ERROR)
			}
		}
	}

	c.handlerManager.RemoveConnect(req.IdSection)

	return nil
}

func (c CoreHandler) Hooks(stream proto.CoreService_HooksServer) error {

	var (
		IdModule string
		token string
		IdSection string
	)

	errorCheck := make(chan error, 2)

	// call time out
	go func() {
		checkTimeout := time.Tick(time.Duration(5) * time.Second)

		select {
			case  <-checkTimeout: {
				errorCheck <- errors.New("timeout_check", errors.TIMEOUT)
				return
			}
			case <- errorCheck: {
				return
			}
		}
	}()

	go func() {

		data, err := stream.Recv()
		if err == io.EOF {
			errorCheck <-  errors.New("timeout_check", errors.TIMEOUT)
			return
		}
		token, err = c.checkValidateAndAuth(data.IdModule, data.IdSection, data.MetaData)
		if err != nil {
			errorCheck <- err
			return
		}
		IdModule = data.IdModule
		IdSection = data.IdSection

		errorCheck <- nil
	}()

	err := <- errorCheck

	if err != nil {
		return err
	}

	// create client
	client := client.NewClientHook(IdSection, token, IdModule, stream)

	c.hookManager.AddConnect(client)

	// loop event
	err = client.Loop()
	if err != nil {
		log.Println(err.Error())
	}

	// client disconnect
	c.hookManager.RemoveConnect(IdSection)

	return nil
}

func (c CoreHandler) Ping(stream proto.CoreService_PingServer) error {

	errorCheck := make(chan error, 2)

	// call time out
	go func() {
		checkTimeout := time.Tick(time.Duration(5) * time.Second)

		select {
		case  <-checkTimeout: {
			errorCheck <- errors.New("timeout_check", errors.TIMEOUT)
			return
		}
		case <- errorCheck: {
			return
		}
		}
	}()

	go func() {

		data, err := stream.Recv()
		if err == io.EOF {
			errorCheck <-  errors.New("timeout_check", errors.TIMEOUT)
			return
		}
		err = c.checkValidateAndExist(data.IdModule, data.IdSection)

		if !c.tokenManger.Check(data.Token) {
			errorCheck <- errors.New("token_not_validate", errors.VALIDATE)
			return
		}
		err = c.moduleManager.Ping(data.IdModule, data.IdSection)
		if err != nil {
			errorCheck <- err
		}

		errorCheck <- nil
	}()

	err := <- errorCheck

	if err != nil {
		return err
	}


	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		err = c.moduleManager.Ping(data.IdModule, data.IdSection)
		if err != nil {
			return  err
		}
	}

	return nil
}

func (c *CoreHandler) Register(ctx context.Context,req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	module := models.ModuleData{}

	// check block
	err := utils.Mapper(req, &module)
	if err != nil {
		return nil, err
	}
	err = module.Validate()
	if err != nil {
		return nil, err
	}

	// check token
	if !c.tokenManger.Check(req.Token) {
		return nil, errors.New("token_not_validate", errors.VALIDATE)
	}

	// check block
	isAccess, err := c.blockManager.CheckAccessModule(req.Id)

	if err != nil {
		log.Println(err.Error())
		return  nil, errors.New("server_error", errors.SERVER_ERROR)
	}
	if !isAccess {
		return nil, errors.New("module_not_permission", errors.PERMISSION)
	}
	sectionId, err := c.moduleManager.Add(module)
	if err != nil {
		return nil, err
	}

	return &proto.RegisterResponse{
		IdSection:            sectionId,
	}, nil
}

func (c *CoreHandler) Action(ct context.Context,req *proto.ActionRequest) (*proto.Response, error) {

	_, err := c.checkValidateAndAuth(req.IdModule, req.IdSection, req.MetaData)
	if err != nil {
		return nil, err
	}
	err = c.hookManager.Process(req.Data)
	if err != nil {
		return nil, err
	}
	err = c.handlerManager.Process(req.Data)
	if err != nil {
		return nil, err
	}

	return &proto.Response {
		Success: true,
	}, nil
}

func (c CoreHandler) checkValidateAndAuth(IdModule string, IdSection string, metaData map[string]string) (string,error) {

	err := c.checkValidateAndExist(IdModule, IdSection)
	if err != nil {
		return "", err
	}

	// check token
	token, exit := metaData["token"]
	if !exit {
		return "", errors.New("token_not_found", errors.NOT_FOUND)
	}
	if !c.tokenManger.Check(token) {
		return "", errors.New("token_not_validate", errors.VALIDATE)
	}


	return token, nil
}

func (c CoreHandler) checkValidateAndExist(IdModule string, IdSection string) error {
	// validate data
	if len(IdModule) == 0 || len(IdSection) == 0 {
		return  errors.New("data_not_validate", errors.VALIDATE)
	}

	// check block
	isAccess, err := c.blockManager.CheckAccessModule(IdModule)

	if err != nil {
		log.Println(err.Error())
		return  errors.New("server_error", errors.SERVER_ERROR)
	}

	if !isAccess {
		return errors.New("module_not_permission", errors.PERMISSION)
	}

	// check module and section
	if !c.moduleManager.Check(IdModule, IdSection) {
		return  errors.New("module_or_section_not_found", errors.NOT_FOUND)
	}

	return nil
}
