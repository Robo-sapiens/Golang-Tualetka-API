package delivery

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"main/internal/models"
	"net/http"
	"main/internal/tools/messages"
	"strconv"
)

func (handlers *Handlers) AddMemberIntoRoom(ctx *fasthttp.RequestCtx) {
	member := models.Member{}
	member.UnmarshalJSON(ctx.PostBody())
	err := handlers.usecase.AddMemberIntoRoom(&member)
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

func (handlers *Handlers) DeleteMemberFromRoom(ctx *fasthttp.RequestCtx) {
	member := models.Member{}
	member.UnmarshalJSON(ctx.PostBody())
	err := handlers.usecase.DeleteMemberFromRoom(&member)
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

func (handlers *Handlers) GetWhoBuy(ctx *fasthttp.RequestCtx) {
	ID := ctx.UserValue("id").(string)
	roomID, err := strconv.Atoi(ID)

	whoBuy, err := handlers.usecase.GetWhoBuy(roomID)

	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(whoBuy)
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}

func (handlers *Handlers) UpdatePaperToilet (ctx *fasthttp.RequestCtx) {
	paper := models.Paper{}
	paper.UnmarshalJSON(ctx.PostBody())
	paperCount, err := handlers.usecase.UpdatePaperToilet(&paper)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		msg := messages.MessageString{}
		msg.Message = "Now user have " + strconv.Itoa(paperCount) + " papers"
		response, _ := msg.MarshalJSON()
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}



func (handlers *Handlers) ChangeValuable (ctx *fasthttp.RequestCtx) {
	valuable := models.Valuable{}
	valuable.UnmarshalJSON(ctx.PostBody())
	valueAbility, err := handlers.usecase.ChangeValuable(&valuable)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		msg := messages.MessageString{}
		msg.Message = "Now user's value is " + strconv.FormatBool(valueAbility)
		response, _ := msg.MarshalJSON()
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}


func (handlers *Handlers) ChangePayAbility (ctx *fasthttp.RequestCtx) {
	payAbility := models.PayAbility{}
	payAbility.UnmarshalJSON(ctx.PostBody())
	canPay, err := handlers.usecase.ChangePayAbility(&payAbility)
	switch err {
	case nil:
		ctx.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		msg := messages.MessageString{}
		msg.Message = "Now user's ability to pay is " + strconv.FormatBool(canPay)
		response, _ := msg.MarshalJSON()
		ctx.Write(response)
	default:
		ctx.SetStatusCode(http.StatusConflict)
		ctx.SetContentType("application/json")
		response, _:= json.Marshal(err.Error())
		ctx.Write(response)
	}
}
