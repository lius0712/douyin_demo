package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"io/ioutil"

	"github.com/RaymondCode/simple-demo/config"
	"github.com/go-redis/redis/v8"
)

var DB *gorm.DB
var RDB *redis.Client

func ConnectDB() error {

	dsn := config.Config.MySQLDSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	return err
}

//连接Redis
func ConnectRDB() error {

	fmt.Println(config.Config.RdbUrl)

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.RdbUrl,
		Password: config.Config.RdbPwd,
		DB:       config.Config.RdbNum,
		PoolSize: config.Config.RdbPoolSize, //连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	RDB = rdb
	ctx = context.Background()
	rdb.Set(ctx, "zwk", 123456, 0)
	return err
}

func ResetDB() error {
	files, err := ioutil.ReadDir("./repository")
	if err != nil {
		return err
	}

	for _, file := range files {
		n := file.Name()
		if len(n) > 6 && n[:2] == "t_" && n[len(n)-4:] == ".txt" {
			err := func() (err error) {
				f, err := os.OpenFile("./repository/"+n, os.O_RDONLY, 0)
				defer f.Close()
				if err != nil {
					return
				}

				b := make([]byte, 65535)
				n, err := f.Read(b)
				if err != nil {
					return
				}

				for _, s := range strings.Split(string(b[:n]), ";") {
					s = strings.TrimSpace(s)
					if len(s) == 0 {
						continue
					}
					err = DB.Exec(s).Error
					if err != nil {
						return
					}
				}

				return
			}()

			if err != nil {
				return err
			}
		}
	}

	return nil
}
