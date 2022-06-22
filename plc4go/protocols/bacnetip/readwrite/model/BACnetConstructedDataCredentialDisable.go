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

// BACnetConstructedDataCredentialDisable is the corresponding interface of BACnetConstructedDataCredentialDisable
type BACnetConstructedDataCredentialDisable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCredentialDisable returns CredentialDisable (property field)
	GetCredentialDisable() BACnetAccessCredentialDisableTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessCredentialDisableTagged
}

// BACnetConstructedDataCredentialDisableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCredentialDisable.
// This is useful for switch cases.
type BACnetConstructedDataCredentialDisableExactly interface {
	isBACnetConstructedDataCredentialDisable() bool
}

// _BACnetConstructedDataCredentialDisable is the data-structure of this message
type _BACnetConstructedDataCredentialDisable struct {
	*_BACnetConstructedData
	CredentialDisable BACnetAccessCredentialDisableTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCredentialDisable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CREDENTIAL_DISABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCredentialDisable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCredentialDisable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetCredentialDisable() BACnetAccessCredentialDisableTagged {
	return m.CredentialDisable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCredentialDisable) GetActualValue() BACnetAccessCredentialDisableTagged {
	return CastBACnetAccessCredentialDisableTagged(m.GetCredentialDisable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCredentialDisable factory function for _BACnetConstructedDataCredentialDisable
func NewBACnetConstructedDataCredentialDisable(credentialDisable BACnetAccessCredentialDisableTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCredentialDisable {
	_result := &_BACnetConstructedDataCredentialDisable{
		CredentialDisable:      credentialDisable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCredentialDisable(structType interface{}) BACnetConstructedDataCredentialDisable {
	if casted, ok := structType.(BACnetConstructedDataCredentialDisable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCredentialDisable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCredentialDisable) GetTypeName() string {
	return "BACnetConstructedDataCredentialDisable"
}

func (m *_BACnetConstructedDataCredentialDisable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCredentialDisable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (credentialDisable)
	lengthInBits += m.CredentialDisable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCredentialDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCredentialDisableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCredentialDisable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCredentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCredentialDisable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (credentialDisable)
	if pullErr := readBuffer.PullContext("credentialDisable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for credentialDisable")
	}
	_credentialDisable, _credentialDisableErr := BACnetAccessCredentialDisableTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _credentialDisableErr != nil {
		return nil, errors.Wrap(_credentialDisableErr, "Error parsing 'credentialDisable' field")
	}
	credentialDisable := _credentialDisable.(BACnetAccessCredentialDisableTagged)
	if closeErr := readBuffer.CloseContext("credentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for credentialDisable")
	}

	// Virtual field
	_actualValue := credentialDisable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCredentialDisable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCredentialDisable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCredentialDisable{
		CredentialDisable:      credentialDisable,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCredentialDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCredentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCredentialDisable")
		}

		// Simple Field (credentialDisable)
		if pushErr := writeBuffer.PushContext("credentialDisable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for credentialDisable")
		}
		_credentialDisableErr := writeBuffer.WriteSerializable(m.GetCredentialDisable())
		if popErr := writeBuffer.PopContext("credentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for credentialDisable")
		}
		if _credentialDisableErr != nil {
			return errors.Wrap(_credentialDisableErr, "Error serializing 'credentialDisable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCredentialDisable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCredentialDisable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCredentialDisable) isBACnetConstructedDataCredentialDisable() bool {
	return true
}

func (m *_BACnetConstructedDataCredentialDisable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
