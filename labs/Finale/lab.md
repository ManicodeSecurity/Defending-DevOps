# Finale
This lab is open-ended and will use all of the skills we learned to launch a hardened Kubernetes cluster:

Launch a completely new application, API, or microservice into a fresh Minikube with the following requirements:

1. Your application must use two or more separate Pods. A simple example would be Wordpress and MySQL
2. Docker images used in the cluster must contain no `Critical` or `High` CVE vulnerabilities 
3. Deploy Vault in the cluster to store secrets
4. All Pods should should have a security policy enforced that forbids the use of root namespaces, has `seccomp` enabled, enforced a read-only container filesystem, and restricts elevation to root privileges
5. The application must exist in three namespaces:
    - `development` 
    - `test` 
    - `qa`

6. The Kubernetes API should have the following groups applied via RBAC and use HTTP Basic Authentication
    - `administrator` role has full administrative access to the cluster in all namespaces
    - `development` role can has access to read/write pods and secrets in the development namespace
    - `qa` role has read-only access to all pods in all namespaces

7. The kubelet API must have authentication 
8. Create an audit log policy that generates logs when secrets are read/written and pods are created
9. Enable [Prometheus](https://github.com/giantswarm/kubernetes-prometheus/) in the cluster to gain deep insight into cluster health and metrics