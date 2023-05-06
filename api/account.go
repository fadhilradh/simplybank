package api

import (
	"net/http"

	db "github.com/fadhilradh/simplybank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateAccountReq struct {
	OwnerName string `json:"owner_name" binding:"required"`
	Currency  string `json:"currency" binding:"required,oneof=USD IDR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data := db.CreateAccountParams{
		OwnerName: req.OwnerName,
		Currency:  req.Currency,
		Balance:   0,
	}

	account, err := server.store.CreateAccount(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
