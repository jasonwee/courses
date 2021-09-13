FROM python:3.9-alpine

WORKDIR /app

RUN apk add --no-cache gcc g++ musl-dev linux-headers geos libc-dev postgresql-dev
COPY . .
RUN pip install -r requirements.txt

EXPOSE 5000

CMD ["flask", "run", "--host", "0.0.0.0"]
