package main

import (
	"adv-demo/configs"
	"adv-demo/internal/hello"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/hello", hello)
	//fmt.Println("Server is starting... on port 8081")
	//err := http.ListenAndServe(":8081", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is starting... on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
