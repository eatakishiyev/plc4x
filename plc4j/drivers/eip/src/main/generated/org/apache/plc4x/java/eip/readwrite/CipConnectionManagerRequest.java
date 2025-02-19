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
package org.apache.plc4x.java.eip.readwrite;

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

public class CipConnectionManagerRequest extends CipService implements Message {

  // Accessors for discriminator values.
  public Byte getService() {
    return (byte) 0x5B;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final PathSegment classSegment;
  protected final PathSegment instanceSegment;
  protected final byte priority;
  protected final byte tickTime;
  protected final short timeoutTicks;
  protected final long otConnectionId;
  protected final long toConnectionId;
  protected final int connectionSerialNumber;
  protected final int originatorVendorId;
  protected final long originatorSerialNumber;
  protected final short timeoutMultiplier;
  protected final long otRpi;
  protected final NetworkConnectionParameters otConnectionParameters;
  protected final long toRpi;
  protected final NetworkConnectionParameters toConnectionParameters;
  protected final TransportType transportType;
  protected final short connectionPathSize;
  protected final List<PathSegment> connectionPaths;

  public CipConnectionManagerRequest(
      PathSegment classSegment,
      PathSegment instanceSegment,
      byte priority,
      byte tickTime,
      short timeoutTicks,
      long otConnectionId,
      long toConnectionId,
      int connectionSerialNumber,
      int originatorVendorId,
      long originatorSerialNumber,
      short timeoutMultiplier,
      long otRpi,
      NetworkConnectionParameters otConnectionParameters,
      long toRpi,
      NetworkConnectionParameters toConnectionParameters,
      TransportType transportType,
      short connectionPathSize,
      List<PathSegment> connectionPaths) {
    super();
    this.classSegment = classSegment;
    this.instanceSegment = instanceSegment;
    this.priority = priority;
    this.tickTime = tickTime;
    this.timeoutTicks = timeoutTicks;
    this.otConnectionId = otConnectionId;
    this.toConnectionId = toConnectionId;
    this.connectionSerialNumber = connectionSerialNumber;
    this.originatorVendorId = originatorVendorId;
    this.originatorSerialNumber = originatorSerialNumber;
    this.timeoutMultiplier = timeoutMultiplier;
    this.otRpi = otRpi;
    this.otConnectionParameters = otConnectionParameters;
    this.toRpi = toRpi;
    this.toConnectionParameters = toConnectionParameters;
    this.transportType = transportType;
    this.connectionPathSize = connectionPathSize;
    this.connectionPaths = connectionPaths;
  }

  public PathSegment getClassSegment() {
    return classSegment;
  }

  public PathSegment getInstanceSegment() {
    return instanceSegment;
  }

  public byte getPriority() {
    return priority;
  }

  public byte getTickTime() {
    return tickTime;
  }

  public short getTimeoutTicks() {
    return timeoutTicks;
  }

  public long getOtConnectionId() {
    return otConnectionId;
  }

  public long getToConnectionId() {
    return toConnectionId;
  }

  public int getConnectionSerialNumber() {
    return connectionSerialNumber;
  }

  public int getOriginatorVendorId() {
    return originatorVendorId;
  }

  public long getOriginatorSerialNumber() {
    return originatorSerialNumber;
  }

  public short getTimeoutMultiplier() {
    return timeoutMultiplier;
  }

  public long getOtRpi() {
    return otRpi;
  }

  public NetworkConnectionParameters getOtConnectionParameters() {
    return otConnectionParameters;
  }

  public long getToRpi() {
    return toRpi;
  }

  public NetworkConnectionParameters getToConnectionParameters() {
    return toConnectionParameters;
  }

  public TransportType getTransportType() {
    return transportType;
  }

  public short getConnectionPathSize() {
    return connectionPathSize;
  }

  public List<PathSegment> getConnectionPaths() {
    return connectionPaths;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipConnectionManagerRequest");

    // Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    short requestPathSize =
        (short)
            ((((getClassSegment().getLengthInBytes()) + (getInstanceSegment().getLengthInBytes())))
                / (2));
    writeImplicitField("requestPathSize", requestPathSize, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (classSegment)
    writeSimpleField("classSegment", classSegment, writeComplex(writeBuffer));

    // Simple Field (instanceSegment)
    writeSimpleField("instanceSegment", instanceSegment, writeComplex(writeBuffer));

    // Simple Field (priority)
    writeSimpleField("priority", priority, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (tickTime)
    writeSimpleField("tickTime", tickTime, writeUnsignedByte(writeBuffer, 4));

    // Simple Field (timeoutTicks)
    writeSimpleField("timeoutTicks", timeoutTicks, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (otConnectionId)
    writeSimpleField("otConnectionId", otConnectionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (toConnectionId)
    writeSimpleField("toConnectionId", toConnectionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (connectionSerialNumber)
    writeSimpleField(
        "connectionSerialNumber", connectionSerialNumber, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (originatorVendorId)
    writeSimpleField("originatorVendorId", originatorVendorId, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (originatorSerialNumber)
    writeSimpleField(
        "originatorSerialNumber", originatorSerialNumber, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (timeoutMultiplier)
    writeSimpleField("timeoutMultiplier", timeoutMultiplier, writeUnsignedShort(writeBuffer, 8));

    // Reserved Field (reserved)
    writeReservedField("reserved", (int) 0x000000, writeUnsignedInt(writeBuffer, 24));

    // Simple Field (otRpi)
    writeSimpleField("otRpi", otRpi, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (otConnectionParameters)
    writeSimpleField("otConnectionParameters", otConnectionParameters, writeComplex(writeBuffer));

    // Simple Field (toRpi)
    writeSimpleField("toRpi", toRpi, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (toConnectionParameters)
    writeSimpleField("toConnectionParameters", toConnectionParameters, writeComplex(writeBuffer));

    // Simple Field (transportType)
    writeSimpleField("transportType", transportType, writeComplex(writeBuffer));

    // Simple Field (connectionPathSize)
    writeSimpleField("connectionPathSize", connectionPathSize, writeUnsignedShort(writeBuffer, 8));

    // Array Field (connectionPaths)
    writeComplexTypeArrayField("connectionPaths", connectionPaths, writeBuffer);

    writeBuffer.popContext("CipConnectionManagerRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipConnectionManagerRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (requestPathSize)
    lengthInBits += 8;

    // Simple field (classSegment)
    lengthInBits += classSegment.getLengthInBits();

    // Simple field (instanceSegment)
    lengthInBits += instanceSegment.getLengthInBits();

    // Simple field (priority)
    lengthInBits += 4;

    // Simple field (tickTime)
    lengthInBits += 4;

    // Simple field (timeoutTicks)
    lengthInBits += 8;

    // Simple field (otConnectionId)
    lengthInBits += 32;

    // Simple field (toConnectionId)
    lengthInBits += 32;

    // Simple field (connectionSerialNumber)
    lengthInBits += 16;

    // Simple field (originatorVendorId)
    lengthInBits += 16;

    // Simple field (originatorSerialNumber)
    lengthInBits += 32;

    // Simple field (timeoutMultiplier)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 24;

    // Simple field (otRpi)
    lengthInBits += 32;

    // Simple field (otConnectionParameters)
    lengthInBits += otConnectionParameters.getLengthInBits();

    // Simple field (toRpi)
    lengthInBits += 32;

    // Simple field (toConnectionParameters)
    lengthInBits += toConnectionParameters.getLengthInBits();

    // Simple field (transportType)
    lengthInBits += transportType.getLengthInBits();

    // Simple field (connectionPathSize)
    lengthInBits += 8;

    // Array field
    if (connectionPaths != null) {
      for (Message element : connectionPaths) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipConnectionManagerRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short requestPathSize = readImplicitField("requestPathSize", readUnsignedShort(readBuffer, 8));

    PathSegment classSegment =
        readSimpleField(
            "classSegment", readComplex(() -> PathSegment.staticParse(readBuffer), readBuffer));

    PathSegment instanceSegment =
        readSimpleField(
            "instanceSegment", readComplex(() -> PathSegment.staticParse(readBuffer), readBuffer));

    byte priority = readSimpleField("priority", readUnsignedByte(readBuffer, 4));

    byte tickTime = readSimpleField("tickTime", readUnsignedByte(readBuffer, 4));

    short timeoutTicks = readSimpleField("timeoutTicks", readUnsignedShort(readBuffer, 8));

    long otConnectionId = readSimpleField("otConnectionId", readUnsignedLong(readBuffer, 32));

    long toConnectionId = readSimpleField("toConnectionId", readUnsignedLong(readBuffer, 32));

    int connectionSerialNumber =
        readSimpleField("connectionSerialNumber", readUnsignedInt(readBuffer, 16));

    int originatorVendorId = readSimpleField("originatorVendorId", readUnsignedInt(readBuffer, 16));

    long originatorSerialNumber =
        readSimpleField("originatorSerialNumber", readUnsignedLong(readBuffer, 32));

    short timeoutMultiplier =
        readSimpleField("timeoutMultiplier", readUnsignedShort(readBuffer, 8));

    Integer reservedField0 =
        readReservedField("reserved", readUnsignedInt(readBuffer, 24), (int) 0x000000);

    long otRpi = readSimpleField("otRpi", readUnsignedLong(readBuffer, 32));

    NetworkConnectionParameters otConnectionParameters =
        readSimpleField(
            "otConnectionParameters",
            readComplex(() -> NetworkConnectionParameters.staticParse(readBuffer), readBuffer));

    long toRpi = readSimpleField("toRpi", readUnsignedLong(readBuffer, 32));

    NetworkConnectionParameters toConnectionParameters =
        readSimpleField(
            "toConnectionParameters",
            readComplex(() -> NetworkConnectionParameters.staticParse(readBuffer), readBuffer));

    TransportType transportType =
        readSimpleField(
            "transportType", readComplex(() -> TransportType.staticParse(readBuffer), readBuffer));

    short connectionPathSize =
        readSimpleField("connectionPathSize", readUnsignedShort(readBuffer, 8));

    List<PathSegment> connectionPaths =
        readTerminatedArrayField(
            "connectionPaths",
            readComplex(() -> PathSegment.staticParse(readBuffer), readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.eip.readwrite.utils.StaticHelper.noMorePathSegments(
                        readBuffer))));

    readBuffer.closeContext("CipConnectionManagerRequest");
    // Create the instance
    return new CipConnectionManagerRequestBuilderImpl(
        classSegment,
        instanceSegment,
        priority,
        tickTime,
        timeoutTicks,
        otConnectionId,
        toConnectionId,
        connectionSerialNumber,
        originatorVendorId,
        originatorSerialNumber,
        timeoutMultiplier,
        otRpi,
        otConnectionParameters,
        toRpi,
        toConnectionParameters,
        transportType,
        connectionPathSize,
        connectionPaths);
  }

  public static class CipConnectionManagerRequestBuilderImpl
      implements CipService.CipServiceBuilder {
    private final PathSegment classSegment;
    private final PathSegment instanceSegment;
    private final byte priority;
    private final byte tickTime;
    private final short timeoutTicks;
    private final long otConnectionId;
    private final long toConnectionId;
    private final int connectionSerialNumber;
    private final int originatorVendorId;
    private final long originatorSerialNumber;
    private final short timeoutMultiplier;
    private final long otRpi;
    private final NetworkConnectionParameters otConnectionParameters;
    private final long toRpi;
    private final NetworkConnectionParameters toConnectionParameters;
    private final TransportType transportType;
    private final short connectionPathSize;
    private final List<PathSegment> connectionPaths;

    public CipConnectionManagerRequestBuilderImpl(
        PathSegment classSegment,
        PathSegment instanceSegment,
        byte priority,
        byte tickTime,
        short timeoutTicks,
        long otConnectionId,
        long toConnectionId,
        int connectionSerialNumber,
        int originatorVendorId,
        long originatorSerialNumber,
        short timeoutMultiplier,
        long otRpi,
        NetworkConnectionParameters otConnectionParameters,
        long toRpi,
        NetworkConnectionParameters toConnectionParameters,
        TransportType transportType,
        short connectionPathSize,
        List<PathSegment> connectionPaths) {
      this.classSegment = classSegment;
      this.instanceSegment = instanceSegment;
      this.priority = priority;
      this.tickTime = tickTime;
      this.timeoutTicks = timeoutTicks;
      this.otConnectionId = otConnectionId;
      this.toConnectionId = toConnectionId;
      this.connectionSerialNumber = connectionSerialNumber;
      this.originatorVendorId = originatorVendorId;
      this.originatorSerialNumber = originatorSerialNumber;
      this.timeoutMultiplier = timeoutMultiplier;
      this.otRpi = otRpi;
      this.otConnectionParameters = otConnectionParameters;
      this.toRpi = toRpi;
      this.toConnectionParameters = toConnectionParameters;
      this.transportType = transportType;
      this.connectionPathSize = connectionPathSize;
      this.connectionPaths = connectionPaths;
    }

    public CipConnectionManagerRequest build() {
      CipConnectionManagerRequest cipConnectionManagerRequest =
          new CipConnectionManagerRequest(
              classSegment,
              instanceSegment,
              priority,
              tickTime,
              timeoutTicks,
              otConnectionId,
              toConnectionId,
              connectionSerialNumber,
              originatorVendorId,
              originatorSerialNumber,
              timeoutMultiplier,
              otRpi,
              otConnectionParameters,
              toRpi,
              toConnectionParameters,
              transportType,
              connectionPathSize,
              connectionPaths);
      return cipConnectionManagerRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipConnectionManagerRequest)) {
      return false;
    }
    CipConnectionManagerRequest that = (CipConnectionManagerRequest) o;
    return (getClassSegment() == that.getClassSegment())
        && (getInstanceSegment() == that.getInstanceSegment())
        && (getPriority() == that.getPriority())
        && (getTickTime() == that.getTickTime())
        && (getTimeoutTicks() == that.getTimeoutTicks())
        && (getOtConnectionId() == that.getOtConnectionId())
        && (getToConnectionId() == that.getToConnectionId())
        && (getConnectionSerialNumber() == that.getConnectionSerialNumber())
        && (getOriginatorVendorId() == that.getOriginatorVendorId())
        && (getOriginatorSerialNumber() == that.getOriginatorSerialNumber())
        && (getTimeoutMultiplier() == that.getTimeoutMultiplier())
        && (getOtRpi() == that.getOtRpi())
        && (getOtConnectionParameters() == that.getOtConnectionParameters())
        && (getToRpi() == that.getToRpi())
        && (getToConnectionParameters() == that.getToConnectionParameters())
        && (getTransportType() == that.getTransportType())
        && (getConnectionPathSize() == that.getConnectionPathSize())
        && (getConnectionPaths() == that.getConnectionPaths())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getClassSegment(),
        getInstanceSegment(),
        getPriority(),
        getTickTime(),
        getTimeoutTicks(),
        getOtConnectionId(),
        getToConnectionId(),
        getConnectionSerialNumber(),
        getOriginatorVendorId(),
        getOriginatorSerialNumber(),
        getTimeoutMultiplier(),
        getOtRpi(),
        getOtConnectionParameters(),
        getToRpi(),
        getToConnectionParameters(),
        getTransportType(),
        getConnectionPathSize(),
        getConnectionPaths());
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
