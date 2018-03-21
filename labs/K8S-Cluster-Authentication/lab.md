# Kubernetes Authentication

The goal of this lab is to enhance the security of our cluster using built in Kubernetes primitives. We will explore several authentication strategies and apply them to our Minikube cluster.

###
First, we delete any dangling clusters and create a new one using the `extra-config` parameter which tells the api-server to use our CSV file as a username/password store. 

This is not the most secure method of authentication for a variety of reasons. 
```
minikube delete
minikube start --extra-config=apiserver.Authorization.Mode=RBAC --extra-config=apiserver.Authentication.PasswordFile.BasicAuthFile=/path/to/Defending-DevOps-Training/labs/K8S-Cluster-Authentication/creds.csv
```

Now we try to access the API endpoint without supplying our credentials:
```
curl  https://<minikubeIP>:8443/ -k
```

If we `base64` our username and password and pass it to the API endpoint as an `Authorization: Basic` HTTP header we should be able to authenticate:
```
echo -n userid:supertopsecretpassword | base64
```
Pass the value from the above command to the `curl` command using the `-H` flag:
```
curl -H "Authorization: Basic bXl1c2VyOnN1cGVydG9wc2VjcmV0cGFzc3dvcmQ=" https://192.168.99.100:8443/ -k
```
If we place an an invalid password in our `base64` value you will see an `401 Unauthorized` response from the API:
```
echo -n userid:hackersgunnahack | base64

curl -H "Authorization: Basic dXNlcmlkOmhhY2tlcnNndW5uYWhhY2s=" https://192.168.99.100:8443/ -k
```

### Applying RBAC
First, let's deploy a Deployment and Service in our cluster and access it using the API directly instead of `kubectl`:
```
# in the manifests directory
kubectl create -f .

# list the secrets using the API
curl -H "Authorization: Basic bXl1c2VyOnN1cGVydG9wc2VjcmV0cGFzc3dvcmQ=" https://192.168.99.100:8443/api/v1/namespaces/default/secrets -k
```
Now, we can create a `Role` to stop this particular user from reading secrets:
```
# in the manifests/role directory 
kubectl create -f .





### Bonus+++++: Implement OIDC authentication using Google on your Minikube cluster