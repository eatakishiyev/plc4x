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

// CIPEncapsulationConnectionResponse is the corresponding interface of CIPEncapsulationConnectionResponse
type CIPEncapsulationConnectionResponse interface {
	utils.LengthAware
	utils.Serializable
	CIPEncapsulationPacket
}

// CIPEncapsulationConnectionResponseExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationConnectionResponse.
// This is useful for switch cases.
type CIPEncapsulationConnectionResponseExactly interface {
	isCIPEncapsulationConnectionResponse() bool
}

// _CIPEncapsulationConnectionResponse is the data-structure of this message
type _CIPEncapsulationConnectionResponse struct {
	*_CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationConnectionResponse) GetCommandType() uint16 {
	return 0x0201
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationConnectionResponse) InitializeParent(parent CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_CIPEncapsulationConnectionResponse) GetParent() CIPEncapsulationPacket {
	return m._CIPEncapsulationPacket
}

// NewCIPEncapsulationConnectionResponse factory function for _CIPEncapsulationConnectionResponse
func NewCIPEncapsulationConnectionResponse(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationConnectionResponse {
	_result := &_CIPEncapsulationConnectionResponse{
		_CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationConnectionResponse(structType interface{}) CIPEncapsulationConnectionResponse {
	if casted, ok := structType.(CIPEncapsulationConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationConnectionResponse) GetTypeName() string {
	return "CIPEncapsulationConnectionResponse"
}

func (m *_CIPEncapsulationConnectionResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CIPEncapsulationConnectionResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_CIPEncapsulationConnectionResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPEncapsulationConnectionResponseParse(readBuffer utils.ReadBuffer) (CIPEncapsulationConnectionResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CIPEncapsulationConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationConnectionResponse")
	}

	// Create a partially initialized instance
	_child := &_CIPEncapsulationConnectionResponse{
		_CIPEncapsulationPacket: &_CIPEncapsulationPacket{},
	}
	_child._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _child
	return _child, nil
}

func (m *_CIPEncapsulationConnectionResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationConnectionResponse")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationConnectionResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CIPEncapsulationConnectionResponse) isCIPEncapsulationConnectionResponse() bool {
	return true
}

func (m *_CIPEncapsulationConnectionResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
