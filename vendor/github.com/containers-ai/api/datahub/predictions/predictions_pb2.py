# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/predictions/predictions.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from datahub.metrics import types_pb2 as datahub_dot_metrics_dot_types__pb2
from datahub.resources import types_pb2 as datahub_dot_resources_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/predictions/predictions.proto',
  package='containersai.datahub.predictions',
  syntax='proto3',
  serialized_options=_b('Z0github.com/containers-ai/api/datahub/predictions'),
  serialized_pb=_b('\n%datahub/predictions/predictions.proto\x12 containersai.datahub.predictions\x1a\x1b\x64\x61tahub/metrics/types.proto\x1a\x1d\x64\x61tahub/resources/types.proto\"\xb1\x05\n\x13\x43ontainerPrediction\x12\x0c\n\x04name\x18\x01 \x01(\t\x12g\n\x12predicted_raw_data\x18\x02 \x03(\x0b\x32K.containersai.datahub.predictions.ContainerPrediction.PredictedRawDataEntry\x12u\n\x19predicted_upperbound_data\x18\x03 \x03(\x0b\x32R.containersai.datahub.predictions.ContainerPrediction.PredictedUpperboundDataEntry\x12u\n\x19predicted_lowerbound_data\x18\x04 \x03(\x0b\x32R.containersai.datahub.predictions.ContainerPrediction.PredictedLowerboundDataEntry\x1a\x61\n\x15PredictedRawDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\x1ah\n\x1cPredictedUpperboundDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\x1ah\n\x1cPredictedLowerboundDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\"\xae\x01\n\rPodPrediction\x12G\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32..containersai.datahub.resources.NamespacedName\x12T\n\x15\x63ontainer_predictions\x18\x02 \x03(\x0b\x32\x35.containersai.datahub.predictions.ContainerPrediction\"\xb3\x05\n\x0eNodePrediction\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0cis_scheduled\x18\x02 \x01(\x08\x12\x62\n\x12predicted_raw_data\x18\x03 \x03(\x0b\x32\x46.containersai.datahub.predictions.NodePrediction.PredictedRawDataEntry\x12p\n\x19predicted_upperbound_data\x18\x04 \x03(\x0b\x32M.containersai.datahub.predictions.NodePrediction.PredictedUpperboundDataEntry\x12p\n\x19predicted_lowerbound_data\x18\x05 \x03(\x0b\x32M.containersai.datahub.predictions.NodePrediction.PredictedLowerboundDataEntry\x1a\x61\n\x15PredictedRawDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\x1ah\n\x1cPredictedUpperboundDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\x1ah\n\x1cPredictedLowerboundDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\x05\x12\x37\n\x05value\x18\x02 \x01(\x0b\x32(.containersai.datahub.metrics.MetricData:\x02\x38\x01\x42\x32Z0github.com/containers-ai/api/datahub/predictionsb\x06proto3')
  ,
  dependencies=[datahub_dot_metrics_dot_types__pb2.DESCRIPTOR,datahub_dot_resources_dot_types__pb2.DESCRIPTOR,])




_CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY = _descriptor.Descriptor(
  name='PredictedRawDataEntry',
  full_name='containersai.datahub.predictions.ContainerPrediction.PredictedRawDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedRawDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedRawDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=516,
  serialized_end=613,
)

_CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY = _descriptor.Descriptor(
  name='PredictedUpperboundDataEntry',
  full_name='containersai.datahub.predictions.ContainerPrediction.PredictedUpperboundDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedUpperboundDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedUpperboundDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=615,
  serialized_end=719,
)

_CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY = _descriptor.Descriptor(
  name='PredictedLowerboundDataEntry',
  full_name='containersai.datahub.predictions.ContainerPrediction.PredictedLowerboundDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedLowerboundDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.ContainerPrediction.PredictedLowerboundDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=721,
  serialized_end=825,
)

_CONTAINERPREDICTION = _descriptor.Descriptor(
  name='ContainerPrediction',
  full_name='containersai.datahub.predictions.ContainerPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.predictions.ContainerPrediction.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.datahub.predictions.ContainerPrediction.predicted_raw_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.datahub.predictions.ContainerPrediction.predicted_upperbound_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.datahub.predictions.ContainerPrediction.predicted_lowerbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY, _CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY, _CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=136,
  serialized_end=825,
)


_PODPREDICTION = _descriptor.Descriptor(
  name='PodPrediction',
  full_name='containersai.datahub.predictions.PodPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containersai.datahub.predictions.PodPrediction.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container_predictions', full_name='containersai.datahub.predictions.PodPrediction.container_predictions', index=1,
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
  serialized_start=828,
  serialized_end=1002,
)


_NODEPREDICTION_PREDICTEDRAWDATAENTRY = _descriptor.Descriptor(
  name='PredictedRawDataEntry',
  full_name='containersai.datahub.predictions.NodePrediction.PredictedRawDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.NodePrediction.PredictedRawDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.NodePrediction.PredictedRawDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=516,
  serialized_end=613,
)

_NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY = _descriptor.Descriptor(
  name='PredictedUpperboundDataEntry',
  full_name='containersai.datahub.predictions.NodePrediction.PredictedUpperboundDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.NodePrediction.PredictedUpperboundDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.NodePrediction.PredictedUpperboundDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=615,
  serialized_end=719,
)

_NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY = _descriptor.Descriptor(
  name='PredictedLowerboundDataEntry',
  full_name='containersai.datahub.predictions.NodePrediction.PredictedLowerboundDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.predictions.NodePrediction.PredictedLowerboundDataEntry.key', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.predictions.NodePrediction.PredictedLowerboundDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=721,
  serialized_end=825,
)

_NODEPREDICTION = _descriptor.Descriptor(
  name='NodePrediction',
  full_name='containersai.datahub.predictions.NodePrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.predictions.NodePrediction.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_scheduled', full_name='containersai.datahub.predictions.NodePrediction.is_scheduled', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.datahub.predictions.NodePrediction.predicted_raw_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.datahub.predictions.NodePrediction.predicted_upperbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.datahub.predictions.NodePrediction.predicted_lowerbound_data', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_NODEPREDICTION_PREDICTEDRAWDATAENTRY, _NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY, _NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1005,
  serialized_end=1696,
)

_CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY.containing_type = _CONTAINERPREDICTION
_CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY.containing_type = _CONTAINERPREDICTION
_CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY.containing_type = _CONTAINERPREDICTION
_CONTAINERPREDICTION.fields_by_name['predicted_raw_data'].message_type = _CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY
_CONTAINERPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = _CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY
_CONTAINERPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = _CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY
_PODPREDICTION.fields_by_name['namespaced_name'].message_type = datahub_dot_resources_dot_types__pb2._NAMESPACEDNAME
_PODPREDICTION.fields_by_name['container_predictions'].message_type = _CONTAINERPREDICTION
_NODEPREDICTION_PREDICTEDRAWDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_NODEPREDICTION_PREDICTEDRAWDATAENTRY.containing_type = _NODEPREDICTION
_NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY.containing_type = _NODEPREDICTION
_NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY.fields_by_name['value'].message_type = datahub_dot_metrics_dot_types__pb2._METRICDATA
_NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY.containing_type = _NODEPREDICTION
_NODEPREDICTION.fields_by_name['predicted_raw_data'].message_type = _NODEPREDICTION_PREDICTEDRAWDATAENTRY
_NODEPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = _NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY
_NODEPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = _NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY
DESCRIPTOR.message_types_by_name['ContainerPrediction'] = _CONTAINERPREDICTION
DESCRIPTOR.message_types_by_name['PodPrediction'] = _PODPREDICTION
DESCRIPTOR.message_types_by_name['NodePrediction'] = _NODEPREDICTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerPrediction = _reflection.GeneratedProtocolMessageType('ContainerPrediction', (_message.Message,), dict(

  PredictedRawDataEntry = _reflection.GeneratedProtocolMessageType('PredictedRawDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.ContainerPrediction.PredictedRawDataEntry)
    ))
  ,

  PredictedUpperboundDataEntry = _reflection.GeneratedProtocolMessageType('PredictedUpperboundDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.ContainerPrediction.PredictedUpperboundDataEntry)
    ))
  ,

  PredictedLowerboundDataEntry = _reflection.GeneratedProtocolMessageType('PredictedLowerboundDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.ContainerPrediction.PredictedLowerboundDataEntry)
    ))
  ,
  DESCRIPTOR = _CONTAINERPREDICTION,
  __module__ = 'datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.ContainerPrediction)
  ))
_sym_db.RegisterMessage(ContainerPrediction)
_sym_db.RegisterMessage(ContainerPrediction.PredictedRawDataEntry)
_sym_db.RegisterMessage(ContainerPrediction.PredictedUpperboundDataEntry)
_sym_db.RegisterMessage(ContainerPrediction.PredictedLowerboundDataEntry)

PodPrediction = _reflection.GeneratedProtocolMessageType('PodPrediction', (_message.Message,), dict(
  DESCRIPTOR = _PODPREDICTION,
  __module__ = 'datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.PodPrediction)
  ))
_sym_db.RegisterMessage(PodPrediction)

NodePrediction = _reflection.GeneratedProtocolMessageType('NodePrediction', (_message.Message,), dict(

  PredictedRawDataEntry = _reflection.GeneratedProtocolMessageType('PredictedRawDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _NODEPREDICTION_PREDICTEDRAWDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.NodePrediction.PredictedRawDataEntry)
    ))
  ,

  PredictedUpperboundDataEntry = _reflection.GeneratedProtocolMessageType('PredictedUpperboundDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.NodePrediction.PredictedUpperboundDataEntry)
    ))
  ,

  PredictedLowerboundDataEntry = _reflection.GeneratedProtocolMessageType('PredictedLowerboundDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY,
    __module__ = 'datahub.predictions.predictions_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.NodePrediction.PredictedLowerboundDataEntry)
    ))
  ,
  DESCRIPTOR = _NODEPREDICTION,
  __module__ = 'datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.predictions.NodePrediction)
  ))
_sym_db.RegisterMessage(NodePrediction)
_sym_db.RegisterMessage(NodePrediction.PredictedRawDataEntry)
_sym_db.RegisterMessage(NodePrediction.PredictedUpperboundDataEntry)
_sym_db.RegisterMessage(NodePrediction.PredictedLowerboundDataEntry)


DESCRIPTOR._options = None
_CONTAINERPREDICTION_PREDICTEDRAWDATAENTRY._options = None
_CONTAINERPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY._options = None
_CONTAINERPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY._options = None
_NODEPREDICTION_PREDICTEDRAWDATAENTRY._options = None
_NODEPREDICTION_PREDICTEDUPPERBOUNDDATAENTRY._options = None
_NODEPREDICTION_PREDICTEDLOWERBOUNDDATAENTRY._options = None
# @@protoc_insertion_point(module_scope)
