package handlers

import (
	"context"
	"fmt"
	"net"
	middleware "server/api/tcp/middlewares"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
)

type RestaurantHandler struct {
	restaurantService services.RestaurantService
}

func NewRestaurantHandler(restaurantService services.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{restaurantService}
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
	createdRestaurant, err := h.restaurantService.CreateRestaurantForOwner(ctx, newRestaurant)

	///response := tcp.CreateRestaurantResponse{}

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	} //else {
	response := tcp.CreateRestaurantResponse{
		Message:    "restaurant created :)",
		Restaurant: createdRestaurant,
	}
	//}

	resData, err := tcp.EncodeCreateRestaurantResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding create restaurant response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *RestaurantHandler) HandleCreateMenu(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeCreateMenuRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding create menu request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newMenu := menu.NewMenu(reqData.Name, reqData.RestaurantID)
	createdMenu, err := h.restaurantService.CreateMenuForRestaurant(ctx, newMenu)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.CreateMenuResponse{
		Message: "menu created",
		Menu:    createdMenu,
	}

	resData, err := tcp.EncodeCreateMenuResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding create menu response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *RestaurantHandler) HandleGetRestaurantMenus(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeGetRestaurantMenusRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding get menus request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	rest := restaurant.NewRestaurantByID(reqData.RestaurantID)
	fetchedMenus, err := h.restaurantService.GetAllRestaurantMenus(ctx, rest)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.GetAllMenusResponse{
		Message: "menus successfully fetched",
		Menus:   fetchedMenus,
	}

	resData, err := tcp.EncodeGetAllMenusResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding create menus response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *RestaurantHandler) HandleAddMenuItemToMenu(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAddMenuItemToMenuRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding add menu item request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newMenuItem := menu.NewMenuItem(reqData.Name, reqData.Price, reqData.PreparationTime, reqData.CancellationPenaltyPercentage, reqData.MenuID)
	createdMenuItem, err := h.restaurantService.AddMenuItemToMenu(ctx, newMenuItem)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.AddMenuItemToMenuResponse{
		Message:  "menu item successfully added",
		MenuItem: createdMenuItem,
	}

	resData, err := tcp.EncodeAddMenuItemToMenuResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding menu item response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *RestaurantHandler) HandleGetMenuItemsOfMenu(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeGetMenuItemsOfMenuRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding get menu items request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	newMenu := menu.NewMenuByID(reqData.MenuID)
	fetchedMenuItems, err := h.restaurantService.GetMenuItemsOfMenu(ctx, newMenu)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.GetMenuItemsOfMenuResponse{
		Message:   "menu items successfully fetched",
		MenuItems: fetchedMenuItems,
	}

	resData, err := tcp.EncodeGetMenuItemsOfMenuResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding menu item response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *RestaurantHandler) HandleAddOperator(ctx context.Context, conn net.Conn, req *tcp.Request) {
}

func (h *RestaurantHandler) ServeTCP(ctx context.Context, conn net.Conn, TCPReq *tcp.Request) {
	firstRoute, _ := utils.RouteSplitter(TCPReq.Location)
	switch firstRoute {
	case "":
		if TCPReq.Header["method"] == tcp.MethodPost {
			createRestaurantHandler := middleware.ApplyMiddlewares(h.HandleCreateRestaurant, middleware.AuthMiddleware)
			createRestaurantHandler(ctx, conn, TCPReq)
			return
		}
	case "menus":
		if TCPReq.Header["method"] == tcp.MethodPost {
			createMenuHandler := middleware.ApplyMiddlewares(h.HandleCreateMenu, middleware.AuthMiddleware)
			createMenuHandler(ctx, conn, TCPReq)
			return
		}
		if TCPReq.Header["method"] == tcp.MethodGet {
			getRestaurantMenus := h.HandleGetRestaurantMenus
			getRestaurantMenus(ctx, conn, TCPReq)
			return
		}
	case "menu-items":
		if TCPReq.Header["method"] == tcp.MethodPost {
			addMenuItemToMenuHandler := middleware.ApplyMiddlewares(h.HandleAddMenuItemToMenu, middleware.AuthMiddleware)
			addMenuItemToMenuHandler(ctx, conn, TCPReq)
			return
		}
		if TCPReq.Header["method"] == tcp.MethodGet {
			getMenuItemsOfMenuHandler := h.HandleGetMenuItemsOfMenu
			getMenuItemsOfMenuHandler(ctx, conn, TCPReq)
			return
		}
	case "withdraw":
		//withdraw_ownership
		fmt.Println("not implemented")

	case "operator":
		// (post, get, delete)
		if TCPReq.Header["method"] == tcp.MethodPost {
			addOperatorHandler := middleware.ApplyMiddlewares(h.HandleAddOperator, middleware.AuthMiddleware)
			addOperatorHandler(ctx, conn, TCPReq)
		}
	case "delivery":
		// add/remove delivery (post, get, delete)
		fmt.Println("not implemented")

	default:
		fmt.Println("bad request")
	}
	tcp.Error(conn, tcp.StatusMethodNotAllowed, nil, "method not allowed.")
}
