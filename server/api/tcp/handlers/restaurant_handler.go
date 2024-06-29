package handlers

import (
	"context"
	"fmt"
	"net"
	"server"
	middleware "server/api/tcp/middlewares"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
	"server/internal/protocol/tcp"
	"server/pkg/utils"
	"server/services"
)

type RestaurantHandler struct {
	restaurantService services.RestaurantService
	userService       services.UserService
}

func NewRestaurantHandler(restaurantService services.RestaurantService, userService services.UserService) *RestaurantHandler {
	return &RestaurantHandler{
		restaurantService: restaurantService,
		userService:       userService,
	}
}

func (h *RestaurantHandler) HandleCreateRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeCreateRestaurantRequest(req.Data)
	if err != nil {
		//logger
		fmt.Println("Error decoding create restaurant request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	reqData.Address.Types = server.RestaurantAddressType
	newRestaurant := restaurant.NewRestaurant(reqData.Name, reqData.Phone, *reqData.Address)
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

func (h *RestaurantHandler) HandleGetOwnerRestaurants(ctx context.Context, conn net.Conn, req *tcp.Request) {
	restaurants, err := h.restaurantService.GetRestaurantsOfAnOwner(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.GetOwnerOperatorRestaurantsResponse{
		Message:     "user wallet cards successfuly fetched.",
		Restaurants: restaurants,
	}
	//}
	resData, err := tcp.EncodeGetOwnerOperatorRestaurantsResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *RestaurantHandler) HandleEditRestaurantName(ctx context.Context, conn net.Conn, req *tcp.Request) {

	reqData, err := tcp.DecodeEditRestaurantNameRequest(req.Data)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	isOwner, err := h.restaurantService.IsRestaurantOwner(ctx, reqData.RestaurantID)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	if !isOwner {
		tcp.Error(conn, tcp.StatusForbidden, nil, "owner not found")
		return
	}

	err = h.restaurantService.EditRestaurantName(ctx, reqData.RestaurantID, reqData.NewName)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	// ** need implementaion
}

func (h *RestaurantHandler) HandleGetOperatorRestaurants(ctx context.Context, conn net.Conn, req *tcp.Request) {
	restaurants, err := h.restaurantService.GetRestaurantsOfAnOperator(ctx)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.GetOwnerOperatorRestaurantsResponse{
		Message:     "user wallet cards successfuly fetched.",
		Restaurants: restaurants,
	}
	//}
	resData, err := tcp.EncodeGetOwnerOperatorRestaurantsResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding get cards response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
}

func (h *RestaurantHandler) HandleAddOperatorToRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAddOperatorToRestaurantRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	// Is the owner of restaurant the requester?
	isOwner, err := h.restaurantService.IsRestaurantOwner(ctx, reqData.RestaurantID)

	if err != nil {
		fmt.Println("Error encoding create restaurant response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	if !isOwner {
		tcp.Error(conn, tcp.StatusForbidden, nil, "owner not found")
		return
	}

	// getuser
	introducedOperatorPhoneOrEmail := reqData.OperatorPhoneOrEmail
	introducedOperator, err := h.userService.GetUserByEmailOrPhone(ctx, introducedOperatorPhoneOrEmail)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	restaurantId := reqData.RestaurantID
	// getrestarant
	restaurant, err := h.restaurantService.GetRestaurantByID(ctx, restaurantId)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	// assign
	_, err = h.restaurantService.AssignOperatorToRestaurant(ctx, introducedOperator, *restaurant)

	if err != nil {
		new_err := fmt.Errorf("failed to assign operator to the restaurant %s", restaurant.Name)
		tcp.Error(conn, tcp.StatusBadRequest, nil, new_err.Error())
		return
	}
	assignOperatorResponse := tcp.AssignOperatorResponse{OperatorPhoneOrEmail: introducedOperatorPhoneOrEmail,
		RestaurantName: restaurant.Name}
	response := tcp.AssignOperatorToRestaurantResponse{
		Message:                fmt.Sprintf("operator %s card added to %s restaurant", introducedOperatorPhoneOrEmail, restaurant.Name),
		AssignOperatorResponse: &assignOperatorResponse}

	resData, err := tcp.EncodeAssignOperatorResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusCreated, nil, resData)
}

func (h *RestaurantHandler) HandleGetAllOperatorsOfRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {

}

func (h *RestaurantHandler) HandleRemoveOperatorFromRestaurant(ctx context.Context, conn net.Conn, req *tcp.Request) {
	reqData, err := tcp.DecodeAddOperatorToRestaurantRequest(req.Data)
	if err != nil {
		//logger.Error("Error decoding register request:", err)
		fmt.Println("Error decoding register request:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	// Is the owner of restaurant the requester?
	isOwner, err := h.restaurantService.IsRestaurantOwner(ctx, reqData.RestaurantID)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	if !isOwner {
		tcp.Error(conn, tcp.StatusForbidden, nil, "owner not found")
		return
	}

	// getUser
	introducedOperatorPhoneOrEmail := reqData.OperatorPhoneOrEmail
	introducedOperator, err := h.userService.GetUserByEmailOrPhone(ctx, introducedOperatorPhoneOrEmail)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}

	restaurantId := reqData.RestaurantID
	// getRestaurant
	restaurant, err := h.restaurantService.GetRestaurantByID(ctx, restaurantId)
	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	_, err = h.restaurantService.AssignOperatorToRestaurant(ctx, introducedOperator, *restaurant)

	if err != nil {
		new_err := fmt.Errorf("failed to assign operator to the restaurant %s", restaurant.Name)
		tcp.Error(conn, tcp.StatusBadRequest, nil, new_err.Error())
		return
	}
	assignOperatorResponse := tcp.AssignOperatorResponse{OperatorPhoneOrEmail: introducedOperatorPhoneOrEmail,
		RestaurantName: restaurant.Name}
	response := tcp.AssignOperatorToRestaurantResponse{
		Message:                fmt.Sprintf("operator %s card added to %s restaurant", introducedOperatorPhoneOrEmail, restaurant.Name),
		AssignOperatorResponse: &assignOperatorResponse}

	resData, err := tcp.EncodeAssignOperatorResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding add to card response:", err)
		tcp.Error(conn, tcp.StatusInternalServerError, nil, err.Error())
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
	newMenuItem := menu.NewMenuItem(reqData.Name, reqData.Price, reqData.PreparationMinutes, reqData.CancellationPenaltyPercentage, reqData.MenuID)
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
func (h *RestaurantHandler)GetRestaurantsToAddCategoryMenuFood(ctx context.Context,conn net.Conn,req *tcp.Request){

	fetchedRestaurants, err := h.restaurantService.GetRestaurantsToAddCategoryMenuFood(ctx)

	if err != nil {
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	response := tcp.RestaurantToAddCategoryMenuFoodResponse{
		Message:   "restaurants items successfully fetched",
		Restaurants:fetchedRestaurants,
	}

	resData, err := tcp.EncodeGetRestaurantToAddCategoryMenuFoodResponse(response)
	if err != nil {
		//logger.Error("Error encoding register response:", err)
		fmt.Println("Error encoding menu item response:", err)
		tcp.Error(conn, tcp.StatusBadRequest, nil, err.Error())
		return
	}
	tcp.SendResponse(conn, tcp.StatusOK, nil, resData)
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
		if TCPReq.Header["method"] == tcp.MethodGet {
			createRestaurantHandler := middleware.ApplyMiddlewares(h.GetRestaurantsToAddCategoryMenuFood,middleware.AuthMiddleware)
			createRestaurantHandler(ctx,conn,TCPReq)
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
			addOperatorHandler := middleware.ApplyMiddlewares(h.HandleAddOperatorToRestaurant, middleware.AuthMiddleware)
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
