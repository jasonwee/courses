## install helm in vm
see https://helm.sh/docs/intro/install/

```sh
vagrant@master:~> curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
vagrant@master:~> chmod 700 get_helm.sh
vagrant@master:~> ./get_helm.sh
Downloading https://get.helm.sh/helm-v3.6.3-linux-amd64.tar.gz
Verifying checksum... Done.
Preparing to install helm into /usr/local/bin
helm installed into /usr/local/bin/helm
vagrant@master:~> helm version
version.BuildInfo{Version:"v3.6.3", GitCommit:"d506314abfb5d21419df8c7e7e68012379db2354", GitTreeState:"clean", GoVersion:"go1.16.5"}
vagrant@master:~> 
```

## install kafka in kubernetes via helm
see https://bitnami.com/stack/kafka/helm
https://www.ignitesol.com/how-to-deploy-kafka-connect-on-kubernetes-using-helm-charts/
```
master:~ # helm repo list
Error: no repositories to show
master:~ # helm repo add bitnami https://charts.bitnami.com/bitnami
"bitnami" has been added to your repositories
master:~ # helm repo list
NAME   	URL                               
bitnami	https://charts.bitnami.com/bitnami
master:~ # 
master:~ # helm search repo kafka
NAME                    	CHART VERSION	APP VERSION	DESCRIPTION                                       
bitnami/kafka           	14.0.5       	2.8.0      	Apache Kafka is a distributed streaming platform. 
bitnami/dataplatform-bp1	7.0.0        	1.0.0      	OCTO Data platform Kafka-Spark-Solr Helm Chart    
bitnami/dataplatform-bp2	6.0.0        	1.0.0      	OCTO Data platform Kafka-Spark-Elasticsearch He...
master:~ # export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
master:~ # helm install my-release bitnami/kafka
NAME: my-release
LAST DEPLOYED: Sun Sep  5 02:18:58 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
** Please be patient while the chart is being deployed **

Kafka can be accessed by consumers via port 9092 on the following DNS name from within your cluster:

    my-release-kafka.default.svc.cluster.local

Each Kafka broker can be accessed by producers via port 9092 on the following DNS name(s) from within your cluster:

    my-release-kafka-0.my-release-kafka-headless.default.svc.cluster.local:9092

To create a pod that you can use as a Kafka client run the following commands:

    kubectl run my-release-kafka-client --restart='Never' --image docker.io/bitnami/kafka:2.8.0-debian-10-r84 --namespace default --command -- sleep infinity
    kubectl exec --tty -i my-release-kafka-client --namespace default -- bash

    PRODUCER:
        kafka-console-producer.sh \
            
            --broker-list my-release-kafka-0.my-release-kafka-headless.default.svc.cluster.local:9092 \
            --topic test

    CONSUMER:
        kafka-console-consumer.sh \
            
            --bootstrap-server my-release-kafka.default.svc.cluster.local:9092 \
            --topic test \
            --from-beginning
master:~ # 

```

verify

```
master:~ # kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
my-release-zookeeper-0            1/1     Running   0          3m17s
my-release-kafka-0                1/1     Running   0          3m17s
master:~ # 
```





