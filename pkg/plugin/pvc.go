package plugin

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllPvc(clientSet *kubernetes.Clientset, namespace string) (map[string]v1.PersistentVolumeClaim, error) {
	list, err := clientSet.CoreV1().PersistentVolumeClaims(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return map[string]v1.PersistentVolumeClaim{}, err
	}
	result := map[string]v1.PersistentVolumeClaim{}
	for _, pvc := range list.Items {
		result[pvc.Name] = pvc
	}
	return result, err
}
