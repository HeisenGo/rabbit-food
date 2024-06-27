package handlers

import (
	"context"
	"fmt"
	"net"
	middleware "server/api/tcp/middlewares"
	creditCard "server/internal/models/wallet/credit_card"
	"server/internal/models/wallet/wallet"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
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
	newCard := creditCard.NewCreditCard(reqData.CardNumber)
	createdCard, err := h.walletService.AddCardToWalletByUserID(ctx, newCard)
	//response := tcp.AddCardToWalletResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.AddCardToWalletResponse{
		Message: "card added.",
		Card:    createdCard,
	}
	//}
	resData, err := tcp.EncodeAddCardToWalletResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *WalletHandler) HandleWalletCards(ctx context.Context, conn net.Conn, req *tcp.Request) {
	userWalletCards, err := h.walletService.GetUserWalletCards(ctx)
	//response := tcp.GetUserWalletCardsResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	} //else {
	response := tcp.GetUserWalletCardsResponse{
		Message: "user wallet cards successfuly fetched.",
		Cards:   userWalletCards,
	}
	//}
	resData, err := tcp.EncodeGetUserWalletCardsResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *WalletHandler) HandleDeposit(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeDepositRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	card := creditCard.NewCreditCard(reqData.CardNumber)
	userWallet, err := h.walletService.Deposit(ctx, card, reqData.Amount)
	//response := tcp.DepositResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.DepositResponse{
		Message: "successful deposit.",
		Wallet:  userWallet,
	}
	//}
	resData, err := tcp.EncodeDepositResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *WalletHandler) HandleWithdraw(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeWithdrawRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	card := creditCard.NewCreditCard(reqData.CardNumber)
	userWallet, err := h.walletService.Withdraw(ctx, card, reqData.Amount)
	//response := tcp.WithdrawResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.WithdrawResponse{
		Message: "successful withdraw.",
		Wallet:  userWallet,
	}
	//}
	resData, err := tcp.EncodeWithdrawResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *WalletHandler) HandleGetWallet(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeGetWalletRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding get wallet request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	chosenWallet := wallet.NewWalletByID(reqData.WalletID)
	userWallet, err := h.walletService.GetWallet(ctx, chosenWallet)
	//response := tcp.WithdrawResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.WithdrawResponse{
		Message: "successful withdraw.",
		Wallet:  userWallet,
	}
	//}
	resData, err := tcp.EncodeWithdrawResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *WalletHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	firstRoute, _ := utils.RouteSplitter(TCPReq.Location)
	switch firstRoute {
	case "wallet":
		if TCPReq.Header["method"] == tcp.MethodGet {
			getWalletHandler := middleware.ApplyMiddlewares(h.HandleGetWallet, middleware.AuthMiddleware)
			getWalletHandler(ctx, conn, TCPReq)
			return
		}
	case "cards":
		if TCPReq.Header["method"] == tcp.MethodPost {
			addToCardHandler := middleware.ApplyMiddlewares(h.HandleAddCardToWallet, middleware.AuthMiddleware)
			addToCardHandler(ctx, conn, TCPReq)
			return
		}
		if TCPReq.Header["method"] == tcp.MethodGet {
			walletCardsHandler := middleware.ApplyMiddlewares(h.HandleWalletCards, middleware.AuthMiddleware)
			walletCardsHandler(ctx, conn, TCPReq)
			return
		}
	case "deposit":
		if TCPReq.Header["method"] == tcp.MethodPost {
			depositHandler := middleware.ApplyMiddlewares(h.HandleDeposit, middleware.AuthMiddleware)
			depositHandler(ctx, conn, TCPReq)
			return
		}
	case "withdraw":
		if TCPReq.Header["method"] == tcp.MethodPost {
			depositHandler := middleware.ApplyMiddlewares(h.HandleWithdraw, middleware.AuthMiddleware)
			depositHandler(ctx, conn, TCPReq)
			return
		}
	}
	tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
}
