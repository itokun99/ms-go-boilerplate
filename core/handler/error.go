package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	// "bitbucket.org/bridce/ms-briapi-brifast-outgoing/models/apply_bg"
	"github.com/itokun99/ms-go-boilerplate/core/constant"
	"github.com/itokun99/ms-go-boilerplate/core/model"
	"github.com/itokun99/ms-go-boilerplate/core/util"

	// "bitbucket.org/bridce/ms-briapi-brifast-outgoing/utils"

	"github.com/Saucon/errcntrct"
	"github.com/gin-gonic/gin"

	"github.com/lib/pq"
)

type ErrorHandlerStruct struct {
}

type ErrorHandlerInterface interface {
	ResponseError(error interface{}) (int, interface{})
}

func NewErrorHandler() ErrorHandlerInterface {
	return &ErrorHandlerStruct{}
}

func (e ErrorHandlerStruct) ResponseError(A interface{}) (int, interface{}) {
	var T interface{}
	var fieldNameErr string
	var serviceCode string
	var caseCode string

	if A.(*gin.Error).Meta != nil {
		fieldNameErr = A.(*gin.Error).Meta.(model.ErrMeta).FieldErr
		serviceCode = A.(*gin.Error).Meta.(model.ErrMeta).ServiceCode
		caseCode = A.(*gin.Error).Meta.(model.ErrMeta).CaseCode
	}

	// Check A is a correct error type and assign to T
	if A.(*gin.Error).Err != nil {
		T = A.(*gin.Error).Err
	}

	switch T.(type) {
	case error:
		if _, ok := T.(*pq.Error); ok {
			switch T.(*pq.Error).Code.Name() {
			case "unique_violation":
				return errcntrct.ErrorMessage(http.StatusBadRequest, "", errors.New(constant.ErrGeneralError))
			}
		}

		switch T.(error).Error() {
		case constant.ErrInvalidFieldFormat:
			return ResponseErrorAdapter(T.(error), http.StatusBadRequest, constant.ErrInvalidFieldFormat, serviceCode, "", fieldNameErr)
		case constant.ErrInvalidFieldMandatory:
			return ResponseErrorAdapter(T.(error), http.StatusBadRequest, constant.ErrInvalidFieldMandatory, serviceCode, "", fieldNameErr)
		case constant.ErrConflict:
			return ResponseErrorAdapter(T.(error), http.StatusConflict, "", serviceCode, "", "")
		case constant.ErrBadRequest:
			return ResponseErrorAdapter(T.(error), http.StatusBadRequest, "", serviceCode, "", "")
		case constant.ErrInternalError:
			return ResponseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "", "")
		case constant.ErrExternalError:
			return ResponseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "", "")
		case constant.ErrGeneralError:
			return ResponseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, caseCode, "")
		case constant.ErrUnauthorized:
			return ResponseErrorAdapter(T.(error), http.StatusUnauthorized, "", serviceCode, caseCode, "")
		case constant.ErrTimeout:
			return ResponseErrorAdapter(T.(error), http.StatusRequestTimeout, "", serviceCode, "", "")
		case constant.ErrDataNotFound:
			return ResponseErrorAdapter(T.(error), http.StatusNotFound, "", serviceCode, caseCode, "")
		default:
			return ResponseErrorAdapter(errors.New(constant.ErrGeneralError), http.StatusInternalServerError, "", serviceCode, "", "")
		}
	}

	return ResponseErrorAdapter(T.(error), http.StatusInternalServerError, "", serviceCode, "", "")
}

func ResponseErrorAdapter(errHttpStatus interface{}, httpStatusCode int, ctr string, serviceCode string, customCaseCode string, fieldErr string) (int, model.BaseResponse) {
	_, errData := errcntrct.ErrorMessage(httpStatusCode, "", errHttpStatus)

	var resp model.BaseResponse
	errCase := strconv.Itoa(httpStatusCode)
	caseCode := strings.Split(errData.Code, "XX")
	if customCaseCode != "" {
		resp.ResponseCode = errCase + serviceCode + customCaseCode
	} else {
		resp.ResponseCode = errCase + serviceCode + caseCode[1]
	}

	if strings.Contains(constant.FieldErr, " ") {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, constant.FieldErr)
	} else if ctr == "400XX01" || ctr == "400XX02" {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg, util.LowerCamelCase(fieldErr))
	} else {
		resp.ResponseMessage = fmt.Sprintf(errData.Msg)
	}
	return httpStatusCode, resp
}
