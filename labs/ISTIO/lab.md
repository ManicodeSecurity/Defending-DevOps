### Task 1: Cluster Prep
Istio is a complex collection of Kubernetes objects. This task will help us prep our cluster for successful installation. Since we will be creating some RBAC rules, we want to first make sure that we are cluster admin (it is ok to run this again to be safe):
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"
```

### Task 2: Install Istio Components and Enable Automatic Sidecar Injection
Istio is huge. Take a look at the install istio-demo-auth.yaml file. These are the Kubernetes objects that are needed to run Istio (including some extra features) in our cluster. Wow. Such yaml. 

```
# In the istio-1.0.0 directory 
kubectl create -f install/kubernetes
```
Each pod in the mesh must be running an Istio compatible sidecar. The sidecar is how all traffic to and from pods in the mesh communicate.

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
The unshorten service will spin up a load balancer. Ensure the API is accessible by running the following:
```
kubectl get svc
# Grab the EXTERNAL-IP from the link-unshorten-service
```
Now, paste the IP address in your browser and 
you will see that there are some issues. Egress is blocked by default
http://35.197.37.188/api/check?url=https://bit.ly/hi
NOT ALLOWED!

Lets build some rules to explictlly allow outboud traffic to only bit.ly
Grafana
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 8080:3000

Then Click "Web Preview" in cloud shell

