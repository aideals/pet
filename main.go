package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"./setting"
	"./routers"
)

func main() {
	// db, err := gorm.Open("mysql", "root:123456@/pets?charset=utf8&&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("DB connect Failed")
	// }

	// defer db.Close()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
