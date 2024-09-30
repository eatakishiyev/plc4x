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

// HistoryData is the corresponding interface of HistoryData
type HistoryData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNoOfDataValues returns NoOfDataValues (property field)
	GetNoOfDataValues() int32
	// GetDataValues returns DataValues (property field)
	GetDataValues() []DataValue
	// IsHistoryData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryData()
	// CreateBuilder creates a HistoryDataBuilder
	CreateHistoryDataBuilder() HistoryDataBuilder
}

// _HistoryData is the data-structure of this message
type _HistoryData struct {
	ExtensionObjectDefinitionContract
	NoOfDataValues int32
	DataValues     []DataValue
}

var _ HistoryData = (*_HistoryData)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryData)(nil)

// NewHistoryData factory function for _HistoryData
func NewHistoryData(noOfDataValues int32, dataValues []DataValue) *_HistoryData {
	_result := &_HistoryData{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NoOfDataValues:                    noOfDataValues,
		DataValues:                        dataValues,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryDataBuilder is a builder for HistoryData
type HistoryDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(noOfDataValues int32, dataValues []DataValue) HistoryDataBuilder
	// WithNoOfDataValues adds NoOfDataValues (property field)
	WithNoOfDataValues(int32) HistoryDataBuilder
	// WithDataValues adds DataValues (property field)
	WithDataValues(...DataValue) HistoryDataBuilder
	// Build builds the HistoryData or returns an error if something is wrong
	Build() (HistoryData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryData
}

// NewHistoryDataBuilder() creates a HistoryDataBuilder
func NewHistoryDataBuilder() HistoryDataBuilder {
	return &_HistoryDataBuilder{_HistoryData: new(_HistoryData)}
}

type _HistoryDataBuilder struct {
	*_HistoryData

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryDataBuilder) = (*_HistoryDataBuilder)(nil)

func (b *_HistoryDataBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_HistoryDataBuilder) WithMandatoryFields(noOfDataValues int32, dataValues []DataValue) HistoryDataBuilder {
	return b.WithNoOfDataValues(noOfDataValues).WithDataValues(dataValues...)
}

func (b *_HistoryDataBuilder) WithNoOfDataValues(noOfDataValues int32) HistoryDataBuilder {
	b.NoOfDataValues = noOfDataValues
	return b
}

func (b *_HistoryDataBuilder) WithDataValues(dataValues ...DataValue) HistoryDataBuilder {
	b.DataValues = dataValues
	return b
}

func (b *_HistoryDataBuilder) Build() (HistoryData, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryData.deepCopy(), nil
}

func (b *_HistoryDataBuilder) MustBuild() HistoryData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_HistoryDataBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_HistoryDataBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryDataBuilder) DeepCopy() any {
	_copy := b.CreateHistoryDataBuilder().(*_HistoryDataBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryDataBuilder creates a HistoryDataBuilder
func (b *_HistoryData) CreateHistoryDataBuilder() HistoryDataBuilder {
	if b == nil {
		return NewHistoryDataBuilder()
	}
	return &_HistoryDataBuilder{_HistoryData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryData) GetIdentifier() string {
	return "658"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryData) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryData) GetNoOfDataValues() int32 {
	return m.NoOfDataValues
}

func (m *_HistoryData) GetDataValues() []DataValue {
	return m.DataValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryData(structType any) HistoryData {
	if casted, ok := structType.(HistoryData); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryData); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryData) GetTypeName() string {
	return "HistoryData"
}

func (m *_HistoryData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (noOfDataValues)
	lengthInBits += 32

	// Array field
	if len(m.DataValues) > 0 {
		for _curItem, element := range m.DataValues {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DataValues), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryData) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyData HistoryData, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfDataValues, err := ReadSimpleField(ctx, "noOfDataValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDataValues' field"))
	}
	m.NoOfDataValues = noOfDataValues

	dataValues, err := ReadCountArrayField[DataValue](ctx, "dataValues", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(noOfDataValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataValues' field"))
	}
	m.DataValues = dataValues

	if closeErr := readBuffer.CloseContext("HistoryData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryData")
	}

	return m, nil
}

func (m *_HistoryData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryData")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDataValues", m.GetNoOfDataValues(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDataValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "dataValues", m.GetDataValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'dataValues' field")
		}

		if popErr := writeBuffer.PopContext("HistoryData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryData")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryData) IsHistoryData() {}

func (m *_HistoryData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryData) deepCopy() *_HistoryData {
	if m == nil {
		return nil
	}
	_HistoryDataCopy := &_HistoryData{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NoOfDataValues,
		utils.DeepCopySlice[DataValue, DataValue](m.DataValues),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryDataCopy
}

func (m *_HistoryData) String() string {
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
