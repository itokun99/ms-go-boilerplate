package model

type BaseResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type BaseHeaderRequest struct {
	ServiceName       string `header:"SERVICE_NAME"`
	ServiceEndpoint   string `header:"SERVICE_ENDPOINT"`
	ServiceExternalId string `header:"SERVICE_EXTERNAL_ID"`
}

type ErrMeta struct {
	ServiceCode string
	FieldErr    string
	CaseCode    string
}
