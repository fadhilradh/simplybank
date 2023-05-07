package api

import (
	"database/sql"
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

type GetAccountByIDReq struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccountByID(ctx *gin.Context) {
	var req GetAccountByIDReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type ListAccountReq struct {
	Offset int32 `form:"offset" binding:"required"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req ListAccountReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Offset: req.Offset - 1,
		Limit:  req.Limit,
	}

	account, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
