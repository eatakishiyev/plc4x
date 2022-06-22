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

// BACnetConstructedDataVTClassesSupported is the corresponding interface of BACnetConstructedDataVTClassesSupported
type BACnetConstructedDataVTClassesSupported interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVtClassesSupported returns VtClassesSupported (property field)
	GetVtClassesSupported() []BACnetVTClassTagged
}

// BACnetConstructedDataVTClassesSupportedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVTClassesSupported.
// This is useful for switch cases.
type BACnetConstructedDataVTClassesSupportedExactly interface {
	isBACnetConstructedDataVTClassesSupported() bool
}

// _BACnetConstructedDataVTClassesSupported is the data-structure of this message
type _BACnetConstructedDataVTClassesSupported struct {
	*_BACnetConstructedData
	VtClassesSupported []BACnetVTClassTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataVTClassesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VT_CLASSES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVTClassesSupported) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVTClassesSupported) GetVtClassesSupported() []BACnetVTClassTagged {
	return m.VtClassesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVTClassesSupported factory function for _BACnetConstructedDataVTClassesSupported
func NewBACnetConstructedDataVTClassesSupported(vtClassesSupported []BACnetVTClassTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataVTClassesSupported {
	_result := &_BACnetConstructedDataVTClassesSupported{
		VtClassesSupported:     vtClassesSupported,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVTClassesSupported(structType interface{}) BACnetConstructedDataVTClassesSupported {
	if casted, ok := structType.(BACnetConstructedDataVTClassesSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVTClassesSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVTClassesSupported) GetTypeName() string {
	return "BACnetConstructedDataVTClassesSupported"
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.VtClassesSupported) > 0 {
		for _, element := range m.VtClassesSupported {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataVTClassesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVTClassesSupportedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVTClassesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVTClassesSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVTClassesSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (vtClassesSupported)
	if pullErr := readBuffer.PullContext("vtClassesSupported", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vtClassesSupported")
	}
	// Terminated array
	vtClassesSupported := make([]BACnetVTClassTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetVTClassTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'vtClassesSupported' field")
			}
			vtClassesSupported = append(vtClassesSupported, _item.(BACnetVTClassTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("vtClassesSupported", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vtClassesSupported")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVTClassesSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVTClassesSupported")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVTClassesSupported{
		VtClassesSupported:     vtClassesSupported,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVTClassesSupported) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVTClassesSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVTClassesSupported")
		}

		// Array Field (vtClassesSupported)
		if m.GetVtClassesSupported() != nil {
			if pushErr := writeBuffer.PushContext("vtClassesSupported", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for vtClassesSupported")
			}
			for _, _element := range m.GetVtClassesSupported() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'vtClassesSupported' field")
				}
			}
			if popErr := writeBuffer.PopContext("vtClassesSupported", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for vtClassesSupported")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVTClassesSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVTClassesSupported")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataVTClassesSupported) isBACnetConstructedDataVTClassesSupported() bool {
	return true
}

func (m *_BACnetConstructedDataVTClassesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
