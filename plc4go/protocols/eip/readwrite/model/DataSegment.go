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

// DataSegment is the corresponding interface of DataSegment
type DataSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PathSegment
	// GetSegmentType returns SegmentType (property field)
	GetSegmentType() DataSegmentType
	// IsDataSegment is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSegment()
	// CreateBuilder creates a DataSegmentBuilder
	CreateDataSegmentBuilder() DataSegmentBuilder
}

// _DataSegment is the data-structure of this message
type _DataSegment struct {
	PathSegmentContract
	SegmentType DataSegmentType
}

var _ DataSegment = (*_DataSegment)(nil)
var _ PathSegmentRequirements = (*_DataSegment)(nil)

// NewDataSegment factory function for _DataSegment
func NewDataSegment(segmentType DataSegmentType) *_DataSegment {
	if segmentType == nil {
		panic("segmentType of type DataSegmentType for DataSegment must not be nil")
	}
	_result := &_DataSegment{
		PathSegmentContract: NewPathSegment(),
		SegmentType:         segmentType,
	}
	_result.PathSegmentContract.(*_PathSegment)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSegmentBuilder is a builder for DataSegment
type DataSegmentBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentType DataSegmentType) DataSegmentBuilder
	// WithSegmentType adds SegmentType (property field)
	WithSegmentType(DataSegmentType) DataSegmentBuilder
	// WithSegmentTypeBuilder adds SegmentType (property field) which is build by the builder
	WithSegmentTypeBuilder(func(DataSegmentTypeBuilder) DataSegmentTypeBuilder) DataSegmentBuilder
	// Build builds the DataSegment or returns an error if something is wrong
	Build() (DataSegment, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSegment
}

// NewDataSegmentBuilder() creates a DataSegmentBuilder
func NewDataSegmentBuilder() DataSegmentBuilder {
	return &_DataSegmentBuilder{_DataSegment: new(_DataSegment)}
}

type _DataSegmentBuilder struct {
	*_DataSegment

	parentBuilder *_PathSegmentBuilder

	err *utils.MultiError
}

var _ (DataSegmentBuilder) = (*_DataSegmentBuilder)(nil)

func (b *_DataSegmentBuilder) setParent(contract PathSegmentContract) {
	b.PathSegmentContract = contract
}

func (b *_DataSegmentBuilder) WithMandatoryFields(segmentType DataSegmentType) DataSegmentBuilder {
	return b.WithSegmentType(segmentType)
}

func (b *_DataSegmentBuilder) WithSegmentType(segmentType DataSegmentType) DataSegmentBuilder {
	b.SegmentType = segmentType
	return b
}

func (b *_DataSegmentBuilder) WithSegmentTypeBuilder(builderSupplier func(DataSegmentTypeBuilder) DataSegmentTypeBuilder) DataSegmentBuilder {
	builder := builderSupplier(b.SegmentType.CreateDataSegmentTypeBuilder())
	var err error
	b.SegmentType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "DataSegmentTypeBuilder failed"))
	}
	return b
}

func (b *_DataSegmentBuilder) Build() (DataSegment, error) {
	if b.SegmentType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'segmentType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataSegment.deepCopy(), nil
}

func (b *_DataSegmentBuilder) MustBuild() DataSegment {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DataSegmentBuilder) Done() PathSegmentBuilder {
	return b.parentBuilder
}

func (b *_DataSegmentBuilder) buildForPathSegment() (PathSegment, error) {
	return b.Build()
}

func (b *_DataSegmentBuilder) DeepCopy() any {
	_copy := b.CreateDataSegmentBuilder().(*_DataSegmentBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataSegmentBuilder creates a DataSegmentBuilder
func (b *_DataSegment) CreateDataSegmentBuilder() DataSegmentBuilder {
	if b == nil {
		return NewDataSegmentBuilder()
	}
	return &_DataSegmentBuilder{_DataSegment: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSegment) GetPathSegment() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSegment) GetParent() PathSegmentContract {
	return m.PathSegmentContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DataSegment) GetSegmentType() DataSegmentType {
	return m.SegmentType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDataSegment(structType any) DataSegment {
	if casted, ok := structType.(DataSegment); ok {
		return casted
	}
	if casted, ok := structType.(*DataSegment); ok {
		return *casted
	}
	return nil
}

func (m *_DataSegment) GetTypeName() string {
	return "DataSegment"
}

func (m *_DataSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PathSegmentContract.(*_PathSegment).getLengthInBits(ctx))

	// Simple field (segmentType)
	lengthInBits += m.SegmentType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DataSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSegment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PathSegment) (__dataSegment DataSegment, err error) {
	m.PathSegmentContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentType, err := ReadSimpleField[DataSegmentType](ctx, "segmentType", ReadComplex[DataSegmentType](DataSegmentTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentType' field"))
	}
	m.SegmentType = segmentType

	if closeErr := readBuffer.CloseContext("DataSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSegment")
	}

	return m, nil
}

func (m *_DataSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSegment")
		}

		if err := WriteSimpleField[DataSegmentType](ctx, "segmentType", m.GetSegmentType(), WriteComplex[DataSegmentType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentType' field")
		}

		if popErr := writeBuffer.PopContext("DataSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSegment")
		}
		return nil
	}
	return m.PathSegmentContract.(*_PathSegment).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSegment) IsDataSegment() {}

func (m *_DataSegment) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSegment) deepCopy() *_DataSegment {
	if m == nil {
		return nil
	}
	_DataSegmentCopy := &_DataSegment{
		m.PathSegmentContract.(*_PathSegment).deepCopy(),
		m.SegmentType.DeepCopy().(DataSegmentType),
	}
	m.PathSegmentContract.(*_PathSegment)._SubType = m
	return _DataSegmentCopy
}

func (m *_DataSegment) String() string {
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
