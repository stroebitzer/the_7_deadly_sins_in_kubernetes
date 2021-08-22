# Setup

## Install chart repo

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm show values ingress-nginx/ingress-nginx > nginx.yaml  
```

## Install chart

```bash
helm install --create-namespace --namespace ingress-controller ingress-nginx ingress-nginx/ingress-nginx
helm upgrade --create-namespace --namespace ingress-controller ingress-nginx ingress-nginx/ingress-nginx
```
