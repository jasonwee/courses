#!/bin/bash

export DB_USERNAME=ct_connections
export DB_PASSWORD=wowimsosecure
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=connections
export PERSON_SERVICE_ENDPOINT="http://svc-api-person:5000/"
flask run --host 0.0.0.0
