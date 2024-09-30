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

// BACnetConstructedDataIPv6DHCPServer is the corresponding interface of BACnetConstructedDataIPv6DHCPServer
type BACnetConstructedDataIPv6DHCPServer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDhcpServer returns DhcpServer (property field)
	GetDhcpServer() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
	// IsBACnetConstructedDataIPv6DHCPServer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6DHCPServer()
	// CreateBuilder creates a BACnetConstructedDataIPv6DHCPServerBuilder
	CreateBACnetConstructedDataIPv6DHCPServerBuilder() BACnetConstructedDataIPv6DHCPServerBuilder
}

// _BACnetConstructedDataIPv6DHCPServer is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPServer struct {
	BACnetConstructedDataContract
	DhcpServer BACnetApplicationTagOctetString
}

var _ BACnetConstructedDataIPv6DHCPServer = (*_BACnetConstructedDataIPv6DHCPServer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6DHCPServer)(nil)

// NewBACnetConstructedDataIPv6DHCPServer factory function for _BACnetConstructedDataIPv6DHCPServer
func NewBACnetConstructedDataIPv6DHCPServer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, dhcpServer BACnetApplicationTagOctetString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPServer {
	if dhcpServer == nil {
		panic("dhcpServer of type BACnetApplicationTagOctetString for BACnetConstructedDataIPv6DHCPServer must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6DHCPServer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DhcpServer:                    dhcpServer,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6DHCPServerBuilder is a builder for BACnetConstructedDataIPv6DHCPServer
type BACnetConstructedDataIPv6DHCPServerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DHCPServerBuilder
	// WithDhcpServer adds DhcpServer (property field)
	WithDhcpServer(BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DHCPServerBuilder
	// WithDhcpServerBuilder adds DhcpServer (property field) which is build by the builder
	WithDhcpServerBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6DHCPServerBuilder
	// Build builds the BACnetConstructedDataIPv6DHCPServer or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6DHCPServer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6DHCPServer
}

// NewBACnetConstructedDataIPv6DHCPServerBuilder() creates a BACnetConstructedDataIPv6DHCPServerBuilder
func NewBACnetConstructedDataIPv6DHCPServerBuilder() BACnetConstructedDataIPv6DHCPServerBuilder {
	return &_BACnetConstructedDataIPv6DHCPServerBuilder{_BACnetConstructedDataIPv6DHCPServer: new(_BACnetConstructedDataIPv6DHCPServer)}
}

type _BACnetConstructedDataIPv6DHCPServerBuilder struct {
	*_BACnetConstructedDataIPv6DHCPServer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6DHCPServerBuilder) = (*_BACnetConstructedDataIPv6DHCPServerBuilder)(nil)

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) WithMandatoryFields(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DHCPServerBuilder {
	return b.WithDhcpServer(dhcpServer)
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) WithDhcpServer(dhcpServer BACnetApplicationTagOctetString) BACnetConstructedDataIPv6DHCPServerBuilder {
	b.DhcpServer = dhcpServer
	return b
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) WithDhcpServerBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetConstructedDataIPv6DHCPServerBuilder {
	builder := builderSupplier(b.DhcpServer.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.DhcpServer, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) Build() (BACnetConstructedDataIPv6DHCPServer, error) {
	if b.DhcpServer == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dhcpServer' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIPv6DHCPServer.deepCopy(), nil
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) MustBuild() BACnetConstructedDataIPv6DHCPServer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIPv6DHCPServerBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIPv6DHCPServerBuilder().(*_BACnetConstructedDataIPv6DHCPServerBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIPv6DHCPServerBuilder creates a BACnetConstructedDataIPv6DHCPServerBuilder
func (b *_BACnetConstructedDataIPv6DHCPServer) CreateBACnetConstructedDataIPv6DHCPServerBuilder() BACnetConstructedDataIPv6DHCPServerBuilder {
	if b == nil {
		return NewBACnetConstructedDataIPv6DHCPServerBuilder()
	}
	return &_BACnetConstructedDataIPv6DHCPServerBuilder{_BACnetConstructedDataIPv6DHCPServer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_SERVER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetDhcpServer() BACnetApplicationTagOctetString {
	return m.DhcpServer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPServer) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetDhcpServer())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPServer(structType any) BACnetConstructedDataIPv6DHCPServer {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPServer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPServer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPServer"
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (dhcpServer)
	lengthInBits += m.DhcpServer.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPServer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6DHCPServer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6DHCPServer BACnetConstructedDataIPv6DHCPServer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dhcpServer, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "dhcpServer", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dhcpServer' field"))
	}
	m.DhcpServer = dhcpServer

	actualValue, err := ReadVirtualField[BACnetApplicationTagOctetString](ctx, "actualValue", (*BACnetApplicationTagOctetString)(nil), dhcpServer)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPServer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPServer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPServer")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "dhcpServer", m.GetDhcpServer(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dhcpServer' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPServer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPServer) IsBACnetConstructedDataIPv6DHCPServer() {}

func (m *_BACnetConstructedDataIPv6DHCPServer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6DHCPServer) deepCopy() *_BACnetConstructedDataIPv6DHCPServer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6DHCPServerCopy := &_BACnetConstructedDataIPv6DHCPServer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DhcpServer.DeepCopy().(BACnetApplicationTagOctetString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6DHCPServerCopy
}

func (m *_BACnetConstructedDataIPv6DHCPServer) String() string {
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
