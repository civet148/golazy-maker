package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/svc"
	"os"

	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
)

const (
	Version     = "0.1.0"
	ProgramName = "main"
)

var (
	BuildTime = "2025-08-14 15:38:50"
	GitCommit = "<N/A>"
)

const (
	CmdFlag_Config = "config"
)

// @title main swagger APIs
// @version 1.0
// @description
// @host localhost:8888
// @BasePath
func main() {

	app := &cli.App{
		Name:    ProgramName,
		Usage:   fmt.Sprintf("%s [options]", ProgramName),
		Version: fmt.Sprintf("v%s %s commit %s", Version, BuildTime, GitCommit),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    CmdFlag_Config,
				Aliases: []string{"c"},
				Usage:   "config file name",
				Value:   "config.yaml",
			},
		},
		Action: func(ctx *cli.Context) error {
			var err error
			var cfg config.Config
			var configFile = ctx.String(CmdFlag_Config)
			viper.SetConfigFile(configFile)
			err = viper.ReadInConfig()
			if err != nil {
				panic(err.Error())
			}
			err = viper.Unmarshal(&cfg)
			if err != nil {
				panic(err.Error())
			}
			log.Json("config", cfg)

			server := gin.Default()
			svcCtx := svc.NewServiceContext(cfg)
			handler.RegisterHandlers(server, svcCtx)

			strListenAddr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
			log.Infof("Starting server on %s mode at %s...\n", cfg.Mode, strListenAddr)
			err = server.Run(strListenAddr)
			if err != nil {
				return log.Errorf(err.Error())
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Errorf("exit in error %s", err)
		os.Exit(1)
		return
	}
}
