package delivery

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)


func (handlers *Handlers) Register(ctx *fasthttp.RequestCtx) {
	user := models.User{}
	user.UnmarshalJSON(ctx.PostBody())

	newUser, err := handlers.usecase.Register(&user)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusCreated)
		ctx.SetContentType("application/json")
		response, _ := newUser.MarshalJSON()
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _ := json.Marshal(err.Error())
		ctx.Write(response)
	}
	return
}

func (handlers *Handlers) DeleteAccount(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue("id").(string)
	userID, err := strconv.Atoi(ID)
	if err != nil {
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
	err = handlers.usecase.DeleteAccount(userID)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
	case errors.UserNotFound:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}



