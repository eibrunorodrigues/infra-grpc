from typing import Dict
import json


class Receiver:
    filters: Dict[str, any]
    destination: str
    content: Dict[str, any]
    is_a_redelivery: bool

    def __init__(self, *args, **kwargs):
        self.filters = args[0].filters
        self.destination = args[0].destination
        self.content = json.loads(args[0].content)
        self.is_a_redelivery = args[0].is_a_redelivery
