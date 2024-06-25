package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/models/user"
	"server/internal/protocol/tcp"
	"server/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}
func (h *AuthHandler) HandleRegister(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeRegisterRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newUser := user.NewUser(reqData.Phone, reqData.Email, reqData.Password)
	_, token, err := h.authService.CreateUser(ctx, newUser)
	// response := tcp.RegisterResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.RegisterResponse{
		Message: "user created",
		Token:   token,
	}
	//}

	resData, err := tcp.EncodeRegisterResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding register response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *AuthHandler) HandleLogin(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeLoginRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding login request:", err)
		return
	}
	authenticatedUserToken, err := h.authService.LoginUser(ctx, reqData.PhoneOrEmail, reqData.Password)
	//response := tcp.LoginResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.LoginResponse{
		Message: "User-Logged-in.",
		Token:   authenticatedUserToken,
	}
	//}
	resData, err := tcp.EncodeLoginResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding Login response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *AuthHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	if TCPReq.Header["method"] == tcp.MethodPost {
		switch TCPReq.Location {
		case "register":
			h.HandleRegister(ctx, conn, TCPReq)
		case "login":
			h.HandleLogin(ctx, conn, TCPReq)
		}
	} else {
		tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
		return
	}
}
