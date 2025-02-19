#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

# Code generated by code-generation. DO NOT EDIT.
from aenum import AutoNumberEnum


class UmasDataType(AutoNumberEnum):
    _init_ = "value, request_size, data_type_size"
    BOOL = (1, int(1), int(1))
    UNKNOWN2 = (2, int(1), int(1))
    UNKNOWN3 = (3, int(1), int(1))
    INT = (4, int(2), int(2))
    UINT = (5, int(2), int(2))
    DINT = (6, int(3), int(4))
    UDINT = (7, int(3), int(4))
    REAL = (8, int(3), int(4))
    STRING = (9, int(17), int(1))
    TIME = (10, int(3), int(4))
    UNKNOWN11 = (11, int(1), int(1))
    UNKNOWN12 = (12, int(1), int(1))
    UNKNOWN13 = (13, int(1), int(1))
    DATE = (14, int(3), int(4))
    TOD = (15, int(3), int(4))
    DT = (16, int(3), int(4))
    UNKNOWN17 = (17, int(1), int(1))
    UNKNOWN18 = (18, int(1), int(1))
    UNKNOWN19 = (19, int(1), int(1))
    UNKNOWN20 = (20, int(1), int(1))
    BYTE = (21, int(1), int(1))
    WORD = (22, int(2), int(2))
    DWORD = (23, int(3), int(4))
    UNKNOWN24 = (24, int(1), int(1))
    EBOOL = (25, int(1), int(1))
