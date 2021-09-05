```
https://kafka.apache.org/quickstart


start zookeeper
user@localhost:~/kafka/kafka_2.13-2.8.0$ export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/zookeeper-server-start.sh config/zookeeper.properties 
[2021-08-30 20:03:00,485] INFO Reading configuration from: config/zookeeper.properties (org.apache.zookeeper.serve


---

start kafka

user@localhost:~/kafka/kafka_2.13-2.8.0$ export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/kafka-server-start.sh config/server.properties 
[2021-08-30 20:04:22,426] INFO Registered kafka:type=kafka.Log4jController MBean (kafka.utils.Log4jControllerRegistration$)
[2021-08-30 20:04:22,705] INFO Setting -D jdk.tls.rejectClientInitiatedRenegotiation=true to disable client-initiated TLS renegotia

---

create topic

user@localhost:~/kafka/kafka_2.13-2.8.0$ export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/kafka-topics.sh --create --topic items --bootstrap-server localhost:9092
Created topic items.

---

test consumer (receive)


user@localhost:~/kafka/kafka_2.13-2.8.0$ export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/kafka-console-consumer.sh --topic items --from-beginning --bootstrap-server localhost:9092
hello, world!
{"value": 1}
exit
quit
Test Message!!!

---

test producer (send)

user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/kafka-console-consumer.sh --topic items --from-beginning --bootstrap-server localhost:9092
hello, world!
{"value": 1}
exit
quit
Test Message!!!

---

user@localhost:~/kafka/kafka_2.13-2.8.0$ export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
user@localhost:~/kafka/kafka_2.13-2.8.0$ bin/kafka-console-producer.sh --topic items --bootstrap-server localhost:9092
>hello, world!
>{"value": 1}
>exit
>quit
>user@localhost:~/kafka/kafka_2.13-2.8.0$ to^C
user@localhost:~/kafka/kafka_2.13-2.8.0$ 


---

how to use python receive

user@localhost:~/courses/kafka_demo$ cat README.md 
## kafka
`python3.9 -m venv py39_env`
`source py39_env/bin/activate`
`pip3 install --upgrade pip`
`pip install `

user@localhost:~/courses/kafka_demo$ python3.9 -m venv py39_env
user@localhost:~/courses/kafka_demo$ source py39_env/bin/activate
(py39_env) user@localhost:~/courses/kafka_demo$ pip3 install --upgrade pip
Requirement already satisfied: pip in ./py39_env/lib/python3.9/site-packages (20.3.4)
Collecting pip
  Using cached pip-21.2.4-py3-none-any.whl (1.6 MB)
Installing collected packages: pip
  Attempting uninstall: pip
    Found existing installation: pip 20.3.4
    Uninstalling pip-20.3.4:
      Successfully uninstalled pip-20.3.4
Successfully installed pip-21.2.4
(py39_env) user@localhost:~/courses/kafka_demo$ pip install kafka-python
Collecting kafka-python
  Downloading kafka_python-2.0.2-py2.py3-none-any.whl (246 kB)
     |████████████████████████████████| 246 kB 1.1 MB/s 
Installing collected packages: kafka-python
Successfully installed kafka-python-2.0.2
(py39_env) user@localhost:~/courses/kafka_demo$ vim README.md 
(py39_env) user@localhost:~/courses/kafka_demo$ ls
total 8.0K
drwxr-xr-x 6 jason jason 4.0K Aug 30 20:36 py39_env
-rw-r--r-- 1 jason jason  126 Aug 30 20:37 README.md
(py39_env) user@localhost:~/courses/kafka_demo$ touch producer.py consumer.py
(py39_env) user@localhost:~/courses/kafka_demo$ vim producer.py 
(py39_env) user@localhost:~/courses/kafka_demo$ vim consumer.py
(py39_env) user@localhost:~/courses/kafka_demo$ python consumer.py 
ConsumerRecord(topic='items', partition=0, offset=5, timestamp=1630327895253, timestamp_type=0, key=None, value=b'Test Message!!!', headers=[], checksum=None, serialized_key_size=-1, serialized_value_size=15, serialized_header_size=-1)



---

test python sending

user@localhost:~/courses/kafka_demo$ source py39_env/bin/activate
(py39_env) user@localhost:~/courses/kafka_demo$ ls
total 12K
drwxr-xr-x 6 jason jason 4.0K Aug 30 20:36 py39_env
-rw-r--r-- 1 jason jason  126 Aug 30 20:37 README.md
-rw-r--r-- 1 jason jason    0 Aug 30 20:37 producer.py
-rw-r--r-- 1 jason jason  136 Aug 30 20:47 consumer.py
(py39_env) user@localhost:~/courses/kafka_demo$ vim producer.py 
(py39_env) user@localhost:~/courses/kafka_demo$ python producer.py 
(py39_env) user@localhost:~/courses/kafka_demo$ 
```
