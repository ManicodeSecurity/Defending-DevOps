# Kubernetes Authentication

The goal of this lab is to enhance the security of our cluster using built in Kubernetes primitaves. We will explore several authentication strategies and apply them to our Minikube cluster.

### Deploying Jenkins to the Cluster
We are now going into full DevOps mode. To build out our pipeline we will need an install of Jenkins. We will use Jenkins to dynamically spin up Pods as slaves in a particular namespace. But first, Jenkins will need a service account to authenticate against the Kubernetes API.

### Task 1: Spin up Jenkins
1. Use the in the `manifests` directory of this lab, you will find a few Jenkins specific .yaml files. Open them up and take a look. The jenkins-deployment.yaml file consists of the following fields:
* The Deployment specifyies a single replica. This ensures one and only one instance will be maintained by the Replication Controller in the event of failure.
* The container image name is jenkins and version is 2.32.2
* The list of ports specified within the spec are a list of ports to expose from the container on the Pods IP address.
Jenkins running on (http) port 8080.
* The Pod exposes the port 8080 of the jenkins container.

2. Create the Deployment by running the following command:
```
kubectl create -f jenkins-deployment.yaml
#Validate the Deployment was created
kubectl get deployments
```

3. We need to expose the Jenkins UI in a similar way to our Unshorten API. Instead of using a command, let's build the Service using a manifest:
```
kubectl create -f jenkins-service.yaml
```

4. We now can access the Jenkins UI using the Minikube IP:
```
minikube service jenkins --url
```

5. Before you put this in your browser, Jenkins is going to need a master password which is provided in the logs as well as in a file on the OS. We can easily grab this value by using the `logs` command:
```
kubectl logs <PodName> | grep -B 3 initialAdminPassword
```

6. Use the IP to access the Jenkins UI in your browser and set up Jenkins. Do not add any plugins yet.


## Exercise 2: Authenticating Jenkins to your Cluster


### Bonus - The version of Jenkins you deployed has plenty of known vulnerabilities. Update your Jenkins version to the latest (2.89.4 at the time of this writing). 
