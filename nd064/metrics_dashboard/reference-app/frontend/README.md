### setup python env
```
python3.9 -m venv py39_env
source py39_env/bin/activate
pip3 install --upgrade pip
pip install -r requirements.txt
```


### to run locally
```
$ source py39_env/bin/activate
$ gunicorn --access-logfile - --error-logfile - -w 4 -b 0.0.0.0:8080 app:app
```


### docker build and push
```
docker build --no-cache -t jasonwee/metrics-dashboard-frontend .
docker push jasonwee/metrics-dashboard-frontend
```

### forward port in k8s
```
vagrant@localhost:~> kubectl port-forward svc/frontend-service 8080:8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
^Cvagrant@localhost:~>
```
