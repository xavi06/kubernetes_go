package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaces func
func (clientset *Clientset) GetNamespaces(labelSelector string) ([]corev1.Namespace, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []corev1.Namespace{}, err
	}
	return namespaces.Items, nil
}
