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

Now, paste the IP address with a shortened link as follows in your browser and you will see that there are some issues. Egress traffic is blocked by default. The API is not able to make outbound connections to follow redirects.
```
http://35.197.37.188/api/check?url=https://bit.ly/hi
# NOT ALLOWED or fail to resolve 
```

### Task 5: Build Egress Rules
Lets build some rules to explicit allow outbound egress traffic to only bit.ly. This can be accomplished by using a `ServiceEntry`. Check out the file `link-unshorten-egress.yaml` located in the `istio-rules` directory and create it as follows:

```
# In the manifests/istio-rules directory
kubectl create -f .
```

Once the rules are created, try to visit the API again and you should be able to successfully unshorten links to `bit.ly` domains only. 


### Task 7: Cleanup

In the `manifests` directory:
```
kubectl delete -f api -f istio-rules
```
(!!) *IMPORTANT* (!!)  Disable auto istio-injection for the `default` namespace:
```
kubectl label namespace default istio-injection= --overwrite
```

Now, disable Istio (just to be safe):
```
gcloud beta container clusters update $(gcloud container clusters list --format json | jq -r '.[].name') --update-addons=Istio=DISABLED --region=us-west1-a
```

### Bonus
[Prometheus](https://istio.io/docs/tasks/telemetry/querying-metrics/) is bundled with Istio in GKE for metrics collection. Can you get the dashboard up and start looking at some metrics from your cluster? You will need to do a `port-forward` similar to earlier labs to use web preview.