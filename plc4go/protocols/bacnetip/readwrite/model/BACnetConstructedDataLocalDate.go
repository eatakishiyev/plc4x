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

// BACnetConstructedDataLocalDate is the corresponding interface of BACnetConstructedDataLocalDate
type BACnetConstructedDataLocalDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLocalDate returns LocalDate (property field)
	GetLocalDate() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
}

// BACnetConstructedDataLocalDateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLocalDate.
// This is useful for switch cases.
type BACnetConstructedDataLocalDateExactly interface {
	isBACnetConstructedDataLocalDate() bool
}

// _BACnetConstructedDataLocalDate is the data-structure of this message
type _BACnetConstructedDataLocalDate struct {
	*_BACnetConstructedData
	LocalDate BACnetApplicationTagDate

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLocalDate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocalDate) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLocalDate) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetLocalDate() BACnetApplicationTagDate {
	return m.LocalDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetActualValue() BACnetApplicationTagDate {
	return CastBACnetApplicationTagDate(m.GetLocalDate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocalDate factory function for _BACnetConstructedDataLocalDate
func NewBACnetConstructedDataLocalDate(localDate BACnetApplicationTagDate, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLocalDate {
	_result := &_BACnetConstructedDataLocalDate{
		LocalDate:              localDate,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocalDate(structType interface{}) BACnetConstructedDataLocalDate {
	if casted, ok := structType.(BACnetConstructedDataLocalDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocalDate) GetTypeName() string {
	return "BACnetConstructedDataLocalDate"
}

func (m *_BACnetConstructedDataLocalDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLocalDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (localDate)
	lengthInBits += m.LocalDate.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLocalDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLocalDateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLocalDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localDate)
	if pullErr := readBuffer.PullContext("localDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localDate")
	}
	_localDate, _localDateErr := BACnetApplicationTagParse(readBuffer)
	if _localDateErr != nil {
		return nil, errors.Wrap(_localDateErr, "Error parsing 'localDate' field")
	}
	localDate := _localDate.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("localDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localDate")
	}

	// Virtual field
	_actualValue := localDate
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLocalDate{
		LocalDate:              localDate,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLocalDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalDate")
		}

		// Simple Field (localDate)
		if pushErr := writeBuffer.PushContext("localDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for localDate")
		}
		_localDateErr := writeBuffer.WriteSerializable(m.GetLocalDate())
		if popErr := writeBuffer.PopContext("localDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for localDate")
		}
		if _localDateErr != nil {
			return errors.Wrap(_localDateErr, "Error serializing 'localDate' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLocalDate) isBACnetConstructedDataLocalDate() bool {
	return true
}

func (m *_BACnetConstructedDataLocalDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
