/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.ads.api.generic.types;

import org.apache.plc4x.java.ads.api.util.ByteValue;

import java.nio.ByteBuffer;

public class DataLength extends ByteValue {

    public static final int NUM_BYTES = 4;

    DataLength(byte... value) {
        super(value);
        assertLength(NUM_BYTES);
    }

    public static DataLength of(long length) {
        checkUnsignedBounds(length, NUM_BYTES);
        return new DataLength(ByteBuffer.allocate(NUM_BYTES)
            // LE
            .put((byte) (length & 0xff))
            .put((byte) (length >> 8 & 0xff))
            .put((byte) (length >> 16 & 0xff))
            .put((byte) (length >> 24 & 0xff))
            .array());
    }

    public static DataLength of(String length) {
        return of(Long.parseLong(length));
    }

    public static DataLength of(byte... values) {
        return new DataLength(values);
    }

    @Override
    public String toString() {
        return "" + (getBytes()[3] << 24 | getBytes()[2] << 16 | getBytes()[1] << 8 | getBytes()[0]);
    }
}
