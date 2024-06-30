package handlers

import (
	"context"
	"encoding/json"
	"net"
	"server/internal/models/address"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) UpdateFirstName(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		FirstName string `json:"first_name"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	user, err := h.userService.UpdateUserFirstName(ctx, userID, reqData.FirstName)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(user)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) UpdateLastName(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		LastName string `json:"last_name"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	user, err := h.userService.UpdateUserLastName(ctx, userID, reqData.LastName)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(user)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) UpdateEmail(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		Email string `json:"email"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	user, err := h.userService.UpdateUserEmail(ctx, userID, reqData.Email)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(user)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) UpdatePhone(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		Phone string `json:"phone"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	user, err := h.userService.UpdateUserPhone(ctx, userID, reqData.Phone)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(user)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) UpdatePassword(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		Password string `json:"password"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	user, err := h.userService.UpdateUserPassword(ctx, userID, reqData.Password)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(user)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *UserHandler) DeleteAddress(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		AddressID uint `json:"address_id"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	err = h.userService.DeleteUserAddress(ctx, userID, reqData.AddressID)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	tcp.SendResponse(conn, tcp.StatusOK, nil, nil)
}

func (h *UserHandler) AddAddress(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var reqData struct {
		AddressLine string              `json:"address_line"`
		Coordinates address.Coordinates `json:"coordinates"`
		Types       string              `json:"types"`
		City        string              `json:"city"`
	}

	if err := json.Unmarshal(req.Data, &reqData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusUnauthorized, nil, err.Error())
		return
	}

	addr := address.NewAddress(reqData.AddressLine, reqData.Coordinates, reqData.Types, reqData.City)
	_, err = h.userService.AddUserAddress(ctx, userID, addr)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}

	resData, _ := json.Marshal(addr)
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}
