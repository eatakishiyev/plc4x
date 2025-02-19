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

// CIPEncapsulationConnectionRequest is the corresponding interface of CIPEncapsulationConnectionRequest
type CIPEncapsulationConnectionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CIPEncapsulationPacket
	// IsCIPEncapsulationConnectionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCIPEncapsulationConnectionRequest()
	// CreateBuilder creates a CIPEncapsulationConnectionRequestBuilder
	CreateCIPEncapsulationConnectionRequestBuilder() CIPEncapsulationConnectionRequestBuilder
}

// _CIPEncapsulationConnectionRequest is the data-structure of this message
type _CIPEncapsulationConnectionRequest struct {
	CIPEncapsulationPacketContract
}

var _ CIPEncapsulationConnectionRequest = (*_CIPEncapsulationConnectionRequest)(nil)
var _ CIPEncapsulationPacketRequirements = (*_CIPEncapsulationConnectionRequest)(nil)

// NewCIPEncapsulationConnectionRequest factory function for _CIPEncapsulationConnectionRequest
func NewCIPEncapsulationConnectionRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *_CIPEncapsulationConnectionRequest {
	_result := &_CIPEncapsulationConnectionRequest{
		CIPEncapsulationPacketContract: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CIPEncapsulationConnectionRequestBuilder is a builder for CIPEncapsulationConnectionRequest
type CIPEncapsulationConnectionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CIPEncapsulationConnectionRequestBuilder
	// Build builds the CIPEncapsulationConnectionRequest or returns an error if something is wrong
	Build() (CIPEncapsulationConnectionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CIPEncapsulationConnectionRequest
}

// NewCIPEncapsulationConnectionRequestBuilder() creates a CIPEncapsulationConnectionRequestBuilder
func NewCIPEncapsulationConnectionRequestBuilder() CIPEncapsulationConnectionRequestBuilder {
	return &_CIPEncapsulationConnectionRequestBuilder{_CIPEncapsulationConnectionRequest: new(_CIPEncapsulationConnectionRequest)}
}

type _CIPEncapsulationConnectionRequestBuilder struct {
	*_CIPEncapsulationConnectionRequest

	parentBuilder *_CIPEncapsulationPacketBuilder

	err *utils.MultiError
}

var _ (CIPEncapsulationConnectionRequestBuilder) = (*_CIPEncapsulationConnectionRequestBuilder)(nil)

func (b *_CIPEncapsulationConnectionRequestBuilder) setParent(contract CIPEncapsulationPacketContract) {
	b.CIPEncapsulationPacketContract = contract
}

func (b *_CIPEncapsulationConnectionRequestBuilder) WithMandatoryFields() CIPEncapsulationConnectionRequestBuilder {
	return b
}

func (b *_CIPEncapsulationConnectionRequestBuilder) Build() (CIPEncapsulationConnectionRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CIPEncapsulationConnectionRequest.deepCopy(), nil
}

func (b *_CIPEncapsulationConnectionRequestBuilder) MustBuild() CIPEncapsulationConnectionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CIPEncapsulationConnectionRequestBuilder) Done() CIPEncapsulationPacketBuilder {
	return b.parentBuilder
}

func (b *_CIPEncapsulationConnectionRequestBuilder) buildForCIPEncapsulationPacket() (CIPEncapsulationPacket, error) {
	return b.Build()
}

func (b *_CIPEncapsulationConnectionRequestBuilder) DeepCopy() any {
	_copy := b.CreateCIPEncapsulationConnectionRequestBuilder().(*_CIPEncapsulationConnectionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCIPEncapsulationConnectionRequestBuilder creates a CIPEncapsulationConnectionRequestBuilder
func (b *_CIPEncapsulationConnectionRequest) CreateCIPEncapsulationConnectionRequestBuilder() CIPEncapsulationConnectionRequestBuilder {
	if b == nil {
		return NewCIPEncapsulationConnectionRequestBuilder()
	}
	return &_CIPEncapsulationConnectionRequestBuilder{_CIPEncapsulationConnectionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationConnectionRequest) GetCommandType() uint16 {
	return 0x0101
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationConnectionRequest) GetParent() CIPEncapsulationPacketContract {
	return m.CIPEncapsulationPacketContract
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationConnectionRequest(structType any) CIPEncapsulationConnectionRequest {
	if casted, ok := structType.(CIPEncapsulationConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationConnectionRequest) GetTypeName() string {
	return "CIPEncapsulationConnectionRequest"
}

func (m *_CIPEncapsulationConnectionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_CIPEncapsulationConnectionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CIPEncapsulationConnectionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CIPEncapsulationPacket) (__cIPEncapsulationConnectionRequest CIPEncapsulationConnectionRequest, err error) {
	m.CIPEncapsulationPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CIPEncapsulationConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationConnectionRequest")
	}

	return m, nil
}

func (m *_CIPEncapsulationConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPEncapsulationConnectionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationConnectionRequest")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationConnectionRequest")
		}
		return nil
	}
	return m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CIPEncapsulationConnectionRequest) IsCIPEncapsulationConnectionRequest() {}

func (m *_CIPEncapsulationConnectionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CIPEncapsulationConnectionRequest) deepCopy() *_CIPEncapsulationConnectionRequest {
	if m == nil {
		return nil
	}
	_CIPEncapsulationConnectionRequestCopy := &_CIPEncapsulationConnectionRequest{
		m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket).deepCopy(),
	}
	m.CIPEncapsulationPacketContract.(*_CIPEncapsulationPacket)._SubType = m
	return _CIPEncapsulationConnectionRequestCopy
}

func (m *_CIPEncapsulationConnectionRequest) String() string {
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
