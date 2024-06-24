package handlers

import (
	"context"
	"net"
	"server/internal/models/user"
	"server/internal/protocol/tcp"
	"server/services"
	"strconv"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) HandleCreateUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeRegisterRequest(req.Data)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newUser := &user.User{
		Phone:     reqData.Phone,
		Email:     reqData.Email,
		FirstName: reqData.FirstName,
		LastName:  reqData.LastName,
		Password:  reqData.Password,
	}
	_, err = h.userService.CreateUser(ctx, newUser)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.RegisterResponse{
		Message: "User created",
		Token:   nil, // Add actual token here if necessary
	}
	resData, err := tcp.EncodeRegisterResponse(response)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *UserHandler) HandleGetUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	id, err := strconv.Atoi(req.Header["id"])
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	user, err := h.userService.GetUserByID(ctx, uint(id))
	if err != nil {
		tcp.Error(conn, tcp.StatusNotFound, nil, err.Error())
		return
	}
	response := tcp.UserResponse{
		Message: "User retrieved",
		User:    user,
	}
	resData, err := tcp.EncodeUserResponse(response)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) HandleUpdateUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeUserRequest(req.Data)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	user := &user.User{
		ID:        reqData.ID,
		Phone:     reqData.Phone,
		Email:     reqData.Email,
		FirstName: reqData.FirstName,
		LastName:  reqData.LastName,
		Password:  reqData.Password,
		BirthDate: reqData.BirthDate,
		Address:   reqData.Address,
	}
	updatedUser, err := h.userService.UpdateUser(ctx, user)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.UserResponse{
		Message: "User updated",
		User:    updatedUser,
	}
	resData, err := tcp.EncodeUserResponse(response)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) HandleDeleteUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	id, err := strconv.Atoi(req.Header["id"])
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	err = h.userService.DeleteUser(ctx, uint(id))
	if err != nil {
		tcp.Error(conn, tcp.StatusNotFound, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusNoContent, nil, nil)
}

func (h *UserHandler) UserRouter(ctx context.Context, conn net.Conn, req *tcp.Request) {
	switch req.Location {
	case "create":
		h.HandleCreateUser(ctx, conn, req)
	case "get":
		h.HandleGetUser(ctx, conn, req)
	case "update":
		h.HandleUpdateUser(ctx, conn, req)
	case "delete":
		h.HandleDeleteUser(ctx, conn, req)
	}
}
