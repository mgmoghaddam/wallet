package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wallet/server"
	"wallet/service/wallet"
)

type WalletHandler struct {
	wallet wallet.UseCase
}

func NewWalletHandler(wallet wallet.UseCase) WalletHandler {
	return WalletHandler{wallet: wallet}
}

func SetupWalletRoutes(s *server.Server, h WalletHandler) {
	g := s.Engine.Group("/wallet")
	g.POST("", h.CreateWallet)
	g.GET("/:walletId", h.GetWallet)
	g.GET("/member/:userId", h.GetWallets)
	g.POST("/gift", h.AddGift)
}

// CreateWallet godoc
// @Summary			Create wallet
// @Description		Create a new wallet.
// @Tags			WalletDTO
// @Accept			json
// @Produce      	json
// @Param        body			body		wallet.CreateRequest		true	"Wallet create request"
// @Success      200			{object}	wallet.DTO
// @Failure      	400  			{object}	Error
// @Failure      	500  			{object}  	Error
// @Router       	/wallet		[post]
func (h WalletHandler) CreateWallet(ctx *gin.Context) {
	var req wallet.CreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		handleError(ctx, err)
		return
	}
	result, err := h.wallet.Create(&req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// GetWallet godoc
// @Summary      Get wallet
// @Description  Get a wallet by id.
// @Tags         WalletDTO
// @Accept       json
// @Produce      json
// @Param        walletId		path		string				true	"Wallet id"
// @Success      200			{object}	wallet.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /wallet/{walletId}	[get]
func (h WalletHandler) GetWallet(ctx *gin.Context) {
	walletId, err := strconv.ParseInt(ctx.Param("walletId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return

	}
	result, err := h.wallet.GetByID(walletId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// GetWallets godoc
// @Summary      Get wallets
// @Description  Get all wallets.
// @Tags         WalletDTO
// @Accept       json
// @Produce      json
// @Success      200			{object}	[]wallet.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /wallets/{userId}		[get]
func (h WalletHandler) GetWallets(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return

	}
	result, err := h.wallet.GetByMemberID(userId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// AddGift godoc
// @Summary      Add gift
// @Description  Add a gift code to wallet.
// @Tags         WalletDTO
// @Accept       json
// @Produce      json
// @Param        body			body		wallet.AddGiftRequest		true	"Add gift request"
// @Success      200			{object}	wallet.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       	/wallet/gift		[post]
func (h WalletHandler) AddGift(ctx *gin.Context) {
	var req wallet.AddGiftRequest
	if err := ctx.ShouldBind(&req); err != nil {
		handleError(ctx, err)
		return
	}
	result, err := h.wallet.AddGift(&req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
