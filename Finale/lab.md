# Finale
This lab will use the skills and techniques we learned throughout the course to launch a hardened Kubernetes cluster:

Use the manifests located in `8-K8S-Cluster-Secrets/manifests/api` and `8-K8S-Cluster-Secrets/manifests/mysql` as our starting point. We have the following requirements for our unshorten-api / mysql deployment. Make sure to copy the .yaml files to this directory for modification.

### Cluster Security Requirements

1. Secrets should be stored in Vault running in the cluster
2. All Pods should should have a security policy that enforces the following:
    - Forbids the use of root namespaces
    - Has `seccomp` enabled
    - Read-only container filesystem
    - Restricts elevation to root privileges

3. The application must exist in three namespaces:
    - `development`
    - `test`
    - `qa`

4. The Kubernetes API should have the following groups applied via RBAC and use HTTP Basic Authentication
    - `administrator` role has full administrative access to the cluster in all namespaces
    - `development` role can has access to read/write pods and secrets in the development namespace
    - `qa` role has read-only access to all pods in all namespaces

5. The kubelet API should be properly authenticated
6. Create an audit log policy that generates logs when secrets are read/written and pods a created
7. Enable [Prometheus](https://github.com/giantswarm/kubernetes-prometheus/) in the cluster to gain deeper insight into cluster health and metrics
