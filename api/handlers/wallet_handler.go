package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"net/http"
	"reservation-api/internal/models"
	"reservation-api/internal/services/domain_services"
	"strconv"
)

type WalletHandler struct {
	service *domain_services.WalletService
	handlerBase
}

func (handler *WalletHandler) Register(service *domain_services.WalletService) {
	handler.service = service
}

type createWalletRequest struct {
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

type createWalletResponse struct {
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

// @Summary Create a new wallet
// @Description Create a new wallet with the specified owner name and balance
// @Tags Wallet
// @Accept json
// @Produce json
// @Param request body createWalletRequest true "Wallet creation request"
// @Success 201 {object} createWalletResponse
// @Router /wallets [post]
func (h *WalletHandler) CreateWallet(c echo.Context) error {
	var req createWalletRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	wallet := &models.Wallet{
		UserId:  req.UserId,
		Balance: req.Balance,
	}
	if _, err := h.service.Create(c.Request().Context(), wallet); err != nil {
		return err
	}
	res := createWalletResponse{
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}
	return c.JSON(http.StatusCreated, res)
}

type getWalletResponse struct {
	ID      uint64          `json:"id"`
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

// @Summary Get a wallet by ID
// @Description Get a wallet with the specified ID
// @Tags Wallet
// @Produce json
// @Param id path uint true "Wallet ID"
// @Success 200 {object} getWalletResponse
// @Router /wallets/{id} [get]
func (h *WalletHandler) GetWallet(c echo.Context) error {
	walletID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid wallet ID")
	}
	wallet, err := h.service.GetWalletByID(c.Request().Context(), walletID)
	if err != nil {
		return err
	}
	res := getWalletResponse{
		ID:      wallet.Id,
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}
	return c.JSON(http.StatusOK, res)
}

type depositRequest struct {
	Amount decimal.Decimal `json:"amount"`
}

type depositResponse struct {
	ID      uint64          `json:"id"`
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

// @Summary Deposit funds into a wallet
// @Description Deposit the specified amount into the wallet with the specified ID
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id path uint true "Wallet ID"
// @Param request body depositRequest true "Deposit request"
// @Success 200 {object} depositResponse
// @Router /wallets/{id}/deposit [post]
func (h *WalletHandler) Deposit(c echo.Context) error {
	walletID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid wallet ID")
	}
	var req depositRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	amount := req.Amount
	wallet, err := h.service.Deposit(c.Request().Context(), walletID, amount)
	if err != nil {
		return err
	}
	res := depositResponse{
		ID:      wallet.Id,
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}
	return c.JSON(http.StatusOK, res)
}

type withdrawRequest struct {
	Amount decimal.Decimal `json:"amount"`
}

type withdrawResponse struct {
	ID      uint64          `json:"id"`
	UserId  uint64          `json:"user_id"`
	Balance decimal.Decimal `json:"balance"`
}

// @Summary Withdraw funds from a wallet
// @Description Withdraw the specified amount from the wallet with the specified ID
// @Tags Wallet
// @Accept json
// @Produce json
// @Param id path uint true "Wallet ID"
// @Param request body withdrawRequest true "Withdrawal request"
// @Success 200 {object} withdrawResponse
// @Router /wallets/{id}/withdraw [post]
func (h *WalletHandler) Withdraw(c echo.Context) error {
	walletID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid wallet ID")
	}
	var req withdrawRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	amount := req.Amount
	wallet, err := h.service.Withdraw(c.Request().Context(), walletID, amount)
	if err != nil {
		return err
	}
	res := withdrawResponse{
		ID:      wallet.Id,
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}
	return c.JSON(http.StatusOK, res)
}

func registerRoutes(handler *WalletHandler) {

	routeGroup := handler.Router.Group("/wallets")
	routeGroup.POST("", handler.CreateWallet)
	routeGroup.GET("/:id", handler.GetWallet)
	routeGroup.POST("/:id/deposit", handler.Deposit)
	routeGroup.POST("/:id/withdraw", handler.Withdraw)
}
