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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableSampling {
  SamplingIntervalDiagnosticsType_SamplingInterval((int) 2166L),
  SamplingIntervalDiagnosticsType_SampledMonitoredItemsCount((int) 11697L),
  SamplingIntervalDiagnosticsType_MaxSampledMonitoredItemsCount((int) 11698L),
  SamplingIntervalDiagnosticsType_DisabledMonitoredItemsSamplingCount((int) 11699L),
  SamplingIntervalDiagnosticsArrayType_SamplingIntervalDiagnostics((int) 12779L),
  SamplingIntervalDiagnosticsArrayType_SamplingIntervalDiagnostics_SamplingInterval((int) 12780L),
  SamplingIntervalDiagnosticsArrayType_SamplingIntervalDiagnostics_SampledMonitoredItemsCount(
      (int) 12781L),
  SamplingIntervalDiagnosticsArrayType_SamplingIntervalDiagnostics_MaxSampledMonitoredItemsCount(
      (int) 12782L),
  SamplingIntervalDiagnosticsArrayType_SamplingIntervalDiagnostics_DisabledMonitoredItemsSamplingCount(
      (int) 12783L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableSampling> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableSampling value : OpcuaNodeIdServicesVariableSampling.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableSampling(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableSampling enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
