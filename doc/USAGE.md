
## Usage
The following assumes you have the plugin installed via

```shell
kubectl krew install volume-reclaim
```

### Scan images in your current kubecontext

```shell
kubectl volume-reclaim
```

### Scan images in another kubecontext

```shell
kubectl volume-reclaim --context=context-name
```

## How it works
Gather all PVC and output all pvcs that doesn't belong to any workloads.