package api

import (
	"database/sql"
	"net/http"

	db "github.com/ericlamnguyen/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// create new account handler function
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// Validate the request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Handle the request
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// get account by ID handler function
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	// Validate the input
	if err := ctx.ShouldBindUri(&req); err != nil {
		// error - invalid input
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Handle the request
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		// error - no account found
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		// error - general server error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// success
	ctx.JSON(http.StatusOK, account)
}

// list accounts handler function
type listAccountRequest struct {
	Owner    string `form:"owner" binding:"required"`
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=2,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest

	// Validate the input
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Handle the request
	arg := db.ListAccountsParams{
		Owner:  req.Owner,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
