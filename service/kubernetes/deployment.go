package kubernetes

import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sCluster) ListK8sDeployments(cid, NameSpace string) (resoult []map[string]interface{}, err error) {
	// 这里加载集群中的configmap数据
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	} else {
		//deployments, err := c.CoreV1().Deployments(NameSpace).List(context.Background(), metav1.ListOptions{})
		deployments, err := c.AppsV1().Deployments(NameSpace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			global.TPLogger.Error("获取deployments失败：", err)
			return nil, errors.New("获取deployments失败")
		}
		resoult = []map[string]interface{}{}
		for _, v := range deployments.Items {
			resoult = append(resoult, map[string]interface{}{
				"name":                v.Name,
				"namespace":           v.Namespace,
				"replicas":            v.Status.Replicas,
				"availableReplicas":   v.Status.AvailableReplicas,
				"updatedReplicas":     v.Status.UpdatedReplicas,
				"unavailableReplicas": v.Status.UnavailableReplicas,
				"readyReplicas":       v.Status.ReadyReplicas,
				"observedGeneration":  v.Status.ObservedGeneration,
				"conditions":          v.Status.Conditions,
			})
		}

	}
	return resoult, nil
}
