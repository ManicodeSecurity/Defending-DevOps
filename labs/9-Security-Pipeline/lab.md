# Security Pipeline and Automation
This lab will spin up Jenkins in our cluster along with a private Docker image repository. We will also kick off a build that runs a vulnerability scan using Clair.

First, we blow away our old cluster and launch a fresh one:
```
minikube delete
# Jenkins is ram-hungry so let's give our cluster a little extra juice
minikube start --memory 4096
```
## Task 1: Build the Internal Registry

1. In the `manifests/registry` directory run the following:
```
kubectl create -f .
```

2. Once all of the Pods and Services are up and healthy, grab the URL for our freshly created registry and visit it in your browser:
```
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
minikube service jenkins --url
```

6. DO NOT INSTALL SUGGESTED PLUGINS! Click manually install plugins and deselect all of them. Only select two plugins from the web UI - `git` and `pipeline`. Search for them.

## Task 2: Build our Pipeline

1. In the Jenkins UI, click `New Item` and select `Pipeline` as the project type. Click `Ok`.

2. You should be redirected to the `General` tab for the project. Under `Pipeline` click `Definition` and select `Pipeline Script from SCM`. This tells the Pipeline to pull from a source code repository.

3. Select `Git` in the SCM dropdown

4. Enter the following repo URL:
paste repo

Run clair

Run static code analysis tool!

