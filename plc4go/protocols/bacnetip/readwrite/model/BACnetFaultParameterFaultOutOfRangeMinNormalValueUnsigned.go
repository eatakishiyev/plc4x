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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedExactly interface {
	isBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValue
	UnsignedValue BACnetApplicationTagUnsignedInteger

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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(unsignedValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		UnsignedValue: unsignedValue,
		_BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned(structType interface{}) BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
	_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned{
		UnsignedValue: unsignedValue,
		_BACnetFaultParameterFaultOutOfRangeMinNormalValue: &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{},
	}
	_child._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}

		// Simple Field (unsignedValue)
		if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for unsignedValue")
		}
		_unsignedValueErr := writeBuffer.WriteSerializable(m.GetUnsignedValue())
		if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for unsignedValue")
		}
		if _unsignedValueErr != nil {
			return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) isBACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
