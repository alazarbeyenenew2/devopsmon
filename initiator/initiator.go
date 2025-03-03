package initiator

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	middleware "github.com/alazarbeyenenew2/devopsmon/internal/handler/middlware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to start logger")
	}

	//initializing config
	logger.Info("initializing config ")
	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}
	initConfig(configName, "config", logger)
	logger.Info("initializing config completed")

	//module
	logger.Info("initializing module")
	module := initModules(logger)
	logger.Info("done initializing module layer ")

	//initializing handler
	logger.Info("initializing handler")
	handler := initHandlers(logger, *module)
	logger.Info("done initializing handler")

	logger.Info("initializing http server")
	gsrv := gin.New()
	gsrv.Use(middleware.GinLogger(*logger))
	gsrv.Use(middleware.CORS())
	gsrv.Use(middleware.ErrorHandler())
	grp := gsrv.Group("api")

	initRoute(grp, handler, logger)

	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")),
		Handler:           gsrv,
		ReadHeaderTimeout: viper.GetDuration("app.timeout"),
		IdleTimeout:       30 * time.Minute,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint
		log.Fatal("HTTP server Shutdown")

	}()
	logger.Info(fmt.Sprintf("http server listening on port : %d", viper.GetInt("app.port")))

	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Could not start HTTP server: %s", err))
	}
}
