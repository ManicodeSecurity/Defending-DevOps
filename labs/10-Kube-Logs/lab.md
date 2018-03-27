## Auditing
A recent addition to Kubernetes, (https://kubernetes.io/docs/tasks/debug-application-cluster/audit/)
[auditing] gives administrators and security teams the ability to log and monitor security-related events occurring on a cluster. By using audit policies, we can create granular rulesets to focus on only on the meaningful events and cut down on the noise. 

## Task 1: We will need to enable some flags in our cluster so we will start off with a clean cluster using the following command:
```
minikube delete
```

## Task 2: To enable audition in Minikube we must utilize the `addons` feature. To do this, we must copy over the `audit-policy.yaml` manifest to the `~/Desktop/lab-tools/.kube/.minikube/addons` directory:
```
cp /path/to/10-Kube-Logs/audit-policy/manifests/audit-policy.yaml ~/Desktop/lab-tools/.kube/.minikube/addons
```

## Task 3: Create the `audit.log` file that we will write our JSON logs out to:
```
touch ~/Desktop/lab-tools/.kube/.minikube/logs/audit.log
```

## Task 4: Launch the cluster with the following flags. Make sure to enter the correct path to the .minikube folder!
```
# Some shells including zsh will prefer this on one line...

minikube start --extra-config=apiserver.Authorization.Mode=RBAC --feature-gates=AdvancedAuditing --extra-config=apiserver.Audit.LogOptions.Path=/Users/jb0ss/Desktop/lab-tools/.kube/.minikube/logs/audit.log --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml
```

## Take a look at the logs coming in If you have [https://stedolan.github.io/jq/download/](jq) installed it will help format the JSON but it is not necessary. 
```
tail -f  ~/Desktop/lab-tools/.kube/.minikube/logs/audit.log 
# with jq 
tail -f  ~/Desktop/lab-tools/.kube/.minikube/logs/audit.log  | jq 
```

## Bonus: Using a webhook, send the Kubernetes logs to an internal logging mechanism if you have an approved one available. 

Hint 1: The `webhook.yaml` file should be copied to the same location as the `audit-policy.yaml` file in `~/Desktop/lab-tools/.kube/minikube/addons` directory.

Hint 2: We need to pass additional flags to Minikube to enable the webhook. Starting a new cluster should look like this:
```
minikube start --extra-config=apiserver.Authorization.Mode=RBAC --feature-gates=AdvancedAuditing --extra-config=apiserver.Audit.LogOptions.Path=/path/to/Desktop/lab-tools/.kube/.minikube/logs/audit.log --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml --extra-config=apiserver.Audit.WebhookOptions.ConfigFile=/etc/kubernetes/addons/webhook.yaml 
```

## Discussion Question: How would you ingest these logs into your current log management systems? What would you alert on?
