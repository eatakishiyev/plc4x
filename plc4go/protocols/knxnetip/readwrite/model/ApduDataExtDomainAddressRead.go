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

// ApduDataExtDomainAddressRead is the corresponding interface of ApduDataExtDomainAddressRead
type ApduDataExtDomainAddressRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtDomainAddressRead is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtDomainAddressRead()
	// CreateBuilder creates a ApduDataExtDomainAddressReadBuilder
	CreateApduDataExtDomainAddressReadBuilder() ApduDataExtDomainAddressReadBuilder
}

// _ApduDataExtDomainAddressRead is the data-structure of this message
type _ApduDataExtDomainAddressRead struct {
	ApduDataExtContract
}

var _ ApduDataExtDomainAddressRead = (*_ApduDataExtDomainAddressRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtDomainAddressRead)(nil)

// NewApduDataExtDomainAddressRead factory function for _ApduDataExtDomainAddressRead
func NewApduDataExtDomainAddressRead(length uint8) *_ApduDataExtDomainAddressRead {
	_result := &_ApduDataExtDomainAddressRead{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtDomainAddressReadBuilder is a builder for ApduDataExtDomainAddressRead
type ApduDataExtDomainAddressReadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtDomainAddressReadBuilder
	// Build builds the ApduDataExtDomainAddressRead or returns an error if something is wrong
	Build() (ApduDataExtDomainAddressRead, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtDomainAddressRead
}

// NewApduDataExtDomainAddressReadBuilder() creates a ApduDataExtDomainAddressReadBuilder
func NewApduDataExtDomainAddressReadBuilder() ApduDataExtDomainAddressReadBuilder {
	return &_ApduDataExtDomainAddressReadBuilder{_ApduDataExtDomainAddressRead: new(_ApduDataExtDomainAddressRead)}
}

type _ApduDataExtDomainAddressReadBuilder struct {
	*_ApduDataExtDomainAddressRead

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtDomainAddressReadBuilder) = (*_ApduDataExtDomainAddressReadBuilder)(nil)

func (b *_ApduDataExtDomainAddressReadBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
}

func (b *_ApduDataExtDomainAddressReadBuilder) WithMandatoryFields() ApduDataExtDomainAddressReadBuilder {
	return b
}

func (b *_ApduDataExtDomainAddressReadBuilder) Build() (ApduDataExtDomainAddressRead, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtDomainAddressRead.deepCopy(), nil
}

func (b *_ApduDataExtDomainAddressReadBuilder) MustBuild() ApduDataExtDomainAddressRead {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ApduDataExtDomainAddressReadBuilder) Done() ApduDataExtBuilder {
	return b.parentBuilder
}

func (b *_ApduDataExtDomainAddressReadBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtDomainAddressReadBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtDomainAddressReadBuilder().(*_ApduDataExtDomainAddressReadBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtDomainAddressReadBuilder creates a ApduDataExtDomainAddressReadBuilder
func (b *_ApduDataExtDomainAddressRead) CreateApduDataExtDomainAddressReadBuilder() ApduDataExtDomainAddressReadBuilder {
	if b == nil {
		return NewApduDataExtDomainAddressReadBuilder()
	}
	return &_ApduDataExtDomainAddressReadBuilder{_ApduDataExtDomainAddressRead: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressRead) GetExtApciType() uint8 {
	return 0x21
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressRead(structType any) ApduDataExtDomainAddressRead {
	if casted, ok := structType.(ApduDataExtDomainAddressRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressRead) GetTypeName() string {
	return "ApduDataExtDomainAddressRead"
}

func (m *_ApduDataExtDomainAddressRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtDomainAddressRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtDomainAddressRead ApduDataExtDomainAddressRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressRead")
	}

	return m, nil
}

func (m *_ApduDataExtDomainAddressRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressRead) IsApduDataExtDomainAddressRead() {}

func (m *_ApduDataExtDomainAddressRead) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtDomainAddressRead) deepCopy() *_ApduDataExtDomainAddressRead {
	if m == nil {
		return nil
	}
	_ApduDataExtDomainAddressReadCopy := &_ApduDataExtDomainAddressRead{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtDomainAddressReadCopy
}

func (m *_ApduDataExtDomainAddressRead) String() string {
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
