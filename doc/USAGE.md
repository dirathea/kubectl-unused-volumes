
## Usage
The following assumes you have the plugin installed via

```shell
kubectl krew install unused-volumes
```

### Scan images in your current kubecontext

```shell
kubectl unused-volumes
```

### Scan images in another kubecontext

```shell
kubectl unused-volumes --context=context-name
```

## How it works
Gather all PVC and output all pvcs that doesn't belong to any workloads. Workloads including : Deployment, Job, Statefulset, DaemonSet. Other than this workload (such as CRD) will be **ignored and the volume will be listed as `No Reference`**.