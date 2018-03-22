### Exploring the Kubernetes Dashboard
Kubernetes Dashboard is a general purpose, web-based UI for Kubernetes clusters. It allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

Since version 1.7 Dashboard uses more secure setup. It means, that by default it has minimal set of privileges and can only be accessed over HTTPS. It is recommended to read Access Control guide before performing any further steps.

## Launch the Dashboard in Minikube
1. Destroy our old cluster and create a fresh one:
```
minikube delete
#kubeadm starts a cluster with the dashboard enabled by default
minikube start --bootstrapper kubeadm
```

2. Use the `minikube` command to view the URL for the dashboard:
```
minikube dashboard --url
```

3. Paste this URL from the above command into your web browser
4. Take a look around the dashboard. What data can you extract from it?

## Bonus: Launch the unshorten-api deployment using only the dashboard

## Discussion Question: Is your Kubernetes dashboard accessible to the internet? What authentication mechanism is enforced?

## Further Reading: [Heptio - Securing K8S Dashboard](https://blog.heptio.com/on-securing-the-kubernetes-dashboard-16b09b1b7aca)