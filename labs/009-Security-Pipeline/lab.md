# Security Pipeline and Automation
This lab will spin up Jenkins in our cluster along with a private Docker image repository. Jenkins will also handle zero-downtime deploys of the unshorten API upon a successful build. The humble beginnings of a self-contained DevSecOps pipeline. 

First, we blow away our old cluster and launch a fresh one:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube delete
# Jenkins is a ram-hungry beast, so let's give our cluster a little extra juice!
minikube start --memory 4096 --bootstrapper=kubeadm
```

## Task 1: Create the Jenkins Service Account
Jenkins does not need full administrative access in our cluster. It is crucial to implement RBAC policies that allow Jenkins to carry out the necessary tasks but nothing more. Take a look at the Jenkins Service Account and associated ClusterRole and ClusterRoleBinding then create the SA as follows:

1. In the `manifests/service-account` directory run the following:

```
kubectl create -f .
```

2. Verify that the ClusterRole was created:
```
 kubectl describe clusterrole jenkins-limited
 ```

## Task 2: Build the Internal Registry

1. In the `manifests/registry` directory run the following:
```
kubectl create -f .
```

2. Once all of the Pods and Services are up and healthy, grab the URL for our freshly created registry and visit it in your browser:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube service registry-ui --url
```

3. Now we deploy Jenkins from the `manifests/jenkins` directory:
```
kubectl create -f .
```

4. Make sure the Jenkins Deployment is up and grab the logs from the Pod:
```
kubectl get pods
kubectl logs <jenkinsPodName> | grep -B 3 initialAdminPassword
```

5. Now we can check out the Jenkins UI:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube service jenkins --url
```

6. DO NOT INSTALL SUGGESTED PLUGINS! Click deselect all, and manually install the plugins needed for this lab. Only select two plugins from the web UI; `git` and `pipeline`. Use the search feature to find them.


## Task 3: Build our Pipeline

1. In the Jenkins UI, click `New Item` and select `Pipeline` as the project type. Click `Ok`.

2. You should be redirected to the `General` tab for the project. Under `Pipeline` click `Definition` and select `Pipeline Script from SCM`. This tells the Pipeline to pull from a source code repository.

3. Select `Git` in the SCM dropdown.

4. Enter the following repo URL. This repo contains our unshorten-api source code and a few other files that help construct our CI/CD build.
```
https://github.com/ManicodeSecurity/unshorten-jenkins-demo
```

5. Inspect the `Jenkinsfile` in the repo. It has the humble beginnings of an AppSec and DevSecOps pipeline. Each stage is meant to apply automation to the process where issues result in failed builds. 

## Task 4: Trigger a Build
Most pipeline setups will trigger builds on a git commit or through some other automated manner. To simulate this, we will tell Jenkins to trigger a build manually:

1. From the Jenkins Dashboard, click on our project name in the table.

2. In the navigation on the left-hand side, click `Build Now`.

3. Click on the actual build and inspect the `Console Output`. You will see each step of the build here running in Jenkins live.

## The Build Broke?!
Use what we have learned so far to debug and fix the issue to run Jenkins with a clean build.

## What Just Happened?

Jenkins is taking the place of us running `kubectl` and `docker` commands locally. Through automation, we were able to build a Docker image, tag it appropriately, and use the Kubernetes rollout feature to ensure a zero-downtime deploy to our cluster. The Pipeline is very bare-bones but can be augmented with your favorite scanning tools for static analysis, dynamic testing, and container vulnerability scanning. This is the beginnings of DevSecOps.
