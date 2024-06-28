package handlers

import (
	"context"
	"fmt"
	"net"
	"server"
	"server/api/tcp/middlewares"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
)

type AddressHandler struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService services.AddressService) *AddressHandler {
	return &AddressHandler{addressService}
}
func (h *AddressHandler) HandleAddAddressToUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAddAddressToUserRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding address request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	createdAddress, err := h.addressService.Create(ctx, reqData.AddressLine, reqData.Coordinates, server.UserAddressType, reqData.City)
	response := tcp.AddressResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.AddressResponse{
			Message: fmt.Sprintf("address created."),
			Address: createdAddress,
		}
	}
	resData, err := tcp.EncodeAddAddressToUserResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding address to response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}
func (h *AddressHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	firstRoute, _ := utils.RouteSplitter(TCPReq.Location)
	switch firstRoute {
	case "addresses":
		if TCPReq.Header["method"] == tcp.MethodPost {
			addToCardHandler := middleware.ApplyMiddlewares(h.HandleAddAddressToUser, middleware.AuthMiddleware)
			addToCardHandler(ctx, conn, TCPReq)
			return
		}
		tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
		return
	}
}
