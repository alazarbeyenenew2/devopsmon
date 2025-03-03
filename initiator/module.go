package initiator

import (
	"github.com/alazarbeyenenew2/devopsmon/internal/module"
	"github.com/alazarbeyenenew2/devopsmon/internal/module/scanner"
	"go.uber.org/zap"
)

type modules struct {
	Scanner module.Scanner
}

func initModules(log *zap.Logger) *modules {
	return &modules{
		Scanner: scanner.Init(log),
	}
}
