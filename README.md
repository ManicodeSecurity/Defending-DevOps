### Defending Modern DevOps Environments 

This repository contains all of the labs for the 2-day "Defending Modern DevOps Course"

# Agenda:
1-Lab-Setup

2-Containerizing-An-Application

3-K8S-Cluster-Setup

4-K8S-Cluster-Authentication

5-K8S-Dashboard

6-Pod-Security-Policy

7-Attacking-Kubelet

8-K8s-Cluster-Secrets

9-Security-Pipeline

10-Kube-Logs

11-Finale


## Useful Debugging Commands
```
minikube logs
kubectl get logs <podname>
kubectl exec -it <podname> /bin/bash
kubectl describe pod|service|deployment <name> 
kubectl get secret <secretname> 
kubectl get events | grep <thething>
```