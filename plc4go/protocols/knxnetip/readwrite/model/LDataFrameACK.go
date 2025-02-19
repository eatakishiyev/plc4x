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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// LDataFrameACK is the corresponding interface of LDataFrameACK
type LDataFrameACK interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	LDataFrame
	// IsLDataFrameACK is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLDataFrameACK()
	// CreateBuilder creates a LDataFrameACKBuilder
	CreateLDataFrameACKBuilder() LDataFrameACKBuilder
}

// _LDataFrameACK is the data-structure of this message
type _LDataFrameACK struct {
	LDataFrameContract
}

var _ LDataFrameACK = (*_LDataFrameACK)(nil)
var _ LDataFrameRequirements = (*_LDataFrameACK)(nil)

// NewLDataFrameACK factory function for _LDataFrameACK
func NewLDataFrameACK(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LDataFrameACK {
	_result := &_LDataFrameACK{
		LDataFrameContract: NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	_result.LDataFrameContract.(*_LDataFrame)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LDataFrameACKBuilder is a builder for LDataFrameACK
type LDataFrameACKBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() LDataFrameACKBuilder
	// Build builds the LDataFrameACK or returns an error if something is wrong
	Build() (LDataFrameACK, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LDataFrameACK
}

// NewLDataFrameACKBuilder() creates a LDataFrameACKBuilder
func NewLDataFrameACKBuilder() LDataFrameACKBuilder {
	return &_LDataFrameACKBuilder{_LDataFrameACK: new(_LDataFrameACK)}
}

type _LDataFrameACKBuilder struct {
	*_LDataFrameACK

	parentBuilder *_LDataFrameBuilder

	err *utils.MultiError
}

var _ (LDataFrameACKBuilder) = (*_LDataFrameACKBuilder)(nil)

func (b *_LDataFrameACKBuilder) setParent(contract LDataFrameContract) {
	b.LDataFrameContract = contract
}

func (b *_LDataFrameACKBuilder) WithMandatoryFields() LDataFrameACKBuilder {
	return b
}

func (b *_LDataFrameACKBuilder) Build() (LDataFrameACK, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LDataFrameACK.deepCopy(), nil
}

func (b *_LDataFrameACKBuilder) MustBuild() LDataFrameACK {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_LDataFrameACKBuilder) Done() LDataFrameBuilder {
	return b.parentBuilder
}

func (b *_LDataFrameACKBuilder) buildForLDataFrame() (LDataFrame, error) {
	return b.Build()
}

func (b *_LDataFrameACKBuilder) DeepCopy() any {
	_copy := b.CreateLDataFrameACKBuilder().(*_LDataFrameACKBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLDataFrameACKBuilder creates a LDataFrameACKBuilder
func (b *_LDataFrameACK) CreateLDataFrameACKBuilder() LDataFrameACKBuilder {
	if b == nil {
		return NewLDataFrameACKBuilder()
	}
	return &_LDataFrameACKBuilder{_LDataFrameACK: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataFrameACK) GetNotAckFrame() bool {
	return bool(false)
}

func (m *_LDataFrameACK) GetPolling() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataFrameACK) GetParent() LDataFrameContract {
	return m.LDataFrameContract
}

// Deprecated: use the interface for direct cast
func CastLDataFrameACK(structType any) LDataFrameACK {
	if casted, ok := structType.(LDataFrameACK); ok {
		return casted
	}
	if casted, ok := structType.(*LDataFrameACK); ok {
		return *casted
	}
	return nil
}

func (m *_LDataFrameACK) GetTypeName() string {
	return "LDataFrameACK"
}

func (m *_LDataFrameACK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LDataFrameContract.(*_LDataFrame).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LDataFrameACK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LDataFrameACK) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LDataFrame) (__lDataFrameACK LDataFrameACK, err error) {
	m.LDataFrameContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataFrameACK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataFrameACK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LDataFrameACK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrameACK")
	}

	return m, nil
}

func (m *_LDataFrameACK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LDataFrameACK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataFrameACK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataFrameACK")
		}

		if popErr := writeBuffer.PopContext("LDataFrameACK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataFrameACK")
		}
		return nil
	}
	return m.LDataFrameContract.(*_LDataFrame).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LDataFrameACK) IsLDataFrameACK() {}

func (m *_LDataFrameACK) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LDataFrameACK) deepCopy() *_LDataFrameACK {
	if m == nil {
		return nil
	}
	_LDataFrameACKCopy := &_LDataFrameACK{
		m.LDataFrameContract.(*_LDataFrame).deepCopy(),
	}
	m.LDataFrameContract.(*_LDataFrame)._SubType = m
	return _LDataFrameACKCopy
}

func (m *_LDataFrameACK) String() string {
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
