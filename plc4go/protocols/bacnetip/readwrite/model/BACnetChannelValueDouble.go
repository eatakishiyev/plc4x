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

// BACnetChannelValueDouble is the corresponding interface of BACnetChannelValueDouble
type BACnetChannelValueDouble interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
}

// BACnetChannelValueDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueDouble.
// This is useful for switch cases.
type BACnetChannelValueDoubleExactly interface {
	isBACnetChannelValueDouble() bool
}

// _BACnetChannelValueDouble is the data-structure of this message
type _BACnetChannelValueDouble struct {
	*_BACnetChannelValue
	DoubleValue BACnetApplicationTagDouble
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueDouble) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueDouble) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueDouble factory function for _BACnetChannelValueDouble
func NewBACnetChannelValueDouble(doubleValue BACnetApplicationTagDouble, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueDouble {
	_result := &_BACnetChannelValueDouble{
		DoubleValue:         doubleValue,
		_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueDouble(structType interface{}) BACnetChannelValueDouble {
	if casted, ok := structType.(BACnetChannelValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueDouble) GetTypeName() string {
	return "BACnetChannelValueDouble"
}

func (m *_BACnetChannelValueDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetChannelValueDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueDoubleParse(readBuffer utils.ReadBuffer) (BACnetChannelValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doubleValue")
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParse(readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field")
	}
	doubleValue := _doubleValue.(BACnetApplicationTagDouble)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doubleValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueDouble")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueDouble{
		DoubleValue:         doubleValue,
		_BACnetChannelValue: &_BACnetChannelValue{},
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueDouble")
		}

		// Simple Field (doubleValue)
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doubleValue")
		}
		_doubleValueErr := writeBuffer.WriteSerializable(m.GetDoubleValue())
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doubleValue")
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueDouble")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetChannelValueDouble) isBACnetChannelValueDouble() bool {
	return true
}

func (m *_BACnetChannelValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
