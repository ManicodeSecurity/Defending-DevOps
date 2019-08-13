# Lab 008 - Istio
The goal of this lab is to enable Istio service mesh in your cluster and enforce an egress policy.

### Ensure you are using the `default` namespace

This lab will work best in the `default` namespace - the following command will ensure that is what we are using.
```
kubectl config set-context $(kubectl config current-context) --namespace default && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Task 1: Cluster Prep
Istio is a complex collection of Kubernetes objects. This task will help us prep our cluster for successful installation. Since we will be creating some RBAC rules, we want to first make sure that we are cluster admin (it is ok to run this again to be safe). Run the following command in Cloud Shell:
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"
```

Now, we Install Istio using the GKE Addon:
```
gcloud beta container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --update-addons=Istio=ENABLED --istio-config=auth=MTLS_STRICT --region=us-west1-a
```

(!)Ensure all cluster operations are labeled `DONE` before continuing(!)
```
gcloud beta container operations list
```

### Task 2: Verify our Istio Installation
Istio is a massive project. Luckily, GKE recently released Istio support out of the box by passing a few beta feature flags upon cluster creation.

First, let's verify that Istio is installed and running properly in our cluster. Ensure the following Kubernetes services are deployed: istio-pilot, istio-ingressgateway, istio-policy, istio-telemetry, prometheus, istio-galley, and, optionally, istio-sidecar-injector.

```
kubectl get svc -n istio-system
```
Ensure the corresponding Kubernetes pods are deployed and all containers are up and running: istio-pilot-*, istio-ingressgateway-*, istio-egressgateway-*, istio-policy-*, istio-telemetry-*, istio-citadel-*, prometheus-*, istio-galley-*, and, optionally, istio-sidecar-injector-*.
```
kubectl get pods -n istio-system
```

### Task 3: Enable Automatic Sidecar Injection

Each pod in the mesh must be running an Istio compatible sidecar. The sidecar is how all traffic to and from pods in the mesh

Manual injection modifies the controller configuration, e.g. deployment. It does this by modifying the pod template spec such that all pods for that deployment are created with the injected sidecar. Adding/Updating/Removing the sidecar requires modifying the entire deployment.

Automatic injection injects at pod creation time. The controller resource is unmodified. Sidecars can be updated selectively by manually deleting a pods or systematically with a deployment rolling update.

The following command will enable automatic injection for the `default` namespace:
```
kubectl label namespace default istio-injection=enabled
```

### Task 4: Launch our API in the Istio Service Mesh
Since we have automatic injection enabled for the `default` namespace, any deployments created in that namespace will now have an extra container aka "sidecar" automatically injected. This now places the pod into the Istio service mesh.
```
# In the manifests/api directory
kubectl create -f .
```
The unshorten service will spin up a load balancer. Ensure the API is accessible. Now that our pod is managed by Istio, we are not going to use the link-unshorten-service IP address as in previous labs. We will use the service provisioned by Istio called `istio-ingressgateway` to grab the routable IP address of the API.
```
kubectl -n istio-system get service istio-ingressgateway
```

Up until version 1.0, Istio’s default behavior was to block access to external endpoints which created connectivity issues and applications were breaking until all endpoints were configured. We are using a version of Istio that newer than 1.0 so egress is not blocked by default.

Paste the IP address with a shortened link as follows in your browser:
```
http://35.197.37.188/api/check?url=https://bit.ly/hi
# This should resolve as expected
```

### Task 5: Build Egress Rules
Lets build some rules to explicit allow outbound egress traffic to only bit.ly and no other endpoints. This can be accomplished by using a `ServiceEntry`. Check out the file `link-unshorten-egress.yaml` located in the `istio-rules` directory and create it as follows:

```
# In the manifests/istio-rules directory
kubectl create -f .
```

Once the rules are created, try to visit the API again and you should be able to successfully unshorten links to `bit.ly` domains only.

```
http://35.197.37.188/api/check?url=https://bit.ly/hi
# This should resolve normally

http://35.197.37.188/api/check?url=https://tinyurl.com/news
# This should NOT resolve
```

### Bonus
[Prometheus](https://istio.io/docs/tasks/telemetry/querying-metrics/) is bundled with Istio in GKE for metrics collection. Can you get the dashboard up and start looking at some metrics from your cluster? You will need to do a `port-forward` similar to earlier labs to use web preview.

### Cleanup

Disable auto istio-injection for the `default` namespace:
```
kubectl label namespace default istio-injection= --overwrite
```

Delete all of the Pods and Istio manifests. In the `/manifests` directory run:
```
kubectl delete -f api -f istio-rules
```

Now, disable Istio:
```
gcloud beta container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --update-addons=Istio=DISABLED --region=us-west1-a
```