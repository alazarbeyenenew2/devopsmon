package initiator

import (
	"github.com/alazarbeyenenew2/devopsmon/internal/glue/scanner"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initRoute(grp *gin.RouterGroup, handler *handlers, log *zap.Logger) {
	scanner.Init(grp, log, handler.portScanner)
}
