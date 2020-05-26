package plugin

import (
	"github.com/dirathea/kubectl-unused-volumes/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetVolumes(clientSet *kubernetes.Clientset, namespace string) (volumes []*api.Volume, err error) {
	list, err := clientSet.CoreV1().PersistentVolumeClaims(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return
	}
	for _, pvc := range list.Items {
		volumes = append(volumes, &api.Volume{
			PersistentVolumeClaim: pvc,
			Reason:                api.NO_RESOURCE_REFFERENCE,
		})
	}
	return
}
