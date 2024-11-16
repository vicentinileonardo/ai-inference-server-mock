# AI inference server mock

## Description

This repository is part of a project: **Carbon-aware workload scheduling in a multi-cloud environment**.
This is a simple mock of an **AI inference server** that provides a scheduling endpoint to simulate the inference process.

## Setup

### Kubernetes deployment

```bash
# Point your shell to minikube's docker-daemon
# this step may vary depending on your setup
eval $(minikube docker-env)

# Check the current context
docker ps
docker images

# Build the image
docker build -t ai-inference-server:latest .

# Apply the deployment and service
kubectl apply -f server-deployment.yaml
kubectl apply -f server-service.yaml

# Check the deployment, pods, and services
kubectl get deployments
kubectl get pods
kubectl get services

# Check the service
kubectl get svc ai-inference-server

# Check detailed pod information including events
kubectl describe pods -l app=ai-inference-server

# If pods aren't appearing or are in error state, check events:
kubectl get events --sort-by='.lastTimestamp'

# Check the pod logs
kubectl logs -l app=ai-inference-server
kubectl logs deploy/ai-inference-server
kubectl logs -f $(kubectl get pods -l app=ai-inference-server -o name)

# Test the service with a test client and curl
kubectl run --rm -it --image=alpine/curl:latest test-client -- /bin/sh
curl http://ai-inference-server:8080/scheduling 

# Test the service with a test client and wget
kubectl run --rm -it --image=busybox:latest test-client -- /bin/sh
wget -O- http://ai-inference-server:8080/scheduling

# Get the pod IP
kubectl get endpoints ai-inference-server
# alternative
kubectl get pods -o wide
wget -O- http://<POD_IP>:8080/scheduling 

# Get the service IP
kubectl get svc ai-inference-server

# Test the pod directly with port-forwarding (replace POD_NAME with actual pod name)
kubectl port-forward pod/POD_NAME 8080:8080
# In another terminal:
curl localhost:8080/health # it works
curl localhost:8080/scheduling # it works

# Clean up
kubectl delete deploy/ai-inference-server
kubectl delete service/ai-inference-server
docker rmi ai-inference-server
```

## TODO

- folder structure organization
- multi stage build in Dockerfile
- helm chart