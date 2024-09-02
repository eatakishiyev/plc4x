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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MeasurementData is the corresponding interface of MeasurementData
type MeasurementData interface {
	MeasurementDataContract
	MeasurementDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// MeasurementDataContract provides a set of functions which can be overwritten by a sub struct
type MeasurementDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MeasurementCommandTypeContainer
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MeasurementCommandType
}

// MeasurementDataRequirements provides a set of functions which need to be implemented by a sub struct
type MeasurementDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() MeasurementCommandType
}

// MeasurementDataExactly can be used when we want exactly this type and not a type which fulfills MeasurementData.
// This is useful for switch cases.
type MeasurementDataExactly interface {
	MeasurementData
	isMeasurementData() bool
}

// _MeasurementData is the data-structure of this message
type _MeasurementData struct {
	_SubType             MeasurementData
	CommandTypeContainer MeasurementCommandTypeContainer
}

var _ MeasurementDataContract = (*_MeasurementData)(nil)

type MeasurementDataChild interface {
	utils.Serializable

	GetParent() *MeasurementData

	GetTypeName() string
	MeasurementData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeasurementData) GetCommandTypeContainer() MeasurementCommandTypeContainer {
	return m.CommandTypeContainer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_MeasurementData) GetCommandType() MeasurementCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastMeasurementCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeasurementData factory function for _MeasurementData
func NewMeasurementData(commandTypeContainer MeasurementCommandTypeContainer) *_MeasurementData {
	return &_MeasurementData{CommandTypeContainer: commandTypeContainer}
}

// Deprecated: use the interface for direct cast
func CastMeasurementData(structType any) MeasurementData {
	if casted, ok := structType.(MeasurementData); ok {
		return casted
	}
	if casted, ok := structType.(*MeasurementData); ok {
		return *casted
	}
	return nil
}

func (m *_MeasurementData) GetTypeName() string {
	return "MeasurementData"
}

func (m *_MeasurementData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MeasurementData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func MeasurementDataParse[T MeasurementData](ctx context.Context, theBytes []byte) (T, error) {
	return MeasurementDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeasurementDataParseWithBufferProducer[T MeasurementData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MeasurementDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func MeasurementDataParseWithBuffer[T MeasurementData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewMeasurementData().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_MeasurementData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__measurementData MeasurementData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeasurementData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeasurementData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMeasurementCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[MeasurementCommandTypeContainer](ctx, "commandTypeContainer", "MeasurementCommandTypeContainer", ReadEnum(MeasurementCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[MeasurementCommandType](ctx, "commandType", (*MeasurementCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MeasurementData
	switch {
	case commandType == MeasurementCommandType_MEASUREMENT_EVENT: // MeasurementDataChannelMeasurementData
		if _child, err = (&_MeasurementDataChannelMeasurementData{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeasurementDataChannelMeasurementData for type-switch of MeasurementData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("MeasurementData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeasurementData")
	}

	return _child, nil
}

func (pm *_MeasurementData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MeasurementData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MeasurementData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MeasurementData")
	}

	if err := WriteSimpleEnumField[MeasurementCommandTypeContainer](ctx, "commandTypeContainer", "MeasurementCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[MeasurementCommandTypeContainer, uint8](MeasurementCommandTypeContainer.GetValue, MeasurementCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MeasurementData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MeasurementData")
	}
	return nil
}

func (m *_MeasurementData) isMeasurementData() bool {
	return true
}
