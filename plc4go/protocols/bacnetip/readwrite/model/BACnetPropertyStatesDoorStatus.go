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

// BACnetPropertyStatesDoorStatus is the corresponding interface of BACnetPropertyStatesDoorStatus
type BACnetPropertyStatesDoorStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetDoorStatus returns DoorStatus (property field)
	GetDoorStatus() BACnetDoorStatusTagged
}

// BACnetPropertyStatesDoorStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesDoorStatus.
// This is useful for switch cases.
type BACnetPropertyStatesDoorStatusExactly interface {
	isBACnetPropertyStatesDoorStatus() bool
}

// _BACnetPropertyStatesDoorStatus is the data-structure of this message
type _BACnetPropertyStatesDoorStatus struct {
	*_BACnetPropertyStates
	DoorStatus BACnetDoorStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesDoorStatus) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesDoorStatus) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesDoorStatus) GetDoorStatus() BACnetDoorStatusTagged {
	return m.DoorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesDoorStatus factory function for _BACnetPropertyStatesDoorStatus
func NewBACnetPropertyStatesDoorStatus(doorStatus BACnetDoorStatusTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesDoorStatus {
	_result := &_BACnetPropertyStatesDoorStatus{
		DoorStatus:            doorStatus,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesDoorStatus(structType interface{}) BACnetPropertyStatesDoorStatus {
	if casted, ok := structType.(BACnetPropertyStatesDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesDoorStatus) GetTypeName() string {
	return "BACnetPropertyStatesDoorStatus"
}

func (m *_BACnetPropertyStatesDoorStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesDoorStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorStatus)
	lengthInBits += m.DoorStatus.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesDoorStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesDoorStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesDoorStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorStatus)
	if pullErr := readBuffer.PullContext("doorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorStatus")
	}
	_doorStatus, _doorStatusErr := BACnetDoorStatusTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _doorStatusErr != nil {
		return nil, errors.Wrap(_doorStatusErr, "Error parsing 'doorStatus' field")
	}
	doorStatus := _doorStatus.(BACnetDoorStatusTagged)
	if closeErr := readBuffer.CloseContext("doorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesDoorStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesDoorStatus{
		DoorStatus:            doorStatus,
		_BACnetPropertyStates: &_BACnetPropertyStates{},
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesDoorStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesDoorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesDoorStatus")
		}

		// Simple Field (doorStatus)
		if pushErr := writeBuffer.PushContext("doorStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorStatus")
		}
		_doorStatusErr := writeBuffer.WriteSerializable(m.GetDoorStatus())
		if popErr := writeBuffer.PopContext("doorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorStatus")
		}
		if _doorStatusErr != nil {
			return errors.Wrap(_doorStatusErr, "Error serializing 'doorStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesDoorStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesDoorStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesDoorStatus) isBACnetPropertyStatesDoorStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
