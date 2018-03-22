# Kubernetes Secrets

The goal of this lab is to utilize the native Kubernetes Secrets functionality to create and consume Kubernetes secrets within our application.

## About Secrets


## Task 1: Create a Mysql Deployment and Service

1. We first must create our Secret using kubectl. This will allow our API to communicate with MySQL. 
```
kubectl create secret generic mysql-secrets --from-literal=password=supertopsecretpassword
```

2. Inspect out the newly created Secret object:
```
kubectl get secrets
kubectl describe secret mysql-secrets
kubectl get secret mysql-secrets -o yaml
```
3. After inspecting the MySQL manifest files located in the `manifests/mysql` directory, launch both the Deployment and the Service using the following command:
```
kubectl create -f .
```

4. launch the API Deployment and Service in the `manifests/api` directory`. This is a new image version which also handles the creation of all necessary MySQL databases and tables.
```
kubectl create -f .
```

5. Ensure our API is running and everything is wired up:
```
minikube service link-unshorten-service --url
```

6. After everything is running and healthy, let's look under the hood at the secret that injected into an environment variable by Kubernetes:
```
kubectl exec -it <podname> /bin/bash
env | grep MYSQL_DB_PASSWORD
```

## Task 2: Using Manifests to Deploy Secrets
Sticking to IaaS principals, we want to avoid creating secrets using the command line. We will now create our MySQL password using a YAML manifest located in the `manifests/secrets` directory.

1. To simplify things, we first tear down our cluster. Run the following command in both the `manifests/api` and `manifests/mysql` directories to delete the running Deployments and Services:
```
kubectl delete -f .
# Delete our existing Secret
kubectl delete secret mysql-secrets
```

2. Take a look at the manifest located at `manifests/mysql-secrets.yaml`. While it is very possible to put our Base64 encoded MySQL password in this file and deploy it, that is not the responsible mechanism for storing sensitive strings such as passwords for a variety of reasons. This lab will rely on a local environment variable and a command to perform variable substitution into our manifest file. 

NOTE: This is just one way to launch secrets and may not be the best for your particular pipeline. Please don't judge the use of `sed` : ) 

3. Create a Base64 representation of our password and store it in a local environment variable:
```
echo -n "supertopsecretpassword" | base64
export MYSQL_PASSWORD="c3VwZXJ0b3BzZWNyZXRwYXNzd29yZA=="
```
4. In the `manifests/secrets` directory, we `cat` our YAML file and pipe it to sed, substituting the `$MYSQL_PASSWORD` placeholder with the value stored in our local environment variable. This is then piped into our `kubectl create -f` command:
```
# hacks ahead!
cat mysql-secrets.yaml | sed s/\$\$MYSQL_PASSWORD/$MYSQL_PASSWORD/ | kubectl create -f -
```

5. Now our Secret is created. We can inspect it just as before:
```
kubectl get secrets
kubectl describe secret mysql-secrets
kubectl get secret mysql-secrets -o yaml
```

6. As before, we launch our MySQL components and the API:
```
# In manifests/mysql
kubectl create -f .
# In manifests/api
kubectl create -f .
```

Everything should now be up and running.

## Deploying Vault in our Cluster
HashiCorp Vault secures, stores, and tightly controls access to tokens, passwords, certificates, API keys, and other secrets in modern computing. Vault handles leasing, key revocation, key rolling, and auditing. Through a unified API, users can access an encrypted Key/Value store and network encryption-as-a-service, or generate AWS IAM/STS credentials, SQL/NoSQL databases, X.509 certificates, SSH credentials, and more.

We can use Vault in our own Kubernetes to store and retrieve a variety of secrets.

1. Inspect the files in the `manifests/vault` directory. You will see a new type of Kubernetes API object called a StatefulSet. StatefulSets are similar to Deployments with the caveat that Pods managed by StatefulSets offer a stable network identity and state. Vault is a good candidate for a StatefulSet for these reasons. This Vault cluster is deployed in Development mode and should NOT be used as-is in production.
```
kubectl create -f .
kubectl get pods
# notice the name of the Vault pod. Because it is managed by a StatefulSet it will always be named vault-0.
kubectl get svc
```
2. In order to use the Vault service from our local machine, we first forward a local port to the Vault port (8200) using `kubectl port-forward`.
```
kubectl port-forward vault-0 8200:8200
```
3. We will need the Vault CLI installed locally in order to interact with our Vault cluster. Download the appropriate Vault binary.

[Vault Project Download Page](https://www.vaultproject.io/downloads.html)

 4. Now we export two environment variables the Vault CLI needs in order to communicate with our Vault deployment:
```
export VAULT_ADDR="http://127.0.0.1:8200"
export VAULT_TOKEN="not-intended-for-production-deployments"
```
5. We can now interact with Vault running in our Kubernetes cluster:
```
vault status
# This should return information about our Vault cluster

vault write secret/test \
  username=admin \
  password=vaultftw!

vault read secret/test
```
# Bonus+: ## Using Vault to Store and our MySQL Password

Hint: Moar Sed!
```
vault_mysql_pass=`vault read -field=password secret/test | base64`; cat mysql-secrets.yaml | sed "s/\$\$MYSQL_PASSWORD/$vault_mysql_pass/" | kubectl create -f -
```

## Discussion Question: What secrets management systems are you using in-house? How could they better plug into DevOps pipelines?



 REMOVE ! ? 
## Bonus+: Encrypting Secrets at Rest

Hint:
minikube delete
minikube start --bootstrapper=kubeadm --mount --mount-string="/Users/jb0ss/Desktop/Defending-DevOps-Training/labs/5-K8S-Cluster-Secrets/manifests/bonus":/encryption-config


