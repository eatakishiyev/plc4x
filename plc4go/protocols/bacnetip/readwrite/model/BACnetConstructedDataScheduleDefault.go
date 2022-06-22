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

// BACnetConstructedDataScheduleDefault is the corresponding interface of BACnetConstructedDataScheduleDefault
type BACnetConstructedDataScheduleDefault interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetScheduleDefault returns ScheduleDefault (property field)
	GetScheduleDefault() BACnetConstructedDataElement
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetConstructedDataElement
}

// BACnetConstructedDataScheduleDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataScheduleDefault.
// This is useful for switch cases.
type BACnetConstructedDataScheduleDefaultExactly interface {
	isBACnetConstructedDataScheduleDefault() bool
}

// _BACnetConstructedDataScheduleDefault is the data-structure of this message
type _BACnetConstructedDataScheduleDefault struct {
	*_BACnetConstructedData
	ScheduleDefault BACnetConstructedDataElement

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataScheduleDefault) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataScheduleDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SCHEDULE_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataScheduleDefault) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataScheduleDefault) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataScheduleDefault) GetScheduleDefault() BACnetConstructedDataElement {
	return m.ScheduleDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataScheduleDefault) GetActualValue() BACnetConstructedDataElement {
	return CastBACnetConstructedDataElement(m.GetScheduleDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataScheduleDefault factory function for _BACnetConstructedDataScheduleDefault
func NewBACnetConstructedDataScheduleDefault(scheduleDefault BACnetConstructedDataElement, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataScheduleDefault {
	_result := &_BACnetConstructedDataScheduleDefault{
		ScheduleDefault:        scheduleDefault,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataScheduleDefault(structType interface{}) BACnetConstructedDataScheduleDefault {
	if casted, ok := structType.(BACnetConstructedDataScheduleDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataScheduleDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataScheduleDefault) GetTypeName() string {
	return "BACnetConstructedDataScheduleDefault"
}

func (m *_BACnetConstructedDataScheduleDefault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataScheduleDefault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (scheduleDefault)
	lengthInBits += m.ScheduleDefault.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataScheduleDefault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataScheduleDefaultParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataScheduleDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataScheduleDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataScheduleDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (scheduleDefault)
	if pullErr := readBuffer.PullContext("scheduleDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for scheduleDefault")
	}
	_scheduleDefault, _scheduleDefaultErr := BACnetConstructedDataElementParse(readBuffer, BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(propertyIdentifierArgument), nil)
	if _scheduleDefaultErr != nil {
		return nil, errors.Wrap(_scheduleDefaultErr, "Error parsing 'scheduleDefault' field")
	}
	scheduleDefault := _scheduleDefault.(BACnetConstructedDataElement)
	if closeErr := readBuffer.CloseContext("scheduleDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for scheduleDefault")
	}

	// Virtual field
	_actualValue := scheduleDefault
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataScheduleDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataScheduleDefault")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataScheduleDefault{
		ScheduleDefault:        scheduleDefault,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataScheduleDefault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataScheduleDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataScheduleDefault")
		}

		// Simple Field (scheduleDefault)
		if pushErr := writeBuffer.PushContext("scheduleDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for scheduleDefault")
		}
		_scheduleDefaultErr := writeBuffer.WriteSerializable(m.GetScheduleDefault())
		if popErr := writeBuffer.PopContext("scheduleDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for scheduleDefault")
		}
		if _scheduleDefaultErr != nil {
			return errors.Wrap(_scheduleDefaultErr, "Error serializing 'scheduleDefault' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataScheduleDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataScheduleDefault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataScheduleDefault) isBACnetConstructedDataScheduleDefault() bool {
	return true
}

func (m *_BACnetConstructedDataScheduleDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
