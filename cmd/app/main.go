package main

import (
	_ "article-service/docs"
	"article-service/internal/app"
	"article-service/internal/config"
	"context"
	"time"
)

const (
	ConfigDevPath = "././config.yaml"
	ConfigPath    = "./config.yaml"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @description article-service API
// @description Это сваггер-документация для сервиса статей: управление статьями, тегами.
// @description Все тела запросов, необходимые токены и возможные ошибки указаны в описании методов.

func main() {
	time.Sleep(10 * time.Second)
	cfg, err := config.LoadConfig(ConfigDevPath)
	if err != nil {
		panic(err)
	}

	application := app.NewApplication(cfg)

	ctx := context.Background()
	err = application.Init(ctx)
	if err != nil {
		panic(err)
	}

	application.Run()
}
