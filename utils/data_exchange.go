package utils

import "gorestAPI/models"

type DataExchange struct {
	StatusCode int     `json:"status_code,omitempty"`
	Header     *Header `json:"header"`
	Body       *Body   `json:"body,omitempty"`
}

type Header struct {
	Accept         string `header:"Accept"`
	AcceptLanguage string `header:"Accept-Language"`
	CacheControl   string `header:"Cache-Control"`
	UserAgent      string `header:"User-Agent"`
}

type Body struct {
	Status  string         `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Test    *[]models.Test `json:"test_data,omitempty"`
}

func (DE *DataExchange) SetState(status string, message string) {
	DE.Body.Status = status
	DE.Body.Message = message
}
