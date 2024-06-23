package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/models/user"
	"server/internal/protocol"
	"server/services"
	"server/pkg/logger"
)

type UserHandler struct {
	userService services.UserService
	logger *logger.CustomLogger
}

func NewUserHandler(userService services.UserService,logger *logger.CustomLogger) *UserHandler {
	return &UserHandler{userService: userService ,
		logger: logger}
}

func (h *UserHandler) HandleRegister(ctx context.Context, conn net.Conn, data []byte) {
	reqData, err := protocol.DecodeRegisterRequest(data)
	if err != nil {
		h.logger.Error("Error decoding register request:", err)
		return
	}
	newUser := user.NewUser(reqData.Phone, reqData.Email, reqData.Password)
	createdUser, err := h.userService.CreateUser(ctx, newUser)
	response := protocol.RegisterResponse{
		Success: err == nil,
		Message: fmt.Sprintf("User with id: %d, phone:%s, email: %s", createdUser.ID, createdUser.Phone, createdUser.Email),
		UserID:  createdUser.ID,
	}
	if err != nil {
		response.Message = err.Error()
		h.logger.Error("In Respones :",response.Message)
	}
	resData, err := protocol.EncodeRegisterResponse(response)
	if err != nil {
		h.logger.Error("In Respones :",response.Message)
		return
	}

	conn.Write(resData)
}
