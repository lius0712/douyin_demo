package main

import (
	"log"
	"os"

	"github.com/RaymondCode/simple-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}

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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Init() error {
	if err := repository.ConnectDB(); err != nil {
		return err
	}
	return nil
}
