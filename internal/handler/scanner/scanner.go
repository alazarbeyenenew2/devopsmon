package scanner

import (
	"github.com/alazarbeyenenew2/devopsmon/internal/handler"
	"github.com/alazarbeyenenew2/devopsmon/internal/module"
	"go.uber.org/zap"
)

type scanner struct {
	log           *zap.Logger
	scannerModule module.Scanner
}

func Init(log *zap.Logger, scannerModule module.Scanner) handler.Scanner {
	return &scanner{
		log:           log,
		scannerModule: scannerModule,
	}
}
