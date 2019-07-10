package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// GetClientset func
func GetClientset(kubeconfigFile string) (*Clientset, error) {
	var kubeconfig *string
	kubeconfig = &kubeconfigFile
	config, _ := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return &Clientset{clientset}, err

}
