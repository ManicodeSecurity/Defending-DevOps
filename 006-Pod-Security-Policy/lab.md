## PodSecurityPolicy
Pod security policies provide a framework to ensure that pods and containers run only with the appropriate privileges and access only a finite set of resources. Security policies also provide a way for cluster administrators to control resource creation, by limiting the capabilities available to specific roles, groups or namespaces.


### Task 1: Define our PodSecurityPolicy
1. In the `manifests/psp` directory, take a look at the `pod-security-policy.yaml` file and launch it into our new cluster:
```
kubectl create -f pod-security-policy.yaml
```

2. Inspect our new PodSecurityPolicy:
```
kubectl get psp
kubectl describe psp restrict-root
```
## Task 2: Authorize Policies using RBAC

You use role-based access control to create a Role or ClusterRole that grants the desired service accounts access to PodSecurityPolicies. A ClusterRole grants cluster-wide permissions, and a Role grants permissions within a namespace that you define.

For simplicity, we will create a ClusterRole and Rolebinding that applies to all service accounts in the default namespace.

```
# in the manifests/role directory run
kubectl create -f .
```

## Task 3: Enable PSP on your Cluster
We need to enable PSP in our GKE cluster. Warning! If you enable the PodSecurityPolicy controller without first defining and authorizing any actual policies, no users, controllers, or service accounts can create or update Pods. If you are working with an existing cluster, you should define and authorize policies before enabling the controller.

In Cloud Shell, run the following command:
```
# Retrieve the name of your cluster using the following command:
gcloud container clusters list

# Enable PSP
gcloud beta container clusters update <CLUSTER-NAME> --enable-pod-security-policy --region=us-west1-a

# Grab a coffee..this will take a few minutes
```

#### Task 3: Launch a Pod That Runs as Root
1. Inspect the modified Unshorten API deployment located in the `manifests` directory and notice the new `runAsUser` field. This field specifies that for any Containers in the Pod, the first process runs with user ID 0 (root). 

2. Launch the Deployment and service:
```
# In the manifests/root-pod directory
kubectl create -f link-unshorten-deployment.yaml
kubectl create -f link-unshorten-service.yaml
```

You will notice that the Pod fails to instantiate:
```
kubectl get pods
# Inspect the event that occurred to cause the failure
kubectl get events
```
3. Delete the Deployment and Service
```
# In the manifests/root-pod directory
kubectl delete -f .
```

Great job! We just stopped a container running as r00t.

### Launch a Pod That Runs as Non-Root
1. Inspect the modified Unshorten API deployment located in the `manifests` directory and notice the new `runAsUser` field. This field specifies that for any Containers in the Pod, the first process runs with user ID 999 (non-root). 

2. Launch the Deployment and service:
```
# In the manifests/non-root-pod directory
kubectl create -f link-unshorten-deployment-non-root.yaml
kubectl create -f link-unshorten-service-non-root.yaml
```

You will notice that the Pod launches successfully:
```
kubectl get pods
```

### Clean Up
1. In the `manifests` directory:
```
kubectl delete -f psp -f role -f non-root-pod -f root-pod
```

2. Disable PSP on your cluster
```
# Retrieve the name of your cluster using the following command:
gcloud container clusters list

# Disable PSP
gcloud beta container clusters update <CLUSTER-NAME> --no-enable-pod-security-policy --region=us-west1-a

# Grab a coffee..this will take a few minutes
```


