package goquots

type QuotsUser struct {
	Id       string  `json:"id"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Credits  float64 `json:"credits"`
	SpentOn  []Spent `json:"spenton"`
}

type Spent struct {
	AppId string                 `json:"appid"`
	Usage map[string]interface{} `json:"usage"`
}

type CanProceed struct {
	UserId  string `json:"userid"`
	Proceed bool   `json:"proceed"`
}

type ErrorReport struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
