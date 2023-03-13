package common

const (
	SUCCESS = 200
	ERROR   = 101
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Success 成功
func Success(data interface{}) Result {
	return Result{SUCCESS, "success", data}
}

// Error 错误
func Error(msg string) Result {
	return Result{ERROR, msg, ""}
}
