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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUSegmentAck is the corresponding interface of APDUSegmentAck
type APDUSegmentAck interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetNegativeAck returns NegativeAck (property field)
	GetNegativeAck() bool
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() uint8
}

// APDUSegmentAckExactly can be used when we want exactly this type and not a type which fulfills APDUSegmentAck.
// This is useful for switch cases.
type APDUSegmentAckExactly interface {
	APDUSegmentAck
	isAPDUSegmentAck() bool
}

// _APDUSegmentAck is the data-structure of this message
type _APDUSegmentAck struct {
	*_APDU
	NegativeAck        bool
	Server             bool
	OriginalInvokeId   uint8
	SequenceNumber     uint8
	ProposedWindowSize uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSegmentAck) GetApduType() ApduType {
	return ApduType_SEGMENT_ACK_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSegmentAck) InitializeParent(parent APDU) {}

func (m *_APDUSegmentAck) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSegmentAck) GetNegativeAck() bool {
	return m.NegativeAck
}

func (m *_APDUSegmentAck) GetServer() bool {
	return m.Server
}

func (m *_APDUSegmentAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSegmentAck) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *_APDUSegmentAck) GetProposedWindowSize() uint8 {
	return m.ProposedWindowSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUSegmentAck factory function for _APDUSegmentAck
func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, proposedWindowSize uint8, apduLength uint16) *_APDUSegmentAck {
	_result := &_APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		_APDU:              NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUSegmentAck(structType interface{}) APDUSegmentAck {
	if casted, ok := structType.(APDUSegmentAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSegmentAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSegmentAck) GetTypeName() string {
	return "APDUSegmentAck"
}

func (m *_APDUSegmentAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_APDUSegmentAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (negativeAck)
	lengthInBits += 1

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Simple field (proposedWindowSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *_APDUSegmentAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUSegmentAckParse(readBuffer utils.ReadBuffer, apduLength uint16) (APDUSegmentAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUSegmentAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSegmentAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
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

	// Simple Field (negativeAck)
	_negativeAck, _negativeAckErr := readBuffer.ReadBit("negativeAck")
	if _negativeAckErr != nil {
		return nil, errors.Wrap(_negativeAckErr, "Error parsing 'negativeAck' field")
	}
	negativeAck := _negativeAck

	// Simple Field (server)
	_server, _serverErr := readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field")
	}
	server := _server

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (proposedWindowSize)
	_proposedWindowSize, _proposedWindowSizeErr := readBuffer.ReadUint8("proposedWindowSize", 8)
	if _proposedWindowSizeErr != nil {
		return nil, errors.Wrap(_proposedWindowSizeErr, "Error parsing 'proposedWindowSize' field")
	}
	proposedWindowSize := _proposedWindowSize

	if closeErr := readBuffer.CloseContext("APDUSegmentAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSegmentAck")
	}

	// Create a partially initialized instance
	_child := &_APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUSegmentAck) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSegmentAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSegmentAck")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (negativeAck)
		negativeAck := bool(m.GetNegativeAck())
		_negativeAckErr := writeBuffer.WriteBit("negativeAck", (negativeAck))
		if _negativeAckErr != nil {
			return errors.Wrap(_negativeAckErr, "Error serializing 'negativeAck' field")
		}

		// Simple Field (server)
		server := bool(m.GetServer())
		_serverErr := writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.GetSequenceNumber())
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (proposedWindowSize)
		proposedWindowSize := uint8(m.GetProposedWindowSize())
		_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, (proposedWindowSize))
		if _proposedWindowSizeErr != nil {
			return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
		}

		if popErr := writeBuffer.PopContext("APDUSegmentAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSegmentAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_APDUSegmentAck) isAPDUSegmentAck() bool {
	return true
}

func (m *_APDUSegmentAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
