## PodSecurityPolicy
Pod security policies provide a framework to ensure that pods and containers run only with the appropriate privileges and access only a finite set of resources. Security policies also provide a way for cluster administrators to control resource creation, by limiting the capabilities available to specific roles, groups or namespaces.

## Task 1: Launch a New Minikube Cluster
1. We need to pass some flags to our Kubernetes API Server in order to use Pod Security Policies. First we delete our cluster:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube delete
```

2. Now we launch a new cluster as follows:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube start \
   --extra-config=apiserver.Authorization.Mode=RBAC \
   --extra-config=apiserver.Admission.PluginNames=PodSecurityPolicy
```

## Task 2: Create our PodSecurityPolicy
1. In the `manifests/podsecuritypolicy` directory, take a look at the `pod-security-policy.yaml` file and launch it into our new cluster:
```
kubectl create -f pod-security-policy.yaml
```

2. Inspect our new PodSecurityPolicy:
```
kubectl get psp
kubectl describe psp restrict-root
```

## Task 3: Launch a Pod That Runs as Root
1. Inspect the modified Unshorten API deployment located in the `manifests` directory and notice the new `runAsUser` field. This field specifies that for any Containers in the Pod, the first process runs with user ID 0 (root). 

2. Launch the Deployment and service:
```
kubectl create -f link-unshorten-deployment.yaml
kubectl create -f link-unshorten-service.yaml
```

3. You will notice that the Pod fails to instantiate:
```
kubectl get pods
# Inspect the event that occurred to cause the failure
kubectl get events
```

Great job! We just stopped a container running as r00t.


