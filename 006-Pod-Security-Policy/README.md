## PodSecurityPolicy
Pod security policies provide a framework to ensure that pods and containers run only with the appropriate privileges and access only a finite set of resources. Security policies also provide a way for cluster administrators to control resource creation, by limiting the capabilities available to specific roles, groups or namespaces.

If you haven't already or have been checking email and Slacking for the past 5 labs, please ensure you are `cluster-admin`. 
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

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
### Task 2: Authorize Policies using RBAC

When a PodSecurityPolicy resource is created, it does nothing. In order to use it, the requesting user or target pod’s service account must be authorized to use the policy, by allowing the use verb on the policy.

RBAC is used to create a Role or ClusterRole that grants the desired service accounts access to PodSecurityPolicies. A ClusterRole grants cluster-wide permissions, and a Role grants permissions within a namespace that you define.

For simplicity, we will create a ClusterRole and ClusterRolebinding that applies to all authenticated users in a default namespace.

```
# in the manifests/role directory run
kubectl create -f .
```

### Task 3: Enable PSP on your Cluster
We need to enable PSP in our GKE cluster. Warning! If you enable the PodSecurityPolicy controller without first defining and authorizing any actual policies, no users, controllers, or service accounts can create or update Pods. If you are working with an existing cluster, you should define and authorize policies before enabling the controller.

In Cloud Shell, run the following command:
```
# Enable PSP
gcloud beta container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --enable-pod-security-policy --region=us-west1-a

# Grab a coffee..this will take a few minutes
```

### Task 4: Launch a Pod That Runs as Root
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

# Events are in non-sequential order by default - use this to order by timestamp
kubectl get events --sort-by='.metadata.creationTimestamp' -o 'go-template={{range .items}}{{.involvedObject.name}}{{"\t"}}{{.involvedObject.kind}}{{"\t"}}{{.message}}{{"\t"}}{{.reason}}{{"\t"}}{{.type}}{{"\t"}}{{.firstTimestamp}}{{"\n"}}{{end}}'
```
3. Delete the Deployment and Service
```
# In the manifests/root-pod directory
kubectl delete -f .
```

Great job! We just stopped a container running as r00t.

### Task 5: Launch a Pod That Runs as Non-Root
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

## Task 6: Restrict Volume Mounts
If you take a close look at the Deployment manifest in the `manifests/non-root-pod` directory you will see that it is requesting to mount a...questionable...directory.

Your task now is to modify the PodSecurityPolicy to whitelist known-valid and safe volume mounts. More info can be found [here](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#example-policies) on writing granular PodSecurityPolicies.

## Bonus
Audit your PodSecurityPolicy using [kube-psp-advisor](https://github.com/sysdiglabs/kube-psp-advisor)

Kubernetes Pod Security Policy Advisor is an opensource tool from Sysdig that can automatically generate the Pod Security Policy for all the resources in the entire cluster.

More info from Sysdig can be found [here](https://sysdig.com/blog/enable-kubernetes-pod-security-policy/).

Note: This is a very new project and at the time of this writing is having problems running in Cloud Shell. You may have to run it on your local machine (proceed with caution).

### Clean Up
1. In the `manifests` directory:
```
kubectl delete -f psp -f role -f non-root-pod -f root-pod
```

2. (!!) *IMPORTANT* (!!) Disable PSP on your cluster
```
# Disable PSP
gcloud beta container clusters update $(gcloud container clusters list --format json | jq -r '.[].name')  --no-enable-pod-security-policy --region=us-west1-a

# Disable Legacy Authorization
gcloud container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --no-enable-legacy-authorization --region=us-west1-a

# Grab another coffee..this will take a few minutes
```


