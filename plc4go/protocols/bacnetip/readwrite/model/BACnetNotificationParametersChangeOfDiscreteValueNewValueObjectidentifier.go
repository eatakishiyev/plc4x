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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
type BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetObjectidentifierValue returns ObjectidentifierValue (property field)
	GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifierExactly interface {
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	ObjectidentifierValue BACnetApplicationTagObjectIdentifier

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

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetObjectidentifierValue() BACnetApplicationTagObjectIdentifier {
	return m.ObjectidentifierValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier(objectidentifierValue BACnetApplicationTagObjectIdentifier, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectidentifierValue)
	lengthInBits += m.ObjectidentifierValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectidentifierValue)
	if pullErr := readBuffer.PullContext("objectidentifierValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectidentifierValue")
	}
	_objectidentifierValue, _objectidentifierValueErr := BACnetApplicationTagParse(readBuffer)
	if _objectidentifierValueErr != nil {
		return nil, errors.Wrap(_objectidentifierValueErr, "Error parsing 'objectidentifierValue' field")
	}
	objectidentifierValue := _objectidentifierValue.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectidentifierValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectidentifierValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier{
		ObjectidentifierValue: objectidentifierValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
		}

		// Simple Field (objectidentifierValue)
		if pushErr := writeBuffer.PushContext("objectidentifierValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectidentifierValue")
		}
		_objectidentifierValueErr := writeBuffer.WriteSerializable(m.GetObjectidentifierValue())
		if popErr := writeBuffer.PopContext("objectidentifierValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectidentifierValue")
		}
		if _objectidentifierValueErr != nil {
			return errors.Wrap(_objectidentifierValueErr, "Error serializing 'objectidentifierValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) isBACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueObjectidentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
