# Kubernetes Secrets

The goal of this lab is to utilize the native Kubernetes Secrets functionality to create and consume Kubernetes secrets within our application.

### Create the `lab009` Namespace and Use as Default

We will create a new Namespace for every lab and switch contexts to ensure it is the default when using `kubectl`.
```
kubectl create ns lab009 && \
kubectl config set-context $(kubectl config current-context) --namespace lab009 && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Task 1: Create a Mysql Deployment and Service (the insecure way)

1. We now must create our Secret using kubectl. This will allow our API to communicate with MySQL.
```
kubectl create secret generic mysql-secrets --from-literal=password=supertopsecretpassword
```

2. Inspect out the newly created Secret object:
```
kubectl get secrets
kubectl describe secret mysql-secrets
kubectl get secret mysql-secrets -o yaml
kubectl get secret mysql-secrets -o jsonpath='{.data.password}' | base64 --decode; echo
```
3. After inspecting the MySQL manifest files located in the `manifests/mysql` directory, launch both the Deployment and the Service using the following command:
```
kubectl create -f .
```

4. launch the API Deployment and Service in the `manifests/api` directory. This is a new image version which also handles the creation of all necessary MySQL databases and tables.
```
kubectl create -f .
```

5. Ensure our API is running and everything is wired up:
```
kubectl get svc
```
In your browser visit the IP address provided at `EXTERNAL-IP`. For example:
```
http://35.197.56.177/api/check?url=bit.ly/test
```
Note: The order in which these pods are deployed matters. Make sure MySQL is up first before launching the API.

6. After everything is running and healthy, let's look under the hood at the environment variable that was injected into the link-unshorten Pod:
```
kubectl exec -it <link-unshorten-podname> /bin/bash
# Once you have a shell in the container run the following
env | grep UNSHORTEN_DB_PASSWORD
```

### Task 2: Using Manifests to Deploy Secrets (the more secure way)
Sticking to DevOps principals, we want to avoid creating secrets using one-off commands. We will now create our MySQL password using a YAML manifest located in the `manifests/secrets` directory.

1. To simplify things, we first tear down our cluster. Run the following command in both the `manifests/api` and `manifests/mysql` directories to delete the running Deployments and Services:
```
# In the manifests/api directory
kubectl delete -f .
# In the manifests/mysql directory
kubectl delete -f .
# Delete our existing Secret
kubectl delete secret mysql-secrets
```

2. Take a look at the manifest located at `manifests/secrets/mysql-secrets.yaml`. While it is very possible to put our Base64 encoded MySQL password in this file and deploy it, that is not the responsible mechanism for storing sensitive strings such as passwords for a variety of reasons. This lab will rely on a local environment variable and a command to perform variable substitution into our manifest file. This is just one way to launch secrets and may not be the best for your particular pipeline.

3. Create a Base64 representation of our password and store it in an environment variable in Cloud Shell:
```
echo -n "supertopsecretpassword" | base64
export MYSQL_PASSWORD="c3VwZXJ0b3BzZWNyZXRwYXNzd29yZA=="
```

4. In the `manifests/secrets` directory, we `cat` our YAML file and pipe it to sed, substituting the `$MYSQL_PASSWORD` placeholder with the value stored in our local environment variable. This is then piped into our `kubectl create -f` command:
```
# sed ahead!
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
# In manifests/mysql (give MySQL a few minutes to start up before running the next command)
kubectl create -f .
# In manifests/api
kubectl create -f .
```

Everything should now be up and running.

### Task 3: Deploying Vault in our Cluster
HashiCorp Vault secures, stores, and tightly controls access to tokens, passwords, certificates, API keys, and other secrets in modern computing. Vault handles leasing, key revocation, key rolling, and auditing. Through a unified API, users can access an encrypted Key/Value store and network encryption-as-a-service, or generate AWS IAM/STS credentials, SQL/NoSQL databases, X.509 certificates, SSH credentials, and more.

We can use Vault in our own Kubernetes to store and retrieve a variety of secrets.

1. Inspect the files in the `manifests/vault` directory. You will see a new type of Kubernetes API object called a StatefulSet. StatefulSets are similar to Deployments with the caveat that Pods managed by StatefulSets offer a stable network identity and state. Vault is a good candidate for a StatefulSet for these reasons. This Vault cluster is deployed in Development mode and should NOT be used as-is in production.
```
kubectl create -f .
kubectl get pods
# notice the name of the Vault pod. Because it is managed by a StatefulSet it will always be named vault-0.
kubectl get svc
```

2. In order to use the Vault service from our Cloud Shell instance, we first forward a local port to the Vault port (8200) using `kubectl port-forward`.
```
kubectl port-forward vault-0 8200:8200 &
# This will open a background process
```

3. We will use `curl` to interact with our newly created Vault cluster. The following commands will write secrets to the cluster:
```
# Run the following command in a separate terminal window or tab
curl \
    -H "X-Vault-Token: not-intended-for-production-deployments" \
    http://127.0.0.1:8200/v1/sys/health
```

4. Now write our MySQL secrets to Vault using the API:
```
curl \
    -H "X-Vault-Token: not-intended-for-production-deployments" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '{"password":"vaultftw!"}' \
    http://127.0.0.1:8200/v1/secret/mysql
```

5. List all secrets in Vault:
```
curl \
    -H "X-Vault-Token: not-intended-for-production-deployments" \
    -X LIST \
    http://127.0.0.1:8200/v1/secret
```

6. Retrieve our MySQL Secrets from the API:
```
curl \
    -H "X-Vault-Token: not-intended-for-production-deployments" \
    http://127.0.0.1:8200/v1/secret/mysql
```

### Task 4: Using Vault to Store and inject our MySQL Password

We can now call the Vault API to inject our secret into our `kubectl create` command on the fly as follows.

First, delete `mysql-secrets` from our cluster:
```
kubectl delete secret mysql-secrets
```

navigate to `manifests/secrets` and run:
```
vault_mysql_pass=`curl -H "X-Vault-Token: not-intended-for-production-deployments" http://127.0.0.1:8200/v1/secret/mysql | jq -r '.data.password' | base64`; cat mysql-secrets.yaml | sed "s/\$\$MYSQL_PASSWORD/$vault_mysql_pass/" | kubectl create -f -
```
### Cleanup

Don't forget to delete the `lab009` namespace when you are done with the Bonuses.
```
kubectl delete ns lab009 && \
kubectl config set-context $(kubectl config current-context) --namespace default && \
echo "Default Namespace Switched:" $(kubectl get sa default -o jsonpath='{.metadata.namespace}')
```

### Discussion Question
What secrets management systems are you using in-house? How could they better plug into DevOps pipelines?