package TuShare

type ApiResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      struct {
		Fields []string   `json:"fields"`
		Items  [][]string `json:"items"`
	} `json:"data"`
}

