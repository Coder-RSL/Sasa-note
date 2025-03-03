package response

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	RequestID string      `json:"requestId"` // 链路追踪标识[5](@ref)
	Code      int         `json:"code"`      // 业务状态码[9](@ref)
	Status    string      `json:"status"`    // 成功/失败状态
	Message   string      `json:"message"`   // 动态提示信息
	Data      interface{} `json:"data"`      // 业务数据主体
}

// 分页扩展结构
type PaginatedResponse struct {
	BaseResponse
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
}

// 成功响应构造器
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, BaseResponse{
		RequestID: c.GetString("requestId"),
		Code:      200,
		Status:    "success",
		Data:      data,
	})
}

// 分页响应专用方法
func PageSuccess(c *gin.Context, data interface{}, total int) {
	c.JSON(200, PaginatedResponse{
		BaseResponse: BaseResponse{},
		Total:        total,
	})
}

// 错误响应模板化
var ErrMap = map[int]string{
	400: "参数校验失败",
	500: "系统繁忙",
}

func Error(c *gin.Context, code int) {
	c.AbortWithStatusJSON(code, BaseResponse{
		Code:    code,
		Message: ErrMap[code],
	})
}
