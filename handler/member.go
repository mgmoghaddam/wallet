package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wallet/server"
	"wallet/service/member"
)

type MemberHandler struct {
	member member.UseCase
}

func NewMemberHandler(member member.UseCase) MemberHandler {
	return MemberHandler{member: member}
}

func SetupMemberRoutes(s *server.Server, h MemberHandler) {
	g := s.Engine.Group("/member")
	g.POST("", h.CreateMember)
	g.GET("/:id", h.GetMember)
	g.PUT("", h.UpdateMember)
	g.GET("/gift/:giftCode", h.GetMembersByGiftCode)
}

// CreateMember godoc
// @Summary			Create member
// @Description		Create a new member.
// @Tags			MemberDTO
// @Accept			json
// @Produce      	json
// @Param        body			body		member.CreateRequest		true	"Member create request"
// @Success      200			{object}	member.DTO
// @Failure      	400  			{object}	Error
// @Failure      	500  			{object}  	Error
// @Router       	/member		[post]
func (h MemberHandler) CreateMember(ctx *gin.Context) {
	var req member.CreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		handleError(ctx, err)
		return
	}
	result, err := h.member.Create(&req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// GetMember godoc
// @Summary      Get member
// @Description  Get a member by id.
// @Tags         MemberDTO
// @Accept       json
// @Produce      json
// @Param        id		path		int64				true	"Member id"
// @Success      200			{object}	member.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /member/{id}	[get]
func (h MemberHandler) GetMember(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		handleError(ctx, err)
		return
	}
	result, err := h.member.GetById(id)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// UpdateMember godoc
// @Summary      Update member
// @Description  Update a member by id.
// @Tags         MemberDTO
// @Accept       json
// @Produce      json
// @Param        body			body		member.DTO				true	"Member update request"
// @Success      200			{object}	member.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /member	[put]
func (h MemberHandler) UpdateMember(ctx *gin.Context) {
	var req member.DTO
	if err := ctx.ShouldBind(&req); err != nil {
		handleError(ctx, err)
		return
	}
	result, err := h.member.Update(&req)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// GetMembersByGiftCode godoc
// @Summary      Get Members by gift code
// @Description  Get Members by gift code.
// @Tags         MemberDTO
// @Accept       json
// @Produce      json
// @Param        giftCode		path		string				true	"Gift code"
// @Param        limit		query		int					false	"Limit"
// @Param        offset		query		int					false	"Offset"
// @Success      200			{object}	[]member.DTO
// @Failure      400  			{object}	Error
// @Failure      500  			{object}  	Error
// @Router       /member/gift/{giftCode}		[get]
func (h MemberHandler) GetMembersByGiftCode(ctx *gin.Context) {
	giftCode := ctx.Param("giftCode")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10000"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	result, err := h.member.GetMembersByGiftCode(giftCode, limit, offset)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}
