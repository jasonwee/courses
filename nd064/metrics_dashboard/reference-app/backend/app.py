from flask import Flask, render_template, request, jsonify

import pymongo
import logging
from flask_pymongo import PyMongo
from flask_cors import CORS
from prometheus_flask_exporter import PrometheusMetrics
from prometheus_flask_exporter.multiprocess import GunicornInternalPrometheusMetrics



app = Flask(__name__)
metrics = GunicornInternalPrometheusMetrics(app)
CORS(app)

app.config['MONGO_DBNAME'] = 'example-mongodb'
app.config['MONGO_URI'] = 'mongodb://example-mongodb-svc.default.svc.cluster.local:27017/example-mongodb'

mongo = PyMongo(app)

logging.basicConfig(level=logging.INFO)
logging.info("enter backend app")



@app.route('/')
def homepage():
    return "Hello World"

@app.route('/api')
def my_api():
    answer = "something"
    return jsonify(response=answer)

@app.route('/star', methods=['POST'])
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
