# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datapipe/scores/services.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common import common_pb2 as common_dot_common__pb2
from datahub.scores import scores_pb2 as datahub_dot_scores_dot_scores__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datapipe/scores/services.proto',
  package='containersai.datapipe.scores',
  syntax='proto3',
  serialized_options=_b('Z,github.com/containers-ai/api/datapipe/scores'),
  serialized_pb=_b('\n\x1e\x64\x61tapipe/scores/services.proto\x12\x1c\x63ontainersai.datapipe.scores\x1a\x13\x63ommon/common.proto\x1a\x1b\x64\x61tahub/scores/scores.proto\x1a\x17google/rpc/status.proto\"d\n$ListSimulatedSchedulingScoresRequest\x12<\n\x0fquery_condition\x18\x01 \x01(\x0b\x32#.containersai.common.QueryCondition\"\x92\x01\n%ListSimulatedSchedulingScoresResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x45\n\x06scores\x18\x02 \x03(\x0b\x32\x35.containersai.datahub.scores.SimulatedSchedulingScore2\xbc\x01\n\rScoresService\x12\xaa\x01\n\x1dListSimulatedSchedulingScores\x12\x42.containersai.datapipe.scores.ListSimulatedSchedulingScoresRequest\x1a\x43.containersai.datapipe.scores.ListSimulatedSchedulingScoresResponse\"\x00\x42.Z,github.com/containers-ai/api/datapipe/scoresb\x06proto3')
  ,
  dependencies=[common_dot_common__pb2.DESCRIPTOR,datahub_dot_scores_dot_scores__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_LISTSIMULATEDSCHEDULINGSCORESREQUEST = _descriptor.Descriptor(
  name='ListSimulatedSchedulingScoresRequest',
  full_name='containersai.datapipe.scores.ListSimulatedSchedulingScoresRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='containersai.datapipe.scores.ListSimulatedSchedulingScoresRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=139,
  serialized_end=239,
)


_LISTSIMULATEDSCHEDULINGSCORESRESPONSE = _descriptor.Descriptor(
  name='ListSimulatedSchedulingScoresResponse',
  full_name='containersai.datapipe.scores.ListSimulatedSchedulingScoresResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.datapipe.scores.ListSimulatedSchedulingScoresResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scores', full_name='containersai.datapipe.scores.ListSimulatedSchedulingScoresResponse.scores', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=242,
  serialized_end=388,
)

_LISTSIMULATEDSCHEDULINGSCORESREQUEST.fields_by_name['query_condition'].message_type = common_dot_common__pb2._QUERYCONDITION
_LISTSIMULATEDSCHEDULINGSCORESRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTSIMULATEDSCHEDULINGSCORESRESPONSE.fields_by_name['scores'].message_type = datahub_dot_scores_dot_scores__pb2._SIMULATEDSCHEDULINGSCORE
DESCRIPTOR.message_types_by_name['ListSimulatedSchedulingScoresRequest'] = _LISTSIMULATEDSCHEDULINGSCORESREQUEST
DESCRIPTOR.message_types_by_name['ListSimulatedSchedulingScoresResponse'] = _LISTSIMULATEDSCHEDULINGSCORESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListSimulatedSchedulingScoresRequest = _reflection.GeneratedProtocolMessageType('ListSimulatedSchedulingScoresRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTSIMULATEDSCHEDULINGSCORESREQUEST,
  __module__ = 'datapipe.scores.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datapipe.scores.ListSimulatedSchedulingScoresRequest)
  ))
_sym_db.RegisterMessage(ListSimulatedSchedulingScoresRequest)

ListSimulatedSchedulingScoresResponse = _reflection.GeneratedProtocolMessageType('ListSimulatedSchedulingScoresResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTSIMULATEDSCHEDULINGSCORESRESPONSE,
  __module__ = 'datapipe.scores.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datapipe.scores.ListSimulatedSchedulingScoresResponse)
  ))
_sym_db.RegisterMessage(ListSimulatedSchedulingScoresResponse)


DESCRIPTOR._options = None

_SCORESSERVICE = _descriptor.ServiceDescriptor(
  name='ScoresService',
  full_name='containersai.datapipe.scores.ScoresService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=391,
  serialized_end=579,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListSimulatedSchedulingScores',
    full_name='containersai.datapipe.scores.ScoresService.ListSimulatedSchedulingScores',
    index=0,
    containing_service=None,
    input_type=_LISTSIMULATEDSCHEDULINGSCORESREQUEST,
    output_type=_LISTSIMULATEDSCHEDULINGSCORESRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SCORESSERVICE)

DESCRIPTOR.services_by_name['ScoresService'] = _SCORESSERVICE

# @@protoc_insertion_point(module_scope)
