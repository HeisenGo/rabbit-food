package handlers

import (
	"context"
	"fmt"
	"net"
	middleware "server/api/tcp/middlewares"
	"server/internal/models/wallet/credit_card"
	"server/internal/protocol/tcp"
	"server/services"
)

type WalletHandler struct {
	walletService services.WalletService
}

func NewWalletHandler(walletService services.WalletService) *WalletHandler {
	return &WalletHandler{walletService}
}
func (h *WalletHandler) HandleAddCardToWallet(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAddCardToWalletRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	newCard := wallet.NewCreditCard(reqData.CardNumber)
	createdCard, err := h.walletService.AddCardToWalletByUserID(ctx, newCard)
	response := tcp.AddCardToWalletResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.AddCardToWalletResponse{
			Message: fmt.Sprintf("card added."),
			Card:    createdCard,
		}
	}
	resData, err := tcp.EncodeAddCardToWalletResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *WalletHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	if TCPReq.Header["method"] == tcp.MethodPost {
		switch TCPReq.Location {
		case "cards":
			addToCardHandler := middleware.ApplyMiddlewares(h.HandleAddCardToWallet, middleware.AuthMiddleware)
			addToCardHandler(ctx, conn, TCPReq)
		}
	} else {
		tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
		return
	}
}
