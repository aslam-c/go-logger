package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)


func main(){

	godotenv.Load()

	fmt.Println("Start the Go Server with Chi Router")
	var appPort string=os.Getenv("PORT") 
	if appPort!=""{
		router:=chi.NewRouter()
	    server:=&http.Server{
		Handler:router,
		Addr: ":"+appPort,
		}
		err:=server.ListenAndServe()
		if err!=nil{
			fmt.Printf("Error occured starting server ",err)
		}
	}else {
		fmt.Printf("Cant start server, Missing configuration")
	}




	}






