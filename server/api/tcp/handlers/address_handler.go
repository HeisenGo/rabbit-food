package handlers

import (
	"context"
	"fmt"
	"net"
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
	reqData, err := tcp.DecodeAdd_AddressToUserRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	reqData.Types , err = utils.GetUserIDFromContext(ctx)
	if err != nil{
		fmt.Println("Error Can't Take Users ID",err)
		tcp.Error(conn,tcp.StatusConflict,nil,err.Error())
	}
	//newAddress := address.NewAddress(reqData.AddressLine,reqData.Cordinates,reqData.Types,reqData.City)
	createdAddress, err := h.addressService.Create(ctx,reqData.AddressLine,reqData.Cordinates,reqData.Types,reqData.City)
	response := tcp.AddressResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.AddressResponse{
			Message: fmt.Sprintf("card added."),
			Address:   createdAddress ,
		}
	}
	resData, err := tcp.EncodeAdd_AddressToUserResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}
func (h *AddressHandler) HandleAddAddressToRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAdd_AddressToUserRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	reqData.Types , err = utils.GetUserIDFromContext(ctx)
	//Check the userID if it exist in the restaurant table and etc...
	if err != nil{
		fmt.Println("Error Can't Take Users ID",err)
		tcp.Error(conn,tcp.StatusConflict,nil,err.Error())
	}
	//newAddress := address.NewAddress(reqData.AddressLine,reqData.Cordinates,reqData.Types,reqData.City)
	createdAddress, err := h.addressService.Create(ctx,reqData.AddressLine,reqData.Cordinates,reqData.Types,reqData.City)
	response := tcp.AddressResponse{}
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} else {
		response = tcp.AddressResponse{
			Message: fmt.Sprintf("card added."),
			Address:   createdAddress ,
		}
	}
	resData, err := tcp.EncodeAdd_AddressToUserResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}