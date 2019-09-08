### Exploring the Kubernetes Dashboard
Kubernetes Dashboard is a general purpose, web-based UI for Kubernetes clusters. It allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

Since version 1.7 Dashboard uses more secure setup. It means, that by default it has minimal set of privileges and can only be accessed over HTTPS. It is recommended to read Access Control guide before performing any further steps.

### Create the `lab005` Namespace and Use as Default

We will create a new Namespace for every Kubernetes lab and switch contexts to ensure it is the default when using `kubectl`.
```
kubectl create ns lab005 && \
kubectl config set-context $(kubectl config current-context) --namespace lab005 && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Task 1: Launch the Dashboard in GKE

IF YOU DIDN'T ALREADY DO THIS IN LAB 004, CREATE THE CLUSTER-ADMIN ROLEBINDING!

```
# You should have done this already but it won't hurt to run again
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

1. The Dashboard UI is not always deployed by default. To deploy it, run the following command in Cloud Shell:

```
# In the manifests directory
kubectl create -f dashboard.yaml
```

2. Since Kubernetes 1.9, authentication to the dashboard is enabled by default. We need to retrieve a service account token that has the appropriate access. In Cloud Shell run the following command:

Note: The token below is auth token `kubectl` uses itself to authenticate as you. You can use any valid token to authenticate to the dashboard.

Since the token may span multiple lines in the Cloud Shell, a newline may have been added *within* the copied token value - just paste it into a text editor and remove whitespace to ensure the token is on one line, and then copy the result.

```
gcloud config config-helper --format=json | jq -r '.credential.access_token'
```
Copy this value to your clipboard.

3. In Cloud Shell, we need to set up a proxy as follows:
```
kubectl proxy --port 8080
```

4. Use `Web Preview` in Cloud shell to view the dashboard. The URL will look like the following (your domain many vary slightly but the path should be the same):
```
https://8080-dot-4279646-dot-devshell.appspot.com/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/login
```

5. Paste your token from the previous command into the dashboard to authenticate.

6. Take a look around the dashboard. What data can you extract from it? Check out the `Secrets` listed in the namespace `kube-system`.

## Bonus 1
Launch and scale the unshorten-api deployment using only the dashboard.

## Bonus 2
The Service Account that the dashboard uses to launch is located in the `manifests/dashboard.yaml` file. This looks suspicious. Can you restrict the dashboard Service Account?

## Bonus 3
Authenticate to the dashboard using a token that is has a more restricted RBAC policy attached (maybe an intern?). Does the dashboard look any different?

### Task 2: Cleanup
Don't forget to delete the `lab005` namespace when you are done with the Bonuses.
```
kubectl delete ns lab005 && \
kubectl config set-context $(kubectl config current-context) --namespace default && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

## Discussion Question
 Is your Kubernetes dashboard accessible to the internet? What authentication mechanism is enforced?

## Further Reading
 [Heptio - Securing K8S Dashboard](https://blog.heptio.com/on-securing-the-kubernetes-dashboard-16b09b1b7aca)

 [Kubernetes Dashboard Wiki - Access Control](https://github.com/kubernetes/dashboard/wiki/Access-control)
