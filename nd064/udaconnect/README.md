

# UdaConnect
## Overview
### Background
Conferences and conventions are hotspots for making connections. Professionals in attendance often share the same interests and can make valuable business and personal connections with one another. At the same time, these events draw a large crowd and it's often hard to make these connections in the midst of all of these events' excitement and energy. To help attendees make connections, we are building the infrastructure for a service that can inform attendees if they have attended the same booths and presentations at an event.

### Goal
You work for a company that is building a app that uses location data from mobile devices. Your company has built a [POC](https://en.wikipedia.org/wiki/Proof_of_concept) application to ingest location data named UdaTracker. This POC was built with the core functionality of ingesting location and identifying individuals who have shared a close geographic proximity.

Management loved the POC so now that there is buy-in, we want to enhance this application. You have been tasked to enhance the POC application into a [MVP](https://en.wikipedia.org/wiki/Minimum_viable_product) to handle the large volume of location data that will be ingested.

### Technologies
* [Flask](https://flask.palletsprojects.com/en/1.1.x/) - API webserver
* [SQLAlchemy](https://www.sqlalchemy.org/) - Database ORM
* [PostgreSQL](https://www.postgresql.org/) - Relational database
* [PostGIS](https://postgis.net/) - Spatial plug-in for PostgreSQL enabling geographic queries]
* [Vagrant](https://www.vagrantup.com/) - Tool for managing virtual deployed environments
* [VirtualBox](https://www.virtualbox.org/) - Hypervisor allowing you to run multiple operating systems
* [K3s](https://k3s.io/) - Lightweight distribution of K8s to easily develop against a local cluster

## Running the app
The project has been set up such that you should be able to have the project up and running with Kubernetes.

### Prerequisites
We will be installing the tools that we'll need to use for getting our environment set up properly.
1. [Install Docker](https://docs.docker.com/get-docker/)
2. [Set up a DockerHub account](https://hub.docker.com/)
3. [Set up `kubectl`](https://rancher.com/docs/rancher/v2.x/en/cluster-admin/cluster-access/kubectl/)
4. [Install VirtualBox](https://www.virtualbox.org/wiki/Downloads) with at least version 6.1
5. [Install Vagrant](https://www.vagrantup.com/docs/installation) with at least version 2.2
6. [Install Helm](https://helm.sh/docs/intro/install/) with at least version 3.6.3

### Environment Setup
To run the application, you will need a K8s cluster running locally and to interface with it via `kubectl`. We will be using Vagrant with VirtualBox to run K3s.

#### Initialize K3s
In this project's root, run `vagrant up`. 
```bash
$ vagrant up
```
The command will take a while and will leverage VirtualBox to load an [openSUSE](https://www.opensuse.org/) OS and automatically install [K3s](https://k3s.io/). When we are taking a break from development, we can run `vagrant suspend` to conserve some ouf our system's resources and `vagrant resume` when we want to bring our resources back up. Some useful vagrant commands can be found in [this cheatsheet](https://gist.github.com/wpscholar/a49594e2e2b918f4d0c4).

#### Set up `kubectl`
After `vagrant up` is done, you will SSH into the Vagrant environment and retrieve the Kubernetes config file used by `kubectl`. We want to copy the contents of this file into our local environment so that `kubectl` knows how to communicate with the K3s cluster.
```bash
$ vagrant ssh
```
You will now be connected inside of the virtual OS. Run `sudo cat /etc/rancher/k3s/k3s.yaml` to print out the contents of the file. You should see output similar to the one that I've shown below. Note that the output below is just for your reference: every configuration is unique and you should _NOT_ copy the output I have below.

Copy the contents from the output issued from your own command into your clipboard -- we will be pasting it somewhere soon!
```bash
master:~ # cat /etc/rancher/k3s/k3s.yaml 
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJlRENDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdGMyVnkKZG1WeUxXTmhRREUyTXpBM09EUXhPRFV3SGhjTk1qRXdPVEEwTVRrek5qSTFXaGNOTXpFd09UQXlNVGt6TmpJMQpXakFqTVNFd0h3WURWUVFEREJock0zTXRjMlZ5ZG1WeUxXTmhRREUyTXpBM09EUXhPRFV3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFRL3lSYlZBL3Y1VUIrVDVoN0RkaXZLMHViS1lDa0pUUTZxRXJDbFhuNUoKd2U0MU5SamNHcW8wWllpYVZUTVFYQXlaeE5mM2dmK3J0dURxTml3ZlR2cHBvMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVUJNNVl5ZW5NUXh3WWN1bmtXYiszCnptdmZoYzB3Q2dZSUtvWkl6ajBFQXdJRFNRQXdSZ0loQUw2c29oR3F1UkVEZEVhdjdzTWpITVNWYXhHNnBMcVkKc01HMndLSGtPaVdiQWlFQS9kdkQzK1BaYWhpZy9xdzFSVkFhZGJFdnFnQkUyRjJtZytQVTd1OTRYZ1U9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://127.0.0.1:6443
  name: default
contexts:
- context:
    cluster: default
    user: default
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: default
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJrVENDQVRlZ0F3SUJBZ0lJVkFON0FuVUEzUlV3Q2dZSUtvWkl6ajBFQXdJd0l6RWhNQjhHQTFVRUF3d1kKYXpOekxXTnNhV1Z1ZEMxallVQXhOak13TnpnME1UZzFNQjRYRFRJeE1Ea3dOREU1TXpZeU5Wb1hEVEl5TURrdwpOREU1TXpZeU5Wb3dNREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhGVEFUQmdOVkJBTVRESE41CmMzUmxiVHBoWkcxcGJqQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlBd0VIQTBJQUJIUFk5ZU5mSE1IYXRGMlQKUEVCcUNTVSt0ZkRRY1Q4UDYxSVRQTjJycnBraVJSSlJwVnNBK1g3ZDRwRzBzVVlIN1I4VEFlRnBKeHM5b2RtcgpHeFg0N2xLalNEQkdNQTRHQTFVZER3RUIvd1FFQXdJRm9EQVRCZ05WSFNVRUREQUtCZ2dyQmdFRkJRY0RBakFmCkJnTlZIU01FR0RBV2dCU05rRENsd1Y4bUY4aE9LNmhMbnRXWlZwQ3J3REFLQmdncWhrak9QUVFEQWdOSUFEQkYKQWlCNEhzY1RGUmxnc2hjQ2tjVnlLNkNTMEswMWhLZUkzK0dObUhWTFBwczhSZ0loQU9ScERFZU9MMm8zYkl5SQpXNkVWVUdEZEphTzdsalpxY0ZxVnBmT2VBTytoCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJkekNDQVIyZ0F3SUJBZ0lCQURBS0JnZ3Foa2pPUFFRREFqQWpNU0V3SHdZRFZRUUREQmhyTTNNdFkyeHAKWlc1MExXTmhRREUyTXpBM09EUXhPRFV3SGhjTk1qRXdPVEEwTVRrek5qSTFXaGNOTXpFd09UQXlNVGt6TmpJMQpXakFqTVNFd0h3WURWUVFEREJock0zTXRZMnhwWlc1MExXTmhRREUyTXpBM09EUXhPRFV3V1RBVEJnY3Foa2pPClBRSUJCZ2dxaGtqT1BRTUJCd05DQUFUc3BsRGE5bHd2aTB0VTFZWlQ5RFV1b2JvWGxXQkxtN0drYWhnL1RrbW0KVEx5RmdmYnJMVnZmcGxwekcvS1RlSDBsZnRtamtsakRwVkxKM0lzOTB5U3NvMEl3UURBT0JnTlZIUThCQWY4RQpCQU1DQXFRd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBZEJnTlZIUTRFRmdRVWpaQXdwY0ZmSmhmSVRpdW9TNTdWCm1WYVFxOEF3Q2dZSUtvWkl6ajBFQXdJRFNBQXdSUUloQU5teC9vRmZLeGlkQzI3QUk4TUFRTFM1SklobGlNWk0KaGpaMytMbU51bUE2QWlCbHFIclRERGtOZkUxdHdOMUUvUlE5UG1KYm5rYVRWOTF5eHJnSUMwSDlSUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBFQyBQUklWQVRFIEtFWS0tLS0tCk1IY0NBUUVFSUJMWFMwMVg3WnoyRnNGR1NBMXJYSTlDci9lQzF1YW15Tm1nTG5rWElVYTFvQW9HQ0NxR1NNNDkKQXdFSG9VUURRZ0FFYzlqMTQxOGN3ZHEwWFpNOFFHb0pKVDYxOE5CeFB3L3JVaE04M2F1dW1TSkZFbEdsV3dENQpmdDNpa2JTeFJnZnRIeE1CNFdrbkd6MmgyYXNiRmZqdVVnPT0KLS0tLS1FTkQgRUMgUFJJVkFURSBLRVktLS0tLQo=

```
Type `exit` to exit the virtual OS and you will find yourself back in your computer's session. Create the file (or replace if it already exists) `~/.kube/config` and paste the contents of the `k3s.yaml` output here.

Afterwards, you can test that `kubectl` works by running a command like `kubectl describe services`. It should not return any errors.


### Steps

1. kafka
######  install helm in vm
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
###### install kafka in kubernetes via helm
```
master:~ # helm repo add bitnami https://charts.bitnami.com/bitnami
"bitnami" has been added to your repositories
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

2. person
```
$ cd modules/person/
$ kubectl apply -f deployment/
```
once the pods are up include postgres person, populate the db

```
bash scripts/run_db_command.sh <POSTGRES POD NAME>
```


3. connection
```
$ cd modules/connection/
$ kubectl apply -f deployment/

```

once the pods are up include postgres connection, populate the db

```
bash scripts/run_db_command.sh <POSTGRES POD NAME>
```


4. location-event-producer
```
$ cd modules/connection/
$ kubectl apply -f deployment/
```

5. location-event-consumer
```
$ cd modules/connection/
$ kubectl apply -f deployment/
```


6. frontend
```
$ cd modules/frontend/
$ kubectl apply -f deployment/
```


### Verifying it Works


#### Deployment Note
You may notice the odd port numbers being served to `localhost`. [By default, Kubernetes services are only exposed to one another in an internal network](https://kubernetes.io/docs/concepts/services-networking/service/). This means that `udaconnect-app` and `udaconnect-api` can talk to one another. For us to connect to the cluster as an "outsider", we need to a way to expose these services to `localhost`.

Connections to the Kubernetes services have been set up through a [NodePort](https://kubernetes.io/docs/concepts/services-networking/service/#nodeport). (While we would use a technology like an [Ingress Controller](https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/) to expose our Kubernetes services in deployment, a NodePort will suffice for development.)

## Development
### New Services
New services can be created inside of the `modules/` subfolder. You can choose to write something new with Flask, copy and rework the `modules/api` service into something new, or just create a very simple Python application.

As a reminder, each module should have:
1. `Dockerfile`
2. Its own corresponding DockerHub repository
3. `requirements.txt` for `pip` packages
4. `__init__.py`

### Docker Images
`udaconnect-app` and `udaconnect-api` use docker images from `isjustintime/udaconnect-app` and `isjustintime/udaconnect-api`. To make changes to the application, build your own Docker image and push it to your own DockerHub repository. Replace the existing container registry path with your own.

## Configs and Secrets
In `deployment/db-secret.yaml`, the secret variable is `d293aW1zb3NlY3VyZQ==`. The value is simply encoded and not encrypted -- this is ***not*** secure! Anyone can decode it to see what it is.
```bash
# Decodes the value into plaintext
echo "d293aW1zb3NlY3VyZQ==" | base64 -d

# Encodes the value to base64 encoding. K8s expects your secrets passed in with base64
echo "hotdogsfordinner" | base64
```
This is okay for development against an exclusively local environment and we want to keep the setup simple so that you can focus on the project tasks. However, in practice we should not commit our code with secret values into our repository. A CI/CD pipeline can help prevent that.

## PostgreSQL Database
The database uses a plug-in named PostGIS that supports geographic queries. It introduces `GEOMETRY` types and functions that we leverage to calculate distance between `ST_POINT`'s which represent latitude and longitude.

_You may find it helpful to be able to connect to the database_. In general, most of the database complexity is abstracted from you. The Docker container in the starter should be configured with PostGIS. Seed scripts are provided to set up the database table and some rows.
### Database Connection
While the Kubernetes service for `postgres` is running (you can use `kubectl get services` to check), you can expose the service to connect locally:
```bash
kubectl port-forward svc/postgres 5432:5432
```
This will enable you to connect to the database at `localhost`. You should then be able to connect to `postgresql://localhost:5432/geoconnections`. This is assuming you use the built-in values in the deployment config map.
### Software
To manually connect to the database, you will need software compatible with PostgreSQL.
* CLI users will find [psql](http://postgresguide.com/utilities/psql.html) to be the industry standard.
* GUI users will find [pgAdmin](https://www.pgadmin.org/) to be a popular open-source solution.

## Architecture Diagrams
Your architecture diagram should focus on the services and how they talk to one another. For our project, we want the diagram in a `.png` format. Some popular free software and tools to create architecture diagrams:
1. [Lucidchart](https://www.lucidchart.com/pages/)
2. [Google Docs](docs.google.com) Drawings (In a Google Doc, _Insert_ - _Drawing_ - _+ New_)
3. [Diagrams.net](https://app.diagrams.net/)

## Tips
* We can access a running Docker container using `kubectl exec -it <pod_id> sh`. From there, we can `curl` an endpoint to debug network issues.
* The starter project uses Python Flask. Flask doesn't work well with `asyncio` out-of-the-box. Consider using `multiprocessing` to create threads for asynchronous behavior in a standard Flask application.
