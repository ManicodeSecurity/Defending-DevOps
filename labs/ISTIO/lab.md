### Task 1: Cluster Prep
Istio is a complex collection of Kubernetes objects. This task will help us prep our cluster for successful installation. Since we will be creating some RBAC rules, we want to first make sure that we are cluster admin (it is ok to run this again to be safe):
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"
```

To interact with Istio, we will use `istioctl` which is bundled as a binary with the Istio install package.
```
# In the istio-1.0.0 directory
export PATH=$PWD/bin:$PATH
# Ensure the binary is available
istio version
```

### Task 2: Install Istio Components and Enable Automatic Sidecar Injection
Istio is a massive project. Check out the yaml file located at `istio-1.0.0/install/kubernetes/istio-demo-auth.yaml`. Wow. Such yaml. Let's install the components necessary in our cluster:
```
# In the istio-1.0.0 directory 
kubectl create -f install/kubernetes
```
Each pod in the mesh must be running an Istio compatible sideecar. The sidecar is how all traffic to and from pods in the mesh

Manual injection modifies the controller configuration, e.g. deployment. It does this by modifying the pod template spec such that all pods for that deployment are created with the injected sidecar. Adding/Updating/Removing the sidecar requires modifying the entire deployment.

Automatic injection injects at pod creation time. The controller resource is unmodified. Sidecars can be updated selectively by manually deleting a pods or systematically with a deployment rolling update.

The following command will enable automatic injection for the `default` namespace:
```
kubectl label namespace default istio-injection=enabled
```

### Task 3: Launch our API in the Istio Service Mesh
Since we have automatic injection enabled for the `default` namespace, any deployments created in that namespace will now have an extra container aka "sidecar" automatically injected. This now places the pod into the Istio service mesh.
```
# In the manifests/api directory
kubectl create -f .
```
The unshorten service will spin up a load balancer. Ensure the API is accessible. Now that our pod is managed by Istio, we are not going to use the link-unshorten-service IP address as in previous labs. We will use the service provisioned by Istio called `istio-ingressgateway` to grab the routable IP address of the API.
```
kubectl -n istio-system get service istio-ingressgateway
```

Now, paste the IP address with a shortened link as follows in your browser and you will see that there are some issues. Egress is blocked by default. The API is not able to make outbound connections by default.
```
http://35.197.37.188/api/check?url=https://bit.ly/hi
# NOT ALLOWED or fail to resolve 
```

### Task 4: Build Egress Rules
Lets build some rules to explictlly allow outbound egress traffic to only bit.ly. This can be accomplished by using a `ServiceEntry`. Check out the file `link-unshorten-egress.yaml` located in the `istio-rules` directory and create it as follows:

```
# In the manifests/istio-rules directory
kubectl create -f .
```

Once the rules are created, try to visit the API again and you should be able to successfully unshorten links to `bit.ly` domains only. 

### Task 5: Logging and Monitoring with Istio




Grafana
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 8080:3000

Then Click "Web Preview" in cloud shell

