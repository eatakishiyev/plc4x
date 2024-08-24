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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ClockAndTimekeepingData is the corresponding interface of ClockAndTimekeepingData
type ClockAndTimekeepingData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() ClockAndTimekeepingCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() ClockAndTimekeepingCommandType
}

// ClockAndTimekeepingDataExactly can be used when we want exactly this type and not a type which fulfills ClockAndTimekeepingData.
// This is useful for switch cases.
type ClockAndTimekeepingDataExactly interface {
	ClockAndTimekeepingData
	isClockAndTimekeepingData() bool
}

// _ClockAndTimekeepingData is the data-structure of this message
type _ClockAndTimekeepingData struct {
	_ClockAndTimekeepingDataChildRequirements
	CommandTypeContainer ClockAndTimekeepingCommandTypeContainer
	Argument             byte
}

type _ClockAndTimekeepingDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetArgument() byte
	GetCommandType() ClockAndTimekeepingCommandType
}

type ClockAndTimekeepingDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ClockAndTimekeepingData, serializeChildFunction func() error) error
	GetTypeName() string
}

type ClockAndTimekeepingDataChild interface {
	utils.Serializable
	InitializeParent(parent ClockAndTimekeepingData, commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte)
	GetParent() *ClockAndTimekeepingData

	GetTypeName() string
	ClockAndTimekeepingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ClockAndTimekeepingData) GetCommandTypeContainer() ClockAndTimekeepingCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_ClockAndTimekeepingData) GetArgument() byte {
	return m.Argument
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ClockAndTimekeepingData) GetCommandType() ClockAndTimekeepingCommandType {
	ctx := context.Background()
	_ = ctx
	return CastClockAndTimekeepingCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewClockAndTimekeepingData factory function for _ClockAndTimekeepingData
func NewClockAndTimekeepingData(commandTypeContainer ClockAndTimekeepingCommandTypeContainer, argument byte) *_ClockAndTimekeepingData {
	return &_ClockAndTimekeepingData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastClockAndTimekeepingData(structType any) ClockAndTimekeepingData {
	if casted, ok := structType.(ClockAndTimekeepingData); ok {
		return casted
	}
	if casted, ok := structType.(*ClockAndTimekeepingData); ok {
		return *casted
	}
	return nil
}

func (m *_ClockAndTimekeepingData) GetTypeName() string {
	return "ClockAndTimekeepingData"
}

func (m *_ClockAndTimekeepingData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ClockAndTimekeepingData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClockAndTimekeepingDataParse(ctx context.Context, theBytes []byte) (ClockAndTimekeepingData, error) {
	return ClockAndTimekeepingDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClockAndTimekeepingDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ClockAndTimekeepingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ClockAndTimekeepingData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsClockAndTimekeepingCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := ClockAndTimekeepingCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of ClockAndTimekeepingData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := ClockAndTimekeepingCommandType(_commandType)
	_ = commandType

	// Simple Field (argument)
	_argument, _argumentErr := readBuffer.ReadByte("argument")
	if _argumentErr != nil {
		return nil, errors.Wrap(_argumentErr, "Error parsing 'argument' field of ClockAndTimekeepingData")
	}
	argument := _argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ClockAndTimekeepingDataChildSerializeRequirement interface {
		ClockAndTimekeepingData
		InitializeParent(ClockAndTimekeepingData, ClockAndTimekeepingCommandTypeContainer, byte)
		GetParent() ClockAndTimekeepingData
	}
	var _childTemp any
	var _child ClockAndTimekeepingDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE && argument == 0x01: // ClockAndTimekeepingDataUpdateTime
		_childTemp, typeSwitchError = ClockAndTimekeepingDataUpdateTimeParseWithBuffer(ctx, readBuffer)
	case commandType == ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE && argument == 0x02: // ClockAndTimekeepingDataUpdateDate
		_childTemp, typeSwitchError = ClockAndTimekeepingDataUpdateDateParseWithBuffer(ctx, readBuffer)
	case commandType == ClockAndTimekeepingCommandType_REQUEST_REFRESH && argument == 0x03: // ClockAndTimekeepingDataRequestRefresh
		_childTemp, typeSwitchError = ClockAndTimekeepingDataRequestRefreshParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ClockAndTimekeepingData")
	}
	_child = _childTemp.(ClockAndTimekeepingDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ClockAndTimekeepingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ClockAndTimekeepingData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, argument)
	return _child, nil
}

func (pm *_ClockAndTimekeepingData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ClockAndTimekeepingData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ClockAndTimekeepingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ClockAndTimekeepingData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (argument)
	argument := byte(m.GetArgument())
	_argumentErr := writeBuffer.WriteByte("argument", (argument))
	if _argumentErr != nil {
		return errors.Wrap(_argumentErr, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ClockAndTimekeepingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ClockAndTimekeepingData")
	}
	return nil
}

func (m *_ClockAndTimekeepingData) isClockAndTimekeepingData() bool {
	return true
}

func (m *_ClockAndTimekeepingData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
