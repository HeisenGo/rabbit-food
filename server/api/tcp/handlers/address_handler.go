package handlers

import (
	"context"
	"fmt"
	"net"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
	middleware "server/api/tcp/middlewares"
)
type AddressHandler struct {
	addressService services.AddressService
}
func NewAddressHandler(addressService services.AddressService) *AddressHandler {
	return &AddressHandler{addressService}
}
func (h *AddressHandler) HandleAddAddressToUser(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAdd_AddressToUserRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding address request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	userID , err := utils.GetUserIDFromContext(ctx)
	if err != nil{
		fmt.Println("Error Can't Take Users ID",err)
		tcp.Error(conn,tcp.StatusConflict,nil,err.Error())
	}
	createdAddress, err := h.addressService.Create(ctx,reqData.AddressLine,reqData.Cordinates,reqData.Types,reqData.City,userID)
	response := tcp.AddressResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.AddressResponse{
			Message: fmt.Sprintf("address created."),
			Address:   createdAddress ,
		}
	}
	resData, err := tcp.EncodeAdd_AddressToUserResponse(response)
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
	case "createaddresses":
		if TCPReq.Header["method"] == tcp.MethodPost {
			addToCardHandler := middleware.ApplyMiddlewares(h.HandleAddAddressToUser, middleware.AuthMiddleware)
			addToCardHandler(ctx, conn, TCPReq)
			return
		}
	tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
	return
}
}