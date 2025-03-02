package scanner

import (
	"github.com/alazarbeyenenew2/devopsmon/internal/module"
	"go.uber.org/zap"
)

type scanner struct {
	log *zap.Logger
}

func Init(log *zap.Logger) module.Scanner {
	return &scanner{
		log: log,
	}
}
