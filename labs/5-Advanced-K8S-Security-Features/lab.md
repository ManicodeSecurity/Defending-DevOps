## PodSecurityPolicy
Pod security policies provide a framework to ensure that pods and containers run only with the appropriate privileges and access only a finite set of resources. Security policies also provide a way for cluster administrators to control resource creation, by limiting the capabilities available to specific roles, groups or namespaces.

## Task 1: Launch a New Minikube Cluster
We need to pass some flags to our Kubernetes API Server in order to use Pod Security Policies. First we delete our cluster:
```
minikube delete
```
Now we launch a new cluster as follows:
```
minikube start \
   --extra-config=apiserver.Authorization.Mode=RBAC \
   --extra-config=apiserver.Admission.PluginNames=PodSecurityPolicy
```

## Task 2: Create our PodSecurityPolicy
1. In the `podsecuritypolicy/manifests` directory, take a look at the `pod-security-policy.yaml` file and launch it into our new cluster:
```
kubectl create -f pod-security-policy.yaml`
```
Now, we can inspect our new PodSecurityPolicy:
```
kubectl get psp
```

## Task 3: Launch a Pod That Runs as Root
1. Inspect the modified Unshorten API deployment located in the `pod-security-policy/manifests` directory and notice the new `runAsUser` field. This field specifies that for any Containers in the Pod, the first process runs with user ID 0 (root). 

2. Launch the Deployment and service:
```
kubectl create -f link-unshorten-deployment.yaml
kubectl create -f link-unshorten-service.yaml
```

3. You will notice that the Pod fails to instantiate:
```
kubectl get pods
# Inspect the event that occured to cause the failure
kubectl get events
```
Great job! We just stopped a container running as r00t.

## Auditing
A recent addition to Kubernetes, (https://kubernetes.io/docs/tasks/debug-application-cluster/audit/)
[auditing] gives adminstrators and security teams the ability to log and monitor security-related events occuring on a cluster. By using audit policies, we can create granular rulesets to focus on only on the meaningful events and cut down on the noise. 

# Task 1: We will need to enable some flags in our cluster so we will start off with a clean cluster using the following command:
```
minikube delete
```

# Task 2: To enable audition in Minikube we must utilize the `addons` feature. To do this, we must copy over the `audit-policy.yaml` manifest to the `.minikube/addons` directory:
```
cp /path/to/5-Advanced-K8S-Security-Features/audit-policy/manifests/audit-policy.yaml ~/.minikube/addons
```

# Task 3: Create the `audit.log` file that we will write our JSON logs out to:
```
touch ~/.minikube/logs/audit.log
```

# Task 4: Launch the cluster with the following flags. Make sure to enter the correct path to the .minikube folder!
```
minikube start \
    --feature-gates=AdvancedAudit=true \
    --extra-config=apiserver.Audit.LogOptions.Path=/path/to/.minikube/logs/audit.log \    --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml
```

# Take a look at the logs coming in. We will filter those from 127.0.0.1 to cut down on the noise. [https://stedolan.github.io/jq/download/](jq) is a command line tool which greatly helps with formatting JSON to the terminal. It's a handy tool to have available when working with JSON.
```
tail -f  ~/.minikube/logs/audit.log  | jq '.| select(.sourceIPs | contains(["127.0.0.1"]) | not)'
```

## Bonus: Using a webhook, send the Kubernetes logs to an API such as requestb.in (this is VERY insecure so only do this for testing purposes). 

Hint 1: The `webhook.yaml` file should be copied to the same location as the `audit-policy.yaml` file in `.minikube/addons` directory.

Hint 2: We need to pass additional flags to Minikube to enable the webhook. Starting a new cluster should look like this:
```
minikube start \
    --feature-gates=AdvancedAudit=true \
    --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml \
    --extra-config=apiserver.Audit.WebhookOptions.ConfigFile=/etc/kubernetes/addons/webhook.yaml 
```

## Discussion Question: How would you ingest these logs into your current log management systems? What would you alert on?


#Kubernetes security tip: if your pod doesn't need API access set `AutomountServiceAccountToken: false` for an extra layer of defense.
Disable the automounting of a default service account token
https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server'

ISTIO - Super Bonus ++
https://github.com/istio/istio/blob/master/tools/minikube.md


