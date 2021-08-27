import json
from flask import Flask, jsonify, request

from .services import retrieve_orders, create_order

app = Flask(__name__)


@app.route('/api/orders/computers', methods=['POST'])
def create():
    """
    {
    "created_by": <user_id: str>,
    "status": <status_enum: str>, // QUEUED, PROCESSING, COMPLETED, FAILED, CANCELLED
    "created_at": <isoformat_timestamp: str>,
    "equipment": [
        <equipment: str>
    ]
    }
    """
    create_order()
    return jsonify({
        "order_id": <order_id: str>,
        "created_by": <user_id: str>,
        "status": <status_enum: str>,
        "created_at": <isoformat_timestamp: str>,
        "equipment": [
            <equipment: str>
        ]
        })

@app.route('/api/orders/computers', methods=['GET'])
def list():
    orders = retrieve_orders()
    return jsonify("computer_orders": [
        {
          "order_id": <order_id: str>,
          "created_by": <user_id: str>,
          "status": <status_enum: str>,
          "created_at": <isoformat_timestamp: str>,
          "equipment": [
            <equipment: str>
          ]
        }
        ])

if __name__ == '__main__':
    app.run()
