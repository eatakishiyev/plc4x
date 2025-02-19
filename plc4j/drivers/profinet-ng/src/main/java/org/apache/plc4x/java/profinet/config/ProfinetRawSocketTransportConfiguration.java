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

package org.apache.plc4x.java.profinet.config;

import org.apache.plc4x.java.transport.rawsocket.DefaultRawSocketTransportConfiguration;
import org.apache.plc4x.java.utils.pcap.netty.handlers.PacketHandler;
import org.pcap4j.packet.*;

public class ProfinetRawSocketTransportConfiguration extends DefaultRawSocketTransportConfiguration {

    public ProfinetRawSocketTransportConfiguration() {
        setResolveMacAddress(true);
    }

    @Override
    public int getDefaultPort() {
        return 34964;
    }

    @Override
    public PacketHandler getPcapPacketHandler() {
        return new PacketHandler() {
            @Override
            public byte[] getData(Packet packet) {
                if(packet instanceof EthernetPacket) {
                    EthernetPacket ethernetPacket = (EthernetPacket) packet;
                    if(ethernetPacket.getPayload() instanceof IpV4Packet) {
                        IpV4Packet ipV4Packet = (IpV4Packet) ethernetPacket.getPayload();
                        if(ipV4Packet.getPayload() instanceof UdpPacket) {
                            UdpPacket udpPacket = (UdpPacket) ipV4Packet.getPayload();
                            if(udpPacket.getHeader().getSrcPort().value() == (short) 49156) {
                                System.out.println(packet);
                            }
                        }
                    }
                }
                return packet.getRawData();
            }
        };
    }

}
