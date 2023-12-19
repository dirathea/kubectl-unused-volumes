package plugin

import (
	"context"

	"github.com/ava-labs/kubectl-unused-volumes/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetVolumes(ctx context.Context, clientSet *kubernetes.Clientset, namespace string) (volumes []*api.Volume, err error) {
	list, err := clientSet.CoreV1().PersistentVolumeClaims(namespace).List(ctx, metaV1.ListOptions{})
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
