import grpc
import receiver.receiver_pb2_grpc as pb2_grpc
import receiver.receiver_pb2 as pb2
from models.receiver import Receiver as ReceiverModel
from models.sender import Message, DestinyType
from typing import List
import time

from sender import sender_pb2_grpc
from sender import sender_pb2


class BrokerClient(object):
    """
    Client for gRPC functionality
    """

    def __init__(self):
        self.host = 'localhost'
        self.server_port = 5052

        # instantiate a channel
        self.channel = grpc.insecure_channel(
            '{}:{}'.format(self.host, self.server_port))

        # bind the client and the server
        self.receiver = pb2_grpc.ReceiverStub(self.channel)
        self.sender = sender_pb2_grpc.SenderStub(self.channel)
        self.__messages: List[Message] = []

    def publish_message(self, message: Message):
        assert isinstance(message, Message)
        self.__messages.append(message)

    def receive_messages(self, queue_name):
        """
        Client function to call the rpc for GetServerResponse
        """
        message = pb2.ReceiverArgs(queue_name=queue_name)
        print(f'{message}')
        return self.receiver.Receive(message)

    def request_generator(self):
        # Initialization code here.
        while True:
            while len(self.__messages) == 0:
                time.sleep(1)
            # Parameters here can be set as arguments to the generator.
            content: Message = self.__messages.pop(0)
            request2 = sender_pb2.SenderArgs(
                header=content.filters,
                destiny_type=sender_pb2.Destiny.Value(content.destiny_type),
                destination=content.destination,
                content=content.content
            )
            yield request2

    def send(self):
        self.sender.Publish(request_iterator=self.request_generator())


def receive():
    from threading import current_thread
    name = current_thread().name
    client = BrokerClient()
    channel = client.receive_messages(queue_name="EXAMPLES")
    for item in channel:
        model = ReceiverModel(item)
        print(f'[THREAD-{name}] : {model.filters}')


def send():
    from threading import Thread
    client = BrokerClient()
    Thread(target=client.send).start()
    for item in range(199999999):
        client.publish_message(message=Message(
            destination="EXAMPLES",
            content={"accx": item},
            filters={},
            destiny_type=DestinyType.QUEUE
        ))


if __name__ == '__main__':
    # receive()
    send()
