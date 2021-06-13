package main


import (
	"puppy-hids/web/common/log"
	"puppy-hids/web/routers"
	"puppy-hids/web/setting"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("11111")
	log.SetLogLevel(log.DEBUG)
	router := routers.InitRouter()

	fmt.Println("2222")
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("333")
	s.ListenAndServe()
}
