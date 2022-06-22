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

// BACnetConstructedDataTimeOfActiveTimeReset is the corresponding interface of BACnetConstructedDataTimeOfActiveTimeReset
type BACnetConstructedDataTimeOfActiveTimeReset interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeOfActiveTimeReset returns TimeOfActiveTimeReset (property field)
	GetTimeOfActiveTimeReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataTimeOfActiveTimeResetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeOfActiveTimeReset.
// This is useful for switch cases.
type BACnetConstructedDataTimeOfActiveTimeResetExactly interface {
	BACnetConstructedDataTimeOfActiveTimeReset
	isBACnetConstructedDataTimeOfActiveTimeReset() bool
}

// _BACnetConstructedDataTimeOfActiveTimeReset is the data-structure of this message
type _BACnetConstructedDataTimeOfActiveTimeReset struct {
	*_BACnetConstructedData
	TimeOfActiveTimeReset BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_ACTIVE_TIME_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTimeOfActiveTimeReset() BACnetDateTime {
	return m.TimeOfActiveTimeReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetTimeOfActiveTimeReset())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfActiveTimeReset factory function for _BACnetConstructedDataTimeOfActiveTimeReset
func NewBACnetConstructedDataTimeOfActiveTimeReset(timeOfActiveTimeReset BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfActiveTimeReset {
	_result := &_BACnetConstructedDataTimeOfActiveTimeReset{
		TimeOfActiveTimeReset:  timeOfActiveTimeReset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfActiveTimeReset(structType interface{}) BACnetConstructedDataTimeOfActiveTimeReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfActiveTimeReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfActiveTimeReset"
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeOfActiveTimeReset)
	lengthInBits += m.TimeOfActiveTimeReset.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeOfActiveTimeResetParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfActiveTimeReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfActiveTimeReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfActiveTimeReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfActiveTimeReset)
	if pullErr := readBuffer.PullContext("timeOfActiveTimeReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeOfActiveTimeReset")
	}
	_timeOfActiveTimeReset, _timeOfActiveTimeResetErr := BACnetDateTimeParse(readBuffer)
	if _timeOfActiveTimeResetErr != nil {
		return nil, errors.Wrap(_timeOfActiveTimeResetErr, "Error parsing 'timeOfActiveTimeReset' field")
	}
	timeOfActiveTimeReset := _timeOfActiveTimeReset.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("timeOfActiveTimeReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeOfActiveTimeReset")
	}

	// Virtual field
	_actualValue := timeOfActiveTimeReset
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfActiveTimeReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfActiveTimeReset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeOfActiveTimeReset{
		TimeOfActiveTimeReset: timeOfActiveTimeReset,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfActiveTimeReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfActiveTimeReset")
		}

		// Simple Field (timeOfActiveTimeReset)
		if pushErr := writeBuffer.PushContext("timeOfActiveTimeReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfActiveTimeReset")
		}
		_timeOfActiveTimeResetErr := writeBuffer.WriteSerializable(m.GetTimeOfActiveTimeReset())
		if popErr := writeBuffer.PopContext("timeOfActiveTimeReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfActiveTimeReset")
		}
		if _timeOfActiveTimeResetErr != nil {
			return errors.Wrap(_timeOfActiveTimeResetErr, "Error serializing 'timeOfActiveTimeReset' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfActiveTimeReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfActiveTimeReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) isBACnetConstructedDataTimeOfActiveTimeReset() bool {
	return true
}

func (m *_BACnetConstructedDataTimeOfActiveTimeReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
