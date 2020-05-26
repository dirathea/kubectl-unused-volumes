package workload

import (
	"github.com/dirathea/kubectl-unused-volumes/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllDeployment(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().Deployments(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return
	}
	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, deploymentWorkload(w))
	}
	return
}

func GetAllDaemonSet(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().DaemonSets(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, daemonsetWorkload(w))
	}
	return
}

func GetAllStatefulset(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().StatefulSets(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, statefulsetWorkload(w))
	}
	return
}

func GetAllJobs(clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.BatchV1().Jobs(namespace).List(metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, jobWorkload(w))
	}
	return
}
