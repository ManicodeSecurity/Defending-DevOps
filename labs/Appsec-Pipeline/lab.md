Basics of web application testing

static code analysis
jenkins running clair? running static code analysis tool?
git pre-commit hook

Jenkins running static scan or ZAP?

Run dependency check
## Build an internal image registry
in the `manifests/registry` directory run the following:
```
kubectl create -f .
```

## Rollout Jenkins
- in /jenkins directory
```
kubectl create -f .
# Make sure Jenkins is up and grab the pod ID
kubectl get pods
# Now copy the admin password passed to the logs
kubectl logs <jenkins-pod-name> | grep -B 3 initialAdminPassword
```
Next, we fire up the Jenkins UI

plugins to install in the Setup Wizard (only install these!):
pipeline
git

Now we build our pipeline
new project
add git as scm
paste repo

Run clair

Run static code analysis tool!

