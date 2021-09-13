
### setup python env
```
python3.9 -m venv py39_env
source py39_env/bin/activate
pip3 install --upgrade pip
pip install -r requirements.txt

```

### run locally
```
$ run-local.sh
```


## Docker commands used to build the application 
docker build --no-cache -t jasonwee/udaconnect-connection .


## Docker commands used to run the application
docker run -e DB_USERNAME=ct_connections -e DB_PASSWORD=wowimsosecure -e DB_HOST=localhost -e DB_PORT=5432 -e DB_NAME=connections -e PERSON_SERVICE_ENDPOINT="http://svc-api-person:5000/" -d -p 5000:5000 --name udaconnect-connection --rm jasonwee/udaconnect-connection


## Docker commands used to get the application logs
docker logs -f udaconnect-connection


## push to dockerhub
docker push jasonwee/udaconnect-connection
