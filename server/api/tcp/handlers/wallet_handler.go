package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/models/wallet/wallet"
	"server/internal/protocol/tcp"
	"server/services"
)

type WalletHandler struct {
	walletService services.WalletService
}

func NewWalletHandler(walletService services.WalletService) *WalletHandler {
	return &WalletHandler{walletService}
}

func (h *WalletHandler) HandleWalletCreation(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeRegisterRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	//userID := ctx.Value(UserID)
	newWallet := wallet.NewWallet(reqData.Phone, reqData.Email, reqData.Password)
	createdUserToken, err := h.authService.CreateUser(ctx, newUser)
	response := tcp.RegisterResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.RegisterResponse{
			Message: fmt.Sprintf("user created"),
			Token:   createdUserToken,
		}
	}
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
	response := tcp.LoginResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.LoginResponse{
			Message:   fmt.Sprintf("User-Logged-in."),
			AuthToken: authenticatedUserToken,
		}
	}
	resData, err := tcp.EncodeLoginResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding Login response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *AuthHandler) AuthRouter(ctx context.Context, conn net.Conn, req *tcp.Request) {
	switch req.Location {
	case "register":
		h.HandleRegister(ctx, conn, req)
	case "login":
		h.HandleLogin(ctx, conn, req)
	}
}
