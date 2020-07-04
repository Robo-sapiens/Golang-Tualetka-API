package delivery

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"main/internal/models"
	"net/http"
	"strconv"
)

func (handlers *Handlers) CreateRoom(ctx *fasthttp.RequestCtx) {
	room := models.Room{}
	room.UnmarshalJSON(ctx.PostBody())
	newRoom, err := handlers.usecase.CreateRoom(&room)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(newRoom)
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}

func (handlers *Handlers) DeleteRoom(ctx *fasthttp.RequestCtx) {
	room := models.Room{}
	room.UnmarshalJSON(ctx.PostBody())

	err := handlers.usecase.DeleteRoom(room.ID)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}

func (handlers *Handlers) GetInfoAboutRoom(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue("id").(string)
	roomID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
	room, err := handlers.usecase.GetInfoAboutRoom(roomID)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(room)
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}



