package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/zecodein/thecarwash.kz/configs"
	"github.com/zecodein/thecarwash.kz/delivery/web"
	"github.com/zecodein/thecarwash.kz/repository"
	"github.com/zecodein/thecarwash.kz/repository/postgres"
	"github.com/zecodein/thecarwash.kz/usecase"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	logger, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}

	config := configs.NewConfig()

	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	db, err := postgres.NewPostgresRepository(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store, err := redis.NewStore(10, "tcp", config.CacheHost+config.CacheAddr, config.CachePassword, []byte("thecarwash"))
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(gin.LoggerWithWriter(logger))
	router.Use(sessions.Sessions("thecarwash", store))

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	web.NewHandler(router, &web.Handler{
		UserUsecase: userUsecase,
	})

	server := &http.Server{
		Addr:         config.BindAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println(server.ListenAndServe())
}
