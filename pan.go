package koyc

type DefaultPanVerifyDataResponse struct {
	FatherName      string   `json:"father_name"`
	PanNumber       string   `json:"pan_number"`
	FullName        string   `json:"full_name"`
}

type DefaultPanVerifyResponse struct {
	MessageCode string                       `json:"message_code"`
	StatusCode  int                          `json:"status_code"`
	Message     string                       `json:"message"`
	Success     bool                         `json:"success"`
	Data        DefaultPanVerifyDataResponse `json:"data"`
}
