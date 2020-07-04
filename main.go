package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	"main/internal/delivery"
	"main/internal/repository"
	"main/internal/usecase"
)


func InitRouter(api *delivery.Handlers) *fasthttprouter.Router {
	r := fasthttprouter.New()
	r.PUT("/api/registration", api.Register)
	r.DELETE("/api/registration", api.DeleteAccount)
	r.PUT("/api/rooms/create", api.CreateRoom)
	r.DELETE("/api/rooms/create", api.DeleteRoom)
	r.GET("/api/rooms/:id/about", api.GetInfoAboutRoom)
	r.GET("/api/rooms/:id/who", api.GetWhoBuy)
	r.PUT("/api/rooms/users", api.AddMemberIntoRoom)
	r.DELETE("/api/rooms/users", api.DeleteMemberFromRoom)
	r.POST("/api/rooms/users/papers/add", api.UpdatePaperToilet)
	r.POST("/api/rooms/users/valuable", api.ChangeValuable)
	r.POST("/api/rooms/users/pay_abilities", api.ChangePayAbility)
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

	log.Println("http server started on 8080 port: ")
	err := fasthttp.ListenAndServe(":8080", router.Handler)
	if err != nil {
		log.Println(err)
		return
	}
}
