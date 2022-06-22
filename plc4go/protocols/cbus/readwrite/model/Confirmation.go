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

// Confirmation is the corresponding interface of Confirmation
type Confirmation interface {
	utils.LengthAware
	utils.Serializable
	// GetConfirmationType returns ConfirmationType (discriminator field)
	GetConfirmationType() byte
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
}

// ConfirmationExactly can be used when we want exactly this type and not a type which fulfills Confirmation.
// This is useful for switch cases.
type ConfirmationExactly interface {
	isConfirmation() bool
}

// _Confirmation is the data-structure of this message
type _Confirmation struct {
	_ConfirmationChildRequirements
	Alpha Alpha
}

type _ConfirmationChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetConfirmationType() byte
}

type ConfirmationParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child Confirmation, serializeChildFunction func() error) error
	GetTypeName() string
}

type ConfirmationChild interface {
	utils.Serializable
	InitializeParent(parent Confirmation, alpha Alpha)
	GetParent() *Confirmation

	GetTypeName() string
	Confirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Confirmation) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewConfirmation factory function for _Confirmation
func NewConfirmation(alpha Alpha) *_Confirmation {
	return &_Confirmation{Alpha: alpha}
}

// Deprecated: use the interface for direct cast
func CastConfirmation(structType interface{}) Confirmation {
	if casted, ok := structType.(Confirmation); ok {
		return casted
	}
	if casted, ok := structType.(*Confirmation); ok {
		return *casted
	}
	return nil
}

func (m *_Confirmation) GetTypeName() string {
	return "Confirmation"
}

func (m *_Confirmation) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (alpha)
	lengthInBits += m.Alpha.GetLengthInBits()
	// Discriminator Field (confirmationType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_Confirmation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConfirmationParse(readBuffer utils.ReadBuffer) (Confirmation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Confirmation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Confirmation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alpha)
	if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alpha")
	}
	_alpha, _alphaErr := AlphaParse(readBuffer)
	if _alphaErr != nil {
		return nil, errors.Wrap(_alphaErr, "Error parsing 'alpha' field")
	}
	alpha := _alpha.(Alpha)
	if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alpha")
	}

	// Discriminator Field (confirmationType) (Used as input to a switch field)
	confirmationType, _confirmationTypeErr := readBuffer.ReadByte("confirmationType")
	if _confirmationTypeErr != nil {
		return nil, errors.Wrap(_confirmationTypeErr, "Error parsing 'confirmationType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ConfirmationChildSerializeRequirement interface {
		Confirmation
		InitializeParent(Confirmation, Alpha)
		GetParent() Confirmation
	}
	var _childTemp interface{}
	var _child ConfirmationChildSerializeRequirement
	var typeSwitchError error
	switch {
	case confirmationType == 0x2E: // ConfirmationSuccessful
		_childTemp, typeSwitchError = ConfirmationSuccessfulParse(readBuffer)
	case confirmationType == 0x23: // NotTransmittedToManyReTransmissions
		_childTemp, typeSwitchError = NotTransmittedToManyReTransmissionsParse(readBuffer)
	case confirmationType == 0x24: // NotTransmittedCorruption
		_childTemp, typeSwitchError = NotTransmittedCorruptionParse(readBuffer)
	case confirmationType == 0x25: // NotTransmittedSyncLoss
		_childTemp, typeSwitchError = NotTransmittedSyncLossParse(readBuffer)
	case confirmationType == 0x27: // NotTransmittedTooLong
		_childTemp, typeSwitchError = NotTransmittedTooLongParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(ConfirmationChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("Confirmation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Confirmation")
	}

	// Finish initializing
	_child.InitializeParent(_child, alpha)
	return _child, nil
}

func (pm *_Confirmation) SerializeParent(writeBuffer utils.WriteBuffer, child Confirmation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Confirmation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Confirmation")
	}

	// Simple Field (alpha)
	if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for alpha")
	}
	_alphaErr := writeBuffer.WriteSerializable(m.GetAlpha())
	if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for alpha")
	}
	if _alphaErr != nil {
		return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
	}

	// Discriminator Field (confirmationType) (Used as input to a switch field)
	confirmationType := byte(child.GetConfirmationType())
	_confirmationTypeErr := writeBuffer.WriteByte("confirmationType", (confirmationType))

	if _confirmationTypeErr != nil {
		return errors.Wrap(_confirmationTypeErr, "Error serializing 'confirmationType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Confirmation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Confirmation")
	}
	return nil
}

func (m *_Confirmation) isConfirmation() bool {
	return true
}

func (m *_Confirmation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
