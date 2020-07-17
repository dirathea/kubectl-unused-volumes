package workload

import (
	"testing"

	appsV1 "k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
)

func Test_deployment_IsEmpty(t *testing.T) {
	type fields struct {
		Deployment appsV1.Deployment
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty deployment",
			fields: fields{
				Deployment: appsV1.Deployment{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := deployment{
				Deployment: tt.fields.Deployment,
			}
			if got := d.IsEmpty(); got != tt.want {
				t.Errorf("deployment.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_deployment_GetVolumeNames(t *testing.T) {
// 	type fields struct {
// 		Deployment appsV1.Deployment
// 	}
// 	tests := []struct {
// 		name            string
// 		fields          fields
// 		wantVolumeNames []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := deployment{
// 				Deployment: tt.fields.Deployment,
// 			}
// 			if gotVolumeNames := d.GetVolumeNames(); !reflect.DeepEqual(gotVolumeNames, tt.wantVolumeNames) {
// 				t.Errorf("deployment.GetVolumeNames() = %v, want %v", gotVolumeNames, tt.wantVolumeNames)
// 			}
// 		})
// 	}
// }

// func Test_daemonSet_GetVolumeNames(t *testing.T) {
// 	type fields struct {
// 		DaemonSet appsV1.DaemonSet
// 	}
// 	tests := []struct {
// 		name            string
// 		fields          fields
// 		wantVolumeNames []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := daemonSet{
// 				DaemonSet: tt.fields.DaemonSet,
// 			}
// 			if gotVolumeNames := d.GetVolumeNames(); !reflect.DeepEqual(gotVolumeNames, tt.wantVolumeNames) {
// 				t.Errorf("daemonSet.GetVolumeNames() = %v, want %v", gotVolumeNames, tt.wantVolumeNames)
// 			}
// 		})
// 	}
// }

func Test_daemonSet_IsEmpty(t *testing.T) {
	type fields struct {
		DaemonSet appsV1.DaemonSet
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty Daemon set",
			fields: fields{
				DaemonSet: appsV1.DaemonSet{},
			},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := daemonSet{
				DaemonSet: tt.fields.DaemonSet,
			}
			if got := d.IsEmpty(); got != tt.want {
				t.Errorf("daemonSet.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_statefulSet_GetVolumeNames(t *testing.T) {
// 	type fields struct {
// 		StatefulSet appsV1.StatefulSet
// 	}
// 	tests := []struct {
// 		name            string
// 		fields          fields
// 		wantVolumeNames []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := statefulSet{
// 				StatefulSet: tt.fields.StatefulSet,
// 			}
// 			if gotVolumeNames := s.GetVolumeNames(); !reflect.DeepEqual(gotVolumeNames, tt.wantVolumeNames) {
// 				t.Errorf("statefulSet.GetVolumeNames() = %v, want %v", gotVolumeNames, tt.wantVolumeNames)
// 			}
// 		})
// 	}
// }

func Test_statefulSet_IsEmpty(t *testing.T) {
	type fields struct {
		StatefulSet appsV1.StatefulSet
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty statefulset",
			fields: fields{
				StatefulSet: appsV1.StatefulSet{},
			},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := statefulSet{
				StatefulSet: tt.fields.StatefulSet,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("statefulSet.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_job_GetVolumeNames(t *testing.T) {
// 	type fields struct {
// 		Job batchV1.Job
// 	}
// 	tests := []struct {
// 		name            string
// 		fields          fields
// 		wantVolumeNames []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			j := job{
// 				Job: tt.fields.Job,
// 			}
// 			if gotVolumeNames := j.GetVolumeNames(); !reflect.DeepEqual(gotVolumeNames, tt.wantVolumeNames) {
// 				t.Errorf("job.GetVolumeNames() = %v, want %v", gotVolumeNames, tt.wantVolumeNames)
// 			}
// 		})
// 	}
// }

func Test_job_IsEmpty(t *testing.T) {
	type fields struct {
		Job batchV1.Job
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty job",
			fields: fields{
				Job: batchV1.Job{
					Spec: batchV1.JobSpec{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := job{
				Job: tt.fields.Job,
			}
			if got := j.IsEmpty(); got != tt.want {
				t.Errorf("job.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
