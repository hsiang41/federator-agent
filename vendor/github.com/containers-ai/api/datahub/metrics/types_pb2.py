# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/metrics/types.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/metrics/types.proto',
  package='containersai.datahub.metrics',
  syntax='proto3',
  serialized_options=_b('Z,github.com/containers-ai/api/datahub/metrics'),
  serialized_pb=_b('\n\x1b\x64\x61tahub/metrics/types.proto\x12\x1c\x63ontainersai.datahub.metrics\x1a\x1fgoogle/protobuf/timestamp.proto\"y\n\x06Sample\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tnum_value\x18\x03 \x01(\t\"U\n\nMetricData\x12\x32\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32$.containersai.datahub.metrics.Sample\x12\x13\n\x0bgranularity\x18\x02 \x01(\x03*a\n\nMetricType\x12\x19\n\x15METRIC_TYPE_UNDEFINED\x10\x00\x12 \n\x1c\x43PU_USAGE_SECONDS_PERCENTAGE\x10\x01\x12\x16\n\x12MEMORY_USAGE_BYTES\x10\x02\x42.Z,github.com/containers-ai/api/datahub/metricsb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_METRICTYPE = _descriptor.EnumDescriptor(
  name='MetricType',
  full_name='containersai.datahub.metrics.MetricType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='METRIC_TYPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_USAGE_SECONDS_PERCENTAGE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_USAGE_BYTES', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=304,
  serialized_end=401,
)
_sym_db.RegisterEnumDescriptor(_METRICTYPE)

MetricType = enum_type_wrapper.EnumTypeWrapper(_METRICTYPE)
METRIC_TYPE_UNDEFINED = 0
CPU_USAGE_SECONDS_PERCENTAGE = 1
MEMORY_USAGE_BYTES = 2



_SAMPLE = _descriptor.Descriptor(
  name='Sample',
  full_name='containersai.datahub.metrics.Sample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.datahub.metrics.Sample.start_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.datahub.metrics.Sample.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_value', full_name='containersai.datahub.metrics.Sample.num_value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=94,
  serialized_end=215,
)


_METRICDATA = _descriptor.Descriptor(
  name='MetricData',
  full_name='containersai.datahub.metrics.MetricData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='containersai.datahub.metrics.MetricData.data', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='granularity', full_name='containersai.datahub.metrics.MetricData.granularity', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=217,
  serialized_end=302,
)

_SAMPLE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SAMPLE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_METRICDATA.fields_by_name['data'].message_type = _SAMPLE
DESCRIPTOR.message_types_by_name['Sample'] = _SAMPLE
DESCRIPTOR.message_types_by_name['MetricData'] = _METRICDATA
DESCRIPTOR.enum_types_by_name['MetricType'] = _METRICTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Sample = _reflection.GeneratedProtocolMessageType('Sample', (_message.Message,), {
  'DESCRIPTOR' : _SAMPLE,
  '__module__' : 'datahub.metrics.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metrics.Sample)
  })
_sym_db.RegisterMessage(Sample)

MetricData = _reflection.GeneratedProtocolMessageType('MetricData', (_message.Message,), {
  'DESCRIPTOR' : _METRICDATA,
  '__module__' : 'datahub.metrics.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metrics.MetricData)
  })
_sym_db.RegisterMessage(MetricData)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
