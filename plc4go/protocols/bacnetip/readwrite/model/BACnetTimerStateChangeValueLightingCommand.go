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

// BACnetTimerStateChangeValueLightingCommand is the corresponding interface of BACnetTimerStateChangeValueLightingCommand
type BACnetTimerStateChangeValueLightingCommand interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
}

// BACnetTimerStateChangeValueLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueLightingCommand.
// This is useful for switch cases.
type BACnetTimerStateChangeValueLightingCommandExactly interface {
	BACnetTimerStateChangeValueLightingCommand
	isBACnetTimerStateChangeValueLightingCommand() bool
}

// _BACnetTimerStateChangeValueLightingCommand is the data-structure of this message
type _BACnetTimerStateChangeValueLightingCommand struct {
	*_BACnetTimerStateChangeValue
	LigthingCommandValue BACnetLightingCommandEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) InitializeParent(parent BACnetTimerStateChangeValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueLightingCommand factory function for _BACnetTimerStateChangeValueLightingCommand
func NewBACnetTimerStateChangeValueLightingCommand(ligthingCommandValue BACnetLightingCommandEnclosed, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueLightingCommand {
	_result := &_BACnetTimerStateChangeValueLightingCommand{
		LigthingCommandValue:         ligthingCommandValue,
		_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueLightingCommand(structType interface{}) BACnetTimerStateChangeValueLightingCommand {
	if casted, ok := structType.(BACnetTimerStateChangeValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetTypeName() string {
	return "BACnetTimerStateChangeValueLightingCommand"
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueLightingCommandParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueLightingCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ligthingCommandValue)
	if pullErr := readBuffer.PullContext("ligthingCommandValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ligthingCommandValue")
	}
	_ligthingCommandValue, _ligthingCommandValueErr := BACnetLightingCommandEnclosedParse(readBuffer, uint8(uint8(3)))
	if _ligthingCommandValueErr != nil {
		return nil, errors.Wrap(_ligthingCommandValueErr, "Error parsing 'ligthingCommandValue' field")
	}
	ligthingCommandValue := _ligthingCommandValue.(BACnetLightingCommandEnclosed)
	if closeErr := readBuffer.CloseContext("ligthingCommandValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ligthingCommandValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueLightingCommand")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueLightingCommand{
		LigthingCommandValue: ligthingCommandValue,
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueLightingCommand")
		}

		// Simple Field (ligthingCommandValue)
		if pushErr := writeBuffer.PushContext("ligthingCommandValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ligthingCommandValue")
		}
		_ligthingCommandValueErr := writeBuffer.WriteSerializable(m.GetLigthingCommandValue())
		if popErr := writeBuffer.PopContext("ligthingCommandValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ligthingCommandValue")
		}
		if _ligthingCommandValueErr != nil {
			return errors.Wrap(_ligthingCommandValueErr, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueLightingCommand")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueLightingCommand) isBACnetTimerStateChangeValueLightingCommand() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
