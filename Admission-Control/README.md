# Admission Control
Admission plugins can be developed as extensions and run as webhooks configured at runtime. In this lab we will build admission webhooks that help improve the security of your cluster. 

Admission webhooks are HTTP callbacks that receive admission requests and do something with them. In this lab we will be checking for two common conditions that can introduce security issues:

- Denying Ingress object types 
- Denying public IP Loadbalancers

The webhooks are hosted using [Cloud Run](https://cloud.google.com/run/) on GCP. The `Dockerfile` used on Cloud Run is located at `webhook/CloudRun.Dockerfile`.

## Block Ingress Objects

An Ingress object in Kubernetes may expose unnecessary external IP addresses to the internet. This can lead to unexpected compromise. 

First, inspect the `ValidatingWebhookConfiguration` located at `deny-ingress/webhook-configuration.yaml`. You will see that we are using an open API endpoint hosted on Cloud Run to inspect the request. 

Create the `ValidatingWebhookConfiguration` object:
```
#in the deny-ingress directory
kubectl create -f webhook-config.yaml
```

Now, to test that the Admission Control is working properly, create a simple Ingress object as follows:

```
kubectl create -f ingress.yaml
```

Due to the rules built in the [DenyIngresses](https://github.com/elithrar/admission-control/blob/master/admit_funcs.go#L64) function of the Admission Controller, the request should be explicitly denied. DenyIngresses denies any kind: Ingress from being deployed to the cluster, except for any explicitly allowed namespaces (e.g. istio-system).

## Block Public Load Balancers

Similar to Ingress objects, spinning up a public Loadbalancer may turn into a security issue. Let's us our Admission Webhook to block any public Loadbalancers from being created in the cluster:

```
#in the deny-public-lb directory 
kubectl create -f webhook-config.yaml
```
Now, to test that the Admission Control is working properly, create a simple Ingress object as follows:

```
kubectl create -f public-lb.yaml
```
Due to the rules built in the [DenyPublicLoadBalancers](https://github.com/elithrar/admission-control/blob/master/admit_funcs.go#L107) function of the Admission Controller, the request should be explicitly denied. DenyIngresses denies any kind: Ingress from being deployed to the cluster, except for any explicitly allowed namespaces (e.g. istio-system).

## Bonus
Successfully enforce one of the other [AdmitFuncs](https://github.com/elithrar/admission-control/blob/master/admit_funcs.go) using the API endpoint that is already running in Cloud Run.

Hint: You will need to modify the `webhook-config.yaml` file. 

## Shout Out
Shout out to the `Admission Control` repo from `elithrar`. Check it out on Github if you are looking for a robust framework for spinning up Admission Control in your clusters.

[https://github.com/elithrar/admission-control](https://github.com/elithrar/admission-control)