package system

import (
	response "backend/model/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContentApi struct{}

func (contentApi *ContentApi) AddContent(c *gin.Context) {
	fmt.Println("Adding content")
}

func (contentApi *ContentApi) ListContent(context *gin.Context) {
	fmt.Println("List content")

	response.Success(context, "asdas")

}

func (contentApi *ContentApi) GetContent(context *gin.Context) {

}

func (contentApi *ContentApi) UpdateContent(context *gin.Context) {

}
