minikube start --extra-config=apiserver.Authorization.Mode=RBAC

https://docs.giantswarm.io/guides/securing-with-rbac-and-psp/
https://www.slideshare.net/Opcito/kubernetesrbacprometheusmonitoring-170626103349-77262683

## Exercise 1: Create a User with Limited Namespace Access

### Step 1: Create the Developer Namespace
kubectl create namespace dev

### Step 2: Create the Credentials
In the certs folder we have provided a certificate for a developer called *dev1.crt*. This certificate was generate using OpenSSL. 

### Step 3: Add a new context to your Kubernetes cluster
Locate your Kubernetes cluster certificate authority (CA). This will be responsible for approving the request and generating the necessary certificate to access the cluster API. Its location is normally /etc/kubernetes/pki/. In the case of Minikube, it would be ~/.minikube/. Check that the files ca.crt and ca.key exist in the location.

We need to tell our cluster about the newly created user by running the following commands in kubectl:

kubectl config set-credentials dev1 --client-certificate=/certs/dev1.crt  --client-key=/certs/dev1.key

kubectl config set-context developer-context --cluster=minikube --namespace=dev --user=dev1

Now we try to perform an action on the dev namespace using our newly create developers-context

kubectl --context=developers-context get pods

This fails because our
### Step 4: Create a Developer Role



https://docs.bitnami.com/kubernetes/how-to/configure-rbac-in-your-kubernetes-cluster/