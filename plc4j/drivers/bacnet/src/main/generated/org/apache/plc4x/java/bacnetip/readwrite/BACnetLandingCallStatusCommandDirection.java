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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetLandingCallStatusCommandDirection extends BACnetLandingCallStatusCommand
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetLiftCarDirectionTagged direction;

  public BACnetLandingCallStatusCommandDirection(
      BACnetTagHeader peekedTagHeader, BACnetLiftCarDirectionTagged direction) {
    super(peekedTagHeader);
    this.direction = direction;
  }

  public BACnetLiftCarDirectionTagged getDirection() {
    return direction;
  }

  @Override
  protected void serializeBACnetLandingCallStatusCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLandingCallStatusCommandDirection");

    // Simple Field (direction)
    writeSimpleField("direction", direction, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetLandingCallStatusCommandDirection");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLandingCallStatusCommandDirection _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (direction)
    lengthInBits += direction.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLandingCallStatusCommandBuilder
      staticParseBACnetLandingCallStatusCommandBuilder(ReadBuffer readBuffer)
          throws ParseException {
    readBuffer.pullContext("BACnetLandingCallStatusCommandDirection");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetLiftCarDirectionTagged direction =
        readSimpleField(
            "direction",
            readComplex(
                () ->
                    BACnetLiftCarDirectionTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetLandingCallStatusCommandDirection");
    // Create the instance
    return new BACnetLandingCallStatusCommandDirectionBuilderImpl(direction);
  }

  public static class BACnetLandingCallStatusCommandDirectionBuilderImpl
      implements BACnetLandingCallStatusCommand.BACnetLandingCallStatusCommandBuilder {
    private final BACnetLiftCarDirectionTagged direction;

    public BACnetLandingCallStatusCommandDirectionBuilderImpl(
        BACnetLiftCarDirectionTagged direction) {
      this.direction = direction;
    }

    public BACnetLandingCallStatusCommandDirection build(BACnetTagHeader peekedTagHeader) {
      BACnetLandingCallStatusCommandDirection bACnetLandingCallStatusCommandDirection =
          new BACnetLandingCallStatusCommandDirection(peekedTagHeader, direction);
      return bACnetLandingCallStatusCommandDirection;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLandingCallStatusCommandDirection)) {
      return false;
    }
    BACnetLandingCallStatusCommandDirection that = (BACnetLandingCallStatusCommandDirection) o;
    return (getDirection() == that.getDirection()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDirection());
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
