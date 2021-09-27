**Note:** For the screenshots, you can store all of your answer images in the `answer-img` directory.


## overview diagram

```
                    +-------------+
                    | +---------+ |
                  +-->| metrics |----+
 +------------+   | | +---------+ |  |   +--------------+
 | sample app |---+ |             |  +-->| visualization|
 +------------+   | | +---------+ |  |   +--------------+
                  +-->| tracing |----+
                    | +---------+ |
                    +-------------+
```

* metrics       : prometheus
* tracing       : jaeger
* visualization : grafana


## allow localhost to access k3s
```
export KUBECONFIG=descriptor/k3s.yaml
kubectl version
```

## install helm
```
vagrant@localhost:~> curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
vagrant@localhost:~> ls
bin  get_helm.sh
vagrant@localhost:~> chmod 700 get_helm.sh
vagrant@localhost:~> ./get_helm.sh
Downloading https://get.helm.sh/helm-v3.7.0-linux-amd64.tar.gz
Verifying checksum... Done.
Preparing to install helm into /usr/local/bin
helm installed into /usr/local/bin/helm
vagrant@localhost:~> helm version
version.BuildInfo{Version:"v3.7.0", GitCommit:"eeac83883cb4014fe60267ec6373570374ce770b", GitTreeState:"clean", GoVersion:"go1.16.8"}
vagrant@localhost:~>

```

## Installing Grafana and Prometheus
```
vagrant@localhost:~> kubectl create namespace monitoring
namespace/monitoring created
vagrant@localhost:~> kubectl get ns
NAME              STATUS   AGE
default           Active   28m
kube-system       Active   28m
kube-public       Active   28m
kube-node-lease   Active   28m
monitoring        Active   4s
vagrant@localhost:~> helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
"prometheus-community" has been added to your repositories
vagrant@localhost:~> helm repo add stable https://charts.helm.sh/stable
"stable" has been added to your repositories
vagrant@localhost:~> helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "prometheus-community" chart repository
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈Happy Helming!⎈
vagrant@localhost:~> helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring --kubeconfig /etc/rancher/k3s/k3s.yaml
WARNING: Kubernetes configuration file is group-readable. This is insecure. Location: /etc/rancher/k3s/k3s.yaml
WARNING: Kubernetes configuration file is world-readable. This is insecure. Location: /etc/rancher/k3s/k3s.yaml
NAME: prometheus
LAST DEPLOYED: Mon Sep 27 14:14:43 2021
NAMESPACE: monitoring
STATUS: deployed
REVISION: 1
NOTES:
kube-prometheus-stack has been installed. Check its status by running:
  kubectl --namespace monitoring get pods -l "release=prometheus"

Visit https://github.com/prometheus-operator/kube-prometheus for instructions on how to create & configure Alertmanager and Prometheus instances using the Operator.
vagrant@localhost:~> kubectl get all -n monitoring
NAME                                                         READY   STATUS    RESTARTS   AGE
pod/prometheus-kube-state-metrics-696cf79768-xnpht           1/1     Running   0          75s
pod/prometheus-kube-prometheus-operator-5c758db547-2hcxx     1/1     Running   0          75s
pod/prometheus-prometheus-node-exporter-4kmn6                1/1     Running   0          75s
pod/alertmanager-prometheus-kube-prometheus-alertmanager-0   2/2     Running   0          60s
pod/prometheus-prometheus-kube-prometheus-prometheus-0       2/2     Running   0          59s
pod/prometheus-grafana-57b5c4fdb6-x7wwp                      2/2     Running   0          75s

NAME                                              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
service/prometheus-kube-prometheus-prometheus     ClusterIP   10.43.174.68    <none>        9090/TCP                     76s
service/prometheus-kube-prometheus-operator       ClusterIP   10.43.144.114   <none>        443/TCP                      76s
service/prometheus-kube-prometheus-alertmanager   ClusterIP   10.43.218.216   <none>        9093/TCP                     75s
service/prometheus-prometheus-node-exporter       ClusterIP   10.43.57.150    <none>        9100/TCP                     75s
service/prometheus-kube-state-metrics             ClusterIP   10.43.113.115   <none>        8080/TCP                     75s
service/prometheus-grafana                        ClusterIP   10.43.188.95    <none>        80/TCP                       75s
service/alertmanager-operated                     ClusterIP   None            <none>        9093/TCP,9094/TCP,9094/UDP   60s
service/prometheus-operated                       ClusterIP   None            <none>        9090/TCP                     60s

NAME                                                 DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/prometheus-prometheus-node-exporter   1         1         1       1            1           <none>          75s

NAME                                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/prometheus-kube-state-metrics         1/1     1            1           75s
deployment.apps/prometheus-kube-prometheus-operator   1/1     1            1           75s
deployment.apps/prometheus-grafana                    1/1     1            1           75s

NAME                                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/prometheus-kube-state-metrics-696cf79768         1         1         1       75s
replicaset.apps/prometheus-kube-prometheus-operator-5c758db547   1         1         1       75s
replicaset.apps/prometheus-grafana-57b5c4fdb6                    1         1         1       75s

NAME                                                                    READY   AGE
statefulset.apps/alertmanager-prometheus-kube-prometheus-alertmanager   1/1     60s
statefulset.apps/prometheus-prometheus-kube-prometheus-prometheus       1/1     60s
vagrant@localhost:~>
```

## Install Jaeger
```
vagrant@localhost:~> kubectl create namespace observability
namespace/observability created
vagrant@localhost:~> kubectl get ns
NAME              STATUS   AGE
default           Active   31m
kube-system       Active   31m
kube-public       Active   31m
kube-node-lease   Active   31m
monitoring        Active   3m17s
observability     Active   4s
vagrant@localhost:~> kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/crds/jaegertracing.io_jaegers_crd.yaml
customresourcedefinition.apiextensions.k8s.io/jaegers.jaegertracing.io created
vagrant@localhost:~> kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/service_account.yaml
serviceaccount/jaeger-operator created
vagrant@localhost:~> kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/role.yaml
role.rbac.authorization.k8s.io/jaeger-operator created
vagrant@localhost:~> kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/jaeger-operator created
vagrant@localhost:~> kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/operator.yaml
deployment.apps/jaeger-operator created
vagrant@localhost:~> kubectl get all -n observability
NAME                                 READY   STATUS    RESTARTS   AGE
pod/jaeger-operator-dbf5767c-jjdcl   1/1     Running   0          41s

NAME                              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)             AGE
service/jaeger-operator-metrics   ClusterIP   10.43.207.133   <none>        8383/TCP,8686/TCP   14s

NAME                              READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/jaeger-operator   1/1     1            1           41s

NAME                                       DESIRED   CURRENT   READY   AGE
replicaset.apps/jaeger-operator-dbf5767c   1         1         1       41s
vagrant@localhost:~>
```

### Cluster wide Jaeger
```
vagrant@localhost:~> kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role.yaml
clusterrole.rbac.authorization.k8s.io/jaeger-operator created
vagrant@localhost:~> kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role_binding.yaml
clusterrolebinding.rbac.authorization.k8s.io/jaeger-operator created
vagrant@localhost:~>
```

## Verify the monitoring installation

*TODO:* run `kubectl` command to show the running pods and services for all components. Take a screenshot of the output and include it here to verify the installation

## Setup the Jaeger and Prometheus source
*TODO:* Expose Grafana to the internet and then setup Prometheus as a data source. Provide a screenshot of the home page after logging into Grafana.

## Create a Basic Dashboard
*TODO:* Create a dashboard in Grafana that shows Prometheus as a source. Take a screenshot and include it here.

## Describe SLO/SLI
*TODO:* Describe, in your own words, what the SLIs are, based on an SLO of *monthly uptime* and *request response time*.

## Creating SLI metrics.
*TODO:* It is important to know why we want to measure certain metrics for our customer. Describe in detail 5 metrics to measure these SLIs. 

## Create a Dashboard to measure our SLIs
*TODO:* Create a dashboard to measure the uptime of the frontend and backend services We will also want to measure to measure 40x and 50x errors. Create a dashboard that show these values over a 24 hour period and take a screenshot.

## Tracing our Flask App
*TODO:*  We will create a Jaeger span to measure the processes on the backend. Once you fill in the span, provide a screenshot of it here.

## Jaeger in Dashboards
*TODO:* Now that the trace is running, let's add the metric to our current Grafana dashboard. Once this is completed, provide a screenshot of it here.

## Report Error
*TODO:* Using the template below, write a trouble ticket for the developers, to explain the errors that you are seeing (400, 500, latency) and to let them know the file that is causing the issue.

TROUBLE TICKET

Name:

Date:

Subject:

Affected Area:

Severity:

Description:


## Creating SLIs and SLOs
*TODO:* We want to create an SLO guaranteeing that our application has a 99.95% uptime per month. Name three SLIs that you would use to measure the success of this SLO.

## Building KPIs for our plan
*TODO*: Now that we have our SLIs and SLOs, create KPIs to accurately measure these metrics. We will make a dashboard for this, but first write them down here.

## Final Dashboard
*TODO*: Create a Dashboard containing graphs that capture all the metrics of your KPIs and adequately representing your SLIs and SLOs. Include a screenshot of the dashboard here, and write a text description of what graphs are represented in the dashboard.

