#!/bin/bash

export DB_USERNAME=ct_person
export DB_PASSWORD=wowimsosecure
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=person
flask run --host 0.0.0.0
