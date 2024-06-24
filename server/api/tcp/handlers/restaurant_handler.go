package handlers

import (
	"context"
	"fmt"
	"net"
	middleware "server/api/tcp/middlewares"
	"server/internal/models/restaurant/restaurant"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
)

type RestaurantHandler struct {
	restauarntService services.RestaurantService
}

func NewRestaurantHandler(restauarntService services.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restauarntService}
}

func (h *RestaurantHandler) HandleCreateRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeCreateRestaurantRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding create restaurant request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newRestaurant := restaurant.NewRestaurant(reqData.Name, reqData.Phone, reqData.City, reqData.Address, reqData.Coordinates)
	createdRestaurant, err := h.restauarntService.CreateResturantForOwner(ctx, newRestaurant)

	response := tcp.CreateRestaurantResponse{}

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.CreateRestaurantResponse{
			Message:    fmt.Sprintf("restaurant created :)"),
			Restaurant: createdRestaurant,
		}
	}

	resData, err := tcp.EncodeCreateRestaurantResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding create restaurn response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

// func (h *WalletHandler) HandleWalletCards(ctx context.Context, conn net.Conn, req *tcp.Request) {
// 	userWalletCards, err := h.walletService.GetUserWalletCards(ctx)
// 	response := tcp.GetUserWalletCardsResponse{}
// 	if err != nil {
// 		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
// 		return
// 	} else {
// 		response = tcp.GetUserWalletCardsResponse{
// 			Message: fmt.Sprintf("user wallet cards successfuly fetched."),
// 			Cards:   userWalletCards,
// 		}
// 	}
// 	resData, err := tcp.EncodeGetUserWalletCardsResponse(response)
// 	if err != nil {
// 		//logger.Error("Error encoding register response:", err)
// 		fmt.Println("Error encoding get cards response:", err)
// 		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
// 		return
// 	}
// 	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
// }

// func (h *WalletHandler) HandleDeposit(ctx context.Context, conn net.Conn, req *tcp.Request) {
// 	reqData, err := tcp.DecodeDepositRequest(req.Data)
// 	if err != nil {
// 		//logger.Error("Error decoding register request:", err)
// 		fmt.Println("Error decoding register request:", err)
// 		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
// 		return
// 	}
// 	card := wallet.NewCreditCard(reqData.CardNumber)
// 	userWallet, err := h.walletService.Deposit(ctx, card, reqData.Amount)
// 	response := tcp.DepositResponse{}
// 	if err != nil {
// 		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
// 		return
// 	} else {
// 		response = tcp.DepositResponse{
// 			Message: fmt.Sprintf("successful deposit."),
// 			Wallet:  userWallet,
// 		}
// 	}
// 	resData, err := tcp.EncodeDepositResponse(response)
// 	if err != nil {
// 		//logger.Error("Error encoding register response:", err)
// 		fmt.Println("Error encoding get cards response:", err)
// 		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
// 		return
// 	}
// 	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
// }

// func (h *WalletHandler) HandleWithdraw(ctx context.Context, conn net.Conn, req *tcp.Request) {
// 	reqData, err := tcp.DecodeWithdrawRequest(req.Data)
// 	if err != nil {
// 		//logger.Error("Error decoding register request:", err)
// 		fmt.Println("Error decoding register request:", err)
// 		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
// 		return
// 	}
// 	card := wallet.NewCreditCard(reqData.CardNumber)
// 	userWallet, err := h.walletService.Withdraw(ctx, card, reqData.Amount)
// 	response := tcp.WithdrawResponse{}
// 	if err != nil {
// 		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
// 		return
// 	} else {
// 		response = tcp.WithdrawResponse{
// 			Message: fmt.Sprintf("seccessful withdraw."),
// 			Wallet:  userWallet,
// 		}
// 	}
// 	resData, err := tcp.EncodeWithdrawResponse(response)
// 	if err != nil {
// 		//logger.Error("Error encoding register response:", err)
// 		fmt.Println("Error encoding get cards response:", err)
// 		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
// 		return
// 	}
// 	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
// }

func (h *RestaurantHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	firstRoute, _ := utils.RouteSplitter(TCPReq.Location)
	switch firstRoute {
	case "restaurant":
		if TCPReq.Header["method"] == tcp.MethodPost {
			createRestaurantHandler := middleware.ApplyMiddlewares(h.HandleCreateRestaurant, middleware.AuthMiddleware)
			createRestaurantHandler(ctx, conn, TCPReq)
			return
		}
		// if TCPReq.Header["method"] == tcp.MethodGet {
		// 	walletCardsHandler := middleware.ApplyMiddlewares(h.HandleWalletCards, middleware.AuthMiddleware)
		// 	walletCardsHandler(ctx, conn, TCPReq)
		// 	return
		// }
		// case "deposit":
		// 	if TCPReq.Header["method"] == tcp.MethodPost {
		// 		depositHandler := middleware.ApplyMiddlewares(h.HandleDeposit, middleware.AuthMiddleware)
		// 		depositHandler(ctx, conn, TCPReq)
		// 		return
		// 	}
		// case "withdraw":
		// 	if TCPReq.Header["method"] == tcp.MethodPost {
		// 		depositHandler := middleware.ApplyMiddlewares(h.HandleWithdraw, middleware.AuthMiddleware)
		// 		depositHandler(ctx, conn, TCPReq)
		// 		return
		// 	}
	}
	tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
	return
}
