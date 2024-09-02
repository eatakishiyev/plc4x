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

// GuidValue is the corresponding interface of GuidValue
type GuidValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetData1 returns Data1 (property field)
	GetData1() uint32
	// GetData2 returns Data2 (property field)
	GetData2() uint16
	// GetData3 returns Data3 (property field)
	GetData3() uint16
	// GetData4 returns Data4 (property field)
	GetData4() []byte
	// GetData5 returns Data5 (property field)
	GetData5() []byte
}

// GuidValueExactly can be used when we want exactly this type and not a type which fulfills GuidValue.
// This is useful for switch cases.
type GuidValueExactly interface {
	GuidValue
	isGuidValue() bool
}

// _GuidValue is the data-structure of this message
type _GuidValue struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 []byte
	Data5 []byte
}

var _ GuidValue = (*_GuidValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GuidValue) GetData1() uint32 {
	return m.Data1
}

func (m *_GuidValue) GetData2() uint16 {
	return m.Data2
}

func (m *_GuidValue) GetData3() uint16 {
	return m.Data3
}

func (m *_GuidValue) GetData4() []byte {
	return m.Data4
}

func (m *_GuidValue) GetData5() []byte {
	return m.Data5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGuidValue factory function for _GuidValue
func NewGuidValue(data1 uint32, data2 uint16, data3 uint16, data4 []byte, data5 []byte) *_GuidValue {
	return &_GuidValue{Data1: data1, Data2: data2, Data3: data3, Data4: data4, Data5: data5}
}

// Deprecated: use the interface for direct cast
func CastGuidValue(structType any) GuidValue {
	if casted, ok := structType.(GuidValue); ok {
		return casted
	}
	if casted, ok := structType.(*GuidValue); ok {
		return *casted
	}
	return nil
}

func (m *_GuidValue) GetTypeName() string {
	return "GuidValue"
}

func (m *_GuidValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (data1)
	lengthInBits += 32

	// Simple field (data2)
	lengthInBits += 16

	// Simple field (data3)
	lengthInBits += 16

	// Array field
	if len(m.Data4) > 0 {
		lengthInBits += 8 * uint16(len(m.Data4))
	}

	// Array field
	if len(m.Data5) > 0 {
		lengthInBits += 8 * uint16(len(m.Data5))
	}

	return lengthInBits
}

func (m *_GuidValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GuidValueParse(ctx context.Context, theBytes []byte) (GuidValue, error) {
	return GuidValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GuidValueParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GuidValue, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GuidValue, error) {
		return GuidValueParseWithBuffer(ctx, readBuffer)
	}
}

func GuidValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GuidValue, error) {
	v, err := NewGuidValue().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_GuidValue) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__guidValue GuidValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GuidValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GuidValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data1, err := ReadSimpleField(ctx, "data1", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data1' field"))
	}
	m.Data1 = data1

	data2, err := ReadSimpleField(ctx, "data2", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data2' field"))
	}
	m.Data2 = data2

	data3, err := ReadSimpleField(ctx, "data3", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data3' field"))
	}
	m.Data3 = data3

	data4, err := readBuffer.ReadByteArray("data4", int(int32(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data4' field"))
	}
	m.Data4 = data4

	data5, err := readBuffer.ReadByteArray("data5", int(int32(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data5' field"))
	}
	m.Data5 = data5

	if closeErr := readBuffer.CloseContext("GuidValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GuidValue")
	}

	return m, nil
}

func (m *_GuidValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GuidValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GuidValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GuidValue")
	}

	if err := WriteSimpleField[uint32](ctx, "data1", m.GetData1(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'data1' field")
	}

	if err := WriteSimpleField[uint16](ctx, "data2", m.GetData2(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'data2' field")
	}

	if err := WriteSimpleField[uint16](ctx, "data3", m.GetData3(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'data3' field")
	}

	if err := WriteByteArrayField(ctx, "data4", m.GetData4(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data4' field")
	}

	if err := WriteByteArrayField(ctx, "data5", m.GetData5(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data5' field")
	}

	if popErr := writeBuffer.PopContext("GuidValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GuidValue")
	}
	return nil
}

func (m *_GuidValue) isGuidValue() bool {
	return true
}

func (m *_GuidValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
