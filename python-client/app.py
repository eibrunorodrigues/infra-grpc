import grpc
from models.receiver import Receiver as ReceiverModel
from models.sender import Message, DestinyType
from models.orchestrator import RouterPrefix, RouterType
from typing import List
import time

import receiver.receiver_pb2_grpc as pb2_grpc
import receiver.receiver_pb2 as pb2
from orchestrator import orchestrator_pb2_grpc
from orchestrator import orchestrator_pb2
from sender import sender_pb2_grpc
from sender import sender_pb2


class BrokerClient(object):
    """
    Client for gRPC functionality
    py -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./receiver/receiver.proto
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
        self.orchestrator = orchestrator_pb2_grpc.OrchestratorStub(
            self.channel)
        self.__messages: List[Message] = []

    def publish_message(self, message: Message):
        assert isinstance(message, Message)
        self.__messages.append(message)

    def acknowledge_message(self, message_id: int):
        return self.receiver.AcknowledgeMessage(
            pb2.MessageId(message_id=message_id)
        ).status

    def reject_message(self, message_id: int, requeue: bool):
        return self.receiver.RejectMessage(
            pb2.Reject(
                id=pb2.MessageId(message_id=message_id),
                requeue=requeue
            )
        ).status

    def check_if_queue_exists(self, name: str):
        return self.orchestrator.CheckIfQueueExists(
            orchestrator_pb2.Name(name=name)
        ).status

    def check_if_router_exists(self, name: str):
        return self.orchestrator.CheckIfRouterExists(
            orchestrator_pb2.Name(name=name)
        ).status

    def create_queue(self, name: str, create_dlq: bool, exclusive: bool):
        return self.orchestrator.CreateQueue(
            orchestrator_pb2.QueueCreation(queue_name=name,
                                           create_dlq=create_dlq,
                                           is_exclusive=exclusive
                                           )).name

    def create_router(self, name: str, prefix: RouterPrefix, router_type: RouterType):
        return self.orchestrator.CreateRouter(
            orchestrator_pb2.RouterCreation(
                router_name=name,
                prefix=orchestrator_pb2.RouterPrefix.Value(prefix),
                router_type=orchestrator_pb2.RouterType.Value(router_type)
            )
        ).name

    def delete_queue(self, name: str):
        return self.orchestrator.DeleteQueue(
            orchestrator_pb2.Name(name=name)
        ).status

    def delete_router(self, name: str):
        return self.orchestrator.DeleteRouter(
            orchestrator_pb2.Name(name=name)
        ).status

    def bind_queue_to_router(self, destination: str, source: str, filters: dict):
        return self.orchestrator.BindQueueToRouter(
            orchestrator_pb2.Bind(
                destination=destination,
                source=source,
                filters=filters
            )
        ).status

    def bind_router_to_router(self, destination: str, source: str, filters: dict):
        return self.orchestrator.BindRouterToRouter(
            orchestrator_pb2.Bind(
                destination=destination,
                source=source,
                filters=filters
            )
        ).status

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


def orchestrator():
    from threading import current_thread
    name = current_thread().name
    client = BrokerClient()
    response = client.create_router("testing_bruno", prefix=RouterPrefix.TOPIC, router_type=RouterType.HEADERS)
    client.bind_queue_to_router(response, "BRUNO_SHOVEL_TEST", {"branch_id": "982"})
    print(str(response))


def receive():
    from threading import current_thread
    from threading import Thread

    name = current_thread().name
    client = BrokerClient()
    channel = client.receive_messages(queue_name="TESTING_PROPERLY")
    queue_created = client.create_queue("EXAMPLE_BR_MIGRATION", True, False)
    Thread(target=client.send).start()

    for item in channel:
        model = ReceiverModel(item)
        client.publish_message(message=Message(
            destination=queue_created,
            content=model.content,
            filters=model.filters,
            destiny_type=DestinyType.QUEUE
        ))
        result = client.acknowledge_message(model.message_id)
        print(f'[THREAD-{name}] : {model.filters} {result}')


def send():
    from threading import Thread
    client = BrokerClient()
    Thread(target=client.send).start()
    for item in range(199999999):
        client.publish_message(message=Message(
            destination="EXAMPLE_BR_MIGRATION",
            content={"accx": item},
            filters={},
            destiny_type=DestinyType.QUEUE
        ))


if __name__ == '__main__':
    receive()
    # send()
    # orchestrator()
