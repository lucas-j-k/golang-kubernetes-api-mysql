#!/bin/bash

# Start Kind cluster in docker with custom config file to allow ingress on ports 80 443 on localhost
# without this, you wont be able to access the cluster over localhost

# apply the ingress yam to add the nginx ingress
kubectl apply --filename https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml

# sleep
echo sleep start
sleep 25
echo sleep finish

echo starting ingress
# wait for the nginx ingress to be ready
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s


echo ______Done setting up ingress________
