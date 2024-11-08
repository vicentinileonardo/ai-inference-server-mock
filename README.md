# AI inference server mock

## Description

## Setup

```bash
# Point your shell to minikube's docker-daemon
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
kubectl describe pods
kubectl describe pods -l app=ai-inference-server

# If pods aren't appearing or are in error state, check events:
kubectl get events --sort-by='.lastTimestamp'

# Check the pod logs
kubectl logs -l app=ai-inference-server
kubectl logs deploy/ai-inference-server
kubectl logs -f $(kubectl get pods -l app=ai-inference-server -o name)



# Test the service
kubectl run --rm -it --image=busybox:latest test-client -- /bin/sh
wget -O- http://ai-inference-server:8080/scheduling
curl http://ai-inference-server:8080/scheduling # curl is not installed in this case

# Get the pod IP and test the service
kubectl get endpoints ai-inference-server
# alternative
kubectl get pods -o wide

wget -O- http://<POD_IP>:8080/scheduling 


# Get the service IP and test the service
kubectl get svc ai-inference-server
wget -O- http://<IP>:8080/scheduling



# Test the pod directly with port-forwarding (replace POD_NAME with actual pod name)
kubectl port-forward pod/POD_NAME 8080:8080
# In another terminal:
curl localhost:8080/health # it works
curl localhost:8080/scheduling # it works

# Clean up
kubectl delete deploy/ai-inference-server
kubectl delete service/ai-inference-server
```