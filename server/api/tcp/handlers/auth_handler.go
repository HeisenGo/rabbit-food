package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/models/user"
	"server/internal/protocol"
	"server/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}
func (h *AuthHandler) HandleRegister(ctx context.Context, conn net.Conn, data []byte) {
	reqData, err := protocol.DecodeRegisterRequest(data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		return
	}
	newUser := user.NewUser(reqData.Phone, reqData.Email, reqData.Password)
	createdUser, err := h.authService.CreateUser(ctx, newUser)
	response := protocol.RegisterResponse{}
	if err != nil {
		response.Message = err.Error()
		// TODO: write a func like http.Error() to return here (else must be removed)
	} else {
		response = protocol.RegisterResponse{
			Success: err == nil,
			Message: fmt.Sprintf("User with id: %d, phone:%s, email: %s", createdUser.ID, createdUser.Phone, createdUser.Email),
			UserID:  createdUser.ID,
		}
	}
	resData, err := protocol.EncodeRegisterResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding register response:", err)
		return
	}

	conn.Write(resData)
}

func (h *AuthHandler) HandleLogin(ctx context.Context, conn net.Conn, data []byte) {
	reqData, err := protocol.DecodeLoginRequest(data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding login request:", err)
		return
	}
	authenticatedUserToken, err := h.authService.LoginUser(ctx, reqData.PhoneOrEmail, reqData.Password)
	response := protocol.LoginResponse{}
	if err != nil {
		response.Message = err.Error()
	} else {
		response = protocol.LoginResponse{
			Success:   err == nil,
			Message:   fmt.Sprintf("User-Logged-in."),
			AuthToken: authenticatedUserToken,
		}
	}
	resData, err := protocol.EncodeLoginReponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding Login response:", err)
		return
	}
	conn.Write(resData)
}
