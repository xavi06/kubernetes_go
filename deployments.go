package kubernetes

import (
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetDeployments func
func (clientset *Clientset) GetDeployments(namespace string, labelSelector string) ([]appv1.Deployment, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []appv1.Deployment{}, err
	}
	return deployments.Items, nil

}
