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

import org.apache.commons.codec.binary.Hex;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class AMSPortTest {

    byte NULL_BYTE = 0x0;

    @Test
    void ofBytes() {
        Assertions.assertEquals("0", AMSPort.of(NULL_BYTE, NULL_BYTE).toString());
        Assertions.assertThrows(IllegalArgumentException.class, () -> AMSPort.of(NULL_BYTE, NULL_BYTE, NULL_BYTE));
    }

    @Test
    void ofInt() {
        assertByte(AMSPort.of(1), "0x0100");
        assertByte(AMSPort.of(65535), "0xffff");
        Assertions.assertThrows(IllegalArgumentException.class, () -> AMSPort.of(-1));
        Assertions.assertThrows(IllegalArgumentException.class, () -> AMSPort.of(65536));
    }

    @Test
    void ofString() {
        assertByte(AMSPort.of("1"), "0x0100");
    }

    @Test
    void testToString() {
        Assertions.assertEquals(AMSPort.of("1").toString(), "1");
    }

    void assertByte(AMSPort actual, String expected) {
        Assertions.assertEquals(expected, "0x" + Hex.encodeHexString(actual.getBytes()));
    }
}