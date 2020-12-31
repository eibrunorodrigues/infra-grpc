from typing import Dict, Union
import json

# header
# destiny_type
# destination
# content


class DestinyType:
    QUEUE = "QUEUE"
    TOPIC = "TOPIC"


class Message:
    filters: Dict[str, any] = {}
    destination: str
    content: Dict[str, any]
    destiny_type: DestinyType = DestinyType.QUEUE

    def __init__(
        self,
        destination: str,
        content: Union[Dict, str],
        filters: Dict[str, any],
        destiny_type: DestinyType = DestinyType.QUEUE
    ):
        assert isinstance(filters, dict)
        assert isinstance(destination, str)
        assert isinstance(content, dict) or isinstance(content, str)
        assert isinstance(destiny_type, str)

        self.filters = filters
        self.destination = destination
        self.content = json.dumps(content, indent=2).encode('utf-8')
        self.destiny_type = destiny_type
