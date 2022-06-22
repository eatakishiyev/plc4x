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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestCreateObject is the corresponding interface of BACnetConfirmedServiceRequestCreateObject
type BACnetConfirmedServiceRequestCreateObject interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectSpecifier returns ObjectSpecifier (property field)
	GetObjectSpecifier() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
}

// BACnetConfirmedServiceRequestCreateObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestCreateObject.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestCreateObjectExactly interface {
	isBACnetConfirmedServiceRequestCreateObject() bool
}

// _BACnetConfirmedServiceRequestCreateObject is the data-structure of this message
type _BACnetConfirmedServiceRequestCreateObject struct {
	*_BACnetConfirmedServiceRequest
	ObjectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
	ListOfValues    BACnetPropertyValues

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CREATE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestCreateObject) GetObjectSpecifier() BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	return m.ObjectSpecifier
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestCreateObject factory function for _BACnetConfirmedServiceRequestCreateObject
func NewBACnetConfirmedServiceRequestCreateObject(objectSpecifier BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, listOfValues BACnetPropertyValues, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestCreateObject {
	_result := &_BACnetConfirmedServiceRequestCreateObject{
		ObjectSpecifier:                objectSpecifier,
		ListOfValues:                   listOfValues,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestCreateObject(structType interface{}) BACnetConfirmedServiceRequestCreateObject {
	if casted, ok := structType.(BACnetConfirmedServiceRequestCreateObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestCreateObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetTypeName() string {
	return "BACnetConfirmedServiceRequestCreateObject"
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectSpecifier)
	lengthInBits += m.ObjectSpecifier.GetLengthInBits()

	// Optional Field (listOfValues)
	if m.ListOfValues != nil {
		lengthInBits += m.ListOfValues.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestCreateObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestCreateObjectParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestCreateObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestCreateObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestCreateObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectSpecifier)
	if pullErr := readBuffer.PullContext("objectSpecifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectSpecifier")
	}
	_objectSpecifier, _objectSpecifierErr := BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParse(readBuffer, uint8(uint8(0)))
	if _objectSpecifierErr != nil {
		return nil, errors.Wrap(_objectSpecifierErr, "Error parsing 'objectSpecifier' field")
	}
	objectSpecifier := _objectSpecifier.(BACnetConfirmedServiceRequestCreateObjectObjectSpecifier)
	if closeErr := readBuffer.CloseContext("objectSpecifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectSpecifier")
	}

	// Optional Field (listOfValues) (Can be skipped, if a given expression evaluates to false)
	var listOfValues BACnetPropertyValues = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
		}
		_val, _err := BACnetPropertyValuesParse(readBuffer, uint8(1), CastBACnetObjectType(utils.InlineIf(objectSpecifier.GetIsObjectType(), func() interface{} { return CastBACnetObjectType(objectSpecifier.GetObjectType()) }, func() interface{} { return CastBACnetObjectType(objectSpecifier.GetObjectIdentifier().GetObjectType()) })))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'listOfValues' field")
		default:
			listOfValues = _val.(BACnetPropertyValues)
			if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestCreateObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestCreateObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestCreateObject{
		ObjectSpecifier:                objectSpecifier,
		ListOfValues:                   listOfValues,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestCreateObject) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestCreateObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestCreateObject")
		}

		// Simple Field (objectSpecifier)
		if pushErr := writeBuffer.PushContext("objectSpecifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectSpecifier")
		}
		_objectSpecifierErr := writeBuffer.WriteSerializable(m.GetObjectSpecifier())
		if popErr := writeBuffer.PopContext("objectSpecifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectSpecifier")
		}
		if _objectSpecifierErr != nil {
			return errors.Wrap(_objectSpecifierErr, "Error serializing 'objectSpecifier' field")
		}

		// Optional Field (listOfValues) (Can be skipped, if the value is null)
		var listOfValues BACnetPropertyValues = nil
		if m.GetListOfValues() != nil {
			if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for listOfValues")
			}
			listOfValues = m.GetListOfValues()
			_listOfValuesErr := writeBuffer.WriteSerializable(listOfValues)
			if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for listOfValues")
			}
			if _listOfValuesErr != nil {
				return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestCreateObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestCreateObject")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestCreateObject) isBACnetConfirmedServiceRequestCreateObject() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestCreateObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
