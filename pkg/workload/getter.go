package workload

import (
	"context"

	"github.com/ava-labs/kubectl-unused-volumes/pkg/api"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetAllDeployment(ctx context.Context, clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().Deployments(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return
	}
	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, deploymentWorkload(w))
	}
	return
}

func GetAllDaemonSet(ctx context.Context, clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().DaemonSets(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, daemonsetWorkload(w))
	}
	return
}

func GetAllStatefulset(ctx context.Context, clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.AppsV1().StatefulSets(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, statefulsetWorkload(w))
	}
	return
}

func GetAllJobs(ctx context.Context, clientSet *kubernetes.Clientset, namespace string) (workloads []api.Workload, err error) {
	workloadList, err := clientSet.BatchV1().Jobs(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return
	}

	workloads = []api.Workload{}
	for _, w := range workloadList.Items {
		workloads = append(workloads, jobWorkload(w))
	}
	return
}
