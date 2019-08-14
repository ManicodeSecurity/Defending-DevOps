# Security Pipeline and Automation
This lab will spin up Jenkins in our cluster along with a private Docker image repository. Jenkins will also handle zero-downtime deploys of the unshorten API upon a successful build. The humble beginnings of a self-contained DevSecOps pipeline.

### Create the `lab010` Namespace and Use as Default

We will create a new Namespace for every lab and switch contexts to ensure it is the default when using `kubectl`.
```
kubectl create ns lab010 && \
kubectl config set-context $(kubectl config current-context) --namespace lab010 && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Task 1: Create the Jenkins Service Account
Jenkins does not need full administrative access in our cluster. It is crucial to implement RBAC policies that allow Jenkins to carry out the necessary tasks but nothing more. Take a look at the Jenkins Service Account and associated ClusterRole and ClusterRoleBinding then create the SA as follows:

First we need to make sure that our user is `cluster-admin`
```
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

1. In the `manifests/service-account` directory run the following:

```
kubectl create -f .
```

2. Verify that the ClusterRole was created:
```
 kubectl describe clusterrole jenkins-limited
 ```

### Task 2: Build the Internal Registry and Launch Jenkins
We need a location to store our versioned Docker images within our Kubernetes cluster. This deployment uses the official Registry image from Docker.

1. In the `manifests/registry` directory run the following:
```
kubectl create -f .
```

2. Once all of the Pods and Services are up and healthy, grab the URL for our freshly created registry and visit it in your browser.

Note: The registry runs on port `8080`.
```
kubectl get svc
# Copy the EXTERNAL-IP value and paste it into your browser
# It will look like this -> http://35.199.183.47:8080/
```

3. Now we deploy Jenkins from the `manifests/jenkins` directory:
```
kubectl create -f .
```

4. Make sure the Jenkins Deployment is up and grab the logs from the Pod:
```
# Grab the pod name and place it in an environment variable
export POD_NAME=$(kubectl get pods -l "app=jenkins" -o jsonpath="{.items[0].metadata.name}")
# Copy the initial Admin password provided by Jenkins
kubectl logs $POD_NAME | grep -B 3 initialAdminPassword
```

5. Now we can check out the Jenkins UI (it also runs on port `8080`):
```
# Run a port forward to expose the Jenkins UI
kubectl port-forward $POD_NAME 8080:8080 >> /dev/null &
```
(!)Use the Web View in Cloud Shell to now visit Jenkins.

6. DO NOT INSTALL SUGGESTED PLUGINS! Click deselect all, and manually install the plugins needed for this lab. Only select two plugins from the web UI; `git` and `pipeline`. Use the search feature to find them.

### Task 3: Build our Pipeline

1. In the Jenkins UI, click `New Item` and select `Pipeline` as the project type. Click `Ok`.

2. You should be redirected to the `General` tab for the project. Under `Pipeline` click `Definition` and select `Pipeline Script from SCM`. This tells the Pipeline to pull from a source code repository.

3. Select `Git` in the SCM dropdown.

4. Enter the following repo URL. This repo contains our unshorten-api source code and a few other files that help construct our CI/CD build.
```
https://github.com/ManicodeSecurity/unshorten-jenkins-demo
```

5. Inspect the [`Jenkinsfile`](https://github.com/ManicodeSecurity/unshorten-jenkins-demo/blob/master/Jenkinsfile) in the repo. It has the humble beginnings of an AppSec and DevSecOps pipeline. Each stage is meant to apply automation to the process where issues result in failed builds.

### Task 4: Trigger a Build
Most pipeline setups will trigger builds on a git commit or through some other automated manner. To simulate this, we will tell Jenkins to trigger a build manually:

1. From the Jenkins Dashboard, click on our project name in the table.

2. In the navigation on the left-hand side, click `Build Now`.

3. Click on the actual build and inspect the `Console Output`. You will see each step of the build here running in Jenkins live.

### The Build Broke?!
Use what we have learned so far to debug and fix the issue to run Jenkins with a clean build.

### What Just Happened?

Jenkins is taking the place of us running `kubectl` and `docker` commands locally. Through automation, we were able to build a Docker image, tag it appropriately, and use the Kubernetes rollout feature to ensure a zero-downtime deploy to our cluster. The Pipeline is very bare-bones but can be augmented with your favorite scanning tools for static analysis, dynamic testing, and container vulnerability scanning. This is the beginnings of DevSecOps.

### Cleanup

Don't forget to delete the `lab010` namespace when you are done with the Bonuses.
```
kubectl delete ns lab010 && \
kubectl config set-context $(kubectl config current-context) --namespace default && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Further Reading
Check out the popular build workflow tool called [Skaffold](https://github.com/GoogleContainerTools/skaffold).
