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

// BACnetConstructedDataEscalatorMode is the corresponding interface of BACnetConstructedDataEscalatorMode
type BACnetConstructedDataEscalatorMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEscalatorMode returns EscalatorMode (property field)
	GetEscalatorMode() BACnetEscalatorModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEscalatorModeTagged
}

// BACnetConstructedDataEscalatorModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEscalatorMode.
// This is useful for switch cases.
type BACnetConstructedDataEscalatorModeExactly interface {
	isBACnetConstructedDataEscalatorMode() bool
}

// _BACnetConstructedDataEscalatorMode is the data-structure of this message
type _BACnetConstructedDataEscalatorMode struct {
	*_BACnetConstructedData
	EscalatorMode BACnetEscalatorModeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEscalatorMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ESCALATOR_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEscalatorMode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEscalatorMode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetEscalatorMode() BACnetEscalatorModeTagged {
	return m.EscalatorMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEscalatorMode) GetActualValue() BACnetEscalatorModeTagged {
	return CastBACnetEscalatorModeTagged(m.GetEscalatorMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEscalatorMode factory function for _BACnetConstructedDataEscalatorMode
func NewBACnetConstructedDataEscalatorMode(escalatorMode BACnetEscalatorModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEscalatorMode {
	_result := &_BACnetConstructedDataEscalatorMode{
		EscalatorMode:          escalatorMode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEscalatorMode(structType interface{}) BACnetConstructedDataEscalatorMode {
	if casted, ok := structType.(BACnetConstructedDataEscalatorMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEscalatorMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEscalatorMode) GetTypeName() string {
	return "BACnetConstructedDataEscalatorMode"
}

func (m *_BACnetConstructedDataEscalatorMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEscalatorMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (escalatorMode)
	lengthInBits += m.EscalatorMode.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEscalatorMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEscalatorModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEscalatorMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEscalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEscalatorMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (escalatorMode)
	if pullErr := readBuffer.PullContext("escalatorMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for escalatorMode")
	}
	_escalatorMode, _escalatorModeErr := BACnetEscalatorModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _escalatorModeErr != nil {
		return nil, errors.Wrap(_escalatorModeErr, "Error parsing 'escalatorMode' field")
	}
	escalatorMode := _escalatorMode.(BACnetEscalatorModeTagged)
	if closeErr := readBuffer.CloseContext("escalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for escalatorMode")
	}

	// Virtual field
	_actualValue := escalatorMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEscalatorMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEscalatorMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEscalatorMode{
		EscalatorMode:          escalatorMode,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEscalatorMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEscalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEscalatorMode")
		}

		// Simple Field (escalatorMode)
		if pushErr := writeBuffer.PushContext("escalatorMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for escalatorMode")
		}
		_escalatorModeErr := writeBuffer.WriteSerializable(m.GetEscalatorMode())
		if popErr := writeBuffer.PopContext("escalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for escalatorMode")
		}
		if _escalatorModeErr != nil {
			return errors.Wrap(_escalatorModeErr, "Error serializing 'escalatorMode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEscalatorMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEscalatorMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEscalatorMode) isBACnetConstructedDataEscalatorMode() bool {
	return true
}

func (m *_BACnetConstructedDataEscalatorMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
