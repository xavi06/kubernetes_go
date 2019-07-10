package kubernetes

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPods func
func (clientset *Clientset) GetPods(namespace string, labelSelector string) ([]corev1.Pod, error) {
	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []corev1.Pod{}, err
	}
	return podList.Items, nil
}

// GetPodsByName func
func (clientset *Clientset) GetPodsByName(namespace string, labelSelector string, name string) (corev1.Pod, error) {
	podList, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	var pod corev1.Pod
	if err != nil {
		return pod, err
	}
	for _, Item := range podList.Items {
		if Item.Name == name {
			pod = Item
			break
		}
	}
	if pod.Name == "" {
		return pod, errors.New("not match")
	}
	return pod, nil
}

//DeletePods func
func (clientset *Clientset) DeletePods(namespace string, name string) (err error) {
	err = clientset.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})
	return
}

// CreatePods func
func (clientset *Clientset) CreatePods(namespace string, pod corev1.Pod) (err error) {
	_, err = clientset.CoreV1().Pods(namespace).Create(&pod)
	return
}

// ListPodContainers func
func (clientset *Clientset) ListPodContainers(namespace string, pod corev1.Pod) ([]corev1.Container, error) {
	return pod.Spec.Containers, nil
}
