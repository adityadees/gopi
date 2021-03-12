package responses

type ResponseGet struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePost struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
