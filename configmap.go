package kubernetes

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetConfigmap func
func (clientset *Clientset) GetConfigmap(namespace string, labelSelector string) ([]corev1.ConfigMap, error) {
	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []corev1.ConfigMap{}, err
	}
	return configmaps.Items, nil
}

// GetConfigmapByName func
func (clientset *Clientset) GetConfigmapByName(namespace string, labelSelector string, name string) (corev1.ConfigMap, error) {
	configmaps, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var configmap corev1.ConfigMap
	if err != nil {
		return configmap, err
	}
	for _, Item := range configmaps.Items {
		if Item.Name == name {
			configmap = Item
		}
	}
	if configmap.Name == "" {
		return configmap, errors.New("not match")
	}
	return configmap, nil
}

// CreateConfigmap func
func (clientset *Clientset) CreateConfigmap(namespace string, configmap corev1.ConfigMap) (err error) {
	_, err = clientset.CoreV1().ConfigMaps(namespace).Create(&configmap)
	return
}

// UpdateConfigmap func
func (clientset *Clientset) UpdateConfigmap(namespace string, configmap corev1.ConfigMap) (err error) {
	_, err = clientset.CoreV1().ConfigMaps(namespace).Update(&configmap)
	return
}

// DeleteConfigmap func
func (clientset *Clientset) DeleteConfigmap(namespace string, name string) (err error) {
	err = clientset.CoreV1().ConfigMaps(namespace).Delete(name, &metav1.DeleteOptions{})
	return
}
