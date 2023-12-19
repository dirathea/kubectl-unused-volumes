package workload

import (
	"fmt"

	"github.com/ava-labs/kubectl-unused-volumes/pkg/api"
	appsV1 "k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
)

type deployment struct {
	appsV1.Deployment
}

func (d deployment) IsEmpty() bool {
	return d.Status.Replicas == 0
}

func (d deployment) GetVolumeNames() (volumeNames []string) {
	volumes := d.Spec.Template.Spec.Volumes
	for _, vol := range volumes {
		if vol.PersistentVolumeClaim != nil {
			volumeNames = append(volumeNames, vol.PersistentVolumeClaim.ClaimName)
		}
	}
	return
}

func (d deployment) GetName() string {
	return d.Name
}

func deploymentWorkload(workload appsV1.Deployment) api.Workload {
	return deployment{workload}
}

type daemonSet struct {
	appsV1.DaemonSet
}

func (d daemonSet) GetVolumeNames() (volumeNames []string) {
	volumes := d.Spec.Template.Spec.Volumes
	for _, vol := range volumes {
		if vol.PersistentVolumeClaim != nil {
			volumeNames = append(volumeNames, vol.PersistentVolumeClaim.ClaimName)
		}
	}
	return
}

func (d daemonSet) IsEmpty() bool {
	return d.Status.DesiredNumberScheduled == 0
}

func (d daemonSet) GetName() string {
	return d.Name
}

func daemonsetWorkload(workload appsV1.DaemonSet) api.Workload {
	return daemonSet{workload}
}

type statefulSet struct {
	appsV1.StatefulSet
}

func (s statefulSet) GetVolumeNames() (volumeNames []string) {
	volumes := s.Spec.Template.Spec.Volumes
	for _, vol := range volumes {
		if vol.PersistentVolumeClaim != nil {
			volumeNames = append(volumeNames, vol.PersistentVolumeClaim.ClaimName)
		}
	}

	// Statefulset has dynamic pvc object
	pvcTemplates := s.Spec.VolumeClaimTemplates
	for _, vol := range pvcTemplates {
		pvcPrefixName := fmt.Sprintf("%s-%s-", vol.GetName(), s.GetName())
		volumeNames = append(volumeNames, pvcPrefixName)
	}
	return
}

func (s statefulSet) IsEmpty() bool {
	return s.Status.Replicas == 0
}

func (s statefulSet) GetName() string {
	return s.Name
}

func statefulsetWorkload(workload appsV1.StatefulSet) api.Workload {
	return statefulSet{workload}
}

type job struct {
	batchV1.Job
}

func (j job) GetVolumeNames() (volumeNames []string) {
	volumes := j.Spec.Template.Spec.Volumes
	for _, vol := range volumes {
		if vol.PersistentVolumeClaim != nil {
			volumeNames = append(volumeNames, vol.PersistentVolumeClaim.ClaimName)
		}
	}
	return
}

func (j job) IsEmpty() bool {
	completions := int32(1)
	if j.Spec.Completions != nil {
		completions = *j.Spec.Completions
	}
	return j.Status.Succeeded == completions
}

func (j job) GetName() string {
	return j.Name
}

func jobWorkload(workload batchV1.Job) api.Workload {
	return job{workload}
}
