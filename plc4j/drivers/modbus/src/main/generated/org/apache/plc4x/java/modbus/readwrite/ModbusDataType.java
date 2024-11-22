/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.modbus.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum ModbusDataType {
  BOOL((short) 1, (short) 1),
  BYTE((short) 2, (short) 2),
  WORD((short) 3, (short) 2),
  DWORD((short) 4, (short) 4),
  LWORD((short) 5, (short) 8),
  SINT((short) 6, (short) 2),
  INT((short) 7, (short) 2),
  DINT((short) 8, (short) 4),
  LINT((short) 9, (short) 8),
  USINT((short) 10, (short) 2),
  UINT((short) 11, (short) 2),
  UDINT((short) 12, (short) 4),
  ULINT((short) 13, (short) 8),
  REAL((short) 14, (short) 4),
  LREAL((short) 15, (short) 8),
  TIME((short) 16, (short) 8),
  LTIME((short) 17, (short) 8),
  DATE((short) 18, (short) 8),
  LDATE((short) 19, (short) 8),
  TIME_OF_DAY((short) 20, (short) 8),
  LTIME_OF_DAY((short) 21, (short) 8),
  DATE_AND_TIME((short) 22, (short) 8),
  LDATE_AND_TIME((short) 23, (short) 8),
  CHAR((short) 24, (short) 1),
  WCHAR((short) 25, (short) 2),
  STRING((short) 26, (short) 1),
  WSTRING((short) 27, (short) 2);
  private static final Map<Short, ModbusDataType> map;

  static {
    map = new HashMap<>();
    for (ModbusDataType value : ModbusDataType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final short dataTypeSize;

  ModbusDataType(short value, short dataTypeSize) {
    this.value = value;
    this.dataTypeSize = dataTypeSize;
  }

  public short getValue() {
    return value;
  }

  public short getDataTypeSize() {
    return dataTypeSize;
  }

  public static ModbusDataType firstEnumForFieldDataTypeSize(short fieldValue) {
    for (ModbusDataType _val : ModbusDataType.values()) {
      if (_val.getDataTypeSize() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<ModbusDataType> enumsForFieldDataTypeSize(short fieldValue) {
    List<ModbusDataType> _values = new ArrayList<>();
    for (ModbusDataType _val : ModbusDataType.values()) {
      if (_val.getDataTypeSize() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static ModbusDataType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
