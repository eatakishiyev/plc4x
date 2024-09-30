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

// BACnetConstructedDataBACnetIPv6UDPPort is the corresponding interface of BACnetConstructedDataBACnetIPv6UDPPort
type BACnetConstructedDataBACnetIPv6UDPPort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6UdpPort returns Ipv6UdpPort (property field)
	GetIpv6UdpPort() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataBACnetIPv6UDPPort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBACnetIPv6UDPPort()
	// CreateBuilder creates a BACnetConstructedDataBACnetIPv6UDPPortBuilder
	CreateBACnetConstructedDataBACnetIPv6UDPPortBuilder() BACnetConstructedDataBACnetIPv6UDPPortBuilder
}

// _BACnetConstructedDataBACnetIPv6UDPPort is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6UDPPort struct {
	BACnetConstructedDataContract
	Ipv6UdpPort BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataBACnetIPv6UDPPort = (*_BACnetConstructedDataBACnetIPv6UDPPort)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBACnetIPv6UDPPort)(nil)

// NewBACnetConstructedDataBACnetIPv6UDPPort factory function for _BACnetConstructedDataBACnetIPv6UDPPort
func NewBACnetConstructedDataBACnetIPv6UDPPort(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6UdpPort BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPv6UDPPort {
	if ipv6UdpPort == nil {
		panic("ipv6UdpPort of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataBACnetIPv6UDPPort must not be nil")
	}
	_result := &_BACnetConstructedDataBACnetIPv6UDPPort{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6UdpPort:                   ipv6UdpPort,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBACnetIPv6UDPPortBuilder is a builder for BACnetConstructedDataBACnetIPv6UDPPort
type BACnetConstructedDataBACnetIPv6UDPPortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6UdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPv6UDPPortBuilder
	// WithIpv6UdpPort adds Ipv6UdpPort (property field)
	WithIpv6UdpPort(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPv6UDPPortBuilder
	// WithIpv6UdpPortBuilder adds Ipv6UdpPort (property field) which is build by the builder
	WithIpv6UdpPortBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBACnetIPv6UDPPortBuilder
	// Build builds the BACnetConstructedDataBACnetIPv6UDPPort or returns an error if something is wrong
	Build() (BACnetConstructedDataBACnetIPv6UDPPort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBACnetIPv6UDPPort
}

// NewBACnetConstructedDataBACnetIPv6UDPPortBuilder() creates a BACnetConstructedDataBACnetIPv6UDPPortBuilder
func NewBACnetConstructedDataBACnetIPv6UDPPortBuilder() BACnetConstructedDataBACnetIPv6UDPPortBuilder {
	return &_BACnetConstructedDataBACnetIPv6UDPPortBuilder{_BACnetConstructedDataBACnetIPv6UDPPort: new(_BACnetConstructedDataBACnetIPv6UDPPort)}
}

type _BACnetConstructedDataBACnetIPv6UDPPortBuilder struct {
	*_BACnetConstructedDataBACnetIPv6UDPPort

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBACnetIPv6UDPPortBuilder) = (*_BACnetConstructedDataBACnetIPv6UDPPortBuilder)(nil)

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) WithMandatoryFields(ipv6UdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPv6UDPPortBuilder {
	return b.WithIpv6UdpPort(ipv6UdpPort)
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) WithIpv6UdpPort(ipv6UdpPort BACnetApplicationTagUnsignedInteger) BACnetConstructedDataBACnetIPv6UDPPortBuilder {
	b.Ipv6UdpPort = ipv6UdpPort
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) WithIpv6UdpPortBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataBACnetIPv6UDPPortBuilder {
	builder := builderSupplier(b.Ipv6UdpPort.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.Ipv6UdpPort, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) Build() (BACnetConstructedDataBACnetIPv6UDPPort, error) {
	if b.Ipv6UdpPort == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6UdpPort' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBACnetIPv6UDPPort.deepCopy(), nil
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) MustBuild() BACnetConstructedDataBACnetIPv6UDPPort {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBACnetIPv6UDPPortBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBACnetIPv6UDPPortBuilder().(*_BACnetConstructedDataBACnetIPv6UDPPortBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBACnetIPv6UDPPortBuilder creates a BACnetConstructedDataBACnetIPv6UDPPortBuilder
func (b *_BACnetConstructedDataBACnetIPv6UDPPort) CreateBACnetConstructedDataBACnetIPv6UDPPortBuilder() BACnetConstructedDataBACnetIPv6UDPPortBuilder {
	if b == nil {
		return NewBACnetConstructedDataBACnetIPv6UDPPortBuilder()
	}
	return &_BACnetConstructedDataBACnetIPv6UDPPortBuilder{_BACnetConstructedDataBACnetIPv6UDPPort: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_UDP_PORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetIpv6UdpPort() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6UdpPort
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6UdpPort())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6UDPPort(structType any) BACnetConstructedDataBACnetIPv6UDPPort {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6UDPPort); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6UDPPort"
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6UdpPort)
	lengthInBits += m.Ipv6UdpPort.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBACnetIPv6UDPPort BACnetConstructedDataBACnetIPv6UDPPort, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6UDPPort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6UDPPort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6UdpPort, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6UdpPort", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6UdpPort' field"))
	}
	m.Ipv6UdpPort = ipv6UdpPort

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipv6UdpPort)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6UDPPort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6UDPPort")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6UDPPort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6UDPPort")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6UdpPort", m.GetIpv6UdpPort(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6UdpPort' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6UDPPort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6UDPPort")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) IsBACnetConstructedDataBACnetIPv6UDPPort() {}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) deepCopy() *_BACnetConstructedDataBACnetIPv6UDPPort {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBACnetIPv6UDPPortCopy := &_BACnetConstructedDataBACnetIPv6UDPPort{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Ipv6UdpPort.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBACnetIPv6UDPPortCopy
}

func (m *_BACnetConstructedDataBACnetIPv6UDPPort) String() string {
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
