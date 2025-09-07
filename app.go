package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aslam-c/go-logger/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main(){

	godotenv.Load()

	fmt.Println("Start the Go Server with Chi Router")
	var appPort string=os.Getenv("PORT") 
	var databaseUrl string=os.Getenv("DATABASE_URL")

	var databaseConfig struct{
		DB *database.Queries
	}




	if appPort!=""{

		postgresConn,connectErr:=sql.Open("postgres",databaseUrl)
		if connectErr!=nil {
			log.Fatal("Cant connect to database: ",connectErr)
		}
		queries,dbError:=database.new(postgresConn)
		if dbError!=nil{
			log.Fatal("Cant create a db conn ",dbError)
		}	

		dbConfig=databaseConfig{DB:queries}

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

		v1Router:=chi.NewRouter()
		v1Router.HandleFunc("/user",handleUserReq)
		
		router.Mount("/v1",v1Router)

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





