package kubernetes

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// GetNodes func
func (clientset *Clientset) GetNodes(labelSelector string) ([]Node, error) {
	nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return []Node{}, err
	}
	var data []Node
	for _, Item := range nodeList.Items {
		Name := Item.Name
		IP := Item.Annotations["flannel.alpha.coreos.com/public-ip"]
		Status := Item.Status.Conditions[3].Status
		CPUByte, err := Item.Status.Allocatable["cpu"].MarshalJSON()
		if err != nil {
			continue
		}
		// byte to string
		CPU := string(CPUByte[:])
		MemoryByte, err := Item.Status.Allocatable["memory"].MarshalJSON()
		if err != nil {
			continue
		}
		Memory := string(MemoryByte[:])
		Labels := Item.Labels
		Annotations := Item.Annotations
		CreationTimestamp := Item.CreationTimestamp
		Address := Item.Status.Addresses
		PodCIDR := Item.Spec.PodCIDR
		Unschedulable := Item.Spec.Unschedulable
		NodeInfo := Item.Status.NodeInfo
		Conditions := Item.Status.Conditions
		Allocatable := Item.Status.Allocatable
		Capacity := Item.Status.Capacity
		Node := Node{Name, IP, Status, CPU, Memory, Labels, Annotations, CreationTimestamp,
			Address, PodCIDR, Unschedulable, NodeInfo, Conditions, Allocatable, Capacity}
		data = append(data, Node)
	}
	return data, nil
}

// GetNodesByName func
func (clientset *Clientset) GetNodesByName(labelSelector string, name string) (Node, error) {
	nodeList, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: labelSelector})
	var node Node
	var item corev1.Node
	if err != nil {
		return node, err
	}
	for _, Item := range nodeList.Items {
		if Item.Name == name {
			item = Item
			break
		}
	}
	if item.Name == "" {
		return node, errors.New("not match")
	}
	node.Name = item.Name
	node.IP = item.Annotations["flannel.alpha.coreos.com/public-ip"]
	node.Status = item.Status.Conditions[3].Status
	CPUByte, _ := item.Status.Allocatable["cpu"].MarshalJSON()
	// byte to string
	node.CPU = string(CPUByte[:])
	MemoryByte, _ := item.Status.Allocatable["memory"].MarshalJSON()
	node.Memory = string(MemoryByte[:])
	node.Labels = item.Labels
	node.Annotations = item.Annotations
	node.CreationTimestamp = item.CreationTimestamp
	node.Address = item.Status.Addresses
	node.PodCIDR = item.Spec.PodCIDR
	node.Unschedulable = item.Spec.Unschedulable
	node.NodeInfo = item.Status.NodeInfo
	node.Conditions = item.Status.Conditions
	node.Allocatable = item.Status.Allocatable
	node.Capacity = item.Status.Capacity
	return node, nil
}
