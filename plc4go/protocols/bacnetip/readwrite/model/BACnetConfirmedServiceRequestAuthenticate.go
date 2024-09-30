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

// BACnetConfirmedServiceRequestAuthenticate is the corresponding interface of BACnetConfirmedServiceRequestAuthenticate
type BACnetConfirmedServiceRequestAuthenticate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
	// IsBACnetConfirmedServiceRequestAuthenticate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestAuthenticate()
	// CreateBuilder creates a BACnetConfirmedServiceRequestAuthenticateBuilder
	CreateBACnetConfirmedServiceRequestAuthenticateBuilder() BACnetConfirmedServiceRequestAuthenticateBuilder
}

// _BACnetConfirmedServiceRequestAuthenticate is the data-structure of this message
type _BACnetConfirmedServiceRequestAuthenticate struct {
	BACnetConfirmedServiceRequestContract
	BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}

var _ BACnetConfirmedServiceRequestAuthenticate = (*_BACnetConfirmedServiceRequestAuthenticate)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestAuthenticate)(nil)

// NewBACnetConfirmedServiceRequestAuthenticate factory function for _BACnetConfirmedServiceRequestAuthenticate
func NewBACnetConfirmedServiceRequestAuthenticate(bytesOfRemovedService []byte, serviceRequestPayloadLength uint32, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestAuthenticate {
	_result := &_BACnetConfirmedServiceRequestAuthenticate{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		BytesOfRemovedService:                 bytesOfRemovedService,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestAuthenticateBuilder is a builder for BACnetConfirmedServiceRequestAuthenticate
type BACnetConfirmedServiceRequestAuthenticateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestAuthenticateBuilder
	// WithBytesOfRemovedService adds BytesOfRemovedService (property field)
	WithBytesOfRemovedService(...byte) BACnetConfirmedServiceRequestAuthenticateBuilder
	// Build builds the BACnetConfirmedServiceRequestAuthenticate or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestAuthenticate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestAuthenticate
}

// NewBACnetConfirmedServiceRequestAuthenticateBuilder() creates a BACnetConfirmedServiceRequestAuthenticateBuilder
func NewBACnetConfirmedServiceRequestAuthenticateBuilder() BACnetConfirmedServiceRequestAuthenticateBuilder {
	return &_BACnetConfirmedServiceRequestAuthenticateBuilder{_BACnetConfirmedServiceRequestAuthenticate: new(_BACnetConfirmedServiceRequestAuthenticate)}
}

type _BACnetConfirmedServiceRequestAuthenticateBuilder struct {
	*_BACnetConfirmedServiceRequestAuthenticate

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestAuthenticateBuilder) = (*_BACnetConfirmedServiceRequestAuthenticateBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) WithMandatoryFields(bytesOfRemovedService []byte) BACnetConfirmedServiceRequestAuthenticateBuilder {
	return b.WithBytesOfRemovedService(bytesOfRemovedService...)
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) WithBytesOfRemovedService(bytesOfRemovedService ...byte) BACnetConfirmedServiceRequestAuthenticateBuilder {
	b.BytesOfRemovedService = bytesOfRemovedService
	return b
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) Build() (BACnetConfirmedServiceRequestAuthenticate, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestAuthenticate.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) MustBuild() BACnetConfirmedServiceRequestAuthenticate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestAuthenticateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestAuthenticateBuilder().(*_BACnetConfirmedServiceRequestAuthenticateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestAuthenticateBuilder creates a BACnetConfirmedServiceRequestAuthenticateBuilder
func (b *_BACnetConfirmedServiceRequestAuthenticate) CreateBACnetConfirmedServiceRequestAuthenticateBuilder() BACnetConfirmedServiceRequestAuthenticateBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestAuthenticateBuilder()
	}
	return &_BACnetConfirmedServiceRequestAuthenticateBuilder{_BACnetConfirmedServiceRequestAuthenticate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_AUTHENTICATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAuthenticate(structType any) BACnetConfirmedServiceRequestAuthenticate {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAuthenticate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAuthenticate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAuthenticate"
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestPayloadLength uint32, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestAuthenticate BACnetConfirmedServiceRequestAuthenticate, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAuthenticate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAuthenticate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bytesOfRemovedService, err := readBuffer.ReadByteArray("bytesOfRemovedService", int(serviceRequestPayloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bytesOfRemovedService' field"))
	}
	m.BytesOfRemovedService = bytesOfRemovedService

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAuthenticate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAuthenticate")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAuthenticate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAuthenticate")
		}

		if err := WriteByteArrayField(ctx, "bytesOfRemovedService", m.GetBytesOfRemovedService(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAuthenticate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAuthenticate")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestAuthenticate) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestAuthenticate) IsBACnetConfirmedServiceRequestAuthenticate() {}

func (m *_BACnetConfirmedServiceRequestAuthenticate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) deepCopy() *_BACnetConfirmedServiceRequestAuthenticate {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestAuthenticateCopy := &_BACnetConfirmedServiceRequestAuthenticate{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.BytesOfRemovedService),
		m.ServiceRequestPayloadLength,
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestAuthenticateCopy
}

func (m *_BACnetConfirmedServiceRequestAuthenticate) String() string {
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
