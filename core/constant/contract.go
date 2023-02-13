package constant

var FieldErr string
var ServiceCode string
var CaseCode string

const (
	ErrInvalidFieldMandatory       = "400XX02"
	ErrInvalidFieldFormatHeader    = "400XX010"
	ErrTokenInvalid                = "401XX01"
	ErrConflict                    = "409XX00"
	ErrGeneralError                = "500XX00"
	ErrInternalError               = "500XX01"
	ErrExternalError               = "500XX06"
	ErrInvalidBill                 = "404XX12"
	ErrInvalidAmount               = "404XX13"
	ErrPaidBill                    = "404XX14"
	ErrInvalidFieldFormat          = "400XX01"
	ErrUnauthorized                = "401XX00"
	ErrBadRequest                  = "400XX00"
	ErrInvalidFieldMandatoryHeader = "400XX020"
	ErrCannotSaveToDB              = "500XX02"
	ErrCannotSaveToFile            = "500XX03"
	ErrTimeout                     = "504XX00"
	ErrDataNotFound                = "404XX00"
	ErrDateSchedule                = "500XX05"
	ErrInvalidMandatoryCustLimitId = "400XX03"
	ErrInvalidMandatoryLimitAmount = "400XX04"
	ErrInvalidMandatoryInsLimit    = "400XX05"
	ErrInvalidMandatorySp3no       = "400XX06"
	ErrInvalidMandatoryholdAccNo   = "400XX07"
	ErrInvalidMandatoryHoldAmount  = "400XX08"
	ErrInvalidMandatoryNrk         = "400XX11"
	ErrInactiveAccount             = "403XX18"
	ErrTrxNotPermitted             = "403XX15"
)
