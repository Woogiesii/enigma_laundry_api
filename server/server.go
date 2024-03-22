package server

import (
	"enigma_laundry_api/config"
	"enigma_laundry_api/controller"
	"enigma_laundry_api/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	apiCfg    config.ApiConfig
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewCustomersController(s.ucManager.UsersUseCase(), rg, s.apiCfg).Route()
	controller.NewServicesController(s.ucManager.ServicesUseCase(), rg, s.apiCfg).Route()
	controller.NewTransactionController(s.ucManager.TransactionUseCase(), s.ucManager.UsersUseCase(), rg, s.apiCfg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, _ := manager.NewInfraManager(cfg)

	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	apiHost := fmt.Sprintf(":%s", cfg.ApiConfig.ApiPort)

	return &Server{
		ucManager: ucManager,
		engine:    engine,
		host:      apiHost,
		apiCfg:    cfg.ApiConfig,
	}
}
