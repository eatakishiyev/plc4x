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

// BACnetCalendarEntryDate is the corresponding interface of BACnetCalendarEntryDate
type BACnetCalendarEntryDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetCalendarEntry
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetContextTagDate
}

// BACnetCalendarEntryDateExactly can be used when we want exactly this type and not a type which fulfills BACnetCalendarEntryDate.
// This is useful for switch cases.
type BACnetCalendarEntryDateExactly interface {
	isBACnetCalendarEntryDate() bool
}

// _BACnetCalendarEntryDate is the data-structure of this message
type _BACnetCalendarEntryDate struct {
	*_BACnetCalendarEntry
	DateValue BACnetContextTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetCalendarEntryDate) InitializeParent(parent BACnetCalendarEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetCalendarEntryDate) GetParent() BACnetCalendarEntry {
	return m._BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCalendarEntryDate) GetDateValue() BACnetContextTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCalendarEntryDate factory function for _BACnetCalendarEntryDate
func NewBACnetCalendarEntryDate(dateValue BACnetContextTagDate, peekedTagHeader BACnetTagHeader) *_BACnetCalendarEntryDate {
	_result := &_BACnetCalendarEntryDate{
		DateValue:            dateValue,
		_BACnetCalendarEntry: NewBACnetCalendarEntry(peekedTagHeader),
	}
	_result._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetCalendarEntryDate(structType interface{}) BACnetCalendarEntryDate {
	if casted, ok := structType.(BACnetCalendarEntryDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCalendarEntryDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCalendarEntryDate) GetTypeName() string {
	return "BACnetCalendarEntryDate"
}

func (m *_BACnetCalendarEntryDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCalendarEntryDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetCalendarEntryDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCalendarEntryDateParse(readBuffer utils.ReadBuffer) (BACnetCalendarEntryDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCalendarEntryDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCalendarEntryDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_DATE))
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field")
	}
	dateValue := _dateValue.(BACnetContextTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetCalendarEntryDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCalendarEntryDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetCalendarEntryDate{
		DateValue:            dateValue,
		_BACnetCalendarEntry: &_BACnetCalendarEntry{},
	}
	_child._BACnetCalendarEntry._BACnetCalendarEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetCalendarEntryDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetCalendarEntryDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetCalendarEntryDate")
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

		if popErr := writeBuffer.PopContext("BACnetCalendarEntryDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetCalendarEntryDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetCalendarEntryDate) isBACnetCalendarEntryDate() bool {
	return true
}

func (m *_BACnetCalendarEntryDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
