# Kubernetes Cluster Setup

The goal of this lab is to successfully setup a local Kubernetes cluster on your machine using Minikube. We will use this cluster throughout the remainder of the labs.

## About Minikube
[Minikube](https://github.com/kubernetes/minikube) is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop. Minikube supports a variety of drivers including:
* virtualbox
* vmwarefusion
* kvm 
* hyperkit

We will be using Virtualbox for all of the labs.

## Task 1: Install Docker
Make sure you have the latest version of [Docker for Mac Community Edition](https://store.docker.com/editions/community/docker-ce-desktop-mac) installed on your local machine. 

## Task 2: Install kubectl
`kubectl` is a command-line tool used to deploy and manage applications on Kubernetes. There are several ways to install kubectl locally. Homebrew and curl are the easiest options for Mac OS.

[Kubectl Installation](https://kubernetes.io/docs/tasks/tools/install-kubectl/
)

Install using `curl`
```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/darwin/amd64/kubectl
chmod +x ./kubectl
# !! Change this to a tools directory !!
sudo mv ./kubectl /usr/local/bin/kubectl
```

Once installed, make sure that the kubectl is available on your machine by running the following command:
```
kubectl version
```

## Task 3: Install Minikube
Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop.

Check out the [latest release](https://github.com/kubernetes/minikube/releases) of Minikube and follow the instructions for downloading and installing.

For Mac OS downloading the binary using `curl` is the easiest option:
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.25.0/minikube-darwin-amd64 && chmod +x minikube && 
# !! Change this to a tools directory !!
sudo mv minikube /usr/local/bin/
```
Feel free to leave off the sudo mv minikube /usr/local/bin if you would like to add Minikube to your path manually.

## Task 4: Launch a Minikube Cluster
Once all of the underlying tools have been installed, it's time to launch our cluster!

We first need to tell Minikube where to put our Kubernetes configuration files. This is done be setting an environment variable.
```
# !! Do we want everything to go on the Desktop? !!
export MINIKUBE_HOME=~/Desktop/.kube
```

Now, we launch the cluster:
```
# !! From the new tools directory !!
./minikube start
```

Make sure Minikube is running:
```
./minikube status
```

Kubectl should now be configured to interact with our newly formed, single-node cluster:
```

# !! From the new tools directory !!
./kubectl get pods --all-namespaces
```