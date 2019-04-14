# Network Policies
For network policies to be at their most effective, we want to ensure that traffic can flow only where it is needed, and nowhere else. We generally want to start with a `DenyAll` default policy that matches all pods with and then take a structured approach to adding network policies which will permit traffic between application pods as necessary.

Suppose we have an application called my-app that stores data in a Postgres database. The following example defines a policy that allows traffic from my-app to my-postgres on the default port for Postgres:

### Create the `lab007` Namespace and Use as Default

We will create a new Namespace for every lab and switch contexts to ensure it is the default when using `kubectl`.
```
kubectl create ns lab007 && \
kubectl config set-context $(kubectl config current-context) --namespace lab007 && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Task 1: Enable Network Policies in our Cluster
Network policy enforcement is only available for clusters running Kubernetes version 1.7.6 or later. GKE uses the popular [Calico](https://www.projectcalico.org/) overlay network when using Network Policies. 

First, update the addons for our cluster to support Network Policies:
```
gcloud container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --update-addons=NetworkPolicy=ENABLED --region=us-west1-a --project=$GOOGLE_CLOUD_PROJECT
```

Next, enable Network Policies:
```
gcloud container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --enable-network-policy --quiet --region=us-west1-a --project=$GOOGLE_CLOUD_PROJECT
```

### Task 2: Create our Network Policy
Go to the `manifests/network-policies` directory and inspect the Network policy named `hello-unshorten.yaml`. This policy simply selects Pods with label `app=unshorten-api` and specifies an ingress policy to allow traffic only from Pods with the label `app=unshorten-fe`. We only want to allow traffic from pods that are acting as frontends to our API. 

In the `manifests/network-policies` directory run:
```
kubectl create -f .

# Check out our new network policy
kubectl get networkpolicies
```

### Task 3: Launch our API
In the `manifests/api` directory run the following:
```
kubectl create -f .
```

### Task 4: Launch a Permitted Pod
First, run a temporary Pod with the label app=unshorten-fe and get a shell in the Pod. This should have access to talk to the `unshorten-api` due to our new Network Policy.

```
kubectl run -l app=unshorten-fe --image=alpine --restart=Never --rm -i -t test-1

```
Once you get a shell, try to run a `wget` hitting the `unshorten-api-service`:
```
# This should be allowed (because our pod has a permitted label name)
wget -qO- --timeout=2 http://unshorten-api:80/api/check?url=bit.ly/test
```
Now type `exit` to exit the shell of the pod.

### Task 5: Launch a Blocked Pod
```
# This should NOT be allowed (wrong label name)
kubectl run -l app=other-teams-app --image=alpine --restart=Never --rm -i -t test-2

wget -qO- --timeout=2 http://unshorten-api:80/api/check?url=bit.ly/test
# timeout
```

### Task 6: Cleanup
Don't forget to delete the `lab007` namespace when you are done with the Bonuses.
```
kubectl delete ns lab007 && \
kubectl config set-context $(kubectl config current-context) --namespace default && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

Disable Network Security Policies in our Cluster (This will take several minutes):
```
gcloud container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --no-enable-network-policy --region=us-west1-a
```

Ensure ALL upgrade operations are complete before moving on to the next lab:
```
gcloud beta container operations list
```