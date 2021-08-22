# Setup

## Prometheus

### Setup chart repo
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

### Setup chart 
```bash
helm install --create-namespace --namespace mla prometheus prometheus-community/prometheus -f 02_mla/prometheus-values.yaml 
helm upgrade --create-namespace --namespace mla prometheus prometheus-community/prometheus -f 02_mla/prometheus-values.yaml 
helm uninstall prometheus
```

## Loki

### Setup chart repo
```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm show values grafana/loki-stack > loki.yaml  
```

### Setup chart 

```bash
helm install --create-namespace --namespace mla loki grafana/loki-stack -f 02_mla/loki-values.yaml
helm upgrade --create-namespace --namespace mla loki grafana/loki-stack -f 02_mla/loki-values.yaml
helm uninstall loki
```

## Grafana

### Setup chart repo
```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
helm show values grafana/grafana > grafana.yaml  
```

### Setup chart 

```bash
helm install --create-namespace --namespace mla grafana grafana/grafana -f 02_mla/grafana-values.yaml
kubectl apply -f 02_mla/ingress.yaml

helm upgrade --create-namespace --namespace mla grafana grafana/grafana -f 02_mla/grafana-values.yaml
helm uninstall grafana
```

### Configure Grafana

Prometheus Datasource: http://prometheus
Loki Datasource: http://loki:3100