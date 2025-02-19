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

// CIPData is the corresponding interface of CIPData
type CIPData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDataType returns DataType (property field)
	GetDataType() CIPDataTypeCode
	// GetData returns Data (property field)
	GetData() []byte
	// IsCIPData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPData()
	// CreateBuilder creates a CIPDataBuilder
	CreateCIPDataBuilder() CIPDataBuilder
}

// _CIPData is the data-structure of this message
type _CIPData struct {
	DataType CIPDataTypeCode
	Data     []byte

	// Arguments.
	PacketLength uint16
}

var _ CIPData = (*_CIPData)(nil)

// NewCIPData factory function for _CIPData
func NewCIPData(dataType CIPDataTypeCode, data []byte, packetLength uint16) *_CIPData {
	return &_CIPData{DataType: dataType, Data: data, PacketLength: packetLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CIPDataBuilder is a builder for CIPData
type CIPDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataType CIPDataTypeCode, data []byte) CIPDataBuilder
	// WithDataType adds DataType (property field)
	WithDataType(CIPDataTypeCode) CIPDataBuilder
	// WithData adds Data (property field)
	WithData(...byte) CIPDataBuilder
	// Build builds the CIPData or returns an error if something is wrong
	Build() (CIPData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CIPData
}

// NewCIPDataBuilder() creates a CIPDataBuilder
func NewCIPDataBuilder() CIPDataBuilder {
	return &_CIPDataBuilder{_CIPData: new(_CIPData)}
}

type _CIPDataBuilder struct {
	*_CIPData

	err *utils.MultiError
}

var _ (CIPDataBuilder) = (*_CIPDataBuilder)(nil)

func (b *_CIPDataBuilder) WithMandatoryFields(dataType CIPDataTypeCode, data []byte) CIPDataBuilder {
	return b.WithDataType(dataType).WithData(data...)
}

func (b *_CIPDataBuilder) WithDataType(dataType CIPDataTypeCode) CIPDataBuilder {
	b.DataType = dataType
	return b
}

func (b *_CIPDataBuilder) WithData(data ...byte) CIPDataBuilder {
	b.Data = data
	return b
}

func (b *_CIPDataBuilder) Build() (CIPData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CIPData.deepCopy(), nil
}

func (b *_CIPDataBuilder) MustBuild() CIPData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CIPDataBuilder) DeepCopy() any {
	_copy := b.CreateCIPDataBuilder().(*_CIPDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCIPDataBuilder creates a CIPDataBuilder
func (b *_CIPData) CreateCIPDataBuilder() CIPDataBuilder {
	if b == nil {
		return NewCIPDataBuilder()
	}
	return &_CIPDataBuilder{_CIPData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPData) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *_CIPData) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCIPData(structType any) CIPData {
	if casted, ok := structType.(CIPData); ok {
		return casted
	}
	if casted, ok := structType.(*CIPData); ok {
		return *casted
	}
	return nil
}

func (m *_CIPData) GetTypeName() string {
	return "CIPData"
}

func (m *_CIPData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataType)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CIPData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPDataParse(ctx context.Context, theBytes []byte, packetLength uint16) (CIPData, error) {
	return CIPDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), packetLength)
}

func CIPDataParseWithBufferProducer(packetLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPData, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (CIPData, error) {
		return CIPDataParseWithBuffer(ctx, readBuffer, packetLength)
	}
}

func CIPDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (CIPData, error) {
	v, err := (&_CIPData{PacketLength: packetLength}).parse(ctx, readBuffer, packetLength)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_CIPData) parse(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (__cIPData CIPData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataType, err := ReadEnumField[CIPDataTypeCode](ctx, "dataType", "CIPDataTypeCode", ReadEnum(CIPDataTypeCodeByValue, ReadUnsignedShort(readBuffer, uint8(16))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataType' field"))
	}
	m.DataType = dataType

	data, err := readBuffer.ReadByteArray("data", int(int32(packetLength)-int32(int32(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("CIPData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPData")
	}

	return m, nil
}

func (m *_CIPData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPData")
	}

	if err := WriteSimpleEnumField[CIPDataTypeCode](ctx, "dataType", "CIPDataTypeCode", m.GetDataType(), WriteEnum[CIPDataTypeCode, uint16](CIPDataTypeCode.GetValue, CIPDataTypeCode.PLC4XEnumName, WriteUnsignedShort(writeBuffer, 16))); err != nil {
		return errors.Wrap(err, "Error serializing 'dataType' field")
	}

	if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("CIPData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPData")
	}
	return nil
}

////
// Arguments Getter

func (m *_CIPData) GetPacketLength() uint16 {
	return m.PacketLength
}

//
////

func (m *_CIPData) IsCIPData() {}

func (m *_CIPData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CIPData) deepCopy() *_CIPData {
	if m == nil {
		return nil
	}
	_CIPDataCopy := &_CIPData{
		m.DataType,
		utils.DeepCopySlice[byte, byte](m.Data),
		m.PacketLength,
	}
	return _CIPDataCopy
}

func (m *_CIPData) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
