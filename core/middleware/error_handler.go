package middleware

import (

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/itokun99/ms-go-boilerplate/core/config/log"
	"github.com/itokun99/ms-go-boilerplate/core/handler"
)

type ErrorMiddleware struct {
	ErrorHandlerUsecase handler.ErrorHandlerInterface
	log                 *log.LogCustom
}

func NewErrorMiddleware(r *gin.RouterGroup, ehus handler.ErrorHandlerInterface, log *log.LogCustom) {
	handler := &ErrorMiddleware{
		ErrorHandlerUsecase: ehus,
		log:                 log,
	}

	r.Use(handler.errorHandler)
}

func (eh *ErrorMiddleware) errorHandler(c *gin.Context) {
	c.Next()
	errorToPrint := c.Errors.Last()
	if errorToPrint != nil {
		// _, v := eh.ErrorHandlerUsecase.ResponseError(errorToPrint)

		c.JSON(eh.ErrorHandlerUsecase.ResponseError(errorToPrint))
		c.Abort()
		return
	}
}
