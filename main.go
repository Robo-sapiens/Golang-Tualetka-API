package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	delivery "main/internal/delivery"
	repository "main/internal/repository"
	"main/internal/usecase"
)


func InitRouter(api *delivery.Handlers) *fasthttprouter.Router {
	r := fasthttprouter.New()
	//r.POST("/api/forum/:slug", api.CreateForum)
	return r
}

func initDatabase() *pgx.ConnPool{
	db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			User:     "docker",
			Password: "docker",
			Port:     5432,
			Database: "docker",
		},
		MaxConnections: 50,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func initLevels()  *delivery.Handlers {
	db := initDatabase()
	useCase := usecase.NewUseCase(repository.NewDBStore(db))
	api := delivery.NewHandlers(useCase)
	return api
}

func main() {
	api := initLevels()
	router := InitRouter(api)

	log.Println("http server started on 5000 port: ")
	err := fasthttp.ListenAndServe(":5000", router.Handler)
	if err != nil {
		log.Println(err)
		return
	}
}
