package service

import (
	"github.com/SmartDuck9000/go-tradex/src/config"
	"github.com/SmartDuck9000/go-tradex/src/db"
	"github.com/SmartDuck9000/go-tradex/src/validator"
	"github.com/gin-gonic/gin"
)

type StatServiceInterface interface {
	Run() error
}

type StatService struct {
	server    *gin.Engine
	db        db.RepoDB
	validator validator.ValidationInterface
	host      string
	port      string
}

func CreateServer(conf config.ServiceConfig) *StatService {
	gin.SetMode(gin.ReleaseMode)

	var api = StatService{
		server:    gin.Default(),
		db:        db.CreateRepoPostgres(conf.DB),
		validator: validator.NewPlaygroundValidator(),
		host:      conf.Host,
		port:      conf.Port,
	}

	api.server.POST("/api/stat/save", api.saveStat)
	api.server.GET("/api/stat/get", api.getStat)
	api.server.DELETE("/api/stat/delete", api.deleteStat)

	return &api
}

func (api StatService) Run() error {
	err := api.db.Open()
	if err != nil {
		return err
	}

	return api.server.Run(api.host + ":" + api.port)
}
