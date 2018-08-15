kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"

  wget https://github.com/istio/istio/releases/download/1.0.0/istio-1.0.0-linux.
tar.gz

tar -xzf istio-1.0.0-linux.tar.gz

export PATH=$PWD/bin:$PATH


kubectl label namespace default istio-injection=enabled

Make sure our API is working..

you will see that there are some issues. Egress is blocked by default
http://35.197.37.188/api/check?url=https://bit.ly/hi
NOT ALLOWED!

Lets build some rules to explictlly allow outboud traffic to only bit.ly
Grafana
kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 8080:3000

Then Click "Web Preview" in cloud shell

