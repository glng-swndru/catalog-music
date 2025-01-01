package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/catalog-music/internal/configs"
	membershipsHandler "github.com/glng-swndru/catalog-music/internal/handler/memberships"
	"github.com/glng-swndru/catalog-music/internal/models/memberships"
	membershipsRepo "github.com/glng-swndru/catalog-music/internal/repository/memberships"
	membershipSvc "github.com/glng-swndru/catalog-music/internal/service/memberships"
	"github.com/glng-swndru/catalog-music/pkg/internalsql"
)

func main() {

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs",
			"./internal/configs",
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to init configs: %v\n", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v\n", err)
	}
	db.AutoMigrate(&memberships.User{})
	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)

	membershipSvc := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
