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

// BACnetConstructedDataIPv6DHCPLeaseTime is the corresponding interface of BACnetConstructedDataIPv6DHCPLeaseTime
type BACnetConstructedDataIPv6DHCPLeaseTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6DhcpLeaseTime returns Ipv6DhcpLeaseTime (property field)
	GetIpv6DhcpLeaseTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIPv6DHCPLeaseTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6DHCPLeaseTime()
	// CreateBuilder creates a BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
	CreateBACnetConstructedDataIPv6DHCPLeaseTimeBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
}

// _BACnetConstructedDataIPv6DHCPLeaseTime is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPLeaseTime struct {
	BACnetConstructedDataContract
	Ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIPv6DHCPLeaseTime = (*_BACnetConstructedDataIPv6DHCPLeaseTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6DHCPLeaseTime)(nil)

// NewBACnetConstructedDataIPv6DHCPLeaseTime factory function for _BACnetConstructedDataIPv6DHCPLeaseTime
func NewBACnetConstructedDataIPv6DHCPLeaseTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPLeaseTime {
	if ipv6DhcpLeaseTime == nil {
		panic("ipv6DhcpLeaseTime of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIPv6DHCPLeaseTime must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6DHCPLeaseTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6DhcpLeaseTime:             ipv6DhcpLeaseTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6DHCPLeaseTimeBuilder is a builder for BACnetConstructedDataIPv6DHCPLeaseTime
type BACnetConstructedDataIPv6DHCPLeaseTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
	// WithIpv6DhcpLeaseTime adds Ipv6DhcpLeaseTime (property field)
	WithIpv6DhcpLeaseTime(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
	// WithIpv6DhcpLeaseTimeBuilder adds Ipv6DhcpLeaseTime (property field) which is build by the builder
	WithIpv6DhcpLeaseTimeBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
	// Build builds the BACnetConstructedDataIPv6DHCPLeaseTime or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6DHCPLeaseTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6DHCPLeaseTime
}

// NewBACnetConstructedDataIPv6DHCPLeaseTimeBuilder() creates a BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
func NewBACnetConstructedDataIPv6DHCPLeaseTimeBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeBuilder {
	return &_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder{_BACnetConstructedDataIPv6DHCPLeaseTime: new(_BACnetConstructedDataIPv6DHCPLeaseTime)}
}

type _BACnetConstructedDataIPv6DHCPLeaseTimeBuilder struct {
	*_BACnetConstructedDataIPv6DHCPLeaseTime

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) = (*_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder)(nil)

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) WithMandatoryFields(ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder {
	return b.WithIpv6DhcpLeaseTime(ipv6DhcpLeaseTime)
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) WithIpv6DhcpLeaseTime(ipv6DhcpLeaseTime BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder {
	b.Ipv6DhcpLeaseTime = ipv6DhcpLeaseTime
	return b
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) WithIpv6DhcpLeaseTimeBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DHCPLeaseTimeBuilder {
	builder := builderSupplier(b.Ipv6DhcpLeaseTime.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.Ipv6DhcpLeaseTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) Build() (BACnetConstructedDataIPv6DHCPLeaseTime, error) {
	if b.Ipv6DhcpLeaseTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'ipv6DhcpLeaseTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6DHCPLeaseTime.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) MustBuild() BACnetConstructedDataIPv6DHCPLeaseTime {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6DHCPLeaseTimeBuilder().(*_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6DHCPLeaseTimeBuilder creates a BACnetConstructedDataIPv6DHCPLeaseTimeBuilder
func (b *_BACnetConstructedDataIPv6DHCPLeaseTime) CreateBACnetConstructedDataIPv6DHCPLeaseTimeBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6DHCPLeaseTimeBuilder()
	}
	return &_BACnetConstructedDataIPv6DHCPLeaseTimeBuilder{_BACnetConstructedDataIPv6DHCPLeaseTime: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_LEASE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetIpv6DhcpLeaseTime() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6DhcpLeaseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6DhcpLeaseTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPLeaseTime(structType any) BACnetConstructedDataIPv6DHCPLeaseTime {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPLeaseTime"
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6DhcpLeaseTime)
	lengthInBits += m.Ipv6DhcpLeaseTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6DHCPLeaseTime BACnetConstructedDataIPv6DHCPLeaseTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPLeaseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6DhcpLeaseTime, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6DhcpLeaseTime", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6DhcpLeaseTime' field"))
	}
	m.Ipv6DhcpLeaseTime = ipv6DhcpLeaseTime

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipv6DhcpLeaseTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPLeaseTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPLeaseTime")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6DhcpLeaseTime", m.GetIpv6DhcpLeaseTime(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6DhcpLeaseTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPLeaseTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) IsBACnetConstructedDataIPv6DHCPLeaseTime() {}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) deepCopy() *_BACnetConstructedDataIPv6DHCPLeaseTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6DHCPLeaseTimeCopy := &_BACnetConstructedDataIPv6DHCPLeaseTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Ipv6DhcpLeaseTime.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6DHCPLeaseTimeCopy
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTime) String() string {
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
