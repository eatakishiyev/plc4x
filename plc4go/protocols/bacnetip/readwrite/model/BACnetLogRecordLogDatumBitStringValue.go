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

// BACnetLogRecordLogDatumBitStringValue is the corresponding interface of BACnetLogRecordLogDatumBitStringValue
type BACnetLogRecordLogDatumBitStringValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetContextTagBitString
}

// BACnetLogRecordLogDatumBitStringValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumBitStringValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumBitStringValueExactly interface {
	BACnetLogRecordLogDatumBitStringValue
	isBACnetLogRecordLogDatumBitStringValue() bool
}

// _BACnetLogRecordLogDatumBitStringValue is the data-structure of this message
type _BACnetLogRecordLogDatumBitStringValue struct {
	*_BACnetLogRecordLogDatum
	BitStringValue BACnetContextTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumBitStringValue) InitializeParent(parent BACnetLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumBitStringValue) GetBitStringValue() BACnetContextTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogRecordLogDatumBitStringValue factory function for _BACnetLogRecordLogDatumBitStringValue
func NewBACnetLogRecordLogDatumBitStringValue(bitStringValue BACnetContextTagBitString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLogRecordLogDatumBitStringValue {
	_result := &_BACnetLogRecordLogDatumBitStringValue{
		BitStringValue:           bitStringValue,
		_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumBitStringValue(structType interface{}) BACnetLogRecordLogDatumBitStringValue {
	if casted, ok := structType.(BACnetLogRecordLogDatumBitStringValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumBitStringValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumBitStringValue"
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLogRecordLogDatumBitStringValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumBitStringValueParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumBitStringValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumBitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumBitStringValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitStringValue")
	}
	_bitStringValue, _bitStringValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(6)), BACnetDataType(BACnetDataType_BIT_STRING))
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field")
	}
	bitStringValue := _bitStringValue.(BACnetContextTagBitString)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumBitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumBitStringValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumBitStringValue{
		BitStringValue: bitStringValue,
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumBitStringValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumBitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumBitStringValue")
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		_bitStringValueErr := writeBuffer.WriteSerializable(m.GetBitStringValue())
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumBitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumBitStringValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetLogRecordLogDatumBitStringValue) isBACnetLogRecordLogDatumBitStringValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumBitStringValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
