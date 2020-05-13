package plugin

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllPod(clientSet *kubernetes.Clientset, namespace string) ([]v1.Pod, error) {
	list, err := clientSet.CoreV1().Pods(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return []v1.Pod{}, err
	}
	return list.Items, err
}

func getPVCNameFromPod(pod v1.Pod) []string {
	allVolumes := []string{}
	for _, vol := range pod.Spec.Volumes {
		if vol.PersistentVolumeClaim != nil {
			allVolumes = append(allVolumes, vol.PersistentVolumeClaim.ClaimName)
		}
	}

	return allVolumes
}
