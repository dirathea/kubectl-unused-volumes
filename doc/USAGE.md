
## Usage
The following assumes you have the plugin installed via

```shell
kubectl krew install kubectl-volume-reclaim
```

### Scan images in your current kubecontext

```shell
kubectl kubectl-volume-reclaim
```

### Scan images in another kubecontext

```shell
kubectl kubectl-volume-reclaim --context=context-name
```

## How it works
Gather all pods and PVC on a namespaces, and output all pvcs that doesn't belong to any pod.