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

// BACnetEventParameterChangeOfLifeSavety is the corresponding interface of BACnetEventParameterChangeOfLifeSavety
type BACnetEventParameterChangeOfLifeSavety interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfLifeSavetyExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfLifeSavety.
// This is useful for switch cases.
type BACnetEventParameterChangeOfLifeSavetyExactly interface {
	isBACnetEventParameterChangeOfLifeSavety() bool
}

// _BACnetEventParameterChangeOfLifeSavety is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavety struct {
	*_BACnetEventParameter
	OpeningTag                  BACnetOpeningTag
	TimeDelay                   BACnetContextTagUnsignedInteger
	ListOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	ListOfAlarmValues           BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	ModePropertyReference       BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag                  BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavety) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavety) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavety factory function for _BACnetEventParameterChangeOfLifeSavety
func NewBACnetEventParameterChangeOfLifeSavety(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterChangeOfLifeSavety {
	_result := &_BACnetEventParameterChangeOfLifeSavety{
		OpeningTag:                  openingTag,
		TimeDelay:                   timeDelay,
		ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues,
		ListOfAlarmValues:           listOfAlarmValues,
		ModePropertyReference:       modePropertyReference,
		ClosingTag:                  closingTag,
		_BACnetEventParameter:       NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavety(structType interface{}) BACnetEventParameterChangeOfLifeSavety {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavety"
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (listOfLifeSavetyAlarmValues)
	lengthInBits += m.ListOfLifeSavetyAlarmValues.GetLengthInBits()

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits()

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfLifeSavetyParse(readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfLifeSavety, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(8)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (listOfLifeSavetyAlarmValues)
	if pullErr := readBuffer.PullContext("listOfLifeSavetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfLifeSavetyAlarmValues")
	}
	_listOfLifeSavetyAlarmValues, _listOfLifeSavetyAlarmValuesErr := BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParse(readBuffer, uint8(uint8(1)))
	if _listOfLifeSavetyAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfLifeSavetyAlarmValuesErr, "Error parsing 'listOfLifeSavetyAlarmValues' field")
	}
	listOfLifeSavetyAlarmValues := _listOfLifeSavetyAlarmValues.(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfLifeSavetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfLifeSavetyAlarmValues")
	}

	// Simple Field (listOfAlarmValues)
	if pullErr := readBuffer.PullContext("listOfAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfAlarmValues")
	}
	_listOfAlarmValues, _listOfAlarmValuesErr := BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParse(readBuffer, uint8(uint8(2)))
	if _listOfAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfAlarmValuesErr, "Error parsing 'listOfAlarmValues' field")
	}
	listOfAlarmValues := _listOfAlarmValues.(BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfAlarmValues")
	}

	// Simple Field (modePropertyReference)
	if pullErr := readBuffer.PullContext("modePropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for modePropertyReference")
	}
	_modePropertyReference, _modePropertyReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(4)))
	if _modePropertyReferenceErr != nil {
		return nil, errors.Wrap(_modePropertyReferenceErr, "Error parsing 'modePropertyReference' field")
	}
	modePropertyReference := _modePropertyReference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("modePropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for modePropertyReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(8)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavety")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfLifeSavety{
		OpeningTag:                  openingTag,
		TimeDelay:                   timeDelay,
		ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues,
		ListOfAlarmValues:           listOfAlarmValues,
		ModePropertyReference:       modePropertyReference,
		ClosingTag:                  closingTag,
		_BACnetEventParameter:       &_BACnetEventParameter{},
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavety")
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeDelay")
		}
		_timeDelayErr := writeBuffer.WriteSerializable(m.GetTimeDelay())
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeDelay")
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (listOfLifeSavetyAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfLifeSavetyAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfLifeSavetyAlarmValues")
		}
		_listOfLifeSavetyAlarmValuesErr := writeBuffer.WriteSerializable(m.GetListOfLifeSavetyAlarmValues())
		if popErr := writeBuffer.PopContext("listOfLifeSavetyAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfLifeSavetyAlarmValues")
		}
		if _listOfLifeSavetyAlarmValuesErr != nil {
			return errors.Wrap(_listOfLifeSavetyAlarmValuesErr, "Error serializing 'listOfLifeSavetyAlarmValues' field")
		}

		// Simple Field (listOfAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfAlarmValues")
		}
		_listOfAlarmValuesErr := writeBuffer.WriteSerializable(m.GetListOfAlarmValues())
		if popErr := writeBuffer.PopContext("listOfAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfAlarmValues")
		}
		if _listOfAlarmValuesErr != nil {
			return errors.Wrap(_listOfAlarmValuesErr, "Error serializing 'listOfAlarmValues' field")
		}

		// Simple Field (modePropertyReference)
		if pushErr := writeBuffer.PushContext("modePropertyReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for modePropertyReference")
		}
		_modePropertyReferenceErr := writeBuffer.WriteSerializable(m.GetModePropertyReference())
		if popErr := writeBuffer.PopContext("modePropertyReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for modePropertyReference")
		}
		if _modePropertyReferenceErr != nil {
			return errors.Wrap(_modePropertyReferenceErr, "Error serializing 'modePropertyReference' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavety")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfLifeSavety) isBACnetEventParameterChangeOfLifeSavety() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfLifeSavety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
