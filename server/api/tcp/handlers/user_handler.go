package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/models/user"
	"server/internal/protocol"
	"server/services"
)


type UserHandler struct {
	userService services.UserService
	
}
type AuthHandler struct {
	authService services.AuthService
}
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService }
}
func (h *UserHandler) HandleRegister(ctx context.Context, conn net.Conn, data []byte) {
	reqData, err := protocol.DecodeRegisterRequest(data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		return
	}
	newUser := user.NewUser(reqData.Phone, reqData.Email, reqData.Password)
	createdUser, err := h.userService.CreateUser(ctx, newUser)
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
func (s *AuthHandler)HandleLogin(ctx context.Context, conn net.Conn, data []byte)  {
	reqData, err := protocol.DecodeLoginRequest(data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding login request:", err)
		return
	}
	authenticatedusertoken,err := s.authService.LoginByEmail(ctx,reqData.Email,reqData.Password)
	response := protocol.LoginResponse{}
	if err != nil {
		if reqData.Email==""{
			fmt.Println("Error By Login With User,Try With Your Phone:")
			s.authService.LoginByPhone(ctx,reqData.Phone,reqData.Password)
		}else { 
		response.Message = err.Error()
		}
		// TODO: write a func like http.Error() to return here (else must be removed)
	} else {
		response = protocol.LoginResponse{
			Success: err == nil,
			Message: fmt.Sprintf("User Logged in with, phone:%s or email: %s", reqData.Phone, reqData.Email),
			Usertoken:authenticatedusertoken,
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
