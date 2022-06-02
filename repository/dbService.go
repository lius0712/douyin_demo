package repository

import (
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"io/ioutil"

	"github.com/RaymondCode/simple-demo/config"
)

var DB *gorm.DB

func ConnectDB() error {

	dsn := config.Config.MySQLDSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

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
