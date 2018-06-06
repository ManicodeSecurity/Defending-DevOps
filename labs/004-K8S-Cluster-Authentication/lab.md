# Kubernetes Authentication

The goal of this lab is to enhance the security of our cluster using built in Kubernetes primitives. We will explore several authentication strategies and apply them to our Minikube cluster.

### Simple Authentication
First, we delete any dangling clusters and create a new one using the `extra-config` parameter which tells the api-server to use our CSV file as a username/password store and RBAC as our authorization mechanism. 

1. This is *not* the most secure method of authentication for a variety of reasons but demonstrates authentication in K8S. Replace the path to `BasicAuthFile` with your own before running the `minikube start` command. 

## You must change the `/path/to` part of the following command to your own absolute path! Things will break if you do not change it.
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube delete
minikube start --extra-config=apiserver.Authorization.Mode=RBAC --extra-config=apiserver.Authentication.PasswordFile.BasicAuthFile=/path/to/Defending-DevOps/labs/004-K8S-Cluster-Authentication/creds.csv
```

2. We have been using `kubectl` for these labs so far, but the Kubernetes API is also accessible using standard REST endpoints on port 8443. If we try to `curl` our API endpoint we will be denied because RBAC is enabled and we did not pass our credentials to the API:
```
export MINIKUBE_HOME=~/Desktop/lab-tools/.kube
minikube ip
curl https://$(minikube ip):8443/ -k
# DENIED
```

3. Now we can spin up our environment along with our Roles and RoleBinding for the new user. First, launch the Deployment, Service, and Namespace located in the `manifests` directory:
```
kubectl create -f .
```

4. Now, we can prove that RBAC is working if we try to list the Secrets in the Development Namespace using valid credentials (this request will be denied because we have not yet created our RBAC rules):
```
echo -n jboss:supertopsecretpassword | base64
# Use this base64 encoded value in our Basic HTTP header below
curl -H "Authorization: Basic amJvc3M6c3VwZXJ0b3BzZWNyZXRwYXNzd29yZA==" https://$(minikube ip):8443/api/v1/namespaces/development/secrets -k
```

5. Access should be denied for our user jboss. 

Poor jboss.

### Applying RBAC
1. We will now create a rule that explicitly allows the user jboss to list secrets in the development namespace and *only* that namespace. In the in the `manifests/role` directory, run the following commands:
```
kubectl create -f .
kubectl describe rolebinding read-secrets-development --namespace=development
```

2. We can now try our `curl` command again and with any luck, jboss will be able to read the secrets in the development namespace:
```
curl -H "Authorization: Basic amJvc3M6c3VwZXJ0b3BzZWNyZXRwYXNzd29yZA==" https://$(minikube ip):8443/api/v1/namespaces/development/secrets -k
```
