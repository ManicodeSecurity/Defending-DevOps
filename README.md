### Defending Modern DevOps Environments 

This repository contains all of the labs for the Manicode "Defending Modern DevOps Course"

# Agenda:
001-Lab-Setup

002-Containerizing-An-Application

003-K8S-Cluster-Setup

004-K8S-Cluster-Authentication

005-K8S-Dashboard

006-Pod-Security-Policy

007-Istio

008-K8S-Cluster-Secrets

009-Kube-Logs

010-Security-Pipeline

DEMO-Attacking-Kubelet

Finale

## Useful `kubectl` Commands
Helpful `kubectl` commands to interact with your cluster and its components:

### Retrieve Info about your cluster
```
# View your cluster credentials and location
kubectl config view

# View list of services running on your cluster
kubectl cluster-info

# View node info
kubectl describe nodes
```

### Interact with running pods 
```
# Display all pods in all namespaces in the cluster
kubectl get pods --all-namespaces

# Use -o wide to show more detail
kubectl get pod -o wide --all-namespaces

# List all services running in the cluster
kubectl get svc --all-namespaces

# Get a shell in a container within the pod
kubectl exec -it <you-pod-name> --namespace=<namespace> /bin/bash
```

### View Logs
```
# View pods logs (first container in pod)
kubectl logs <your-pod-name>

# View pod logs (specific container)
kubectl logs <your-pod-name> -c <your-container-name>
```
### Misc. commands
```
kubectl get logs <podname>
kubectl exec -it <podname> /bin/bash
kubectl describe pod|service|deployment <name> 
kubectl get secret <secretname> 
kubectl get events | grep <thething>
kubectl create --v 10 -f .
```

For more `kubectl` commands check out the [kubectl cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/#interacting-with-nodes-and-cluster)

