package scanner

import (
	"net/http"

	"github.com/alazarbeyenenew2/devopsmon/internal/glue/routing"
	"github.com/alazarbeyenenew2/devopsmon/internal/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	grp *gin.RouterGroup,
	log *zap.Logger,
	scannerHandler handler.Scanner,
) {
	scannerRouts := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/scan/port",
			Handler:    scannerHandler.PortScanner,
			Middleware: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoute(grp, scannerRouts, *log)
}
