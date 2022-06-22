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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// HPAIControlEndpoint is the corresponding interface of HPAIControlEndpoint
type HPAIControlEndpoint interface {
	utils.LengthAware
	utils.Serializable
	// GetHostProtocolCode returns HostProtocolCode (property field)
	GetHostProtocolCode() HostProtocolCode
	// GetIpAddress returns IpAddress (property field)
	GetIpAddress() IPAddress
	// GetIpPort returns IpPort (property field)
	GetIpPort() uint16
}

// HPAIControlEndpointExactly can be used when we want exactly this type and not a type which fulfills HPAIControlEndpoint.
// This is useful for switch cases.
type HPAIControlEndpointExactly interface {
	isHPAIControlEndpoint() bool
}

// _HPAIControlEndpoint is the data-structure of this message
type _HPAIControlEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        IPAddress
	IpPort           uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HPAIControlEndpoint) GetHostProtocolCode() HostProtocolCode {
	return m.HostProtocolCode
}

func (m *_HPAIControlEndpoint) GetIpAddress() IPAddress {
	return m.IpAddress
}

func (m *_HPAIControlEndpoint) GetIpPort() uint16 {
	return m.IpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewHPAIControlEndpoint factory function for _HPAIControlEndpoint
func NewHPAIControlEndpoint(hostProtocolCode HostProtocolCode, ipAddress IPAddress, ipPort uint16) *_HPAIControlEndpoint {
	return &_HPAIControlEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

// Deprecated: use the interface for direct cast
func CastHPAIControlEndpoint(structType interface{}) HPAIControlEndpoint {
	if casted, ok := structType.(HPAIControlEndpoint); ok {
		return casted
	}
	if casted, ok := structType.(*HPAIControlEndpoint); ok {
		return *casted
	}
	return nil
}

func (m *_HPAIControlEndpoint) GetTypeName() string {
	return "HPAIControlEndpoint"
}

func (m *_HPAIControlEndpoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_HPAIControlEndpoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits()

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_HPAIControlEndpoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HPAIControlEndpointParse(readBuffer utils.ReadBuffer) (HPAIControlEndpoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HPAIControlEndpoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HPAIControlEndpoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pullErr := readBuffer.PullContext("hostProtocolCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hostProtocolCode")
	}
	_hostProtocolCode, _hostProtocolCodeErr := HostProtocolCodeParse(readBuffer)
	if _hostProtocolCodeErr != nil {
		return nil, errors.Wrap(_hostProtocolCodeErr, "Error parsing 'hostProtocolCode' field")
	}
	hostProtocolCode := _hostProtocolCode
	if closeErr := readBuffer.CloseContext("hostProtocolCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hostProtocolCode")
	}

	// Simple Field (ipAddress)
	if pullErr := readBuffer.PullContext("ipAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipAddress")
	}
	_ipAddress, _ipAddressErr := IPAddressParse(readBuffer)
	if _ipAddressErr != nil {
		return nil, errors.Wrap(_ipAddressErr, "Error parsing 'ipAddress' field")
	}
	ipAddress := _ipAddress.(IPAddress)
	if closeErr := readBuffer.CloseContext("ipAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipAddress")
	}

	// Simple Field (ipPort)
	_ipPort, _ipPortErr := readBuffer.ReadUint16("ipPort", 16)
	if _ipPortErr != nil {
		return nil, errors.Wrap(_ipPortErr, "Error parsing 'ipPort' field")
	}
	ipPort := _ipPort

	if closeErr := readBuffer.CloseContext("HPAIControlEndpoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HPAIControlEndpoint")
	}

	// Create the instance
	return NewHPAIControlEndpoint(hostProtocolCode, ipAddress, ipPort), nil
}

func (m *_HPAIControlEndpoint) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("HPAIControlEndpoint"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for HPAIControlEndpoint")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pushErr := writeBuffer.PushContext("hostProtocolCode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hostProtocolCode")
	}
	_hostProtocolCodeErr := writeBuffer.WriteSerializable(m.GetHostProtocolCode())
	if popErr := writeBuffer.PopContext("hostProtocolCode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hostProtocolCode")
	}
	if _hostProtocolCodeErr != nil {
		return errors.Wrap(_hostProtocolCodeErr, "Error serializing 'hostProtocolCode' field")
	}

	// Simple Field (ipAddress)
	if pushErr := writeBuffer.PushContext("ipAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ipAddress")
	}
	_ipAddressErr := writeBuffer.WriteSerializable(m.GetIpAddress())
	if popErr := writeBuffer.PopContext("ipAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ipAddress")
	}
	if _ipAddressErr != nil {
		return errors.Wrap(_ipAddressErr, "Error serializing 'ipAddress' field")
	}

	// Simple Field (ipPort)
	ipPort := uint16(m.GetIpPort())
	_ipPortErr := writeBuffer.WriteUint16("ipPort", 16, (ipPort))
	if _ipPortErr != nil {
		return errors.Wrap(_ipPortErr, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIControlEndpoint"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for HPAIControlEndpoint")
	}
	return nil
}

func (m *_HPAIControlEndpoint) isHPAIControlEndpoint() bool {
	return true
}

func (m *_HPAIControlEndpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
