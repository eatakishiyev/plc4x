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

// BACnetPropertyAccessResultAccessResult is the corresponding interface of BACnetPropertyAccessResultAccessResult
type BACnetPropertyAccessResultAccessResult interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyAccessResultAccessResultExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyAccessResultAccessResult.
// This is useful for switch cases.
type BACnetPropertyAccessResultAccessResultExactly interface {
	isBACnetPropertyAccessResultAccessResult() bool
}

// _BACnetPropertyAccessResultAccessResult is the data-structure of this message
type _BACnetPropertyAccessResultAccessResult struct {
	_BACnetPropertyAccessResultAccessResultChildRequirements
	PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument         BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	PropertyArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

type _BACnetPropertyAccessResultAccessResultChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetPropertyAccessResultAccessResultParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetPropertyAccessResultAccessResult, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetPropertyAccessResultAccessResultChild interface {
	utils.Serializable
	InitializeParent(parent BACnetPropertyAccessResultAccessResult, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetPropertyAccessResultAccessResult

	GetTypeName() string
	BACnetPropertyAccessResultAccessResult
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResult) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResult) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResult factory function for _BACnetPropertyAccessResultAccessResult
func NewBACnetPropertyAccessResultAccessResult(peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResult {
	return &_BACnetPropertyAccessResultAccessResult{PeekedTagHeader: peekedTagHeader, ObjectTypeArgument: objectTypeArgument, PropertyIdentifierArgument: propertyIdentifierArgument, PropertyArrayIndexArgument: propertyArrayIndexArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResult(structType interface{}) BACnetPropertyAccessResultAccessResult {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResult) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResult"
}

func (m *_BACnetPropertyAccessResultAccessResult) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyAccessResultAccessResultParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetPropertyAccessResultAccessResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetPropertyAccessResultAccessResultChildSerializeRequirement interface {
		BACnetPropertyAccessResultAccessResult
		InitializeParent(BACnetPropertyAccessResultAccessResult, BACnetTagHeader)
		GetParent() BACnetPropertyAccessResultAccessResult
	}
	var _childTemp interface{}
	var _child BACnetPropertyAccessResultAccessResultChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(4): // BACnetPropertyAccessResultAccessResultPropertyValue
		_childTemp, typeSwitchError = BACnetPropertyAccessResultAccessResultPropertyValueParse(readBuffer, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
	case peekedTagNumber == uint8(5): // BACnetPropertyAccessResultAccessResultPropertyAccessError
		_childTemp, typeSwitchError = BACnetPropertyAccessResultAccessResultPropertyAccessErrorParse(readBuffer, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(BACnetPropertyAccessResultAccessResultChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResult")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetPropertyAccessResultAccessResult) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetPropertyAccessResultAccessResult, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResult")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResult")
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResult) isBACnetPropertyAccessResultAccessResult() bool {
	return true
}

func (m *_BACnetPropertyAccessResultAccessResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
