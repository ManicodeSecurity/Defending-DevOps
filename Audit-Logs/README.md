# Auditing
A recent addition to Kubernetes, [auditing](https://kubernetes.io/docs/tasks/debug-application-cluster/audit/), gives administrators and security teams the ability to log and monitor security-related events occurring on a cluster. By using audit policies, we can create granular rulesets to focus on only on the meaningful events and cut down on the noise.

### Task 1: Inspect the Default Audit Policy
1. Since GKE uses a shared master node, we can't easily apply our own audit policies to a cluster. GKE clusters are bootstrapped with a default sane policy that can be found in the `manifests` directory. The policy summary can be found in the [https://cloud.google.com/kubernetes-engine/docs/concepts/audit-policy](Kubernetes docs).


### Task 2: Open Audit Logs in Stackdriver
In the main navigation in GCP, navigate to `Stackdriver` -> `Logging` -> `Logs`
Under the `Resources` drop-down select `Kubernetes Cluster` then the name of your cluster

 Here you will see a list of audit events associated with your cluster.

### Task 3: Trigger an Event
In Cloud Shell run a command that will trigger an audit event and inspect the event in Stackdriver.
```
# Creating a secret will add a log entry
kubectl create secret generic mysql-secrets --from-literal=password=supertopsecretpassword
```
Back in Stackdriver logs page, filter for the word `secret` and click the `Jump to Now` button. You should be able to see a log entry associated with the secret creation kubectl command. You can expand the log entry to view more detail.

## Discussion Question:
How would you ingest these logs into your current log management systems? What would you alert on?
