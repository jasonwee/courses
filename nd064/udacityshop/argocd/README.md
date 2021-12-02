# ArgoCD Manifests

ArgoCD manifests to deploy the Uda'CityShop application.

```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f argocd-service-nodeport.yaml
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

```
kubectl apply -f udacityshop-dev.yaml 
kubectl apply -f udacityshop-prod.yaml 
```
