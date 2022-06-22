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

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the corresponding interface of BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues.
// This is useful for switch cases.
type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesExactly interface {
	isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool
}

// _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues struct {
	OpeningTag                  BACnetOpeningTag
	ListOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged
	ClosingTag                  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetListOfLifeSavetyAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues factory function for _BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
func NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(openingTag BACnetOpeningTag, listOfLifeSavetyAlarmValues []BACnetLifeSafetyStateTagged, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return &_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues{OpeningTag: openingTag, ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(structType interface{}) BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfLifeSavetyAlarmValues) > 0 {
		for _, element := range m.ListOfLifeSavetyAlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfLifeSavetyAlarmValues)
	if pullErr := readBuffer.PullContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfLifeSavetyAlarmValues")
	}
	// Terminated array
	listOfLifeSavetyAlarmValues := make([]BACnetLifeSafetyStateTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfLifeSavetyAlarmValues' field")
			}
			listOfLifeSavetyAlarmValues = append(listOfLifeSavetyAlarmValues, _item.(BACnetLifeSafetyStateTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfLifeSavetyAlarmValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}

	// Create the instance
	return NewBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(openingTag, listOfLifeSavetyAlarmValues, closingTag, tagNumber), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfLifeSavetyAlarmValues)
	if m.GetListOfLifeSavetyAlarmValues() != nil {
		if pushErr := writeBuffer.PushContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfLifeSavetyAlarmValues")
		}
		for _, _element := range m.GetListOfLifeSavetyAlarmValues() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfLifeSavetyAlarmValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfLifeSavetyAlarmValues", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfLifeSavetyAlarmValues")
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues")
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) isBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
