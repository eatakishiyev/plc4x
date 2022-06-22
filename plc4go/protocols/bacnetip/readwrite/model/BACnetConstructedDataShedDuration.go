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

// BACnetConstructedDataShedDuration is the corresponding interface of BACnetConstructedDataShedDuration
type BACnetConstructedDataShedDuration interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetShedDuration returns ShedDuration (property field)
	GetShedDuration() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataShedDurationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataShedDuration.
// This is useful for switch cases.
type BACnetConstructedDataShedDurationExactly interface {
	isBACnetConstructedDataShedDuration() bool
}

// _BACnetConstructedDataShedDuration is the data-structure of this message
type _BACnetConstructedDataShedDuration struct {
	*_BACnetConstructedData
	ShedDuration BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataShedDuration) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SHED_DURATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataShedDuration) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataShedDuration) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetShedDuration() BACnetApplicationTagUnsignedInteger {
	return m.ShedDuration
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataShedDuration) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetShedDuration())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataShedDuration factory function for _BACnetConstructedDataShedDuration
func NewBACnetConstructedDataShedDuration(shedDuration BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataShedDuration {
	_result := &_BACnetConstructedDataShedDuration{
		ShedDuration:           shedDuration,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataShedDuration(structType interface{}) BACnetConstructedDataShedDuration {
	if casted, ok := structType.(BACnetConstructedDataShedDuration); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataShedDuration); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataShedDuration) GetTypeName() string {
	return "BACnetConstructedDataShedDuration"
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (shedDuration)
	lengthInBits += m.ShedDuration.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataShedDuration) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataShedDurationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataShedDuration, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataShedDuration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataShedDuration")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (shedDuration)
	if pullErr := readBuffer.PullContext("shedDuration"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for shedDuration")
	}
	_shedDuration, _shedDurationErr := BACnetApplicationTagParse(readBuffer)
	if _shedDurationErr != nil {
		return nil, errors.Wrap(_shedDurationErr, "Error parsing 'shedDuration' field")
	}
	shedDuration := _shedDuration.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("shedDuration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for shedDuration")
	}

	// Virtual field
	_actualValue := shedDuration
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataShedDuration"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataShedDuration")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataShedDuration{
		ShedDuration:           shedDuration,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataShedDuration) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataShedDuration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataShedDuration")
		}

		// Simple Field (shedDuration)
		if pushErr := writeBuffer.PushContext("shedDuration"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for shedDuration")
		}
		_shedDurationErr := writeBuffer.WriteSerializable(m.GetShedDuration())
		if popErr := writeBuffer.PopContext("shedDuration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for shedDuration")
		}
		if _shedDurationErr != nil {
			return errors.Wrap(_shedDurationErr, "Error serializing 'shedDuration' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataShedDuration"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataShedDuration")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataShedDuration) isBACnetConstructedDataShedDuration() bool {
	return true
}

func (m *_BACnetConstructedDataShedDuration) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
