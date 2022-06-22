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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const TPKTPacket_PROTOCOLID uint8 = 0x03

// TPKTPacket is the corresponding interface of TPKTPacket
type TPKTPacket interface {
	utils.LengthAware
	utils.Serializable
	// GetPayload returns Payload (property field)
	GetPayload() COTPPacket
}

// TPKTPacketExactly can be used when we want exactly this type and not a type which fulfills TPKTPacket.
// This is useful for switch cases.
type TPKTPacketExactly interface {
	isTPKTPacket() bool
}

// _TPKTPacket is the data-structure of this message
type _TPKTPacket struct {
	Payload COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TPKTPacket) GetPayload() COTPPacket {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_TPKTPacket) GetProtocolId() uint8 {
	return TPKTPacket_PROTOCOLID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTPKTPacket factory function for _TPKTPacket
func NewTPKTPacket(payload COTPPacket) *_TPKTPacket {
	return &_TPKTPacket{Payload: payload}
}

// Deprecated: use the interface for direct cast
func CastTPKTPacket(structType interface{}) TPKTPacket {
	if casted, ok := structType.(TPKTPacket); ok {
		return casted
	}
	if casted, ok := structType.(*TPKTPacket); ok {
		return *casted
	}
	return nil
}

func (m *_TPKTPacket) GetTypeName() string {
	return "TPKTPacket"
}

func (m *_TPKTPacket) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TPKTPacket) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (protocolId)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Implicit Field (len)
	lengthInBits += 16

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}

func (m *_TPKTPacket) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TPKTPacketParse(readBuffer utils.ReadBuffer) (TPKTPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TPKTPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TPKTPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (protocolId)
	protocolId, _protocolIdErr := readBuffer.ReadUint8("protocolId", 8)
	if _protocolIdErr != nil {
		return nil, errors.Wrap(_protocolIdErr, "Error parsing 'protocolId' field")
	}
	if protocolId != TPKTPacket_PROTOCOLID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", TPKTPacket_PROTOCOLID) + " but got " + fmt.Sprintf("%d", protocolId))
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Implicit Field (len) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	len, _lenErr := readBuffer.ReadUint16("len", 16)
	_ = len
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field")
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := COTPPacketParse(readBuffer, uint16(uint16(len)-uint16(uint16(4))))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(COTPPacket)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	if closeErr := readBuffer.CloseContext("TPKTPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TPKTPacket")
	}

	// Create the instance
	return NewTPKTPacket(payload), nil
}

func (m *_TPKTPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("TPKTPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TPKTPacket")
	}

	// Const Field (protocolId)
	_protocolIdErr := writeBuffer.WriteUint8("protocolId", 8, 0x03)
	if _protocolIdErr != nil {
		return errors.Wrap(_protocolIdErr, "Error serializing 'protocolId' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 8, uint8(0x00))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (len) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	len := uint16(uint16(m.GetPayload().GetLengthInBytes()) + uint16(uint16(4)))
	_lenErr := writeBuffer.WriteUint16("len", 16, (len))
	if _lenErr != nil {
		return errors.Wrap(_lenErr, "Error serializing 'len' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}

	if popErr := writeBuffer.PopContext("TPKTPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TPKTPacket")
	}
	return nil
}

func (m *_TPKTPacket) isTPKTPacket() bool {
	return true
}

func (m *_TPKTPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
