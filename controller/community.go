package controller

import (
	"bluebell01/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// ---- 跟社区相关的 ----

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id" )
	communityID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c,CodeInvalidParam)
		return
	}
	// 2. 根据id获取社区详情
	data, err := logic.GetCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}