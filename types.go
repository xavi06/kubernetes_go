package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Clientset struct
type Clientset struct {
	*kubernetes.Clientset
}

// Node struct
type Node struct {
	Name              string
	IP                string
	Status            corev1.ConditionStatus
	CPU               string
	Memory            string
	Labels            map[string]string
	Annotations       map[string]string
	CreationTimestamp metav1.Time
	Address           []corev1.NodeAddress
	PodCIDR           string
	Unschedulable     bool
	NodeInfo          corev1.NodeSystemInfo
	Conditions        []corev1.NodeCondition
	Allocatable       corev1.ResourceList
	Capacity          corev1.ResourceList
}
