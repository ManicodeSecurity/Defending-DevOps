### Exploring the Kubernetes Dashboard
Kubernetes Dashboard is a general purpose, web-based UI for Kubernetes clusters. It allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

Since version 1.7 Dashboard uses more secure setup. It means, that by default it has minimal set of privileges and can only be accessed over HTTPS. It is recommended to read Access Control guide before performing any further steps.

## Launch the Dashboard in GKE
1. The Dashboard UI is not deployed by default. To deploy it, run the following command in Cloud Shell:

```
# In the manifests directory
kubectl apply -f dashboard.yaml
```

2. Since Kubernetes 1.9, authentication to the dashboard is enabled by default. We need to retrieve a service account token that has the appropriate access. In Cloud Shell run the following command:

Note: All secrets in the `kube-system` namespace have full access. The `clusterrole-aggregation-controller` is one of those but others will work.

```
kubectl -n kube-system describe secrets `kubectl -n kube-system get secrets | awk '/clusterrole-aggregation-controller/ {print $1}'` | awk '/token:/ {print $2}'
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

6. Take a look around the dashboard. What data can you extract from it?

## Bonus: Launch the unshorten-api deployment using only the dashboard

## Discussion Question: Is your Kubernetes dashboard accessible to the internet? What authentication mechanism is enforced?

## Further Reading: [Heptio - Securing K8S Dashboard](https://blog.heptio.com/on-securing-the-kubernetes-dashboard-16b09b1b7aca)