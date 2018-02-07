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
package org.apache.plc4x.java.ads.api.commands.types;

import org.apache.commons.codec.binary.Hex;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class SampleSizeTest {

    byte NULL_BYTE = 0x0;

    @Test
    void ofBytes() {
        Assertions.assertEquals("0", SampleSize.of(NULL_BYTE, NULL_BYTE, NULL_BYTE, NULL_BYTE).toString());
        Assertions.assertThrows(IllegalArgumentException.class, () -> SampleSize.of(NULL_BYTE, NULL_BYTE, NULL_BYTE, NULL_BYTE, NULL_BYTE));
    }

    @Test
    void ofLong() {
        assertByte(SampleSize.of(1), "0x01000000");
        assertByte(SampleSize.of(65535), "0xffff0000");
        Assertions.assertThrows(IllegalArgumentException.class, () -> SampleSize.of(-1));
        Assertions.assertThrows(IllegalArgumentException.class, () -> SampleSize.of(4294967296L));
    }

    @Test
    void ofString() {
        assertByte(SampleSize.of("1"), "0x01000000");
    }

    @Test
    void testToString() {
        Assertions.assertEquals(SampleSize.of("1").toString(), "1");
    }

    void assertByte(SampleSize actual, String expected) {
        Assertions.assertEquals(expected, "0x" + Hex.encodeHexString(actual.getBytes()));
    }


}