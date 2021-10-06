from flask import Flask, render_template, request, jsonify

import pymongo
import logging
from flask_pymongo import PyMongo
from flask_cors import CORS
from prometheus_flask_exporter import PrometheusMetrics
from prometheus_flask_exporter.multiprocess import GunicornInternalPrometheusMetrics

from jaeger_client import Config
from jaeger_client.metrics.prometheus import PrometheusMetricsFactory
from flask_opentracing import FlaskTracing

## jaeger
def config_tracer():
    config = Config(
           config = {
                'sampler': {
                'type': 'const',
                'param': 1,
            },
            'logging': True,
        },
        service_name="service_backend",
        validate=True,
        metrics_factory=PrometheusMetricsFactory(service_name_label="service_backend")
    )
    return config.initialize_tracer()
## jaeger

app = Flask(__name__)
metrics = GunicornInternalPrometheusMetrics(app)
CORS(app)

jaeger_tracer = config_tracer()
tracing = FlaskTracing(jaeger_tracer, False, app)


app.config['MONGO_DBNAME'] = 'example-mongodb'
app.config['MONGO_URI'] = 'mongodb://example-mongodb-svc.default.svc.cluster.local:27017/example-mongodb'

mongo = PyMongo(app)

logging.basicConfig(level=logging.INFO)
logging.info("enter backend app")



@app.route('/')
@tracing.trace()
def homepage():
    return "Hello World"

@app.route('/api')
@tracing.trace()
def my_api():
    answer = "something"
    return jsonify(response=answer)

@app.route('/star', methods=['POST'])
@tracing.trace()
def add_star():
  star = mongo.db.stars
  name = request.json['name']
  distance = request.json['distance']
  star_id = star.insert({'name': name, 'distance': distance})
  new_star = star.find_one({'_id': star_id })
  output = {'name' : new_star['name'], 'distance' : new_star['distance']}
  return jsonify({'result' : output})

if __name__ == "__main__":
    app.run(debug=True)
