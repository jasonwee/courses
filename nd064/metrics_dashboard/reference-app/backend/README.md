### setup python env
```
python3.9 -m venv py39_env
source py39_env/bin/activate
pip3 install --upgrade pip
pip install -r requirements.txt
```


### to run locally
```
$ python3.9 -m venv py39_env
$ gunicorn --access-logfile - --error-logfile - -w 4 -b 0.0.0.0:8080 app:app
```
