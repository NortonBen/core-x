package main

import (
	"context"
	"core_x/contract"
	"fmt"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"sort"
	"syscall"
)

type Bootstrap struct {
	application *contract.AppContext
	contextApp  context.Context
	cancelFun   context.CancelFunc
	modules     []contract.IModule

	port int64
	lis  net.Listener
}

func NewBootstrap() Bootstrap {
	contextApp, cancelFun := context.WithCancel(context.Background())
	return Bootstrap{
		application: nil,
		contextApp:  contextApp,
		cancelFun:   cancelFun,
		modules:     make([]contract.IModule, 0),
	}
}

func (b *Bootstrap) createServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", b.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	log.Printf("Listen tcp %d", b.port)
	b.lis = lis
	return nil
}

func (b *Bootstrap) Modules(modules ...contract.IModule) {
	b.modules = append(b.modules, modules...)
}

func (b *Bootstrap) Flags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "root",
			Usage: "folder root core-x",
		},
		&cli.Int64Flag{
			Name:        "port",
			Usage:       "Port Listener Core",
			DefaultText: "5346",
			Value:       5346,
		},
	}
	for _, module := range b.modules {
		if module == nil {
			continue
		}
		flags = append(flags, module.Flag()...)
	}

	return flags
}

func (b *Bootstrap) Command() []*cli.Command {
	return []*cli.Command{}
}

func (b *Bootstrap) Application(application *contract.AppContext) {
	b.application = application
}

func (b *Bootstrap) Action(ctx *cli.Context) error {
	b.port = ctx.Int64("port")

	b.createServer()

	if b.application == nil {
		b.application = &contract.AppContext{
			GrpcServer: grpc.NewServer(),
		}
	}

	sort.Slice(b.modules, func(i, j int) bool {
		return b.modules[i].Priority() < b.modules[j].Priority()
	})

	for _, module := range b.modules {
		if module == nil {
			continue
		}
		module.Context(b.contextApp)
		module.Load()(b.application)
	}

	for _, module := range b.modules {
		if module == nil {
			continue
		}
		err := module.Init(*ctx)
		if err != nil {
			return err
		}
	}

	fmt.Print("Run Core-x \n")

	b.application.GrpcServer.Serve(b.lis)

	b.listenSignal()

	return nil
}

func (b *Bootstrap) listenSignal() {
	chanSignal := make(chan os.Signal)
	exit := make(chan bool, 1)

	signal.Notify(chanSignal, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-chanSignal
		log.Print("Stop core-x \n")
		b.cancelFun()
		b.lis.Close()
		exit <- true
	}()

	// stop core-x
	<-exit
}
