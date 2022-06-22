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

// BACnetConstructedDataGroupMode is the corresponding interface of BACnetConstructedDataGroupMode
type BACnetConstructedDataGroupMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetGroupMode returns GroupMode (property field)
	GetGroupMode() BACnetLiftGroupModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftGroupModeTagged
}

// BACnetConstructedDataGroupModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGroupMode.
// This is useful for switch cases.
type BACnetConstructedDataGroupModeExactly interface {
	isBACnetConstructedDataGroupMode() bool
}

// _BACnetConstructedDataGroupMode is the data-structure of this message
type _BACnetConstructedDataGroupMode struct {
	*_BACnetConstructedData
	GroupMode BACnetLiftGroupModeTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupMode) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataGroupMode) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetGroupMode() BACnetLiftGroupModeTagged {
	return m.GroupMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetActualValue() BACnetLiftGroupModeTagged {
	return CastBACnetLiftGroupModeTagged(m.GetGroupMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGroupMode factory function for _BACnetConstructedDataGroupMode
func NewBACnetConstructedDataGroupMode(groupMode BACnetLiftGroupModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupMode {
	_result := &_BACnetConstructedDataGroupMode{
		GroupMode:              groupMode,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupMode(structType interface{}) BACnetConstructedDataGroupMode {
	if casted, ok := structType.(BACnetConstructedDataGroupMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupMode) GetTypeName() string {
	return "BACnetConstructedDataGroupMode"
}

func (m *_BACnetConstructedDataGroupMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataGroupMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (groupMode)
	lengthInBits += m.GroupMode.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGroupModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGroupMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupMode)
	if pullErr := readBuffer.PullContext("groupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupMode")
	}
	_groupMode, _groupModeErr := BACnetLiftGroupModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _groupModeErr != nil {
		return nil, errors.Wrap(_groupModeErr, "Error parsing 'groupMode' field")
	}
	groupMode := _groupMode.(BACnetLiftGroupModeTagged)
	if closeErr := readBuffer.CloseContext("groupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupMode")
	}

	// Virtual field
	_actualValue := groupMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataGroupMode{
		GroupMode:              groupMode,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataGroupMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupMode")
		}

		// Simple Field (groupMode)
		if pushErr := writeBuffer.PushContext("groupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for groupMode")
		}
		_groupModeErr := writeBuffer.WriteSerializable(m.GetGroupMode())
		if popErr := writeBuffer.PopContext("groupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for groupMode")
		}
		if _groupModeErr != nil {
			return errors.Wrap(_groupModeErr, "Error serializing 'groupMode' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupMode) isBACnetConstructedDataGroupMode() bool {
	return true
}

func (m *_BACnetConstructedDataGroupMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
