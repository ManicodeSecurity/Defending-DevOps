# Defending Modern DevOps Environments Agenda

Timeframes are estimated

## Day 1:
### Section 1: Introduction (2 Hours)
- Logistics and Introductions
- Tool Overview
- DevSecOps Introduction 
- Container Basics
- Container Security
LAB: 1-Containerizing-An-Application

### Section 2: Kuberenetes Basics (2 Hours)
- History
- Introduction
https://vimeo.com/173090241
- Technical Overview and Primitives
- Security Features
    - Built for CI/CD pipelines
    - Upgrading is simple (again, less snowflakes)
- Security Pitfalls
 -"separation of concerns" using separate environments

- Deployment Options (AWS, GCP, etc.)
Lab: Minikube Installation and Configuration

### Section 3: Interacting with your Kubernetes Cluster (2 Hours)
- Manifests
- API resources
- kubectl
  - 
Lab: Working with Your Cluster
- Scaling and Updating Code

### Section 4: Kuberenetes Cluster Hardening
- Overview
- RBAC
- API Authentication
Lab: Authenticating to your Cluster
- Authorization
Lab: Cluster Authorization


## Section x: Advanced Kubernetes Security Features
seccomp
- PodSecurityPolicy
Lab: Implement PSP in Your Cluster
https://docs.bitnami.com/kubernetes/how-to/secure-kubernetes-cluster-psp/
- Network Policy
Lab: Implement NetworkPolicy in your Cluster
https://speakerdeck.com/ianlewis/kubernetes-security-best-practices
https://ftp.osuosl.org/pub/fosdem/2018/UD2.120%20(Chavanne)/containers_kubernetes_security.mp4
- istio
    - sidecar proxy
- RBAC for Kubelet
- Rotate Kubelet Certs

## Day 2


### Section 5: Secrets Management
Lab 4: Secrets Everywhere!
Lab Secrets using Vault


### Section 6: Logging and Monitoring
- Logging 
    - ElasticSearch (free like a puppy)
    - rsyslog
    - auditd (go-audit)
    - OSQuery
    - nginx logs
    - docker logs
Lab: https://kubernetes.io/docs/tasks/debug-application-cluster/audit/

- Alerting
    - ElastAlert
    - File watch
    - ssh connections
    - slack bots
Elk stack monitoring your application 

### Section 7: DevSecOps Pipelines
- 
Lab: kubeaudit
https://github.com/Shopify/kubeaudit
Lab: Static Code Analysis 
Lab: Third Party Dependency Analysis

https://github.com/sebgoa/oreilly-kubernetes