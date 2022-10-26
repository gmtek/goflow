package runtime

import (
	"github.com/gmtek/goflow/core/sdk/executor"
)

type Runtime interface {
	Init() error
	CreateExecutor(*Request) (executor.Executor, error)
}
