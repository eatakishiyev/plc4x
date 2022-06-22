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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateExactly interface {
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	DateValue BACnetApplicationTagDate

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate(dateValue BACnetApplicationTagDate, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate{
		DateValue: dateValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDateParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParse(readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field")
	}
	dateValue := _dateValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate{
		DateValue: dateValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		_dateValueErr := writeBuffer.WriteSerializable(m.GetDateValue())
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) isBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
