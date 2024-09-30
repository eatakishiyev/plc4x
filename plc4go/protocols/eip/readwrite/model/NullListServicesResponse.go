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

// NullListServicesResponse is the corresponding interface of NullListServicesResponse
type NullListServicesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EipPacket
	// IsNullListServicesResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNullListServicesResponse()
	// CreateBuilder creates a NullListServicesResponseBuilder
	CreateNullListServicesResponseBuilder() NullListServicesResponseBuilder
}

// _NullListServicesResponse is the data-structure of this message
type _NullListServicesResponse struct {
	EipPacketContract
}

var _ NullListServicesResponse = (*_NullListServicesResponse)(nil)
var _ EipPacketRequirements = (*_NullListServicesResponse)(nil)

// NewNullListServicesResponse factory function for _NullListServicesResponse
func NewNullListServicesResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_NullListServicesResponse {
	_result := &_NullListServicesResponse{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NullListServicesResponseBuilder is a builder for NullListServicesResponse
type NullListServicesResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NullListServicesResponseBuilder
	// Build builds the NullListServicesResponse or returns an error if something is wrong
	Build() (NullListServicesResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NullListServicesResponse
}

// NewNullListServicesResponseBuilder() creates a NullListServicesResponseBuilder
func NewNullListServicesResponseBuilder() NullListServicesResponseBuilder {
	return &_NullListServicesResponseBuilder{_NullListServicesResponse: new(_NullListServicesResponse)}
}

type _NullListServicesResponseBuilder struct {
	*_NullListServicesResponse

	parentBuilder *_EipPacketBuilder

	err *utils.MultiError
}

var _ (NullListServicesResponseBuilder) = (*_NullListServicesResponseBuilder)(nil)

func (b *_NullListServicesResponseBuilder) setParent(contract EipPacketContract) {
	b.EipPacketContract = contract
}

func (b *_NullListServicesResponseBuilder) WithMandatoryFields() NullListServicesResponseBuilder {
	return b
}

func (b *_NullListServicesResponseBuilder) Build() (NullListServicesResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NullListServicesResponse.deepCopy(), nil
}

func (b *_NullListServicesResponseBuilder) MustBuild() NullListServicesResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NullListServicesResponseBuilder) Done() EipPacketBuilder {
	return b.parentBuilder
}

func (b *_NullListServicesResponseBuilder) buildForEipPacket() (EipPacket, error) {
	return b.Build()
}

func (b *_NullListServicesResponseBuilder) DeepCopy() any {
	_copy := b.CreateNullListServicesResponseBuilder().(*_NullListServicesResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNullListServicesResponseBuilder creates a NullListServicesResponseBuilder
func (b *_NullListServicesResponse) CreateNullListServicesResponseBuilder() NullListServicesResponseBuilder {
	if b == nil {
		return NewNullListServicesResponseBuilder()
	}
	return &_NullListServicesResponseBuilder{_NullListServicesResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullListServicesResponse) GetCommand() uint16 {
	return 0x0004
}

func (m *_NullListServicesResponse) GetResponse() bool {
	return bool(true)
}

func (m *_NullListServicesResponse) GetPacketLength() uint16 {
	return uint16(0)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullListServicesResponse) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// Deprecated: use the interface for direct cast
func CastNullListServicesResponse(structType any) NullListServicesResponse {
	if casted, ok := structType.(NullListServicesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NullListServicesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NullListServicesResponse) GetTypeName() string {
	return "NullListServicesResponse"
}

func (m *_NullListServicesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullListServicesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullListServicesResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__nullListServicesResponse NullListServicesResponse, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullListServicesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullListServicesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullListServicesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullListServicesResponse")
	}

	return m, nil
}

func (m *_NullListServicesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullListServicesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullListServicesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullListServicesResponse")
		}

		if popErr := writeBuffer.PopContext("NullListServicesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullListServicesResponse")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullListServicesResponse) IsNullListServicesResponse() {}

func (m *_NullListServicesResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NullListServicesResponse) deepCopy() *_NullListServicesResponse {
	if m == nil {
		return nil
	}
	_NullListServicesResponseCopy := &_NullListServicesResponse{
		m.EipPacketContract.(*_EipPacket).deepCopy(),
	}
	m.EipPacketContract.(*_EipPacket)._SubType = m
	return _NullListServicesResponseCopy
}

func (m *_NullListServicesResponse) String() string {
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
