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

// BACnetConstructedDataMaxInfoFrames is the corresponding interface of BACnetConstructedDataMaxInfoFrames
type BACnetConstructedDataMaxInfoFrames interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMaxInfoFramesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaxInfoFrames.
// This is useful for switch cases.
type BACnetConstructedDataMaxInfoFramesExactly interface {
	BACnetConstructedDataMaxInfoFrames
	isBACnetConstructedDataMaxInfoFrames() bool
}

// _BACnetConstructedDataMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataMaxInfoFrames struct {
	*_BACnetConstructedData
	MaxInfoFrames BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_INFO_FRAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaxInfoFrames factory function for _BACnetConstructedDataMaxInfoFrames
func NewBACnetConstructedDataMaxInfoFrames(maxInfoFrames BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxInfoFrames {
	_result := &_BACnetConstructedDataMaxInfoFrames{
		MaxInfoFrames:          maxInfoFrames,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxInfoFrames(structType interface{}) BACnetConstructedDataMaxInfoFrames {
	if casted, ok := structType.(BACnetConstructedDataMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataMaxInfoFrames"
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxInfoFrames) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxInfoFramesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaxInfoFrames, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxInfoFrames)
	if pullErr := readBuffer.PullContext("maxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxInfoFrames")
	}
	_maxInfoFrames, _maxInfoFramesErr := BACnetApplicationTagParse(readBuffer)
	if _maxInfoFramesErr != nil {
		return nil, errors.Wrap(_maxInfoFramesErr, "Error parsing 'maxInfoFrames' field")
	}
	maxInfoFrames := _maxInfoFrames.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxInfoFrames")
	}

	// Virtual field
	_actualValue := maxInfoFrames
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxInfoFrames")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaxInfoFrames{
		MaxInfoFrames: maxInfoFrames,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaxInfoFrames) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxInfoFrames")
		}

		// Simple Field (maxInfoFrames)
		if pushErr := writeBuffer.PushContext("maxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxInfoFrames")
		}
		_maxInfoFramesErr := writeBuffer.WriteSerializable(m.GetMaxInfoFrames())
		if popErr := writeBuffer.PopContext("maxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxInfoFrames")
		}
		if _maxInfoFramesErr != nil {
			return errors.Wrap(_maxInfoFramesErr, "Error serializing 'maxInfoFrames' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxInfoFrames")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxInfoFrames) isBACnetConstructedDataMaxInfoFrames() bool {
	return true
}

func (m *_BACnetConstructedDataMaxInfoFrames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
