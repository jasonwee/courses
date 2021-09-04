```
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl apply -f deployment/db-configmap.yaml
configmap/db-env created
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl apply -f deployment/db-secret.yaml
secret/db-secret created
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl apply -f deployment/postgres.yaml
persistentvolume/postgres-volume created
persistentvolumeclaim/postgres-pv-claim created
service/postgres created
deployment.apps/postgres created
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl apply -f deployment/udaconnect-api.yaml
service/udaconnect-api created
deployment.apps/udaconnect-api created
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl apply -f deployment/udaconnect-app.yaml
service/udaconnect-app created
deployment.apps/udaconnect-app created
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get pods
NAME                              READY   STATUS              RESTARTS   AGE
postgres-5f676c995d-jvdrp         0/1     Pending             0          17s
udaconnect-api-89dbffbf9-qhxzx    0/1     ContainerCreating   0          11s
udaconnect-app-6ddf7c79fb-kqfmb   0/1     ContainerCreating   0          6s
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get pods
NAME                              READY   STATUS                 RESTARTS   AGE
udaconnect-app-6ddf7c79fb-kqfmb   0/1     ContainerCreating      0          52s
udaconnect-api-89dbffbf9-qhxzx    0/1     CreateContainerError   0          57s
postgres-5f676c995d-jvdrp         0/1     CreateContainerError   0          63s
jason@localhost:~/oss/courses/nd064/udaconnect$ 
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get pods
NAME                              READY   STATUS                 RESTARTS   AGE
udaconnect-app-6ddf7c79fb-kqfmb   0/1     ContainerCreating      0          52s
udaconnect-api-89dbffbf9-qhxzx    0/1     CreateContainerError   0          57s
postgres-5f676c995d-jvdrp         0/1     CreateContainerError   0          63s
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl describe pod postgres-5f676c995d-jvdrp
Name:         postgres-5f676c995d-jvdrp
Namespace:    default
Priority:     0
Node:         master/10.0.2.15
Start Time:   Sun, 05 Sep 2021 03:52:21 +0800
Labels:       app=postgres
              pod-template-hash=5f676c995d
Annotations:  <none>
Status:       Pending
IP:           10.42.0.9
IPs:
  IP:           10.42.0.9
Controlled By:  ReplicaSet/postgres-5f676c995d
Containers:
  postgres:
    Container ID:   
    Image:          postgis/postgis:12-2.5-alpine
    Image ID:       
    Port:           5432/TCP
    Host Port:      0/TCP
    State:          Waiting
      Reason:       CreateContainerError
    Ready:          False
    Restart Count:  0
    Environment:
      POSTGRES_USER:      <set to the key 'DB_USERNAME' of config map 'db-env'>  Optional: false
      POSTGRES_PASSWORD:  <set to the key 'DB_PASSWORD' in secret 'db-secret'>   Optional: false
      POSTGRES_DB:        <set to the key 'DB_NAME' of config map 'db-env'>      Optional: false
    Mounts:
      /var/lib/postgresql/data from postgresdb (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-gmll8 (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             False 
  ContainersReady   False 
  PodScheduled      True 
Volumes:
  postgresdb:
    Type:       PersistentVolumeClaim (a reference to a PersistentVolumeClaim in the same namespace)
    ClaimName:  postgres-pv-claim
    ReadOnly:   false
  kube-api-access-gmll8:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason            Age                From               Message
  ----     ------            ----               ----               -------
  Warning  FailedScheduling  85s                default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.
  Warning  FailedScheduling  83s                default-scheduler  0/1 nodes are available: 1 pod has unbound immediate PersistentVolumeClaims.
  Normal   Scheduled         65s                default-scheduler  Successfully assigned default/postgres-5f676c995d-jvdrp to master
  Normal   Pulling           65s                kubelet            Pulling image "postgis/postgis:12-2.5-alpine"
  Normal   Pulled            28s                kubelet            Successfully pulled image "postgis/postgis:12-2.5-alpine" in 37.501507403s
  Normal   Pulled            11s (x2 over 27s)  kubelet            Container image "postgis/postgis:12-2.5-alpine" already present on machine
  Warning  Failed            11s (x3 over 27s)  kubelet            Error: failed to create containerd container: get apparmor_parser version: exec: "apparmor_parser": executable file not found in $PATH
jason@localhost:~/oss/courses/nd064/udaconnect$ vagrant ssh
Last login: Sat Sep  4 19:39:33 2021 from 10.0.2.2
vagrant@master:~> sudo su -
master:~ # zypper install apparmor-utils
Loading repository data...
Warning: Repository 'openSUSE-Leap-42.3-Update' appears to be outdated. Consider using a different mirror or server.
Warning: Repository 'openSUSE-Leap-42.3-Update-Non-Oss' appears to be outdated. Consider using a different mirror or server.
Reading installed packages...
Resolving package dependencies...

The following 25 NEW packages are going to be installed:
  apparmor-utils libnotify-tools libpython3_4m1_0 perl-Clone perl-DBD-SQLite perl-DBI perl-MLDBM perl-Math-Base-Convert perl-Module-Runtime perl-Net-Daemon perl-Params-Util perl-PlRPC
  perl-RPC-XML perl-SQL-Statement perl-Term-ReadKey perl-XML-LibXML perl-apparmor python3 python3-apparmor python3-base python3-gobject python3-gobject-Gdk python3-gobject-cairo python3-pip
  python3-setuptools

The following 8 recommended packages were automatically selected:
  libnotify-tools perl-Clone perl-MLDBM perl-Math-Base-Convert perl-Net-Daemon perl-PlRPC perl-SQL-Statement python3-pip

25 new packages to install.
Overall download size: 13.2 MiB. Already cached: 0 B. After the operation, additional 50.2 MiB will be used.
Continue? [y/n/...? shows all options] (y): y
Retrieving package perl-MLDBM-2.05-10.1.noarch                                                                                                          (1/25),  20.0 KiB ( 29.6 KiB unpacked)
Retrieving: perl-MLDBM-2.05-10.1.noarch.rpm ............................................................................................................................................[done]
Retrieving package perl-Math-Base-Convert-0.11-3.11.noarch                                                                                              (2/25),  38.0 KiB ( 95.3 KiB unpacked)
Retrieving: perl-Math-Base-Convert-0.11-3.11.noarch.rpm ................................................................................................................................[done]
Retrieving package perl-Module-Runtime-0.014-10.1.noarch                                                                                                (3/25),  20.3 KiB ( 29.0 KiB unpacked)
Retrieving: perl-Module-Runtime-0.014-10.1.noarch.rpm ..................................................................................................................................[done]
Retrieving package perl-Net-Daemon-0.48-18.1.noarch                                                                                                     (4/25),  42.6 KiB ( 97.0 KiB unpacked)
Retrieving: perl-Net-Daemon-0.48-18.1.noarch.rpm .......................................................................................................................................[done]
Retrieving package libnotify-tools-0.7.6-8.3.x86_64                                                                                                     (5/25),  11.0 KiB ( 14.5 KiB unpacked)
Retrieving: libnotify-tools-0.7.6-8.3.x86_64.rpm .......................................................................................................................................[done]
Retrieving package perl-Clone-0.38-6.1.x86_64                                                                                                           (6/25),  16.1 KiB ( 23.2 KiB unpacked)
Retrieving: perl-Clone-0.38-6.1.x86_64.rpm .............................................................................................................................................[done]
Retrieving package perl-Params-Util-1.07-14.1.x86_64                                                                                                    (7/25),  36.1 KiB ( 88.4 KiB unpacked)
Retrieving: perl-Params-Util-1.07-14.1.x86_64.rpm ......................................................................................................................................[done]
Retrieving package perl-Term-ReadKey-2.31-6.12.x86_64                                                                                                   (8/25),  28.0 KiB ( 61.0 KiB unpacked)
Retrieving: perl-Term-ReadKey-2.31-6.12.x86_64.rpm .....................................................................................................................................[done]
Retrieving package perl-PlRPC-0.2020-28.1.x86_64                                                                                                        (9/25),  32.7 KiB ( 67.1 KiB unpacked)
Retrieving: perl-PlRPC-0.2020-28.1.x86_64.rpm ..........................................................................................................................................[done]
Retrieving package perl-SQL-Statement-1.410-3.1.noarch                                                                                                 (10/25), 204.9 KiB (539.1 KiB unpacked)
Retrieving: perl-SQL-Statement-1.410-3.1.noarch.rpm .......................................................................................................................[done (95.6 KiB/s)]
Retrieving package perl-DBI-1.628-8.11.x86_64                                                                                                          (11/25), 765.3 KiB (  2.1 MiB unpacked)
Retrieving: perl-DBI-1.628-8.11.x86_64.rpm .............................................................................................................................................[done]
Retrieving package perl-DBD-SQLite-1.50-3.11.x86_64                                                                                                    (12/25),   1.7 MiB (  7.9 MiB unpacked)
Retrieving: perl-DBD-SQLite-1.50-3.11.x86_64.rpm .........................................................................................................................[done (879.6 KiB/s)]
Retrieving package libpython3_4m1_0-3.4.6-12.10.1.x86_64                                                                                               (13/25), 781.0 KiB (  2.3 MiB unpacked)
Retrieving: libpython3_4m1_0-3.4.6-12.10.1.x86_64.rpm .....................................................................................................................[done (15.6 KiB/s)]
Retrieving package perl-XML-LibXML-2.0019-10.1.x86_64                                                                                                  (14/25), 353.8 KiB (950.7 KiB unpacked)
Retrieving: perl-XML-LibXML-2.0019-10.1.x86_64.rpm .....................................................................................................................................[done]
Retrieving package python3-base-3.4.6-12.10.1.x86_64                                                                                                   (15/25),   4.7 MiB ( 22.9 MiB unpacked)
Retrieving: python3-base-3.4.6-12.10.1.x86_64.rpm ......................................................................................................................................[done]
Retrieving package python3-3.4.6-12.10.2.x86_64                                                                                                        (16/25),   1.8 MiB (  2.9 MiB unpacked)
Retrieving: python3-3.4.6-12.10.2.x86_64.rpm ...........................................................................................................................................[done]
Retrieving package perl-RPC-XML-0.77-9.14.noarch                                                                                                       (17/25), 202.0 KiB (664.2 KiB unpacked)
Retrieving: perl-RPC-XML-0.77-9.14.noarch.rpm .............................................................................................................................[done (15.6 KiB/s)]
Retrieving package python3-gobject-3.20.1-5.4.x86_64                                                                                                   (18/25), 501.0 KiB (  1.9 MiB unpacked)
Retrieving: python3-gobject-3.20.1-5.4.x86_64.rpm ......................................................................................................................................[done]
Retrieving package python3-setuptools-18.3.2-4.4.noarch                                                                                               (19/25), 293.3 KiB (1010.4 KiB unpacked)
Retrieving: python3-setuptools-18.3.2-4.4.noarch.rpm ...................................................................................................................................[done]
Retrieving package python3-gobject-cairo-3.20.1-5.4.x86_64                                                                                             (20/25),  65.3 KiB ( 14.4 KiB unpacked)
Retrieving: python3-gobject-cairo-3.20.1-5.4.x86_64.rpm ................................................................................................................................[done]
Retrieving package python3-pip-7.1.2-7.1.noarch                                                                                                        (21/25),   1.1 MiB (  5.2 MiB unpacked)
Retrieving: python3-pip-7.1.2-7.1.noarch.rpm ...........................................................................................................................................[done]
Retrieving package python3-gobject-Gdk-3.20.1-5.4.x86_64                                                                                               (22/25),  75.4 KiB ( 71.4 KiB unpacked)
Retrieving: python3-gobject-Gdk-3.20.1-5.4.x86_64.rpm ..................................................................................................................................[done]
Retrieving package perl-apparmor-2.10.4-19.1.x86_64                                                                                                    (23/25), 112.9 KiB (380.0 KiB unpacked)
Retrieving: perl-apparmor-2.10.4-19.1.x86_64.rpm .......................................................................................................................................[done]
Retrieving package python3-apparmor-2.10.4-19.1.x86_64                                                                                                 (24/25), 216.1 KiB (790.3 KiB unpacked)
Retrieving: python3-apparmor-2.10.4-19.1.x86_64.rpm ....................................................................................................................................[done]
Retrieving package apparmor-utils-2.10.4-19.1.noarch                                                                                                   (25/25), 132.6 KiB (225.8 KiB unpacked)
Retrieving: apparmor-utils-2.10.4-19.1.noarch.rpm ......................................................................................................................................[done]
Checking for file conflicts: ...........................................................................................................................................................[done]
( 1/25) Installing: perl-MLDBM-2.05-10.1.noarch ........................................................................................................................................[done]
( 2/25) Installing: perl-Math-Base-Convert-0.11-3.11.noarch ............................................................................................................................[done]
( 3/25) Installing: perl-Module-Runtime-0.014-10.1.noarch ..............................................................................................................................[done]
( 4/25) Installing: perl-Net-Daemon-0.48-18.1.noarch ...................................................................................................................................[done]
( 5/25) Installing: libnotify-tools-0.7.6-8.3.x86_64 ...................................................................................................................................[done]
( 6/25) Installing: perl-Clone-0.38-6.1.x86_64 .........................................................................................................................................[done]
( 7/25) Installing: perl-Params-Util-1.07-14.1.x86_64 ..................................................................................................................................[done]
( 8/25) Installing: perl-Term-ReadKey-2.31-6.12.x86_64 .................................................................................................................................[done]
( 9/25) Installing: perl-PlRPC-0.2020-28.1.x86_64 ......................................................................................................................................[done]
(10/25) Installing: perl-SQL-Statement-1.410-3.1.noarch ................................................................................................................................[done]
(11/25) Installing: perl-DBI-1.628-8.11.x86_64 .........................................................................................................................................[done]
(12/25) Installing: perl-DBD-SQLite-1.50-3.11.x86_64 ...................................................................................................................................[done]
(13/25) Installing: libpython3_4m1_0-3.4.6-12.10.1.x86_64 ..............................................................................................................................[done]
(14/25) Installing: perl-XML-LibXML-2.0019-10.1.x86_64 .................................................................................................................................[done]
(15/25) Installing: python3-base-3.4.6-12.10.1.x86_64 ..................................................................................................................................[done]
(16/25) Installing: python3-3.4.6-12.10.2.x86_64 .......................................................................................................................................[done]
(17/25) Installing: perl-RPC-XML-0.77-9.14.noarch ......................................................................................................................................[done]
(18/25) Installing: python3-gobject-3.20.1-5.4.x86_64 ..................................................................................................................................[done]
(19/25) Installing: python3-setuptools-18.3.2-4.4.noarch ...............................................................................................................................[done]
Additional rpm output:
update-alternatives: using /usr/bin/easy_install-3.4 to provide /usr/bin/easy_install (easy_install) in auto mode


(20/25) Installing: python3-gobject-cairo-3.20.1-5.4.x86_64 ............................................................................................................................[done]
(21/25) Installing: python3-pip-7.1.2-7.1.noarch .......................................................................................................................................[done]
Additional rpm output:
update-alternatives: using /usr/bin/pip3.4 to provide /usr/bin/pip (pip) in auto mode


(22/25) Installing: python3-gobject-Gdk-3.20.1-5.4.x86_64 ..............................................................................................................................[done]
(23/25) Installing: perl-apparmor-2.10.4-19.1.x86_64 ...................................................................................................................................[done]
(24/25) Installing: python3-apparmor-2.10.4-19.1.x86_64 ................................................................................................................................[done]
(25/25) Installing: apparmor-utils-2.10.4-19.1.noarch ..................................................................................................................................[done]
master:~ # zypper install -t pattern apparmor
Loading repository data...
Warning: Repository 'openSUSE-Leap-42.3-Update' appears to be outdated. Consider using a different mirror or server.
Warning: Repository 'openSUSE-Leap-42.3-Update-Non-Oss' appears to be outdated. Consider using a different mirror or server.
Reading installed packages...
Resolving package dependencies...

The following 7 NEW packages are going to be installed:
  apparmor-abstractions apparmor-docs apparmor-parser apparmor-profiles patterns-openSUSE-apparmor patterns-openSUSE-apparmor_opt yast2-apparmor

The following 2 NEW patterns are going to be installed:
  apparmor apparmor_opt

The following 2 recommended packages were automatically selected:
  patterns-openSUSE-apparmor_opt yast2-apparmor

7 new packages to install.
Overall download size: 989.1 KiB. Already cached: 0 B. After the operation, additional 1.9 MiB will be used.
Continue? [y/n/...? shows all options] (y): y
Retrieving package yast2-apparmor-3.2.0-1.1.noarch                                                                                                       (1/7),  49.7 KiB (218.3 KiB unpacked)
Retrieving: yast2-apparmor-3.2.0-1.1.noarch.rpm ........................................................................................................................................[done]
Retrieving package apparmor-docs-2.10.4-19.1.noarch                                                                                                      (2/7), 314.3 KiB (457.2 KiB unpacked)
Retrieving: apparmor-docs-2.10.4-19.1.noarch.rpm .......................................................................................................................................[done]
Retrieving package apparmor-parser-2.10.4-19.1.x86_64                                                                                                    (3/7), 360.5 KiB (951.4 KiB unpacked)
Retrieving: apparmor-parser-2.10.4-19.1.x86_64.rpm .....................................................................................................................................[done]
Retrieving package apparmor-abstractions-2.10.4-19.1.noarch                                                                                              (4/7),  72.1 KiB (100.8 KiB unpacked)
Retrieving: apparmor-abstractions-2.10.4-19.1.noarch.rpm ...............................................................................................................................[done]
Retrieving package apparmor-profiles-2.10.4-19.1.noarch                                                                                                  (5/7),  84.2 KiB (177.4 KiB unpacked)
Retrieving: apparmor-profiles-2.10.4-19.1.noarch.rpm ...................................................................................................................................[done]
Retrieving package patterns-openSUSE-apparmor_opt-20170518-6.1.x86_64                                                                                    (6/7),  54.0 KiB (   58   B unpacked)
Retrieving: patterns-openSUSE-apparmor_opt-20170518-6.1.x86_64.rpm ........................................................................................................[done (15.6 KiB/s)]
Retrieving package patterns-openSUSE-apparmor-20170518-6.1.x86_64                                                                                        (7/7),  54.2 KiB (   54   B unpacked)
Retrieving: patterns-openSUSE-apparmor-20170518-6.1.x86_64.rpm .........................................................................................................................[done]
Checking for file conflicts: ...........................................................................................................................................................[done]
(1/7) Installing: yast2-apparmor-3.2.0-1.1.noarch ......................................................................................................................................[done]
(2/7) Installing: apparmor-docs-2.10.4-19.1.noarch .....................................................................................................................................[done]
(3/7) Installing: apparmor-parser-2.10.4-19.1.x86_64 ...................................................................................................................................[done]
Additional rpm output:
Created symlink from /etc/systemd/system/multi-user.target.wants/apparmor.service to /usr/lib/systemd/system/apparmor.service.


(4/7) Installing: apparmor-abstractions-2.10.4-19.1.noarch .............................................................................................................................[done]
(5/7) Installing: apparmor-profiles-2.10.4-19.1.noarch .................................................................................................................................[done]
(6/7) Installing: patterns-openSUSE-apparmor_opt-20170518-6.1.x86_64 ...................................................................................................................[done]
(7/7) Installing: patterns-openSUSE-apparmor-20170518-6.1.x86_64 .......................................................................................................................[done]
master:~ # 
master:~ # exit
logout
vagrant@master:~> exit
logout
Connection to 127.0.0.1 closed.
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
udaconnect-api-89dbffbf9-qhxzx    1/1     Running   0          3m14s
postgres-5f676c995d-jvdrp         1/1     Running   0          3m20s
udaconnect-app-6ddf7c79fb-kqfmb   1/1     Running   0          3m9s
jason@localhost:~/oss/courses/nd064/udaconnect$ 
jason@localhost:~/oss/courses/nd064/udaconnect$ sh scripts/run_db_command.sh postgres-5f676c995d-jvdrp
CREATE TABLE
CREATE TABLE
CREATE INDEX
CREATE INDEX
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
INSERT 0 1
jason@localhost:~/oss/courses/nd064/udaconnect$ 
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get pods
NAME                              READY   STATUS    RESTARTS   AGE
udaconnect-api-89dbffbf9-qhxzx    1/1     Running   0          5m38s
postgres-5f676c995d-jvdrp         1/1     Running   0          5m44s
udaconnect-app-6ddf7c79fb-kqfmb   1/1     Running   0          5m33s
jason@localhost:~/oss/courses/nd064/udaconnect$ kubectl get services
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes       ClusterIP   10.43.0.1       <none>        443/TCP          21m
postgres         NodePort    10.43.242.130   <none>        5432:30915/TCP   5m45s
udaconnect-api   NodePort    10.43.159.0     <none>        5000:30001/TCP   5m39s
udaconnect-app   NodePort    10.43.65.125    <none>        3000:30000/TCP   5m34s
jason@localhost:~/oss/courses/nd064/udaconnect$ 
```
