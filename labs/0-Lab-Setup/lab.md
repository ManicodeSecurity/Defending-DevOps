# Kubernetes Cluster Setup

The goal of this lab is to successfully setup a local Kubernetes cluster on your machine using minikube. We will use this cluster throughout the remainder of the labs.

## About Minikube
[Minikube](https://github.com/kubernetes/minikube) is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop. Minikube supports a variety of drivers including:
* virtualbox
* vmwarefusion
* kvm 
* hyperkit

For this lab we will be using Virtualbox, but feel free to experiment with others. Minikube does not offer everything a full cluster would such as provisioning load balancers but it can help us get started with Kubernetes quickly without spinning up costly cloud infrastructure. 

## Task 1: Install Docker
Make sure you have the latest version of [Docker for Mac Community Edition](https://store.docker.com/editions/community/docker-ce-desktop-mac) installed on your local machine. 

## Task 2: Install Kubectl
Kubectl is a command-line tool used to deploy and manage applications on Kubernetes. There are several ways to install kubectl locally. Homebrew and curl are the easiest options for Mac OS.

[Kubectl Installation](https://kubernetes.io/docs/tasks/tools/install-kubectl/
)

Once installed, make sure that the kubectl is available on your machine by running the following command:
```
kubectl version
```

## Task 3: Install Minikube
Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop.

Check out the [latest release](https://github.com/kubernetes/minikube/releases) of Minikube and follow the instructions for downloading and installing.

For Mac OS downloading the binary using `curl` is the easiest option:
```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.25.0/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
```
Feel free to leave off the sudo mv minikube /usr/local/bin if you would like to add Minikube to your path manually.

## Task 4: Launch a Minikube Cluster
Once all of the underlying tools have been installed, it's time to launch our cluster!
```
minikube start
```
Make sure Minikube is running:
```
minikube status
```
Kubectl should now be configured to interact with our newly formed, single-node cluster:
```
kubectl get pods --all-namespaces
```

## Experimental - Docker for Mac: Deploy to Kubernetes
Docker for Mac recently released a "Deploy to Kubernetes" feature which is worth taking a look at. These labs have *not* been tested using these tools so use at your own risk!

[Deploy to Kubernetes for Docker For Mac](https://docs.docker.com/docker-for-mac/kubernetes/)