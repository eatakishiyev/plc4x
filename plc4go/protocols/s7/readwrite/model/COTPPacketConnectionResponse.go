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

// COTPPacketConnectionResponse is the corresponding interface of COTPPacketConnectionResponse
type COTPPacketConnectionResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
	// IsCOTPPacketConnectionResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPPacketConnectionResponse()
	// CreateBuilder creates a COTPPacketConnectionResponseBuilder
	CreateCOTPPacketConnectionResponseBuilder() COTPPacketConnectionResponseBuilder
}

// _COTPPacketConnectionResponse is the data-structure of this message
type _COTPPacketConnectionResponse struct {
	COTPPacketContract
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass
}

var _ COTPPacketConnectionResponse = (*_COTPPacketConnectionResponse)(nil)
var _ COTPPacketRequirements = (*_COTPPacketConnectionResponse)(nil)

// NewCOTPPacketConnectionResponse factory function for _COTPPacketConnectionResponse
func NewCOTPPacketConnectionResponse(parameters []COTPParameter, payload S7Message, destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, cotpLen uint16) *_COTPPacketConnectionResponse {
	_result := &_COTPPacketConnectionResponse{
		COTPPacketContract:   NewCOTPPacket(parameters, payload, cotpLen),
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
	}
	_result.COTPPacketContract.(*_COTPPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPPacketConnectionResponseBuilder is a builder for COTPPacketConnectionResponse
type COTPPacketConnectionResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketConnectionResponseBuilder
	// WithDestinationReference adds DestinationReference (property field)
	WithDestinationReference(uint16) COTPPacketConnectionResponseBuilder
	// WithSourceReference adds SourceReference (property field)
	WithSourceReference(uint16) COTPPacketConnectionResponseBuilder
	// WithProtocolClass adds ProtocolClass (property field)
	WithProtocolClass(COTPProtocolClass) COTPPacketConnectionResponseBuilder
	// Build builds the COTPPacketConnectionResponse or returns an error if something is wrong
	Build() (COTPPacketConnectionResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPPacketConnectionResponse
}

// NewCOTPPacketConnectionResponseBuilder() creates a COTPPacketConnectionResponseBuilder
func NewCOTPPacketConnectionResponseBuilder() COTPPacketConnectionResponseBuilder {
	return &_COTPPacketConnectionResponseBuilder{_COTPPacketConnectionResponse: new(_COTPPacketConnectionResponse)}
}

type _COTPPacketConnectionResponseBuilder struct {
	*_COTPPacketConnectionResponse

	parentBuilder *_COTPPacketBuilder

	err *utils.MultiError
}

var _ (COTPPacketConnectionResponseBuilder) = (*_COTPPacketConnectionResponseBuilder)(nil)

func (b *_COTPPacketConnectionResponseBuilder) setParent(contract COTPPacketContract) {
	b.COTPPacketContract = contract
}

func (b *_COTPPacketConnectionResponseBuilder) WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketConnectionResponseBuilder {
	return b.WithDestinationReference(destinationReference).WithSourceReference(sourceReference).WithProtocolClass(protocolClass)
}

func (b *_COTPPacketConnectionResponseBuilder) WithDestinationReference(destinationReference uint16) COTPPacketConnectionResponseBuilder {
	b.DestinationReference = destinationReference
	return b
}

func (b *_COTPPacketConnectionResponseBuilder) WithSourceReference(sourceReference uint16) COTPPacketConnectionResponseBuilder {
	b.SourceReference = sourceReference
	return b
}

func (b *_COTPPacketConnectionResponseBuilder) WithProtocolClass(protocolClass COTPProtocolClass) COTPPacketConnectionResponseBuilder {
	b.ProtocolClass = protocolClass
	return b
}

func (b *_COTPPacketConnectionResponseBuilder) Build() (COTPPacketConnectionResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPPacketConnectionResponse.deepCopy(), nil
}

func (b *_COTPPacketConnectionResponseBuilder) MustBuild() COTPPacketConnectionResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_COTPPacketConnectionResponseBuilder) Done() COTPPacketBuilder {
	return b.parentBuilder
}

func (b *_COTPPacketConnectionResponseBuilder) buildForCOTPPacket() (COTPPacket, error) {
	return b.Build()
}

func (b *_COTPPacketConnectionResponseBuilder) DeepCopy() any {
	_copy := b.CreateCOTPPacketConnectionResponseBuilder().(*_COTPPacketConnectionResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPPacketConnectionResponseBuilder creates a COTPPacketConnectionResponseBuilder
func (b *_COTPPacketConnectionResponse) CreateCOTPPacketConnectionResponseBuilder() COTPPacketConnectionResponseBuilder {
	if b == nil {
		return NewCOTPPacketConnectionResponseBuilder()
	}
	return &_COTPPacketConnectionResponseBuilder{_COTPPacketConnectionResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketConnectionResponse) GetTpduCode() uint8 {
	return 0xD0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketConnectionResponse) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketConnectionResponse) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketConnectionResponse) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *_COTPPacketConnectionResponse) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPPacketConnectionResponse(structType any) COTPPacketConnectionResponse {
	if casted, ok := structType.(COTPPacketConnectionResponse); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketConnectionResponse); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketConnectionResponse) GetTypeName() string {
	return "COTPPacketConnectionResponse"
}

func (m *_COTPPacketConnectionResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPPacketConnectionResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketConnectionResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketConnectionResponse COTPPacketConnectionResponse, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketConnectionResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketConnectionResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}
	m.DestinationReference = destinationReference

	sourceReference, err := ReadSimpleField(ctx, "sourceReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceReference' field"))
	}
	m.SourceReference = sourceReference

	protocolClass, err := ReadEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", ReadEnum(COTPProtocolClassByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolClass' field"))
	}
	m.ProtocolClass = protocolClass

	if closeErr := readBuffer.CloseContext("COTPPacketConnectionResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketConnectionResponse")
	}

	return m, nil
}

func (m *_COTPPacketConnectionResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketConnectionResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketConnectionResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketConnectionResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationReference", m.GetDestinationReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationReference' field")
		}

		if err := WriteSimpleField[uint16](ctx, "sourceReference", m.GetSourceReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceReference' field")
		}

		if err := WriteSimpleEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", m.GetProtocolClass(), WriteEnum[COTPProtocolClass, uint8](COTPProtocolClass.GetValue, COTPProtocolClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolClass' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketConnectionResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketConnectionResponse")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketConnectionResponse) IsCOTPPacketConnectionResponse() {}

func (m *_COTPPacketConnectionResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPPacketConnectionResponse) deepCopy() *_COTPPacketConnectionResponse {
	if m == nil {
		return nil
	}
	_COTPPacketConnectionResponseCopy := &_COTPPacketConnectionResponse{
		m.COTPPacketContract.(*_COTPPacket).deepCopy(),
		m.DestinationReference,
		m.SourceReference,
		m.ProtocolClass,
	}
	m.COTPPacketContract.(*_COTPPacket)._SubType = m
	return _COTPPacketConnectionResponseCopy
}

func (m *_COTPPacketConnectionResponse) String() string {
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
