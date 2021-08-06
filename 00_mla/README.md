# Prometheus


## Add the prometheus-community helm repo

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

## Prometheus

### Search for a prometheus chart via helm

```bash
helm search repo prometheus
```

### Install prometheus

```bash
helm install prometheus prometheus-community/prometheus -f prometheus-values.yaml 
```

### Take a look at the installed resources

```bash
helm list
kubectl get all
```

### Visit the prometheus UI in the browser

```bash
# Get the External IP of the service
kubectl get svc my-prometheus-server
```

### Execute some query over the "graph" page in your browser

* http requests: `container_memory_usage_bytes`
* API server requests count: `container_memory_usage_bytes{pod="my-pod",container="nginx"}`
* per-second rate of increase, messured over the last 5 min: `rate(container_memory_usage_bytes{pod="my-pod",container="nginx"}[5m])`
* predict memory usage in 24 hours based on the rate of the last 5 minutes: `predict_linear((container_memory_usage_bytes{pod="my-pod",container="nginx"}[5m]),24*3600)`
* more examples: [https://prometheus.io/docs/prometheus/latest/querying/examples](https://prometheus.io/docs/prometheus/latest/querying/examples)

## Grafana

## Add the grafana helm repo

```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```

### Install grafana

```bash
helm install my-grafana grafana/grafana -f grafana-values.yaml
```

### Visit the grafana UI in the browser

```bash
# Get the External IP of the service
kubectl get svc
```

## Add Prometheus as data source to Grafana

* Choose the Option `Configuration`/`Data Sources`
* Choose type `Prometheus`
* Insert URL `http://my-prometheus-server` and click `Save & Test`
* Choose the Option `Dashboards`/`Manage`/`Import`
* Paste the number 1860 into the field `Grafana.com Dashboard` and click `Load` (Visit [https://grafana.com/grafana/dashboards](https://grafana.com/grafana/dashboards) for other Dashboards)
* Choose the Datasource Prometheus in the field `Prometheus` and `Import` afterwards
* Investigate the Dashboard; take care to a proper time range via the button on the right above

## Clean up

```bash
helm delete my-grafana
helm delete my-prometheus
kubectl delete pod my-pod
```