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

// MeteringData is the corresponding interface of MeteringData
type MeteringData interface {
	MeteringDataContract
	MeteringDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// MeteringDataContract provides a set of functions which can be overwritten by a sub struct
type MeteringDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MeteringCommandTypeContainer
	// GetArgument returns Argument (property field)
	GetArgument() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MeteringCommandType
}

// MeteringDataRequirements provides a set of functions which need to be implemented by a sub struct
type MeteringDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetArgument returns Argument (discriminator field)
	GetArgument() byte
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() MeteringCommandType
}

// MeteringDataExactly can be used when we want exactly this type and not a type which fulfills MeteringData.
// This is useful for switch cases.
type MeteringDataExactly interface {
	MeteringData
	isMeteringData() bool
}

// _MeteringData is the data-structure of this message
type _MeteringData struct {
	_SubType             MeteringData
	CommandTypeContainer MeteringCommandTypeContainer
	Argument             byte
}

var _ MeteringDataContract = (*_MeteringData)(nil)

type MeteringDataChild interface {
	utils.Serializable

	GetParent() *MeteringData

	GetTypeName() string
	MeteringData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringData) GetCommandTypeContainer() MeteringCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_MeteringData) GetArgument() byte {
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

func (pm *_MeteringData) GetCommandType() MeteringCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastMeteringCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMeteringData factory function for _MeteringData
func NewMeteringData(commandTypeContainer MeteringCommandTypeContainer, argument byte) *_MeteringData {
	return &_MeteringData{CommandTypeContainer: commandTypeContainer, Argument: argument}
}

// Deprecated: use the interface for direct cast
func CastMeteringData(structType any) MeteringData {
	if casted, ok := structType.(MeteringData); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringData); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringData) GetTypeName() string {
	return "MeteringData"
}

func (m *_MeteringData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (argument)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MeteringData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func MeteringDataParse[T MeteringData](ctx context.Context, theBytes []byte) (T, error) {
	return MeteringDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataParseWithBufferProducer[T MeteringData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MeteringDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func MeteringDataParseWithBuffer[T MeteringData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewMeteringData().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_MeteringData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__meteringData MeteringData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMeteringCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[MeteringCommandTypeContainer](ctx, "commandTypeContainer", "MeteringCommandTypeContainer", ReadEnum(MeteringCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[MeteringCommandType](ctx, "commandType", (*MeteringCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	argument, err := ReadSimpleField(ctx, "argument", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'argument' field"))
	}
	m.Argument = argument

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MeteringData
	switch {
	case commandType == MeteringCommandType_EVENT && argument == 0x01: // MeteringDataMeasureElectricity
		if _child, err = (&_MeteringDataMeasureElectricity{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataMeasureElectricity for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x02: // MeteringDataMeasureGas
		if _child, err = (&_MeteringDataMeasureGas{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataMeasureGas for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x03: // MeteringDataMeasureDrinkingWater
		if _child, err = (&_MeteringDataMeasureDrinkingWater{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataMeasureDrinkingWater for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x04: // MeteringDataMeasureOtherWater
		if _child, err = (&_MeteringDataMeasureOtherWater{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataMeasureOtherWater for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x05: // MeteringDataMeasureOil
		if _child, err = (&_MeteringDataMeasureOil{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataMeasureOil for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x81: // MeteringDataElectricityConsumption
		if _child, err = (&_MeteringDataElectricityConsumption{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataElectricityConsumption for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x82: // MeteringDataGasConsumption
		if _child, err = (&_MeteringDataGasConsumption{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataGasConsumption for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x83: // MeteringDataDrinkingWaterConsumption
		if _child, err = (&_MeteringDataDrinkingWaterConsumption{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataDrinkingWaterConsumption for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x84: // MeteringDataOtherWaterConsumption
		if _child, err = (&_MeteringDataOtherWaterConsumption{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataOtherWaterConsumption for type-switch of MeteringData")
		}
	case commandType == MeteringCommandType_EVENT && argument == 0x85: // MeteringDataOilConsumption
		if _child, err = (&_MeteringDataOilConsumption{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MeteringDataOilConsumption for type-switch of MeteringData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v, argument=%v]", commandType, argument)
	}

	if closeErr := readBuffer.CloseContext("MeteringData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringData")
	}

	return _child, nil
}

func (pm *_MeteringData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MeteringData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MeteringData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MeteringData")
	}

	if err := WriteSimpleEnumField[MeteringCommandTypeContainer](ctx, "commandTypeContainer", "MeteringCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[MeteringCommandTypeContainer, uint8](MeteringCommandTypeContainer.GetValue, MeteringCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "argument", m.GetArgument(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'argument' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MeteringData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MeteringData")
	}
	return nil
}

func (m *_MeteringData) isMeteringData() bool {
	return true
}
