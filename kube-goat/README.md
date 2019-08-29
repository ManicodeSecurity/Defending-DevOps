
# Kube-Goat

Your turn to hack a cluster. You have come across a kubeconfig file on a public Github repo. What can you do?

## The instructor will provide a link to the kubeconfig file for you to download.

In Google Cloud Shell, we will use the kubeconfig for the kube-goat cluster. Copy the kubeconfig and paste it in Cloud Shell at `~/.kube/goat-config` (You will need to create the `.kube` directory).

Now, we need to tell `kubectl` to use the new kubeconfig file:
```
export KUBECONFIG=~/.kube/goat-config
```

### Attack Categories
- Sensitive Volume Mount
- Open Dashboard
- Revealing Secrets
- Sensitive Data Exposure
- API Authentication Issues
