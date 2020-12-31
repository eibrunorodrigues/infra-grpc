# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from sender import sender_pb2 as sender_dot_sender__pb2


class SenderStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Publish = channel.stream_unary(
                '/sender.Sender/Publish',
                request_serializer=sender_dot_sender__pb2.SenderArgs.SerializeToString,
                response_deserializer=sender_dot_sender__pb2.Bool.FromString,
                )


class SenderServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Publish(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SenderServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Publish': grpc.stream_unary_rpc_method_handler(
                    servicer.Publish,
                    request_deserializer=sender_dot_sender__pb2.SenderArgs.FromString,
                    response_serializer=sender_dot_sender__pb2.Bool.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'sender.Sender', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Sender(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Publish(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/sender.Sender/Publish',
            sender_dot_sender__pb2.SenderArgs.SerializeToString,
            sender_dot_sender__pb2.Bool.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)