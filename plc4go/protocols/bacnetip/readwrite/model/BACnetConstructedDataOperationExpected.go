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

// BACnetConstructedDataOperationExpected is the corresponding interface of BACnetConstructedDataOperationExpected
type BACnetConstructedDataOperationExpected interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLifeSafetyOperations returns LifeSafetyOperations (property field)
	GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyOperationTagged
}

// BACnetConstructedDataOperationExpectedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOperationExpected.
// This is useful for switch cases.
type BACnetConstructedDataOperationExpectedExactly interface {
	isBACnetConstructedDataOperationExpected() bool
}

// _BACnetConstructedDataOperationExpected is the data-structure of this message
type _BACnetConstructedDataOperationExpected struct {
	*_BACnetConstructedData
	LifeSafetyOperations BACnetLifeSafetyOperationTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOperationExpected) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_EXPECTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOperationExpected) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOperationExpected) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetLifeSafetyOperations() BACnetLifeSafetyOperationTagged {
	return m.LifeSafetyOperations
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOperationExpected) GetActualValue() BACnetLifeSafetyOperationTagged {
	return CastBACnetLifeSafetyOperationTagged(m.GetLifeSafetyOperations())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOperationExpected factory function for _BACnetConstructedDataOperationExpected
func NewBACnetConstructedDataOperationExpected(lifeSafetyOperations BACnetLifeSafetyOperationTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOperationExpected {
	_result := &_BACnetConstructedDataOperationExpected{
		LifeSafetyOperations:   lifeSafetyOperations,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOperationExpected(structType interface{}) BACnetConstructedDataOperationExpected {
	if casted, ok := structType.(BACnetConstructedDataOperationExpected); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationExpected); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOperationExpected) GetTypeName() string {
	return "BACnetConstructedDataOperationExpected"
}

func (m *_BACnetConstructedDataOperationExpected) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOperationExpected) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lifeSafetyOperations)
	lengthInBits += m.LifeSafetyOperations.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOperationExpected) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOperationExpectedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOperationExpected, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationExpected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationExpected")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lifeSafetyOperations)
	if pullErr := readBuffer.PullContext("lifeSafetyOperations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lifeSafetyOperations")
	}
	_lifeSafetyOperations, _lifeSafetyOperationsErr := BACnetLifeSafetyOperationTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _lifeSafetyOperationsErr != nil {
		return nil, errors.Wrap(_lifeSafetyOperationsErr, "Error parsing 'lifeSafetyOperations' field")
	}
	lifeSafetyOperations := _lifeSafetyOperations.(BACnetLifeSafetyOperationTagged)
	if closeErr := readBuffer.CloseContext("lifeSafetyOperations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lifeSafetyOperations")
	}

	// Virtual field
	_actualValue := lifeSafetyOperations
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationExpected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationExpected")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOperationExpected{
		LifeSafetyOperations:   lifeSafetyOperations,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOperationExpected) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationExpected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationExpected")
		}

		// Simple Field (lifeSafetyOperations)
		if pushErr := writeBuffer.PushContext("lifeSafetyOperations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lifeSafetyOperations")
		}
		_lifeSafetyOperationsErr := writeBuffer.WriteSerializable(m.GetLifeSafetyOperations())
		if popErr := writeBuffer.PopContext("lifeSafetyOperations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lifeSafetyOperations")
		}
		if _lifeSafetyOperationsErr != nil {
			return errors.Wrap(_lifeSafetyOperationsErr, "Error serializing 'lifeSafetyOperations' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationExpected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationExpected")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOperationExpected) isBACnetConstructedDataOperationExpected() bool {
	return true
}

func (m *_BACnetConstructedDataOperationExpected) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
