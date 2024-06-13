package model

type Errors struct {
	Code    string `json:"code"`
	Message error  `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param,omitempty"`
	Error   string `json:"error"`
}
