package initiator

import (
	"github.com/alazarbeyenenew2/devopsmon/internal/handler"
	"github.com/alazarbeyenenew2/devopsmon/internal/handler/scanner"
	"go.uber.org/zap"
)

type handlers struct {
	portScanner handler.Scanner
}

func initHandlers(log *zap.Logger, module modules) *handlers {
	return &handlers{
		portScanner: scanner.Init(log, module.Scanner),
	}
}
