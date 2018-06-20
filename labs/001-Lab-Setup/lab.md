# Kubernetes Cluster Setup

The goal of this lab is to successfully setup a local Kubernetes cluster on your machine using Minikube. We will use this cluster throughout the remainder of the labs.

## Requirements
Please ensure you have an approved version of Virtualbox installed on your laptop:

[Virtualbox Download](https://www.virtualbox.org/wiki/Downloads)

Once downloaded, follow the instructions to install Virtualbox.

## About Minikube
[Minikube](https://storage.googleapis.com/minikube/releases/v0.25.2/minikube-darwin-amd64) is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop. Please use version 0.25.2 for the labs.

## Task 1: Install Docker
Make sure you have the latest version of [Docker for Mac Community Edition](https://store.docker.com/editions/community/docker-ce-desktop-mac) installed on your local machine. 

Once downloaded, follow the instructions to install Docker.

## Task 2: Create a lab-tools directory on your Desktop
We will be downloading a few binaries to spin up and interact with our Kubernetes cluster throughout the labs. Please *only* download binaries from the approved Oath artifact repository and into a `lab-tools` directory that we will later destroy upon completion of the class:
```
cd ~/Desktop
mkdir lab-tools && cd lab-tools
```

## Task 3: Install kubectl
`kubectl` is a command-line tool used to deploy and manage applications on Kubernetes. There are several ways to install kubectl locally. We will be using `curl` to install the binary from the Oath artifact repo.

```
cd ~/Desktop/lab-tools
curl -o kubectl https://storage.googleapis.com/kubernetes-release/release/v1.9.0/bin/darwin/amd64/kubectl
chmod +x ./kubectl
# This creates a symlink to our PATH so that we can use the kubectl command in other locations throughout the lab
sudo ln -s ~/Desktop/lab-tools/kubectl /usr/local/bin/kubectl
```

Once installed, make sure that the kubectl is available on your machine by running the following command:
```
kubectl version
```

## Task 4: Install Minikube
Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop.

Download the binary using `curl` from the official Oath artifactory repo:
```
cd ~/Desktop/lab-tools
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.25.2/minikube-darwin-amd64 +x minikube
# This creates a symlink to our PATH so that we can use the minikube command in other locations throughout the lab
sudo ln -s ~/Desktop/lab-tools/minikube /usr/local/bin/minikube
```

## Task 5: Launch a Minikube Cluster
Once all of the underlying tools have been installed, it's time to launch our cluster!

We first need to tell Minikube where to put our Kubernetes configuration files. This is done be setting an environment variable. DO NOT SKIP THIS STEP!
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
```

Now, we bootstrap the cluster:
```
# from the lab-tools directory
minikube start
# Grab a coffee, this bootstrap process may take a few minutes
```

Ensure the configs are located in the right location:
```
cd ~/Desktop/lab-tools
ls -a
# you should see the following two binaries and the .kube directory
.kube    kubectl  minikube
# Feel free to browse around the .kube directory and inspect the Kubernetes configurations
```

Make sure Minikube is running:
```
minikube status
```

Kubectl should now be configured to interact with our newly formed, single-node cluster:
```
kubectl get pods --all-namespaces
```

# Important! All binaries should be installed in the ~/Desktop/lab-tools directory only