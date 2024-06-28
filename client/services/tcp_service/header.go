package tcp_service

var token *string

func GetToken() string {
	return *token
}

func SetToken(tk string) {
	token = &tk
}

func UnSetToken(){
	token = nil
}
