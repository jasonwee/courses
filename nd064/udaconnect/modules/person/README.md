
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
docker build -t jasonwee/udaconnect-person .


## Docker commands used to run the application
docker run -e DB_USERNAME=ct_person -e DB_PASSWORD=wowimsosecure -e DB_HOST=localhost -e DB_PORT=5432 -e DB_NAME=person -d -p 3111:5000 --name udaconnect-person --rm jasonwee/udaconnect-person


## Docker commands used to get the application logs
docker logs -f $(docker ps | grep udaperson-person | tr " " "\\n" | tail -1)


## push to dockerhub
docker push jasonwee/udaconnect-person
