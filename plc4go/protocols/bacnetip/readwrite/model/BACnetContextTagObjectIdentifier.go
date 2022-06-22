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

// BACnetContextTagObjectIdentifier is the corresponding interface of BACnetContextTagObjectIdentifier
type BACnetContextTagObjectIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadObjectIdentifier
	// GetObjectType returns ObjectType (virtual field)
	GetObjectType() BACnetObjectType
	// GetInstanceNumber returns InstanceNumber (virtual field)
	GetInstanceNumber() uint32
}

// BACnetContextTagObjectIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagObjectIdentifier.
// This is useful for switch cases.
type BACnetContextTagObjectIdentifierExactly interface {
	BACnetContextTagObjectIdentifier
	isBACnetContextTagObjectIdentifier() bool
}

// _BACnetContextTagObjectIdentifier is the data-structure of this message
type _BACnetContextTagObjectIdentifier struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagObjectIdentifier) GetDataType() BACnetDataType {
	return BACnetDataType_BACNET_OBJECT_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagObjectIdentifier) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagObjectIdentifier) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagObjectIdentifier) GetPayload() BACnetTagPayloadObjectIdentifier {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagObjectIdentifier) GetObjectType() BACnetObjectType {
	return CastBACnetObjectType(m.GetPayload().GetObjectType())
}

func (m *_BACnetContextTagObjectIdentifier) GetInstanceNumber() uint32 {
	return uint32(m.GetPayload().GetInstanceNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagObjectIdentifier factory function for _BACnetContextTagObjectIdentifier
func NewBACnetContextTagObjectIdentifier(payload BACnetTagPayloadObjectIdentifier, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagObjectIdentifier {
	_result := &_BACnetContextTagObjectIdentifier{
		Payload:           payload,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagObjectIdentifier(structType interface{}) BACnetContextTagObjectIdentifier {
	if casted, ok := structType.(BACnetContextTagObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagObjectIdentifier) GetTypeName() string {
	return "BACnetContextTagObjectIdentifier"
}

func (m *_BACnetContextTagObjectIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagObjectIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagObjectIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagObjectIdentifierParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagObjectIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadObjectIdentifierParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadObjectIdentifier)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_objectType := payload.GetObjectType()
	objectType := BACnetObjectType(_objectType)
	_ = objectType

	// Virtual field
	_instanceNumber := payload.GetInstanceNumber()
	instanceNumber := uint32(_instanceNumber)
	_ = instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetContextTagObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagObjectIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagObjectIdentifier{
		Payload: payload,
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagObjectIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagObjectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagObjectIdentifier")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _objectTypeErr := writeBuffer.WriteVirtual("objectType", m.GetObjectType()); _objectTypeErr != nil {
			return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
		}
		// Virtual field
		if _instanceNumberErr := writeBuffer.WriteVirtual("instanceNumber", m.GetInstanceNumber()); _instanceNumberErr != nil {
			return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagObjectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagObjectIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagObjectIdentifier) isBACnetContextTagObjectIdentifier() bool {
	return true
}

func (m *_BACnetContextTagObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
