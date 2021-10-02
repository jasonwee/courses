from flask import Flask, render_template, request, jsonify

import pymongo
from flask_pymongo import PyMongo
from prometheus_flask_exporter import PrometheusMetrics

app = Flask(__name__)

app.config['MONGO_DBNAME'] = 'example-mongodb'
app.config['MONGO_URI'] = 'mongodb://example-mongodb-svc.default.svc.cluster.local:27017/example-mongodb'

mongo = PyMongo(app)
metrics = PrometheusMetrics(app)
metrics.info("app_info", "Backend App", version="1.0.0")



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
    app.run()
