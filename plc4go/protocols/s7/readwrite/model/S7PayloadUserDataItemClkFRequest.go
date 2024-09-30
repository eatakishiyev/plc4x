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

// S7PayloadUserDataItemClkFRequest is the corresponding interface of S7PayloadUserDataItemClkFRequest
type S7PayloadUserDataItemClkFRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7PayloadUserDataItem
	// IsS7PayloadUserDataItemClkFRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadUserDataItemClkFRequest()
	// CreateBuilder creates a S7PayloadUserDataItemClkFRequestBuilder
	CreateS7PayloadUserDataItemClkFRequestBuilder() S7PayloadUserDataItemClkFRequestBuilder
}

// _S7PayloadUserDataItemClkFRequest is the data-structure of this message
type _S7PayloadUserDataItemClkFRequest struct {
	S7PayloadUserDataItemContract
}

var _ S7PayloadUserDataItemClkFRequest = (*_S7PayloadUserDataItemClkFRequest)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadUserDataItemClkFRequest)(nil)

// NewS7PayloadUserDataItemClkFRequest factory function for _S7PayloadUserDataItemClkFRequest
func NewS7PayloadUserDataItemClkFRequest(returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadUserDataItemClkFRequest {
	_result := &_S7PayloadUserDataItemClkFRequest{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadUserDataItemClkFRequestBuilder is a builder for S7PayloadUserDataItemClkFRequest
type S7PayloadUserDataItemClkFRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() S7PayloadUserDataItemClkFRequestBuilder
	// Build builds the S7PayloadUserDataItemClkFRequest or returns an error if something is wrong
	Build() (S7PayloadUserDataItemClkFRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadUserDataItemClkFRequest
}

// NewS7PayloadUserDataItemClkFRequestBuilder() creates a S7PayloadUserDataItemClkFRequestBuilder
func NewS7PayloadUserDataItemClkFRequestBuilder() S7PayloadUserDataItemClkFRequestBuilder {
	return &_S7PayloadUserDataItemClkFRequestBuilder{_S7PayloadUserDataItemClkFRequest: new(_S7PayloadUserDataItemClkFRequest)}
}

type _S7PayloadUserDataItemClkFRequestBuilder struct {
	*_S7PayloadUserDataItemClkFRequest

	parentBuilder *_S7PayloadUserDataItemBuilder

	err *utils.MultiError
}

var _ (S7PayloadUserDataItemClkFRequestBuilder) = (*_S7PayloadUserDataItemClkFRequestBuilder)(nil)

func (b *_S7PayloadUserDataItemClkFRequestBuilder) setParent(contract S7PayloadUserDataItemContract) {
	b.S7PayloadUserDataItemContract = contract
}

func (b *_S7PayloadUserDataItemClkFRequestBuilder) WithMandatoryFields() S7PayloadUserDataItemClkFRequestBuilder {
	return b
}

func (b *_S7PayloadUserDataItemClkFRequestBuilder) Build() (S7PayloadUserDataItemClkFRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadUserDataItemClkFRequest.deepCopy(), nil
}

func (b *_S7PayloadUserDataItemClkFRequestBuilder) MustBuild() S7PayloadUserDataItemClkFRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_S7PayloadUserDataItemClkFRequestBuilder) Done() S7PayloadUserDataItemBuilder {
	return b.parentBuilder
}

func (b *_S7PayloadUserDataItemClkFRequestBuilder) buildForS7PayloadUserDataItem() (S7PayloadUserDataItem, error) {
	return b.Build()
}

func (b *_S7PayloadUserDataItemClkFRequestBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadUserDataItemClkFRequestBuilder().(*_S7PayloadUserDataItemClkFRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadUserDataItemClkFRequestBuilder creates a S7PayloadUserDataItemClkFRequestBuilder
func (b *_S7PayloadUserDataItemClkFRequest) CreateS7PayloadUserDataItemClkFRequestBuilder() S7PayloadUserDataItemClkFRequestBuilder {
	if b == nil {
		return NewS7PayloadUserDataItemClkFRequestBuilder()
	}
	return &_S7PayloadUserDataItemClkFRequestBuilder{_S7PayloadUserDataItemClkFRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuFunctionGroup() uint8 {
	return 0x07
}

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuFunctionType() uint8 {
	return 0x04
}

func (m *_S7PayloadUserDataItemClkFRequest) GetCpuSubfunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemClkFRequest) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemClkFRequest(structType any) S7PayloadUserDataItemClkFRequest {
	if casted, ok := structType.(S7PayloadUserDataItemClkFRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemClkFRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemClkFRequest) GetTypeName() string {
	return "S7PayloadUserDataItemClkFRequest"
}

func (m *_S7PayloadUserDataItemClkFRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_S7PayloadUserDataItemClkFRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadUserDataItemClkFRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadUserDataItemClkFRequest S7PayloadUserDataItemClkFRequest, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemClkFRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemClkFRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemClkFRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemClkFRequest")
	}

	return m, nil
}

func (m *_S7PayloadUserDataItemClkFRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemClkFRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemClkFRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemClkFRequest")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemClkFRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemClkFRequest")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemClkFRequest) IsS7PayloadUserDataItemClkFRequest() {}

func (m *_S7PayloadUserDataItemClkFRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadUserDataItemClkFRequest) deepCopy() *_S7PayloadUserDataItemClkFRequest {
	if m == nil {
		return nil
	}
	_S7PayloadUserDataItemClkFRequestCopy := &_S7PayloadUserDataItemClkFRequest{
		m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).deepCopy(),
	}
	m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = m
	return _S7PayloadUserDataItemClkFRequestCopy
}

func (m *_S7PayloadUserDataItemClkFRequest) String() string {
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
