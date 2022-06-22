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

// BACnetConstructedDataLifeSafetyPointPresentValue is the corresponding interface of BACnetConstructedDataLifeSafetyPointPresentValue
type BACnetConstructedDataLifeSafetyPointPresentValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetLifeSafetyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyStateTagged
}

// BACnetConstructedDataLifeSafetyPointPresentValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLifeSafetyPointPresentValue.
// This is useful for switch cases.
type BACnetConstructedDataLifeSafetyPointPresentValueExactly interface {
	isBACnetConstructedDataLifeSafetyPointPresentValue() bool
}

// _BACnetConstructedDataLifeSafetyPointPresentValue is the data-structure of this message
type _BACnetConstructedDataLifeSafetyPointPresentValue struct {
	*_BACnetConstructedData
	PresentValue BACnetLifeSafetyStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_POINT
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetPresentValue() BACnetLifeSafetyStateTagged {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetActualValue() BACnetLifeSafetyStateTagged {
	return CastBACnetLifeSafetyStateTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyPointPresentValue factory function for _BACnetConstructedDataLifeSafetyPointPresentValue
func NewBACnetConstructedDataLifeSafetyPointPresentValue(presentValue BACnetLifeSafetyStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyPointPresentValue {
	_result := &_BACnetConstructedDataLifeSafetyPointPresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyPointPresentValue(structType interface{}) BACnetConstructedDataLifeSafetyPointPresentValue {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyPointPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyPointPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyPointPresentValue"
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLifeSafetyPointPresentValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLifeSafetyPointPresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyPointPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyPointPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for presentValue")
	}
	_presentValue, _presentValueErr := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _presentValueErr != nil {
		return nil, errors.Wrap(_presentValueErr, "Error parsing 'presentValue' field")
	}
	presentValue := _presentValue.(BACnetLifeSafetyStateTagged)
	if closeErr := readBuffer.CloseContext("presentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for presentValue")
	}

	// Virtual field
	_actualValue := presentValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyPointPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyPointPresentValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLifeSafetyPointPresentValue{
		PresentValue:           presentValue,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyPointPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyPointPresentValue")
		}

		// Simple Field (presentValue)
		if pushErr := writeBuffer.PushContext("presentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for presentValue")
		}
		_presentValueErr := writeBuffer.WriteSerializable(m.GetPresentValue())
		if popErr := writeBuffer.PopContext("presentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for presentValue")
		}
		if _presentValueErr != nil {
			return errors.Wrap(_presentValueErr, "Error serializing 'presentValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyPointPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyPointPresentValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) isBACnetConstructedDataLifeSafetyPointPresentValue() bool {
	return true
}

func (m *_BACnetConstructedDataLifeSafetyPointPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
