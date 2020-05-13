package plugin

import (
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func RunPlugin(configFlags *genericclioptions.ConfigFlags, outputCh chan string) error {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read kubeconfig")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "failed to create clientset")
	}

	// get all pvcs on namespace
	pvcList, err := GetAllPvc(clientset, "default")

	// namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to get all pvc in namespaces")
	}

	podList, err := GetAllPod(clientset, "default")
	if err != nil {
		return errors.Wrap(err, "failed to get all pod in namespaces")
	}

	allClaimsFromPod := []string{}

	for _, pod := range podList {
		allClaimsFromPod = append(allClaimsFromPod, getPVCNameFromPod(pod)...)
	}

	for _, claims := range allClaimsFromPod {
		delete(pvcList, claims)
	}

	table := uitable.New()
	table.AddRow("Name", "Volume Name", "Size", "Storage Class")

	for _, p := range pvcList {
		table.AddRow(p.Name, p.Spec.VolumeName, p.Spec.Resources.Requests[v1.ResourceStorage], p.Spec.StorageClassName)
	}

	fmt.Println(table.String())
	outputCh <- table.String()

	return nil
}
