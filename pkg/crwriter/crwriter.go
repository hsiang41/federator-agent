package crwriter

import (
	"context"
	"errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	K8SErrors                               "k8s.io/apimachinery/pkg/api/errors"
	logUtil                                 "github.com/containers-ai/alameda/pkg/utils/log"
	OperatorAPIsAutoScalingV1Alpha1         "github.com/containers-ai/alameda/operator/pkg/apis/autoscaling/v1alpha1"
	OperatorReconcilerAlamedaRecommendation "github.com/containers-ai/alameda/operator/pkg/reconciler/alamedarecommendation"
	DatahubV1Alpha1                         "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	OperatorAPIs                            "github.com/containers-ai/alameda/operator/pkg/apis"
	"k8s.io/apimachinery/pkg/types"
	"fmt"
)

type CrWriter struct {
	K8sClient client.Client
	Scope     *logUtil.Scope
}

func NewCrWriter(scope *logUtil.Scope) (*CrWriter, error) {
	var (
		err error
		k8sCli client.Client
	)

	k8sClientConfig, err := config.GetConfig()
	if err != nil {
		return nil, errors.New("Get kubernetes configuration failed: " + err.Error())
	}

	if k8sCli, err = client.New(k8sClientConfig, client.Options{}); err != nil {
		return nil, errors.New("Create kubernetes client failed: " + err.Error())
	}

	mgr, err := manager.New(k8sClientConfig, manager.Options{})
	if err != nil {
		scope.Error(err.Error())
	}

	if err := OperatorAPIs.AddToScheme(mgr.GetScheme()); err != nil {
		scope.Error(err.Error())
	}

	return &CrWriter{k8sCli, scope}, nil
}

func (c *CrWriter) CreatePodRecommendations(ctx context.Context, in []*DatahubV1Alpha1.PodRecommendation) {
	for _, podRecommendation := range in {
		podNS := podRecommendation.GetNamespacedName().Namespace
		podName := podRecommendation.GetNamespacedName().Name
		// alamedaRecommendation := &OperatorAPIsAutoScalingV1Alpha1.AlamedaRecommendation{}
		alamedaRecommendation := &OperatorAPIsAutoScalingV1Alpha1.AlamedaRecommendation{}

		c.Scope.Debugf(fmt.Sprintf("CreatePodRecommendations: namespaces: %s, name: %s", podNS, podName))
		err := c.K8sClient.Get(context.TODO(), types.NamespacedName{
			Namespace: podNS,
			Name:      podName,
		}, alamedaRecommendation)
		if err != nil {
			c.Scope.Errorf(fmt.Sprintf("Failed to CreatePodRecommendations: %v", err))
		}
		if err == nil {
			alamedarecommendationReconciler := OperatorReconcilerAlamedaRecommendation.NewReconciler(c.K8sClient, alamedaRecommendation)
			if alamedaRecommendation, err = alamedarecommendationReconciler.UpdateResourceRecommendation(podRecommendation); err == nil {
				if err = c.K8sClient.Update(context.TODO(), alamedaRecommendation); err != nil {
					c.Scope.Error(err.Error())
				} else {
					c.Scope.Debugf(fmt.Sprintf("Succeed to CreatePodRecommendations %s %s, recommendation: %v", podNS, podName, podRecommendation))
				}
			} else {
				c.Scope.Errorf(fmt.Sprintf("Failed to update resource recommendation: %v", err))
			}
		} else if !K8SErrors.IsNotFound(err) {
			c.Scope.Error(err.Error())
		}
	}
}
