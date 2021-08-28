from .enums import Status


def create_order(order_data):
    """
    This is a stubbed method of retrieving a resource. It doesn't actually do anything.
    """
    # Do something to create the resource
    return order_data


def retrieve_orders():
    """
    This is a stubbed method of retrieving multiple resources. It doesn't actually do anything.
    """
    return [
        {
            "order_id": 1,
            "created_by": "justin",
            "status": Status.Completed.value,
            "created_at": "2020-09-28T08:56:44",
            "equipment": [
                "KEYBOARD"
            ]
        },
        {
            "order_id": 2,
            "created_by": "tom",
            "status": Status.Queued.value,
            "created_at": "2020-09-28T09:56:44",
            "equipment": [
               "MOUSE",
               "WEBCAM"
            ]
        }
    ]


