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

// BACnetServiceAckVTOpen is the corresponding interface of BACnetServiceAckVTOpen
type BACnetServiceAckVTOpen interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetRemoteVtSessionIdentifier returns RemoteVtSessionIdentifier (property field)
	GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger
}

// BACnetServiceAckVTOpenExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckVTOpen.
// This is useful for switch cases.
type BACnetServiceAckVTOpenExactly interface {
	isBACnetServiceAckVTOpen() bool
}

// _BACnetServiceAckVTOpen is the data-structure of this message
type _BACnetServiceAckVTOpen struct {
	*_BACnetServiceAck
	RemoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceAckLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_OPEN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckVTOpen) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckVTOpen) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckVTOpen) GetRemoteVtSessionIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.RemoteVtSessionIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckVTOpen factory function for _BACnetServiceAckVTOpen
func NewBACnetServiceAckVTOpen(remoteVtSessionIdentifier BACnetApplicationTagUnsignedInteger, serviceAckLength uint16) *_BACnetServiceAckVTOpen {
	_result := &_BACnetServiceAckVTOpen{
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
		_BACnetServiceAck:         NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckVTOpen(structType interface{}) BACnetServiceAckVTOpen {
	if casted, ok := structType.(BACnetServiceAckVTOpen); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckVTOpen); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckVTOpen) GetTypeName() string {
	return "BACnetServiceAckVTOpen"
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (remoteVtSessionIdentifier)
	lengthInBits += m.RemoteVtSessionIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetServiceAckVTOpen) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckVTOpenParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (BACnetServiceAckVTOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckVTOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckVTOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (remoteVtSessionIdentifier)
	if pullErr := readBuffer.PullContext("remoteVtSessionIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for remoteVtSessionIdentifier")
	}
	_remoteVtSessionIdentifier, _remoteVtSessionIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _remoteVtSessionIdentifierErr != nil {
		return nil, errors.Wrap(_remoteVtSessionIdentifierErr, "Error parsing 'remoteVtSessionIdentifier' field")
	}
	remoteVtSessionIdentifier := _remoteVtSessionIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("remoteVtSessionIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for remoteVtSessionIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckVTOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckVTOpen")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckVTOpen{
		RemoteVtSessionIdentifier: remoteVtSessionIdentifier,
		_BACnetServiceAck:         &_BACnetServiceAck{},
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckVTOpen) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckVTOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckVTOpen")
		}

		// Simple Field (remoteVtSessionIdentifier)
		if pushErr := writeBuffer.PushContext("remoteVtSessionIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for remoteVtSessionIdentifier")
		}
		_remoteVtSessionIdentifierErr := writeBuffer.WriteSerializable(m.GetRemoteVtSessionIdentifier())
		if popErr := writeBuffer.PopContext("remoteVtSessionIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for remoteVtSessionIdentifier")
		}
		if _remoteVtSessionIdentifierErr != nil {
			return errors.Wrap(_remoteVtSessionIdentifierErr, "Error serializing 'remoteVtSessionIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckVTOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckVTOpen")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckVTOpen) isBACnetServiceAckVTOpen() bool {
	return true
}

func (m *_BACnetServiceAckVTOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
