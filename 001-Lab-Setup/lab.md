# Kubernetes Cluster Setup

The goal of this lab is to successfully setup a Kubernetes cluster using Google Kubernetes Engine (GKE). GKE is a managed Kubernetes service offered in Google Cloud Platform (GCP). We will use this cluster throughout the remainder of the labs.

## Requirements
To minimize software and configuration on your local machine, will be using a service called [Google Cloud Shell](https://cloud.google.com/shell/docs/) which provides you with command-line access to your cloud resources directly from your browser. Cloud Shell supports the latest versions of Google Chrome, Mozilla Firefox, Microsoft Edge, Microsoft Internet Explorer 11+ and Apple Safari 8+. Safari in private browser mode is not supported.

## About GKE
Kubernetes Engine is a managed, production-ready environment for deploying containerized applications in Google Cloud Platform. Each student has been provisioned with a Google account under the `manicode.us` domain as well as a cluster with the following specs:
```
Cluster Size: 4 Nodes
Total Ram: 15GB
Total Cores: 4vCPUs
Location: US-West1-a
```
The cluster master runs the Kubernetes control plane processes, including the Kubernetes API server, scheduler, and core resource controllers. The master's lifecycle is managed by Kubernetes Engine when you create or delete a cluster. This includes upgrades to the Kubernetes version running on the cluster master, which Kubernetes Engine performs automatically, or manually at your request if you prefer to upgrade earlier than the automatic schedule.

## Task 1: Authenticate to Google Cloud Platform
Navigate to the [GCP Console](https://console.cloud.google.com/) and enter the credentials you received on the slip of paper at your desk. You must select your project before beginning. The `Select a Project` link in the upper nav bar will take you to a screen where you can choose your project. You will need to be in the `MANICODE.US` organization then you will be able to select your project. It will look as follows:
```
<org>-<org>123xmanicodexus
```

Note: You may need to refresh the page a few times before seeing your Kubernetes cluster.

## Task 2: Explore Your Pre-Provisioned Kubernetes Cluster
In the navigation on the left side of the console, click `Kubernetes Engine`. Here you will find the details about the cluster and a GUI for accessing and administering workloads and services.

## Task 3: Launch Cloud Shell
There is a button titled `Activate Google Cloud Shell` located in the top-bar navigation of the console. When clicked, a terminal will appear in the lower half of the console. This gives you direct command-line access to your Kubernetes cluster. 

Cloud shell comes packaged with a beta feature called `code editor` which gives you a minimal IDE for viewing and editing files. This will be used throughout the remainder of the labs. The link is found in the upper-right hand corner of the terminal.

## Task 4: Clone the Git Repository
In your home directory, we are going to pull in the documentation and source code used for the course labs. We can do this by running the following command:
```
git clone https://github.com/ManicodeSecurity/Defending-DevOps/ 
```

## Task 5: Explore your Kubernetes Cluster
Most of the tools necessary to complete the labs come pre-installed in Google Cloud Shell including `kubectl` which is used extensively to interact with your cluster. Ensure your cluster is operational by running the following commands. 

First, we need to use `gcloud` to retrieve the credentials from the cluster in order to start interacting with it. 

```
gcloud container clusters get-credentials <YOUR-CLUSTER-NAME> --zone us-west1-a --project <YOUR-PROJECT-NAME>
```

```
kubectl get pods --all-namespaces
kubectl get svc --all-namespaces
```

We are now set up with an operational cluster.