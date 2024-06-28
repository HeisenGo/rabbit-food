package server

type contextKey string

const UserIDKey contextKey = "userID"

type RoleType string

const Owner RoleType = "owner"
const Operator RoleType = "operator"
const UserAddressType = "user"
const RestaurantAddressType = "restaurant"