# AI inference server mock

## Description

This repository is part of a project: **Carbon-aware workload scheduling in a multi-cloud environment**.
This is a simple mock of an **AI inference server** that provides a scheduling endpoint to simulate the inference process.

It is possible that a modified version of this will be used as proxy for the real AI inference server in the project, if needed.

## Setup

### Kubernetes deployment

The deployment is done using Helm.

```bash
helm install test-ai-mock ./chart --create-namespace --namespace ai-inference-server-mock
```
or:
```bash
# to be modified properly
helm repo add ...
helm install ...
```

Check the deployment, pod, and service:
```bash
kubectl get all -n ai-inference-server-mock
```

Check detailed pod information including events:
```bash
kubectl describe pods -l app=ai-inference-server-mock -n ai-inference-server-mock
```

If pods aren't appearing or are in error state, check events:
```bash
kubectl get events --sort-by='.lastTimestamp'
```

Check the pod logs:
```bash
kubectl logs -l app=ai-inference-server-mock -n ai-inference-server-mock
kubectl logs deploy/ai-inference-server-mock -n ai-inference-server-mock
kubectl logs -f $(kubectl get pods -l app=ai-inference-server-mock -o name -n ai-inference-server-mock) -n ai-inference-server-mock
```

Test the service with a test client and `curl`:
```bash
kubectl run --rm -it --image=alpine/curl:latest test-client -- /bin/sh
curl http://ai-inference-server-mock.ai-inference-server-mock.svc.cluster.local:8080/scheduling  
```

Test the service with a test client and `wget`:
```bash
kubectl run --rm -it --image=busybox:latest test-client -- /bin/sh
wget -O- http://ai-inference-server-mock.ai-inference-server-mock.svc.cluster.local:8080/scheduling 
```

Get the pod IP (if needed for debugging purposes):
```bash
kubectl get endpoints ai-inference-server-mock -n ai-inference-server-mock
# alternative
kubectl get pods -o wide -n ai-inference-server-mock 
```

Get the service IP (if needed for debugging purposes):
```bash
kubectl get svc ai-inference-server-mock -n ai-inference-server-mock
```

Test the pod directly with port-forwarding (replace POD_NAME with actual pod name):
```bash
kubectl port-forward pod/POD_NAME 8080:8080
# In another terminal:
curl localhost:8080/health
curl localhost:8080/scheduling
```

Clean up:
```bash
helm uninstall test-ai-mock
```

## TODO
- folder structure organization
- multi stage build in Dockerfile
- release helm chart (for now there is no published chart)