package kubernetes

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetCronJob func
func (clientset *Clientset) GetCronJob(namespace string, labelSelector string) ([]batchv1beta1.CronJob, error) {
	cronjobs, err := clientset.BatchV1beta1().CronJobs(namespace).List(metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return nil, err
	}
	return cronjobs.Items, nil
}
