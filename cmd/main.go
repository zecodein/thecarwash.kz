package main

import (
	"flag"
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
	flag.StringVar(&configPath, "config-path", "configs/docker.toml", "path to config file")
}

func main() {
	// request := fmt.Sprintf(`https://data.egov.kz/api/v4/gbd_ul/v1?apiKey=1a241b0789c74223842f63c5013262a3&source={size:100,query:{bool:{must:[{match:{bin:160340014245}}]}}}`)
	// fmt.Println(request)
	// resp, err := http.Get(request)
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

	db, err := postgres.NewPostgresRepository(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	store, err := redis.NewStore(10, "tcp", config.CacheHost+config.CacheAddr, config.CachePassword, []byte(web.Key))
	if err != nil {
		log.Fatal(err)
	}

	router.Use(gin.LoggerWithWriter(logger))
	router.Use(sessions.Sessions(web.Key, store))

	userRepo := repository.NewUserRepository(db)
	washingRepo := repository.NewWashingRepostiroty(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	washingUsecase := usecase.NewWashingUsecase(washingRepo)

	web.NewHandler(router, &web.Handler{
		UserUsecase:    userUsecase,
		WashingUsecase: washingUsecase,
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
