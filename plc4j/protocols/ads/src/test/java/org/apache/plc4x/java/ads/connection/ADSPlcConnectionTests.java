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

package org.apache.plc4x.java.ads.connection;

import org.apache.plc4x.java.ads.api.generic.types.AMSNetId;
import org.apache.plc4x.java.ads.api.generic.types.AMSPort;
import org.apache.plc4x.java.ads.model.ADSAddress;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ADSPlcConnectionTests {

    private ADSPlcConnection adsPlcConnection;

    @BeforeEach
    void setUp() {
        adsPlcConnection = new ADSPlcConnection("localhost", AMSNetId.of("0.0.0.0.0.0"), AMSPort.of(13));
    }

    @AfterEach
    void tearDown() {
        adsPlcConnection = null;
    }

    @Test
    void initialState() {
        assertTrue(adsPlcConnection.getHostName().equalsIgnoreCase("localhost"), "Hostname is incorrect");
        assertEquals(adsPlcConnection.getTargetAmsNetId().toString(), "0.0.0.0.0.0");
        assertEquals(adsPlcConnection.getTargetAmsPort().toString(), "13");
    }

    @Test
    void emptyParseAddress() throws Exception {
        try {
            adsPlcConnection.parseAddress("");
        } catch (IllegalArgumentException exception) {
            assertTrue(exception.getMessage().startsWith("address  doesn't match "), "Unexpected exception");
        }
    }

    @Test
    void parseAddress() throws Exception {
        try {
            ADSAddress address = (ADSAddress) adsPlcConnection.parseAddress("1/1");
            assertEquals(address.getIndexGroup(), 1);
            assertEquals(address.getIndexOffset(), 1);
        } catch (IllegalArgumentException exception) {
            fail("valid data block address");
        }
    }
}