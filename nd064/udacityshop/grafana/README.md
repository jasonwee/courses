# Grafana Dashboards

Grafana Dashboards for the Uda'CityShop application.


```
kubectl create namespace monitoring
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring  --kubeconfig /etc/rancher/k3s/k3s.yaml
kubectl --namespace monitoring get pods -l "release=prometheus"
kubectl get all -n monitoring
```
