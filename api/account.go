package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/husamettinarabaci/tinyurl/db/sqlc"
	"github.com/husamettinarabaci/tinyurl/token"
)

type createTransformRequest struct {
	LongUrl string `json:"longurl" binding:"required"`
}

func (server *Server) createTransform(ctx *gin.Context) {
	var req createTransformRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTransformParams{
		Owner:   authPayload.Username,
		LongUrl: req.LongUrl,
	}

	transform, err := server.store.CreateTransformTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transform)
}

type getTransformRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTransform(ctx *gin.Context) {
	var req getTransformRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transform, err := server.store.GetTransform(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if transform.Owner != authPayload.Username {
		err := errors.New("transform doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transform)
}

type listTransformRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTransforms(ctx *gin.Context) {
	var req listTransformRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTransformsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	transforms, err := server.store.ListTransforms(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transforms)
}
