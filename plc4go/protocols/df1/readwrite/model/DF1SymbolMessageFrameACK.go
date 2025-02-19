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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1SymbolMessageFrameACK is the corresponding interface of DF1SymbolMessageFrameACK
type DF1SymbolMessageFrameACK interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1Symbol
	// IsDF1SymbolMessageFrameACK is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1SymbolMessageFrameACK()
	// CreateBuilder creates a DF1SymbolMessageFrameACKBuilder
	CreateDF1SymbolMessageFrameACKBuilder() DF1SymbolMessageFrameACKBuilder
}

// _DF1SymbolMessageFrameACK is the data-structure of this message
type _DF1SymbolMessageFrameACK struct {
	DF1SymbolContract
}

var _ DF1SymbolMessageFrameACK = (*_DF1SymbolMessageFrameACK)(nil)
var _ DF1SymbolRequirements = (*_DF1SymbolMessageFrameACK)(nil)

// NewDF1SymbolMessageFrameACK factory function for _DF1SymbolMessageFrameACK
func NewDF1SymbolMessageFrameACK() *_DF1SymbolMessageFrameACK {
	_result := &_DF1SymbolMessageFrameACK{
		DF1SymbolContract: NewDF1Symbol(),
	}
	_result.DF1SymbolContract.(*_DF1Symbol)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1SymbolMessageFrameACKBuilder is a builder for DF1SymbolMessageFrameACK
type DF1SymbolMessageFrameACKBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DF1SymbolMessageFrameACKBuilder
	// Build builds the DF1SymbolMessageFrameACK or returns an error if something is wrong
	Build() (DF1SymbolMessageFrameACK, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1SymbolMessageFrameACK
}

// NewDF1SymbolMessageFrameACKBuilder() creates a DF1SymbolMessageFrameACKBuilder
func NewDF1SymbolMessageFrameACKBuilder() DF1SymbolMessageFrameACKBuilder {
	return &_DF1SymbolMessageFrameACKBuilder{_DF1SymbolMessageFrameACK: new(_DF1SymbolMessageFrameACK)}
}

type _DF1SymbolMessageFrameACKBuilder struct {
	*_DF1SymbolMessageFrameACK

	parentBuilder *_DF1SymbolBuilder

	err *utils.MultiError
}

var _ (DF1SymbolMessageFrameACKBuilder) = (*_DF1SymbolMessageFrameACKBuilder)(nil)

func (b *_DF1SymbolMessageFrameACKBuilder) setParent(contract DF1SymbolContract) {
	b.DF1SymbolContract = contract
}

func (b *_DF1SymbolMessageFrameACKBuilder) WithMandatoryFields() DF1SymbolMessageFrameACKBuilder {
	return b
}

func (b *_DF1SymbolMessageFrameACKBuilder) Build() (DF1SymbolMessageFrameACK, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1SymbolMessageFrameACK.deepCopy(), nil
}

func (b *_DF1SymbolMessageFrameACKBuilder) MustBuild() DF1SymbolMessageFrameACK {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DF1SymbolMessageFrameACKBuilder) Done() DF1SymbolBuilder {
	return b.parentBuilder
}

func (b *_DF1SymbolMessageFrameACKBuilder) buildForDF1Symbol() (DF1Symbol, error) {
	return b.Build()
}

func (b *_DF1SymbolMessageFrameACKBuilder) DeepCopy() any {
	_copy := b.CreateDF1SymbolMessageFrameACKBuilder().(*_DF1SymbolMessageFrameACKBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1SymbolMessageFrameACKBuilder creates a DF1SymbolMessageFrameACKBuilder
func (b *_DF1SymbolMessageFrameACK) CreateDF1SymbolMessageFrameACKBuilder() DF1SymbolMessageFrameACKBuilder {
	if b == nil {
		return NewDF1SymbolMessageFrameACKBuilder()
	}
	return &_DF1SymbolMessageFrameACKBuilder{_DF1SymbolMessageFrameACK: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1SymbolMessageFrameACK) GetSymbolType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1SymbolMessageFrameACK) GetParent() DF1SymbolContract {
	return m.DF1SymbolContract
}

// Deprecated: use the interface for direct cast
func CastDF1SymbolMessageFrameACK(structType any) DF1SymbolMessageFrameACK {
	if casted, ok := structType.(DF1SymbolMessageFrameACK); ok {
		return casted
	}
	if casted, ok := structType.(*DF1SymbolMessageFrameACK); ok {
		return *casted
	}
	return nil
}

func (m *_DF1SymbolMessageFrameACK) GetTypeName() string {
	return "DF1SymbolMessageFrameACK"
}

func (m *_DF1SymbolMessageFrameACK) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1SymbolContract.(*_DF1Symbol).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DF1SymbolMessageFrameACK) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1SymbolMessageFrameACK) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Symbol) (__dF1SymbolMessageFrameACK DF1SymbolMessageFrameACK, err error) {
	m.DF1SymbolContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1SymbolMessageFrameACK"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1SymbolMessageFrameACK")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DF1SymbolMessageFrameACK"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1SymbolMessageFrameACK")
	}

	return m, nil
}

func (m *_DF1SymbolMessageFrameACK) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1SymbolMessageFrameACK) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1SymbolMessageFrameACK"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1SymbolMessageFrameACK")
		}

		if popErr := writeBuffer.PopContext("DF1SymbolMessageFrameACK"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1SymbolMessageFrameACK")
		}
		return nil
	}
	return m.DF1SymbolContract.(*_DF1Symbol).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1SymbolMessageFrameACK) IsDF1SymbolMessageFrameACK() {}

func (m *_DF1SymbolMessageFrameACK) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1SymbolMessageFrameACK) deepCopy() *_DF1SymbolMessageFrameACK {
	if m == nil {
		return nil
	}
	_DF1SymbolMessageFrameACKCopy := &_DF1SymbolMessageFrameACK{
		m.DF1SymbolContract.(*_DF1Symbol).deepCopy(),
	}
	m.DF1SymbolContract.(*_DF1Symbol)._SubType = m
	return _DF1SymbolMessageFrameACKCopy
}

func (m *_DF1SymbolMessageFrameACK) String() string {
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
