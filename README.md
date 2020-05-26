# kubectl-unused-volumes

A `kubectl` plugin to gather all PVC and output all pvcs that doesn't belong to any workloads.
This plugins checks for standard kubernetes workloads : 
- DaemonSet
- Deployment
- Job
- StatefulSet


## Quick Start

```
kubectl krew install unused-volumes
kubectl unused-volumes
```

