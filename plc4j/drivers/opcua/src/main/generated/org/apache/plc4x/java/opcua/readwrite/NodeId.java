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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class NodeId implements Message {

  // Properties.
  protected final NodeIdTypeDefinition nodeId;

  public NodeId(NodeIdTypeDefinition nodeId) {
    super();
    this.nodeId = nodeId;
  }

  public NodeIdTypeDefinition getNodeId() {
    return nodeId;
  }

  public String getId() {
    return String.valueOf(getNodeId().getIdentifier());
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NodeId");

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeSignedByte(writeBuffer, 2));

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, writeComplex(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    String id = getId();
    writeBuffer.writeVirtual("id", id);

    writeBuffer.popContext("NodeId");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    NodeId _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 2;

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static NodeId staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("NodeId");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readSignedByte(readBuffer, 2), (byte) 0x00);

    NodeIdTypeDefinition nodeId =
        readSimpleField(
            "nodeId", readComplex(() -> NodeIdTypeDefinition.staticParse(readBuffer), readBuffer));
    String id = readVirtualField("id", String.class, nodeId.getIdentifier());

    readBuffer.closeContext("NodeId");
    // Create the instance
    NodeId _nodeId;
    _nodeId = new NodeId(nodeId);
    return _nodeId;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NodeId)) {
      return false;
    }
    NodeId that = (NodeId) o;
    return (getNodeId() == that.getNodeId()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNodeId());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
