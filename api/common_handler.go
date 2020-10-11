package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	CreateRequest(c *gin.Context)(interface{},error)
	Handle(ctx *gin.Context,request interface{})(interface{},error)
	CreateResponse(resp interface{})(interface{},error)
}

func CommonHandler(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Create a standard context.Context from the gin context
		//ctx := contextutils.CreateContextUsingGinContext(c)
		//logger := logutils.GetCtxLogEntry(ctx, "CommonHandler")
		req, err := handler.CreateRequest(c)
		if err != nil {
			//logger.WithField("error", err).Error()
			//bytes := contextutils.GetBaseFailureResponse(constants.STATUS_CODE_INVALID_REQUEST, err)
			c.Data(http.StatusBadRequest, c.ContentType(), []byte(err.Error()))
			return
		}

		//logger.WithField("request", jsonutils.MustGetJson(req)).Info()
		response, err := handler.Handle(c, req)
		if err != nil {
			//logger.WithField("error", err).Error()
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		//logger.WithField("response", jsonutils.MustGetJson(response)).Info()
		c.JSON(http.StatusOK, response)

	}
}
