package handlers

import (
	"fmt"
	"net"
	"rabbit-food/internal/models/user"
	"rabbit-food/internal/protocol"
	"rabbit-food/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) HandleRegister(conn net.Conn, data []byte) {
	reqData, err := protocol.DecodeRegisterRequest(data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		return
	}
	newUser := user.NewUser(reqData.Phone, reqData.Email, reqData.Password)
	createdUser, err := h.userService.CreateUser(newUser)
	response := protocol.RegisterResponse{
		Success: err == nil,
		Message: fmt.Sprintf("User with id: %d, phone:%s, email: %s", createdUser.ID, createdUser.Phone, createdUser.Email),
		UserID:  createdUser.ID,
	}
	if err != nil {
		response.Message = err.Error()
	}

	resData, err := protocol.EncodeRegisterResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding register response:", err)
		return
	}

	conn.Write(resData)
}
