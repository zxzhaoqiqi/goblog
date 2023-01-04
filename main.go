package main

import (
	"net/http"

	"github.com/zxzhaoqiqi/goblog/app/http/middlewares"
	"github.com/zxzhaoqiqi/goblog/bootstrap"
	"github.com/zxzhaoqiqi/goblog/pkg/logger"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.Error(err)
}
