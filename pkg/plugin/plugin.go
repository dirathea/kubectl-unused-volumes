package plugin

import (
	"strings"
	"sync"

	"github.com/dirathea/kubectl-volume-reclaim/pkg/api"
	"github.com/dirathea/kubectl-volume-reclaim/pkg/workload"
	"github.com/gosuri/uitable"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func RunPlugin(configFlags *genericclioptions.ConfigFlags) (string, error) {
	// log := logger.NewLogger()
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return "", errors.Wrap(err, "failed to read kubeconfig")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", errors.Wrap(err, "failed to create clientset")
	}

	// get all pvcs on namespace
	pvcList, err := GetAllPvc(clientset, "default")

	// namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return "", errors.Wrap(err, "failed to get all pvc in namespaces")
	}

	var wg sync.WaitGroup

	allResources := []func(client *kubernetes.Clientset, namespace string) ([]api.Workload, error){
		workload.GetAllDaemonSet,
		workload.GetAllDeployment,
		workload.GetAllJobs,
		workload.GetAllStatefulset,
	}

	workloads := []api.Workload{}

	for _, f := range allResources {
		wg.Add(1)
		go func(getWorkloadFunc func(client *kubernetes.Clientset, namespace string) ([]api.Workload, error)) {
			lists, err := getWorkloadFunc(clientset, v1.NamespaceDefault)
			if err != nil {
				return
			}
			workloads = append(workloads, lists...)
			wg.Done()
		}(f)
	}

	wg.Wait()

	for _, wk := range workloads {
		if !wk.IsEmpty() {
			removeVolume(pvcList, wk)
		} else {
			markVolumeAsZeroReplica(pvcList, wk)
		}
	}

	table := uitable.New()
	table.AddRow("Name", "Volume Name", "Size", "Reason", "Used By")

	for _, p := range pvcList {
		if p != nil {
			storageSize := p.Spec.Resources.Requests[v1.ResourceStorage]
			table.AddRow(p.Name, p.Spec.VolumeName, storageSize.String(), p.Reason, workload.Join(p.Workloads, ","))
		}
	}

	return table.String(), nil
}

func removeVolume(volumes []*api.Volume, workload api.Workload) {
	for _, claim := range workload.GetVolumeNames() {
		for idx, vol := range volumes {
			if vol != nil {
				if strings.HasPrefix(vol.GetName(), claim) {
					volumes[idx] = nil
				}
			}
		}
	}
}

func markVolumeAsZeroReplica(volumes []*api.Volume, workload api.Workload) {
	for _, claim := range workload.GetVolumeNames() {
		for idx, vol := range volumes {
			if vol != nil {
				if strings.HasPrefix(vol.GetName(), claim) {
					volumes[idx].Reason = api.WORKLOAD_HAS_ZERO_REPLICAS
					volumes[idx].Workloads = append(volumes[idx].Workloads, workload)
				}
			}
		}
	}
}
