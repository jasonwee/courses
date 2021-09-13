FROM python:3.9-alpine

WORKDIR /app

RUN apk add --no-cache gcc g++ musl-dev linux-headers libc-dev postgresql-dev
COPY . .
RUN pip install -r requirements.txt

ENTRYPOINT ["python", "location_consumer_service.py"]
