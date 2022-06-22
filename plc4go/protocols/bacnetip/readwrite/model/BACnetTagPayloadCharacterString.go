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

// BACnetTagPayloadCharacterString is the corresponding interface of BACnetTagPayloadCharacterString
type BACnetTagPayloadCharacterString interface {
	utils.LengthAware
	utils.Serializable
	// GetEncoding returns Encoding (property field)
	GetEncoding() BACnetCharacterEncoding
	// GetValue returns Value (property field)
	GetValue() string
	// GetActualLengthInBit returns ActualLengthInBit (virtual field)
	GetActualLengthInBit() uint16
}

// BACnetTagPayloadCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadCharacterString.
// This is useful for switch cases.
type BACnetTagPayloadCharacterStringExactly interface {
	isBACnetTagPayloadCharacterString() bool
}

// _BACnetTagPayloadCharacterString is the data-structure of this message
type _BACnetTagPayloadCharacterString struct {
	Encoding BACnetCharacterEncoding
	Value    string

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadCharacterString) GetEncoding() BACnetCharacterEncoding {
	return m.Encoding
}

func (m *_BACnetTagPayloadCharacterString) GetValue() string {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadCharacterString) GetActualLengthInBit() uint16 {
	return uint16(uint16(uint16(m.ActualLength)*uint16(uint16(8))) - uint16(uint16(8)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadCharacterString factory function for _BACnetTagPayloadCharacterString
func NewBACnetTagPayloadCharacterString(encoding BACnetCharacterEncoding, value string, actualLength uint32) *_BACnetTagPayloadCharacterString {
	return &_BACnetTagPayloadCharacterString{Encoding: encoding, Value: value, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadCharacterString(structType interface{}) BACnetTagPayloadCharacterString {
	if casted, ok := structType.(BACnetTagPayloadCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadCharacterString) GetTypeName() string {
	return "BACnetTagPayloadCharacterString"
}

func (m *_BACnetTagPayloadCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (encoding)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.GetActualLengthInBit())

	return lengthInBits
}

func (m *_BACnetTagPayloadCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadCharacterStringParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (encoding)
	if pullErr := readBuffer.PullContext("encoding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for encoding")
	}
	_encoding, _encodingErr := BACnetCharacterEncodingParse(readBuffer)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	encoding := _encoding
	if closeErr := readBuffer.CloseContext("encoding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for encoding")
	}

	// Virtual field
	_actualLengthInBit := uint16(uint16(actualLength)*uint16(uint16(8))) - uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)
	_ = actualLengthInBit

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadCharacterString")
	}

	// Create the instance
	return NewBACnetTagPayloadCharacterString(encoding, value, actualLength), nil
}

func (m *_BACnetTagPayloadCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadCharacterString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadCharacterString")
	}

	// Simple Field (encoding)
	if pushErr := writeBuffer.PushContext("encoding"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for encoding")
	}
	_encodingErr := writeBuffer.WriteSerializable(m.GetEncoding())
	if popErr := writeBuffer.PopContext("encoding"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for encoding")
	}
	if _encodingErr != nil {
		return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
	}
	// Virtual field
	if _actualLengthInBitErr := writeBuffer.WriteVirtual("actualLengthInBit", m.GetActualLengthInBit()); _actualLengthInBitErr != nil {
		return errors.Wrap(_actualLengthInBitErr, "Error serializing 'actualLengthInBit' field")
	}

	// Simple Field (value)
	value := string(m.GetValue())
	_valueErr := writeBuffer.WriteString("value", uint32(m.GetActualLengthInBit()), "UTF-8", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadCharacterString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadCharacterString")
	}
	return nil
}

func (m *_BACnetTagPayloadCharacterString) isBACnetTagPayloadCharacterString() bool {
	return true
}

func (m *_BACnetTagPayloadCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
