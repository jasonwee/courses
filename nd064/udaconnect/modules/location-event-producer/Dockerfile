FROM python:3.9-alpine

WORKDIR /app

RUN apk add --no-cache gcc g++ musl-dev linux-headers libc-dev
COPY . .
RUN pip install -r requirements.txt

EXPOSE 5005

ENTRYPOINT ["python", "grpc_server_startup.py"] 
