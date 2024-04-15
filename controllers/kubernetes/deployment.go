package kubernetes

import (
	"genbu/common/global"
	"genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

// ListK8sDeployments 获取集群中的deployment列表
func ListK8sDeploymentMap(ctx *gin.Context) {
	cid := ctx.Param("cid")
	namespace := ctx.DefaultQuery("namespace", "default")
	res, err := kubernetes.NewK8sInterface().ListK8sDeployments(cid, namespace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", res)
}
