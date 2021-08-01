package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/smallpaes/go-blog-backend/global"
	"github.com/smallpaes/go-blog-backend/internal/model"
	"github.com/smallpaes/go-blog-backend/internal/routers"
	"github.com/smallpaes/go-blog-backend/pkg/setting"
)

func init() {
	// init config: map setting to app struct
	err := setupSetting()
	if err != nil {
		log.Fatalf("Init setup err: %v", err)
	}

	// init DBEngine
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("Init setupDBEngine err: %v", err)
	}
}

func main() {
	// set gin mode for execution
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()

	// create custom server
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println(global.ServerSetting)
	fmt.Println(global.AppSetting)
	fmt.Println(global.DatabaseSetting)
	// start the HTTP server
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	// unmarshal config to struct by specifying key
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return nil
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
