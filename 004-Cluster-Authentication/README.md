# Authentication and Authorization

The goal of this lab is to enhance the security of our cluster using built-in Role Based Access Control (RBAC). We will explore authentication and authorization strategies and apply them to our GKE cluster.

We will be setting up our cluster with two namespaces - `development` and `production`. User 1 (our admin user) will have full-access to administer the objects in the cluster and we will restrict access to only the `development` namespace for User 2 ("the intern") by using custom RBAC policies.

### Lab Setup
Your GCP Project has two end users provisioned with the following permissions:
```
User 1: <your-username>@manicode.us
Roles:
Kubernetes Engine Admin
Editor

User 2: <your-intern-email>@manicode.us
Roles:
Minimal GKE Role
Browser
```
The role titled `Minimal GKE Role` is a custom role in GCP. It includes the bare-minimum permissions to be able to access the cluster but not the resources within. The `Minimal GKE Role` has only the following permissions:
```
container.apiServices.get
container.apiServices.list
container.clusters.get
container.clusters.getCredentials
```

### Task 1: Launch Your Infrastructure
First, we will spin up our application in both a `development` and `production` namespace.

Note: You should be logged in to Cloud Shell using the admin account provided at the beginning of class to run the following commands, NOT `<your-intern-email>@manicode.us`.

We need to retrieve the credentials of our running cluster using the following `gcloud` command. This command updates our kubeconfig in Cloud Shell file with appropriate credentials and endpoint information to point kubectl at a specific cluster in Google Kubernetes Engine.

```
# Use gcloud get-credentials to retrieve the cert
gcloud container clusters get-credentials $(gcloud container clusters list --format json | jq -r '.[].name') --zone us-west1-a --project $GOOGLE_CLOUD_PROJECT
```
Now we launch our pods and services for each Namespace:
```
# in the manifests/development directory

kubectl create -f link-unshorten-ns.yaml
kubectl create -f link-unshorten-service.yaml
kubectl create -f link-unshorten-deployment.yaml
```
Make sure our pods are running in the `development` namespace:
```
kubectl get pods --namespace=development
```

Do the same for the production namespace:
```
# in the manifests/production directory

kubectl create -f link-unshorten-ns.yaml
kubectl create -f link-unshorten-service.yaml
kubectl create -f link-unshorten-deployment.yaml
```
Ensure pods are running without error in both namespaces:
```
kubectl get pods --all-namespaces
```

Take note of this process. Our user has full administrative access to our cluster due to being provisioned with the `Kubernetes Engine Admin` role. We will now see how RBAC helps give us granular access control at the object-level within our cluster.

### Task 2: Authenticate as a Restricted User
We will now log in using a separate user who has very locked down access to the entire project. In an incognito window browse to `cloud.google.com` and authenticate with the user `<your-intern-email>@manicode.us` and the same password that was provided to you for the admin user.

Note: *Using the same password for multiple accounts is bad. Don't do this at home.*

Now open up Cloud Shell and use the following `gcloud get-credentials` command to retrieve the credentials for your user so we can start interacting with the cluster. This is the same cluster you just launched the `production` and `development` infrastructure in.

```
# Authenticate to the cluster

gcloud container clusters get-credentials $(gcloud container clusters list --format json | jq -r '.[].name') --zone us-west1-a --project $GOOGLE_CLOUD_PROJECT
```
Now, attempt to run some `kubectl` queries on the cluster.
```
kubectl get pods --namespace=production
kubectl get pods --namespace=development
kubectl get secrets
kubectl run link-unshorten --image=jmbmxer/link-unshorten:0.1 --port=8080
```
These should all fail with a `Forbidden` error. While <your-intern-email>@manicode.us does technically have an account on the cluster, RBAC is stopping it from accessing any of the objects.

One simple way to quickly check what permissions your user has is to use `can-i`:
```
# Some examples
kubectl auth can-i create deployment --namespace production
kubectl auth can-i list secrets --namespace default
```

### Task 3: Add Yourself as `cluster-admin`
By default, User 1 will not be able to create the `roles` or `rolebindings` needed to begin building our RBAC policies. We need to ensure User 1 (our Administrator) has the appropriate access to the cluster by granting the user `cluster-admin` rights.

`cluster-admin` is one of several Default User-facing roles included with every Kubernetes installation. They should be used with caution as many of these roles grant excessive privileges and are often abused for a quick fix.

[Default System and User-Facing Roles and Role and RoleBindings](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#default-roles-and-role-bindings)

 Switch back to your Cloud Shell for User 1 (the administrative user) and run the following commands:

```
# Inspect what exactly the `cluster-admin` role will be granting the user
kubectl get clusterrole/cluster-admin -o yaml

# Compare this to the second most privileged role, `admin`
kubectl get clusterrole/admin -o yaml
```

Now, add User 1 to the Default ClusterRole and ClusterRoleBinding called `cluster-admin`
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

`cluster-admin` is the most elevated role in Kubernetes. Use with extreme caution!

Use the helpful `kubectl auth can-i` command to verify you are able to create roles:
```
kubectl auth can-i create roles --as=root --as-group=system:authenticated --as-group=system:masters
yes
```

### Task 4: Create RBAC Rules
Our user `<your-intern-email>@manicode.us` is a restricted user so we only want to grant access to read pods in the `development` namespace and nothing more. We will use RBAC to enforce a policy

Now, open the file `user-role-binding.yaml` in the `manifests/role` directory and replace <your-intern-email> with the one provided to you. It will be the same as your admin account but with the word `intern` at the end (eg. `manicode0003intern@manicode.us`).
```
# In the manifests/role directory
kubectl create -f .
kubectl get role --all-namespaces
```

### Task 5: Verify Pods can be Accessed by the Intern

Switch back to the Cloud Shell for `<your-intern-email>@manicode.us` and run the following commands:
```
kubectl get pods --namespace=development
# success
kubectl get pods --namespace=production
# fail
```

We have successfully limited access using RBAC.

### Bonus 1
You are tasked with performing an audit of the `Minimal GKE` role in GCP below. What would you change?

```
container.apiServices.get
container.apiServices.list
container.clusters.delete
container.clusters.get
container.clusters.getCredentials
container.clusters.list
resourcemanager.projects.get
resourcemanager.projects.list
```

### Bonus 2
Our intern just got promoted to Jr. DevSecOpsSysAdminNinja! Change the permissions to allow `get`, `watch`, `list`, `update`, and `delete`, on `secrets` objects in the `development` and `production` namespaces?

### Task 6: Cleanup
Don't forget to delete the `development` and `production` namespace when you are done with the Bonuses.
```
kubectl delete ns development production
```
