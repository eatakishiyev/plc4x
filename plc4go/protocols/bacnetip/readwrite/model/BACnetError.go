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

// BACnetError is the corresponding interface of BACnetError
type BACnetError interface {
	utils.LengthAware
	utils.Serializable
	// GetErrorChoice returns ErrorChoice (discriminator field)
	GetErrorChoice() BACnetConfirmedServiceChoice
}

// BACnetErrorExactly can be used when we want exactly this type and not a type which fulfills BACnetError.
// This is useful for switch cases.
type BACnetErrorExactly interface {
	isBACnetError() bool
}

// _BACnetError is the data-structure of this message
type _BACnetError struct {
	_BACnetErrorChildRequirements
}

type _BACnetErrorChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetErrorChoice() BACnetConfirmedServiceChoice
}

type BACnetErrorParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetError, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetErrorChild interface {
	utils.Serializable
	InitializeParent(parent BACnetError)
	GetParent() *BACnetError

	GetTypeName() string
	BACnetError
}

// NewBACnetError factory function for _BACnetError
func NewBACnetError() *_BACnetError {
	return &_BACnetError{}
}

// Deprecated: use the interface for direct cast
func CastBACnetError(structType interface{}) BACnetError {
	if casted, ok := structType.(BACnetError); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetError); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetError) GetTypeName() string {
	return "BACnetError"
}

func (m *_BACnetError) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_BACnetError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (BACnetError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetErrorChildSerializeRequirement interface {
		BACnetError
		InitializeParent(BACnetError)
		GetParent() BACnetError
	}
	var _childTemp interface{}
	var _child BACnetErrorChildSerializeRequirement
	var typeSwitchError error
	switch {
	case errorChoice == BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE: // SubscribeCOVPropertyMultipleError
		_childTemp, typeSwitchError = SubscribeCOVPropertyMultipleErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT: // ChangeListAddError
		_childTemp, typeSwitchError = ChangeListAddErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT: // ChangeListRemoveError
		_childTemp, typeSwitchError = ChangeListRemoveErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_CREATE_OBJECT: // CreateObjectError
		_childTemp, typeSwitchError = CreateObjectErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_WRITE_PROPERTY_MULTIPLE: // WritePropertyMultipleError
		_childTemp, typeSwitchError = WritePropertyMultipleErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_CONFIRMED_PRIVATE_TRANSFER: // ConfirmedPrivateTransferError
		_childTemp, typeSwitchError = ConfirmedPrivateTransferErrorParse(readBuffer, errorChoice)
	case errorChoice == BACnetConfirmedServiceChoice_VT_CLOSE: // VTCloseError
		_childTemp, typeSwitchError = VTCloseErrorParse(readBuffer, errorChoice)
	case true: // BACnetErrorGeneral
		_childTemp, typeSwitchError = BACnetErrorGeneralParse(readBuffer, errorChoice)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(BACnetErrorChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetError")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_BACnetError) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetError, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetError"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetError")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetError"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetError")
	}
	return nil
}

func (m *_BACnetError) isBACnetError() bool {
	return true
}

func (m *_BACnetError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
