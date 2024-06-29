package handlers

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/mock"
	"server/internal/models/user"
	"server/internal/protocol/tcp"
)

// MockAuthService is a mock implementation of the AuthService interface
type MockAuthService struct {
	mock.Mock
}

var email = "test@example.com"

func (m *MockAuthService) CreateUser(ctx context.Context, newUser *user.User) (string, string, error) {
	args := m.Called(ctx, newUser)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockAuthService) LoginUser(ctx context.Context, phoneOrEmail, password string) (string, error) {
	args := m.Called(ctx, phoneOrEmail, password)
	return args.String(0), args.Error(1)
}

func TestHandleRegister(t *testing.T) {
	mockAuthService := new(MockAuthService)
	handler := NewAuthHandler(mockAuthService)

	ctx := context.Background()
	conn := new(net.Conn) // Use a mock connection if necessary
	reqData := tcp.RegisterRequest{Phone: "1234567890", Email: &email, Password: "password"}
	req := &tcp.Request{Data: tcp.EncodeRegisterRequest(reqData)}

	mockAuthService.On("CreateUser", ctx, mock.AnythingOfType("*user.User")).Return("userID", "token", nil)

	handler.HandleRegister(ctx, *conn, req)

	// Add assertions to verify the behavior
	mockAuthService.AssertExpectations(t)
}

func TestHandleLogin(t *testing.T) {
	mockAuthService := new(MockAuthService)
	handler := NewAuthHandler(mockAuthService)

	ctx := context.Background()
	conn := new(net.Conn) // Use a mock connection if necessary
	reqData := tcp.LoginRequest{PhoneOrEmail: "test@example.com", Password: "password"}
	req := &tcp.Request{Data: tcp.EncodeLoginRequest(reqData)}

	mockAuthService.On("LoginUser", ctx, "test@example.com", "password").Return("token", nil)

	handler.HandleLogin(ctx, *conn, req)

	// Add assertions to verify the behavior
	mockAuthService.AssertExpectations(t)
}

func TestServeTCP(t *testing.T) {
	mockAuthService := new(MockAuthService)
	handler := NewAuthHandler(mockAuthService)

	ctx := context.Background()
	conn := new(net.Conn) // Use a mock connection if necessary

	// Test Register
	reqData := tcp.RegisterRequest{Phone: "1234567890", Email: &email, Password: "password"}
	req := &tcp.Request{Header: map[string]string{"method": tcp.MethodPost}, Location: "register", Data: tcp.EncodeRegisterRequest(reqData)}

	mockAuthService.On("CreateUser", ctx, mock.AnythingOfType("*user.User")).Return("userID", "token", nil)

	handler.ServeTCP(ctx, *conn, req)

	// Test Login
	reqDataLogin := tcp.LoginRequest{PhoneOrEmail: "test@example.com", Password: "password"}
	reqLogin := &tcp.Request{Header: map[string]string{"method": tcp.MethodPost}, Location: "login", Data: tcp.EncodeLoginRequest(reqDataLogin)}

	mockAuthService.On("LoginUser", ctx, "test@example.com", "password").Return("token", nil)

	handler.ServeTCP(ctx, *conn, reqLogin)

	// Add assertions to verify the behavior
	mockAuthService.AssertExpectations(t)
}
