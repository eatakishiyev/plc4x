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

// BACnetConfirmedServiceRequestDeleteObject is the corresponding interface of BACnetConfirmedServiceRequestDeleteObject
type BACnetConfirmedServiceRequestDeleteObject interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
}

// BACnetConfirmedServiceRequestDeleteObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestDeleteObject.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestDeleteObjectExactly interface {
	isBACnetConfirmedServiceRequestDeleteObject() bool
}

// _BACnetConfirmedServiceRequestDeleteObject is the data-structure of this message
type _BACnetConfirmedServiceRequestDeleteObject struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier BACnetApplicationTagObjectIdentifier

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_DELETE_OBJECT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestDeleteObject factory function for _BACnetConfirmedServiceRequestDeleteObject
func NewBACnetConfirmedServiceRequestDeleteObject(objectIdentifier BACnetApplicationTagObjectIdentifier, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestDeleteObject {
	_result := &_BACnetConfirmedServiceRequestDeleteObject{
		ObjectIdentifier:               objectIdentifier,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestDeleteObject(structType interface{}) BACnetConfirmedServiceRequestDeleteObject {
	if casted, ok := structType.(BACnetConfirmedServiceRequestDeleteObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestDeleteObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetTypeName() string {
	return "BACnetConfirmedServiceRequestDeleteObject"
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestDeleteObjectParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestDeleteObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestDeleteObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestDeleteObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := _objectIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestDeleteObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestDeleteObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestDeleteObject{
		ObjectIdentifier:               objectIdentifier,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestDeleteObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestDeleteObject")
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestDeleteObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestDeleteObject")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) isBACnetConfirmedServiceRequestDeleteObject() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestDeleteObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
