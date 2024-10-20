package userGetContent

import (
	"go-admin/constant"
	"go-admin/model"
	"go-admin/tools"

	"github.com/gin-gonic/gin"
)

func Controllers(g *gin.Context) {

	json, _ := g.Get("json")

	// 执行base64解码
	content, err := tools.ToolsBase64Decode(json.(model.ModelData).Content)
	if err != nil {
		g.String(constant.CODE500, "数据无法正常解析")
		return
	}

	g.String(constant.CODE200, content)
}
