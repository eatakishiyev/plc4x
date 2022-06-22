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

// BACnetConstructedDataExitPoints is the corresponding interface of BACnetConstructedDataExitPoints
type BACnetConstructedDataExitPoints interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetExitPoints returns ExitPoints (property field)
	GetExitPoints() []BACnetDeviceObjectReference
}

// BACnetConstructedDataExitPointsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataExitPoints.
// This is useful for switch cases.
type BACnetConstructedDataExitPointsExactly interface {
	isBACnetConstructedDataExitPoints() bool
}

// _BACnetConstructedDataExitPoints is the data-structure of this message
type _BACnetConstructedDataExitPoints struct {
	*_BACnetConstructedData
	ExitPoints []BACnetDeviceObjectReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExitPoints) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExitPoints) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXIT_POINTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExitPoints) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataExitPoints) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExitPoints) GetExitPoints() []BACnetDeviceObjectReference {
	return m.ExitPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExitPoints factory function for _BACnetConstructedDataExitPoints
func NewBACnetConstructedDataExitPoints(exitPoints []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExitPoints {
	_result := &_BACnetConstructedDataExitPoints{
		ExitPoints:             exitPoints,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExitPoints(structType interface{}) BACnetConstructedDataExitPoints {
	if casted, ok := structType.(BACnetConstructedDataExitPoints); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExitPoints); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExitPoints) GetTypeName() string {
	return "BACnetConstructedDataExitPoints"
}

func (m *_BACnetConstructedDataExitPoints) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataExitPoints) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ExitPoints) > 0 {
		for _, element := range m.ExitPoints {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataExitPoints) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataExitPointsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataExitPoints, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExitPoints"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExitPoints")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (exitPoints)
	if pullErr := readBuffer.PullContext("exitPoints", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exitPoints")
	}
	// Terminated array
	exitPoints := make([]BACnetDeviceObjectReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'exitPoints' field")
			}
			exitPoints = append(exitPoints, _item.(BACnetDeviceObjectReference))

		}
	}
	if closeErr := readBuffer.CloseContext("exitPoints", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exitPoints")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExitPoints"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExitPoints")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataExitPoints{
		ExitPoints:             exitPoints,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataExitPoints) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExitPoints"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExitPoints")
		}

		// Array Field (exitPoints)
		if m.GetExitPoints() != nil {
			if pushErr := writeBuffer.PushContext("exitPoints", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for exitPoints")
			}
			for _, _element := range m.GetExitPoints() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'exitPoints' field")
				}
			}
			if popErr := writeBuffer.PopContext("exitPoints", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for exitPoints")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExitPoints"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExitPoints")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExitPoints) isBACnetConstructedDataExitPoints() bool {
	return true
}

func (m *_BACnetConstructedDataExitPoints) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
