package contract

import (
	"context"
	"github.com/urfave/cli/v2"
)

type IModule interface {
	Init(c cli.Context) error
	Load() AppLoad
	Flag() []cli.Flag
	Priority() int
	Context(ctx context.Context)
}
