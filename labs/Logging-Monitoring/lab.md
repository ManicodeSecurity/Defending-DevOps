https://itnext.io/kubernetes-monitoring-with-prometheus-in-15-minutes-8e54d1de2e13


## Auditing
A recent addition to Kubernetes, (https://kubernetes.io/docs/tasks/debug-application-cluster/audit/)
[auditing] gives administrators and security teams the ability to log and monitor security-related events occurring on a cluster. By using audit policies, we can create granular rulesets to focus on only on the meaningful events and cut down on the noise. 

# Task 1: We will need to enable some flags in our cluster so we will start off with a clean cluster using the following command:
```
minikube delete
```

# Task 2: To enable audition in Minikube we must utilize the `addons` feature. To do this, we must copy over the `audit-policy.yaml` manifest to the `.minikube/addons` directory:
```
cp /path/to/5-Advanced-K8S-Security-Features/audit-policy/manifests/audit-policy.yaml ~/.minikube/addons
```

# Task 3: Create the `audit.log` file that we will write our JSON logs out to:
```
touch ~/.minikube/logs/audit.log
```

# Task 4: Launch the cluster with the following flags. Make sure to enter the correct path to the .minikube folder!
```
minikube start \
    --feature-gates=AdvancedAudit=true \
    --extra-config=apiserver.Audit.LogOptions.Path=/path/to/.minikube/logs/audit.log \    --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml
```

# Take a look at the logs coming in. We will filter those from 127.0.0.1 to cut down on the noise. [https://stedolan.github.io/jq/download/](jq) is a command line tool which greatly helps with formatting JSON to the terminal. It's a handy tool to have available when working with JSON.
```
tail -f  ~/.minikube/logs/audit.log  | jq '.| select(.sourceIPs | contains(["127.0.0.1"]) | not)'
```

## Bonus: Using a webhook, send the Kubernetes logs to an API such as requestb.in (This is obviously VERY insecure and should only be used for testing or learning purposes. Always ship logs to an approved log aggregation system.

Hint 1: The `webhook.yaml` file should be copied to the same location as the `audit-policy.yaml` file in `.minikube/addons` directory.

Hint 2: We need to pass additional flags to Minikube to enable the webhook. Starting a new cluster should look like this:
```
minikube start \
    --feature-gates=AdvancedAudit=true \
    --extra-config=apiserver.Audit.PolicyFile=/etc/kubernetes/addons/audit-policy.yaml \
    --extra-config=apiserver.Audit.WebhookOptions.ConfigFile=/etc/kubernetes/addons/webhook.yaml 
```

## Discussion Question: How would you ingest these logs into your current log management systems? What would you alert on?
