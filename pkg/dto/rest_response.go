package dto

type RestResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}
