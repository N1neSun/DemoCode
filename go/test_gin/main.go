package main

import (
	"fmt"
	"net/http"
	"time"

	"test_gin/app/router"
	"test_gin/common/setting"
	"test_gin/common/validator"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	binding.Validator = new(validator.DefaultValidator)
	router := router.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
