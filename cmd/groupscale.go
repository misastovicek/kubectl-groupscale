package groupscale

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type deploymentRelease struct {
	name      string
	namespace string
}

func (d *deploymentRelease) scaleDeployment(clientset *kubernetes.Clientset, scaleTo int32) {
	scale, err := clientset.AppsV1().Deployments(d.namespace).GetScale(
		context.TODO(), d.name, metav1.GetOptions{})
	genericErrorHandler(err)
	scale.Spec.Replicas = scaleTo
	_, err = clientset.AppsV1().Deployments(d.namespace).UpdateScale(
		context.TODO(), d.name, scale, metav1.UpdateOptions{})
	genericErrorHandler(err)
}

func genericErrorHandler(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func scaleLabelExists(mResourceLables map[string]string, labelKey string, labelValue string) bool {
	if val, ok := mResourceLables[labelKey]; ok && val == labelValue {
		return true
	}
	return false
}

// GroupScale scales Kubernetes deployment based on label key/value pair to number of replicas
func GroupScale(kubeconfigPath *string, labelKey string, labelValue string, replicas int) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	genericErrorHandler(err)

	clientset, err := kubernetes.NewForConfig(config)
	genericErrorHandler(err)

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	genericErrorHandler(err)

	var deploymentsToScale []deploymentRelease

	for _, deployment := range deployments.Items {
		if scaleLabelExists(deployment.Labels, labelKey, labelValue) {
			deploymentsToScale = append(deploymentsToScale, deploymentRelease{
				name:      deployment.Name,
				namespace: deployment.Namespace,
			})
		}
	}
	for _, deployment := range deploymentsToScale {
		deployment.scaleDeployment(clientset, int32(replicas))
	}
}
