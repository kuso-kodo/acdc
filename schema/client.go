package schema

type ClientRegisterRequest struct {
	Room string `json:"room"`
}

type ClientRegisterResponse struct {
	Status  bool   `json:"status"`
	Period  int    `json:"period"`
	Message string `json:"msg"`
}

type ClientHeartBeatRequest struct {
	Room               string  `json:"room"`
	Power              bool    `json:"power"`
	Mode               int     `json:"mode"`
	TargetTemperature  float32 `json:"target"`
	CurrentTemperature float32 `json:"current"`
	FanSpeed           uint    `json:"wind"`
}

const (
	ACK = 0
	RST = 1
	ERR = -1
)

type ClientHeartBeatResponse struct {
	Status  int     `json:"status"`
	Serve   bool    `json:"wind"`
	Fee     float32 `json:"cost"`
	Message string  `json:"msg"`
}
