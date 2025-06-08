package main

import (
	"os"

	"github.com/lalo64/payment_domain/internal/server"
)

var (
	HOST = os.Getenv("HOST_SERVER")
	PORT = os.Getenv("PORT_SERVER")

)

func main(){
	srv := server.NewServer(HOST, PORT)
	srv.Run()
}