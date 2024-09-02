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

// LightingLabelOptions is the corresponding interface of LightingLabelOptions
type LightingLabelOptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLabelFlavour returns LabelFlavour (property field)
	GetLabelFlavour() LightingLabelFlavour
	// GetLabelType returns LabelType (property field)
	GetLabelType() LightingLabelType
}

// LightingLabelOptionsExactly can be used when we want exactly this type and not a type which fulfills LightingLabelOptions.
// This is useful for switch cases.
type LightingLabelOptionsExactly interface {
	LightingLabelOptions
	isLightingLabelOptions() bool
}

// _LightingLabelOptions is the data-structure of this message
type _LightingLabelOptions struct {
	LabelFlavour LightingLabelFlavour
	LabelType    LightingLabelType
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

var _ LightingLabelOptions = (*_LightingLabelOptions)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingLabelOptions) GetLabelFlavour() LightingLabelFlavour {
	return m.LabelFlavour
}

func (m *_LightingLabelOptions) GetLabelType() LightingLabelType {
	return m.LabelType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingLabelOptions factory function for _LightingLabelOptions
func NewLightingLabelOptions(labelFlavour LightingLabelFlavour, labelType LightingLabelType) *_LightingLabelOptions {
	return &_LightingLabelOptions{LabelFlavour: labelFlavour, LabelType: labelType}
}

// Deprecated: use the interface for direct cast
func CastLightingLabelOptions(structType any) LightingLabelOptions {
	if casted, ok := structType.(LightingLabelOptions); ok {
		return casted
	}
	if casted, ok := structType.(*LightingLabelOptions); ok {
		return *casted
	}
	return nil
}

func (m *_LightingLabelOptions) GetTypeName() string {
	return "LightingLabelOptions"
}

func (m *_LightingLabelOptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelFlavour)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (labelType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_LightingLabelOptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LightingLabelOptionsParse(ctx context.Context, theBytes []byte) (LightingLabelOptions, error) {
	return LightingLabelOptionsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LightingLabelOptionsParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (LightingLabelOptions, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (LightingLabelOptions, error) {
		return LightingLabelOptionsParseWithBuffer(ctx, readBuffer)
	}
}

func LightingLabelOptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LightingLabelOptions, error) {
	v, err := NewLightingLabelOptions().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_LightingLabelOptions) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__lightingLabelOptions LightingLabelOptions, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingLabelOptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingLabelOptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	labelFlavour, err := ReadEnumField[LightingLabelFlavour](ctx, "labelFlavour", "LightingLabelFlavour", ReadEnum(LightingLabelFlavourByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelFlavour' field"))
	}
	m.LabelFlavour = labelFlavour

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField2 = reservedField2

	labelType, err := ReadEnumField[LightingLabelType](ctx, "labelType", "LightingLabelType", ReadEnum(LightingLabelTypeByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'labelType' field"))
	}
	m.LabelType = labelType

	reservedField3, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField3 = reservedField3

	if closeErr := readBuffer.CloseContext("LightingLabelOptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingLabelOptions")
	}

	return m, nil
}

func (m *_LightingLabelOptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingLabelOptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("LightingLabelOptions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LightingLabelOptions")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleEnumField[LightingLabelFlavour](ctx, "labelFlavour", "LightingLabelFlavour", m.GetLabelFlavour(), WriteEnum[LightingLabelFlavour, uint8](LightingLabelFlavour.GetValue, LightingLabelFlavour.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'labelFlavour' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	if err := WriteSimpleEnumField[LightingLabelType](ctx, "labelType", "LightingLabelType", m.GetLabelType(), WriteEnum[LightingLabelType, uint8](LightingLabelType.GetValue, LightingLabelType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'labelType' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 4")
	}

	if popErr := writeBuffer.PopContext("LightingLabelOptions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LightingLabelOptions")
	}
	return nil
}

func (m *_LightingLabelOptions) isLightingLabelOptions() bool {
	return true
}

func (m *_LightingLabelOptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
