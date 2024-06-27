package handlers

import (
	"context"
	"encoding/json"
	"net"
	"server/internal/models/user"
	"server/internal/protocol/tcp"
	"server/services"
	"strconv"
)

type UserProfileHandler struct {
	userProfileService services.UserProfileService
}

func NewUserProfileHandler(userProfileService services.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{userProfileService}
}

func (h *UserProfileHandler) HandleGetProfile(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var requestData map[string]string
	if err := json.Unmarshal(req.Data, &requestData); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, "Invalid request data")
		return
	}

	userIDStr, ok := requestData["user_id"]
	if !ok {
		tcp.Error(conn, tcp.StatusBadRequest, nil, "User ID not provided")
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, "Invalid user ID")
		return
	}

	user, err := h.userProfileService.GetUserProfile(ctx, uint(userID))
	if err != nil {
		tcp.Error(conn, tcp.StatusNotFound, nil, "User not found")
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, "Failed to encode response")
		return
	}

	tcp.SendResponse(conn, tcp.StatusOK, nil, response)
}

func (h *UserProfileHandler) HandleUpdateProfile(ctx context.Context, conn net.Conn, req *tcp.Request) {
	var user user.User
	if err := json.Unmarshal(req.Data, &user); err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, "Invalid request data")
		return
	}

	updatedUser, err := h.userProfileService.UpdateUserProfile(ctx, &user)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, "Failed to update user profile")
		return
	}

	response, err := json.Marshal(updatedUser)
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, "Failed to encode response")
		return
	}

	tcp.SendResponse(conn, tcp.StatusOK, nil, response)
}

func (h *UserProfileHandler) ServeTCP(ctx context.Context, conn net.Conn, req *tcp.Request) {
	switch req.Header["method"] {
	case tcp.MethodGet:
		h.HandleGetProfile(ctx, conn, req)
	case tcp.MethodPost:
		h.HandleUpdateProfile(ctx, conn, req)
	default:
		tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "Method not allowed")
	}
}
