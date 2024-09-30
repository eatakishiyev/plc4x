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

// SetAttributeListResponse is the corresponding interface of SetAttributeListResponse
type SetAttributeListResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsSetAttributeListResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSetAttributeListResponse()
	// CreateBuilder creates a SetAttributeListResponseBuilder
	CreateSetAttributeListResponseBuilder() SetAttributeListResponseBuilder
}

// _SetAttributeListResponse is the data-structure of this message
type _SetAttributeListResponse struct {
	CipServiceContract
}

var _ SetAttributeListResponse = (*_SetAttributeListResponse)(nil)
var _ CipServiceRequirements = (*_SetAttributeListResponse)(nil)

// NewSetAttributeListResponse factory function for _SetAttributeListResponse
func NewSetAttributeListResponse(serviceLen uint16) *_SetAttributeListResponse {
	_result := &_SetAttributeListResponse{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SetAttributeListResponseBuilder is a builder for SetAttributeListResponse
type SetAttributeListResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SetAttributeListResponseBuilder
	// Build builds the SetAttributeListResponse or returns an error if something is wrong
	Build() (SetAttributeListResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SetAttributeListResponse
}

// NewSetAttributeListResponseBuilder() creates a SetAttributeListResponseBuilder
func NewSetAttributeListResponseBuilder() SetAttributeListResponseBuilder {
	return &_SetAttributeListResponseBuilder{_SetAttributeListResponse: new(_SetAttributeListResponse)}
}

type _SetAttributeListResponseBuilder struct {
	*_SetAttributeListResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (SetAttributeListResponseBuilder) = (*_SetAttributeListResponseBuilder)(nil)

func (b *_SetAttributeListResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_SetAttributeListResponseBuilder) WithMandatoryFields() SetAttributeListResponseBuilder {
	return b
}

func (b *_SetAttributeListResponseBuilder) Build() (SetAttributeListResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SetAttributeListResponse.deepCopy(), nil
}

func (b *_SetAttributeListResponseBuilder) MustBuild() SetAttributeListResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_SetAttributeListResponseBuilder) Done() CipServiceBuilder {
	return b.parentBuilder
}

func (b *_SetAttributeListResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_SetAttributeListResponseBuilder) DeepCopy() any {
	_copy := b.CreateSetAttributeListResponseBuilder().(*_SetAttributeListResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSetAttributeListResponseBuilder creates a SetAttributeListResponseBuilder
func (b *_SetAttributeListResponse) CreateSetAttributeListResponseBuilder() SetAttributeListResponseBuilder {
	if b == nil {
		return NewSetAttributeListResponseBuilder()
	}
	return &_SetAttributeListResponseBuilder{_SetAttributeListResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeListResponse) GetService() uint8 {
	return 0x04
}

func (m *_SetAttributeListResponse) GetResponse() bool {
	return bool(true)
}

func (m *_SetAttributeListResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeListResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastSetAttributeListResponse(structType any) SetAttributeListResponse {
	if casted, ok := structType.(SetAttributeListResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeListResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeListResponse) GetTypeName() string {
	return "SetAttributeListResponse"
}

func (m *_SetAttributeListResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeListResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SetAttributeListResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__setAttributeListResponse SetAttributeListResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeListResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeListResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeListResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeListResponse")
	}

	return m, nil
}

func (m *_SetAttributeListResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeListResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeListResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeListResponse")
		}

		if popErr := writeBuffer.PopContext("SetAttributeListResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeListResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeListResponse) IsSetAttributeListResponse() {}

func (m *_SetAttributeListResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SetAttributeListResponse) deepCopy() *_SetAttributeListResponse {
	if m == nil {
		return nil
	}
	_SetAttributeListResponseCopy := &_SetAttributeListResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _SetAttributeListResponseCopy
}

func (m *_SetAttributeListResponse) String() string {
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
