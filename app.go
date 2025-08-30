package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){

	godotenv.Load()

	fmt.Println("Start the Go Server with Chi Router")
	var appPort string=os.Getenv("PORT") 
	if appPort!=""{
		router:=chi.NewRouter()
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:  []string{"http://*","https://*"},
			AllowedMethods: []string{"GET","POST","OPTIONS"},
			AllowedHeaders:[]string{"*"},
			ExposedHeaders:[]string{"Link"},
			AllowCredentials:false,
			MaxAge:300,
		}))
		router.HandleFunc("/",HandleReadiness)
		router.HandleFunc("/user",handleUserReq)
		server:=&http.Server{
		Handler:router,
		Addr: ":"+appPort,
		}
		err:=server.ListenAndServe()
		if err!=nil{
			fmt.Printf("Error occured starting server ",err)
			log.Fatal("err ",err)
		}
	}else {
		fmt.Printf("Cant start server, Missing configuration")
	}



	}






