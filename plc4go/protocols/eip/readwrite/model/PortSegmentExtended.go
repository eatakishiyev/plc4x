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

// PortSegmentExtended is the corresponding interface of PortSegmentExtended
type PortSegmentExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	PortSegmentType
	// GetPort returns Port (property field)
	GetPort() uint8
	// GetLinkAddressSize returns LinkAddressSize (property field)
	GetLinkAddressSize() uint8
	// GetAddress returns Address (property field)
	GetAddress() string
	// GetPaddingByte returns PaddingByte (virtual field)
	GetPaddingByte() uint8
	// IsPortSegmentExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPortSegmentExtended()
	// CreateBuilder creates a PortSegmentExtendedBuilder
	CreatePortSegmentExtendedBuilder() PortSegmentExtendedBuilder
}

// _PortSegmentExtended is the data-structure of this message
type _PortSegmentExtended struct {
	PortSegmentTypeContract
	Port            uint8
	LinkAddressSize uint8
	Address         string
}

var _ PortSegmentExtended = (*_PortSegmentExtended)(nil)
var _ PortSegmentTypeRequirements = (*_PortSegmentExtended)(nil)

// NewPortSegmentExtended factory function for _PortSegmentExtended
func NewPortSegmentExtended(port uint8, linkAddressSize uint8, address string) *_PortSegmentExtended {
	_result := &_PortSegmentExtended{
		PortSegmentTypeContract: NewPortSegmentType(),
		Port:                    port,
		LinkAddressSize:         linkAddressSize,
		Address:                 address,
	}
	_result.PortSegmentTypeContract.(*_PortSegmentType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PortSegmentExtendedBuilder is a builder for PortSegmentExtended
type PortSegmentExtendedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(port uint8, linkAddressSize uint8, address string) PortSegmentExtendedBuilder
	// WithPort adds Port (property field)
	WithPort(uint8) PortSegmentExtendedBuilder
	// WithLinkAddressSize adds LinkAddressSize (property field)
	WithLinkAddressSize(uint8) PortSegmentExtendedBuilder
	// WithAddress adds Address (property field)
	WithAddress(string) PortSegmentExtendedBuilder
	// Build builds the PortSegmentExtended or returns an error if something is wrong
	Build() (PortSegmentExtended, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PortSegmentExtended
}

// NewPortSegmentExtendedBuilder() creates a PortSegmentExtendedBuilder
func NewPortSegmentExtendedBuilder() PortSegmentExtendedBuilder {
	return &_PortSegmentExtendedBuilder{_PortSegmentExtended: new(_PortSegmentExtended)}
}

type _PortSegmentExtendedBuilder struct {
	*_PortSegmentExtended

	parentBuilder *_PortSegmentTypeBuilder

	err *utils.MultiError
}

var _ (PortSegmentExtendedBuilder) = (*_PortSegmentExtendedBuilder)(nil)

func (b *_PortSegmentExtendedBuilder) setParent(contract PortSegmentTypeContract) {
	b.PortSegmentTypeContract = contract
}

func (b *_PortSegmentExtendedBuilder) WithMandatoryFields(port uint8, linkAddressSize uint8, address string) PortSegmentExtendedBuilder {
	return b.WithPort(port).WithLinkAddressSize(linkAddressSize).WithAddress(address)
}

func (b *_PortSegmentExtendedBuilder) WithPort(port uint8) PortSegmentExtendedBuilder {
	b.Port = port
	return b
}

func (b *_PortSegmentExtendedBuilder) WithLinkAddressSize(linkAddressSize uint8) PortSegmentExtendedBuilder {
	b.LinkAddressSize = linkAddressSize
	return b
}

func (b *_PortSegmentExtendedBuilder) WithAddress(address string) PortSegmentExtendedBuilder {
	b.Address = address
	return b
}

func (b *_PortSegmentExtendedBuilder) Build() (PortSegmentExtended, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PortSegmentExtended.deepCopy(), nil
}

func (b *_PortSegmentExtendedBuilder) MustBuild() PortSegmentExtended {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_PortSegmentExtendedBuilder) Done() PortSegmentTypeBuilder {
	return b.parentBuilder
}

func (b *_PortSegmentExtendedBuilder) buildForPortSegmentType() (PortSegmentType, error) {
	return b.Build()
}

func (b *_PortSegmentExtendedBuilder) DeepCopy() any {
	_copy := b.CreatePortSegmentExtendedBuilder().(*_PortSegmentExtendedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePortSegmentExtendedBuilder creates a PortSegmentExtendedBuilder
func (b *_PortSegmentExtended) CreatePortSegmentExtendedBuilder() PortSegmentExtendedBuilder {
	if b == nil {
		return NewPortSegmentExtendedBuilder()
	}
	return &_PortSegmentExtendedBuilder{_PortSegmentExtended: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PortSegmentExtended) GetExtendedLinkAddress() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PortSegmentExtended) GetParent() PortSegmentTypeContract {
	return m.PortSegmentTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PortSegmentExtended) GetPort() uint8 {
	return m.Port
}

func (m *_PortSegmentExtended) GetLinkAddressSize() uint8 {
	return m.LinkAddressSize
}

func (m *_PortSegmentExtended) GetAddress() string {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_PortSegmentExtended) GetPaddingByte() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(uint8(m.GetLinkAddressSize()) % uint8(uint8(2)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPortSegmentExtended(structType any) PortSegmentExtended {
	if casted, ok := structType.(PortSegmentExtended); ok {
		return casted
	}
	if casted, ok := structType.(*PortSegmentExtended); ok {
		return *casted
	}
	return nil
}

func (m *_PortSegmentExtended) GetTypeName() string {
	return "PortSegmentExtended"
}

func (m *_PortSegmentExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.PortSegmentTypeContract.(*_PortSegmentType).getLengthInBits(ctx))

	// Simple field (port)
	lengthInBits += 4

	// Simple field (linkAddressSize)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (address)
	lengthInBits += uint16(int32((int32(m.GetLinkAddressSize()) * int32(int32(8)))) + int32((int32(m.GetPaddingByte()) * int32(int32(8)))))

	return lengthInBits
}

func (m *_PortSegmentExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PortSegmentExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_PortSegmentType) (__portSegmentExtended PortSegmentExtended, err error) {
	m.PortSegmentTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PortSegmentExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PortSegmentExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	port, err := ReadSimpleField(ctx, "port", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'port' field"))
	}
	m.Port = port

	linkAddressSize, err := ReadSimpleField(ctx, "linkAddressSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkAddressSize' field"))
	}
	m.LinkAddressSize = linkAddressSize

	paddingByte, err := ReadVirtualField[uint8](ctx, "paddingByte", (*uint8)(nil), uint8(linkAddressSize)%uint8(uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'paddingByte' field"))
	}
	_ = paddingByte

	address, err := ReadSimpleField(ctx, "address", ReadString(readBuffer, uint32(int32((int32(linkAddressSize)*int32(int32(8))))+int32((int32(paddingByte)*int32(int32(8)))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("PortSegmentExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PortSegmentExtended")
	}

	return m, nil
}

func (m *_PortSegmentExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PortSegmentExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PortSegmentExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PortSegmentExtended")
		}

		if err := WriteSimpleField[uint8](ctx, "port", m.GetPort(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'port' field")
		}

		if err := WriteSimpleField[uint8](ctx, "linkAddressSize", m.GetLinkAddressSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'linkAddressSize' field")
		}
		// Virtual field
		paddingByte := m.GetPaddingByte()
		_ = paddingByte
		if _paddingByteErr := writeBuffer.WriteVirtual(ctx, "paddingByte", m.GetPaddingByte()); _paddingByteErr != nil {
			return errors.Wrap(_paddingByteErr, "Error serializing 'paddingByte' field")
		}

		if err := WriteSimpleField[string](ctx, "address", m.GetAddress(), WriteString(writeBuffer, int32(int32((int32(m.GetLinkAddressSize())*int32(int32(8))))+int32((int32(m.GetPaddingByte())*int32(int32(8))))))); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("PortSegmentExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PortSegmentExtended")
		}
		return nil
	}
	return m.PortSegmentTypeContract.(*_PortSegmentType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PortSegmentExtended) IsPortSegmentExtended() {}

func (m *_PortSegmentExtended) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PortSegmentExtended) deepCopy() *_PortSegmentExtended {
	if m == nil {
		return nil
	}
	_PortSegmentExtendedCopy := &_PortSegmentExtended{
		m.PortSegmentTypeContract.(*_PortSegmentType).deepCopy(),
		m.Port,
		m.LinkAddressSize,
		m.Address,
	}
	m.PortSegmentTypeContract.(*_PortSegmentType)._SubType = m
	return _PortSegmentExtendedCopy
}

func (m *_PortSegmentExtended) String() string {
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
