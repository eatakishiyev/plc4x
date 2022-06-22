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

// BACnetConstructedDataEnable is the corresponding interface of BACnetConstructedDataEnable
type BACnetConstructedDataEnable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEnable returns Enable (property field)
	GetEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEnable.
// This is useful for switch cases.
type BACnetConstructedDataEnableExactly interface {
	isBACnetConstructedDataEnable() bool
}

// _BACnetConstructedDataEnable is the data-structure of this message
type _BACnetConstructedDataEnable struct {
	*_BACnetConstructedData
	Enable BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetEnable() BACnetApplicationTagBoolean {
	return m.Enable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnable) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEnable factory function for _BACnetConstructedDataEnable
func NewBACnetConstructedDataEnable(enable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnable {
	_result := &_BACnetConstructedDataEnable{
		Enable:                 enable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnable(structType interface{}) BACnetConstructedDataEnable {
	if casted, ok := structType.(BACnetConstructedDataEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnable) GetTypeName() string {
	return "BACnetConstructedDataEnable"
}

func (m *_BACnetConstructedDataEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enable)
	lengthInBits += m.Enable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEnableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enable)
	if pullErr := readBuffer.PullContext("enable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enable")
	}
	_enable, _enableErr := BACnetApplicationTagParse(readBuffer)
	if _enableErr != nil {
		return nil, errors.Wrap(_enableErr, "Error parsing 'enable' field")
	}
	enable := _enable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("enable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enable")
	}

	// Virtual field
	_actualValue := enable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEnable{
		Enable:                 enable,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEnable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnable")
		}

		// Simple Field (enable)
		if pushErr := writeBuffer.PushContext("enable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enable")
		}
		_enableErr := writeBuffer.WriteSerializable(m.GetEnable())
		if popErr := writeBuffer.PopContext("enable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enable")
		}
		if _enableErr != nil {
			return errors.Wrap(_enableErr, "Error serializing 'enable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnable) isBACnetConstructedDataEnable() bool {
	return true
}

func (m *_BACnetConstructedDataEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
