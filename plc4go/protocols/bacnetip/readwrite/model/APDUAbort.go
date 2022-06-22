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

// APDUAbort is the corresponding interface of APDUAbort
type APDUAbort interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetServer returns Server (property field)
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetAbortReason returns AbortReason (property field)
	GetAbortReason() BACnetAbortReasonTagged
}

// APDUAbortExactly can be used when we want exactly this type and not a type which fulfills APDUAbort.
// This is useful for switch cases.
type APDUAbortExactly interface {
	isAPDUAbort() bool
}

// _APDUAbort is the data-structure of this message
type _APDUAbort struct {
	*_APDU
	Server           bool
	OriginalInvokeId uint8
	AbortReason      BACnetAbortReasonTagged

	// Arguments.
	ApduLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUAbort) GetApduType() ApduType {
	return ApduType_ABORT_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUAbort) InitializeParent(parent APDU) {}

func (m *_APDUAbort) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUAbort) GetServer() bool {
	return m.Server
}

func (m *_APDUAbort) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUAbort) GetAbortReason() BACnetAbortReasonTagged {
	return m.AbortReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUAbort factory function for _APDUAbort
func NewAPDUAbort(server bool, originalInvokeId uint8, abortReason BACnetAbortReasonTagged, apduLength uint16) *_APDUAbort {
	_result := &_APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
		_APDU:            NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUAbort(structType interface{}) APDUAbort {
	if casted, ok := structType.(APDUAbort); ok {
		return casted
	}
	if casted, ok := structType.(*APDUAbort); ok {
		return *casted
	}
	return nil
}

func (m *_APDUAbort) GetTypeName() string {
	return "APDUAbort"
}

func (m *_APDUAbort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_APDUAbort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 3

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (abortReason)
	lengthInBits += m.AbortReason.GetLengthInBits()

	return lengthInBits
}

func (m *_APDUAbort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUAbortParse(readBuffer utils.ReadBuffer, apduLength uint16) (APDUAbort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUAbort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUAbort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 3)
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

	// Simple Field (abortReason)
	if pullErr := readBuffer.PullContext("abortReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for abortReason")
	}
	_abortReason, _abortReasonErr := BACnetAbortReasonTaggedParse(readBuffer, uint32(uint32(1)))
	if _abortReasonErr != nil {
		return nil, errors.Wrap(_abortReasonErr, "Error parsing 'abortReason' field")
	}
	abortReason := _abortReason.(BACnetAbortReasonTagged)
	if closeErr := readBuffer.CloseContext("abortReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for abortReason")
	}

	if closeErr := readBuffer.CloseContext("APDUAbort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUAbort")
	}

	// Create a partially initialized instance
	_child := &_APDUAbort{
		Server:           server,
		OriginalInvokeId: originalInvokeId,
		AbortReason:      abortReason,
		_APDU:            &_APDU{},
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUAbort) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUAbort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUAbort")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 3, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
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

		// Simple Field (abortReason)
		if pushErr := writeBuffer.PushContext("abortReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for abortReason")
		}
		_abortReasonErr := writeBuffer.WriteSerializable(m.GetAbortReason())
		if popErr := writeBuffer.PopContext("abortReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for abortReason")
		}
		if _abortReasonErr != nil {
			return errors.Wrap(_abortReasonErr, "Error serializing 'abortReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUAbort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUAbort")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_APDUAbort) isAPDUAbort() bool {
	return true
}

func (m *_APDUAbort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
