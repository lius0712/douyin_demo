package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/middleware"
	"log"
	"os"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}

	//for example
	//cache.Set("name", "wwww", 0)

	if len(os.Args) > 1 && os.Args[1] == "reset" {
		err := repository.ResetDB()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("reset db success")
		os.Exit(0)
	}

	r := gin.Default()
	initRouter(r)
	// run on port config
	r.Run(fmt.Sprintf(":%d", config.Config.DbPort))
}

func Init() error {
	// connect Mysql
	if err := repository.ConnectDB(); err != nil {
		return err
	}
	// init logger
	if err := middleware.InitLogger(config.Config.LogConfig, config.Config.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}

	// connect Redis
	//if err := repository.ConnectRDB(); err != nil {
	//	return err
	//}
	return nil
}
