
## Setting up the cluster

```console
kind create cluster --config kind.yaml

# Deploy nginx ingress
kubectl apply -f deploy-nginx.yaml

# Wait until is ready to process requests running
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```


## Install metal-lb on kind
[Install docs](https://kind.sigs.k8s.io/docs/user/loadbalancer/)

```console
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/namespace.yaml
kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="$(openssl rand -base64 128)" 
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/master/manifests/metallb.yaml

# Wait for metallb pods to come up
kubectl get pods -n metallb-system --watch 
```

MetalLB is a L2 load balancer. It needs a range of ip addresses which it can control
Once pods are up we can now setup address pool for our load balancer. 
For this we need to know the range of docker kind network.

```console
aaqa@altair ~/p/s/g/a/h/i/kind infra [!±] % docker network inspect -f '{{.IPAM.Config}}' kind
[{172.21.0.0/16  172.21.0.1 map[]} {fc00:f853:ccd:e793::/64  fc00:f853:ccd:e793::1 map[]}]
```
Here `172.21.0.0/16` is the cidr block for docker kind network.
We can allot some ip addresses from this range

