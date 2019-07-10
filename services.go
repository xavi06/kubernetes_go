package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetServices func
func (clientset *Clientset) GetServices(namespace string, labelSelector string) ([]corev1.Service, error) {
	services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []corev1.Service{}, err
	}
	return services.Items, nil
}
