# Flask Starter
This is a bare-bones Flask application that showcases how simple it can be to set up a REST API. It also serves as a starter template for users to begin experimenting with writing their own REST API endpoints.

Flask applications can look very different depending on how they are structured and implemented. This application is built to have the minimum number of dependencies to run.

## Running the app
1. Install Flask: `pip install Flask kafka-python`
2. Run the app: `flask run`

The application should be available at `localhost:5000`.

## setup topic
export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
bin/kafka-topics.sh --create --topic items --bootstrap-server localhost:9092


## kafka consumer
export JAVA_HOME=/usr/lib/jvm/jdk-11.0.5/
bin/kafka-console-consumer.sh --topic items --from-beginning --bootstrap-server localhost:9092

## test send via postman

http://localhost:5000/api/orders/computers

POST

{
  "id": "3",
  "status": "QUEUED",
  "created_at": "2020-10-16T10:31:10.123456",
  "created_by": "USER13",
  "equipment": [
     "KEYBOARD",
     "MOUSE"
  ]
}
