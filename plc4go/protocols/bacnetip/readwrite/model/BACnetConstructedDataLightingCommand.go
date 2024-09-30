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

// BACnetConstructedDataLightingCommand is the corresponding interface of BACnetConstructedDataLightingCommand
type BACnetConstructedDataLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLightingCommand returns LightingCommand (property field)
	GetLightingCommand() BACnetLightingCommand
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingCommand
	// IsBACnetConstructedDataLightingCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLightingCommand()
	// CreateBuilder creates a BACnetConstructedDataLightingCommandBuilder
	CreateBACnetConstructedDataLightingCommandBuilder() BACnetConstructedDataLightingCommandBuilder
}

// _BACnetConstructedDataLightingCommand is the data-structure of this message
type _BACnetConstructedDataLightingCommand struct {
	BACnetConstructedDataContract
	LightingCommand BACnetLightingCommand
}

var _ BACnetConstructedDataLightingCommand = (*_BACnetConstructedDataLightingCommand)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLightingCommand)(nil)

// NewBACnetConstructedDataLightingCommand factory function for _BACnetConstructedDataLightingCommand
func NewBACnetConstructedDataLightingCommand(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lightingCommand BACnetLightingCommand, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingCommand {
	if lightingCommand == nil {
		panic("lightingCommand of type BACnetLightingCommand for BACnetConstructedDataLightingCommand must not be nil")
	}
	_result := &_BACnetConstructedDataLightingCommand{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LightingCommand:               lightingCommand,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLightingCommandBuilder is a builder for BACnetConstructedDataLightingCommand
type BACnetConstructedDataLightingCommandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lightingCommand BACnetLightingCommand) BACnetConstructedDataLightingCommandBuilder
	// WithLightingCommand adds LightingCommand (property field)
	WithLightingCommand(BACnetLightingCommand) BACnetConstructedDataLightingCommandBuilder
	// WithLightingCommandBuilder adds LightingCommand (property field) which is build by the builder
	WithLightingCommandBuilder(func(BACnetLightingCommandBuilder) BACnetLightingCommandBuilder) BACnetConstructedDataLightingCommandBuilder
	// Build builds the BACnetConstructedDataLightingCommand or returns an error if something is wrong
	Build() (BACnetConstructedDataLightingCommand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLightingCommand
}

// NewBACnetConstructedDataLightingCommandBuilder() creates a BACnetConstructedDataLightingCommandBuilder
func NewBACnetConstructedDataLightingCommandBuilder() BACnetConstructedDataLightingCommandBuilder {
	return &_BACnetConstructedDataLightingCommandBuilder{_BACnetConstructedDataLightingCommand: new(_BACnetConstructedDataLightingCommand)}
}

type _BACnetConstructedDataLightingCommandBuilder struct {
	*_BACnetConstructedDataLightingCommand

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLightingCommandBuilder) = (*_BACnetConstructedDataLightingCommandBuilder)(nil)

func (b *_BACnetConstructedDataLightingCommandBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLightingCommandBuilder) WithMandatoryFields(lightingCommand BACnetLightingCommand) BACnetConstructedDataLightingCommandBuilder {
	return b.WithLightingCommand(lightingCommand)
}

func (b *_BACnetConstructedDataLightingCommandBuilder) WithLightingCommand(lightingCommand BACnetLightingCommand) BACnetConstructedDataLightingCommandBuilder {
	b.LightingCommand = lightingCommand
	return b
}

func (b *_BACnetConstructedDataLightingCommandBuilder) WithLightingCommandBuilder(builderSupplier func(BACnetLightingCommandBuilder) BACnetLightingCommandBuilder) BACnetConstructedDataLightingCommandBuilder {
	builder := builderSupplier(b.LightingCommand.CreateBACnetLightingCommandBuilder())
	var err error
	b.LightingCommand, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingCommandBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLightingCommandBuilder) Build() (BACnetConstructedDataLightingCommand, error) {
	if b.LightingCommand == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lightingCommand' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLightingCommand.deepCopy(), nil
}

func (b *_BACnetConstructedDataLightingCommandBuilder) MustBuild() BACnetConstructedDataLightingCommand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLightingCommandBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLightingCommandBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLightingCommandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLightingCommandBuilder().(*_BACnetConstructedDataLightingCommandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLightingCommandBuilder creates a BACnetConstructedDataLightingCommandBuilder
func (b *_BACnetConstructedDataLightingCommand) CreateBACnetConstructedDataLightingCommandBuilder() BACnetConstructedDataLightingCommandBuilder {
	if b == nil {
		return NewBACnetConstructedDataLightingCommandBuilder()
	}
	return &_BACnetConstructedDataLightingCommandBuilder{_BACnetConstructedDataLightingCommand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLightingCommand) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIGHTING_COMMAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetLightingCommand() BACnetLightingCommand {
	return m.LightingCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLightingCommand) GetActualValue() BACnetLightingCommand {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLightingCommand(m.GetLightingCommand())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingCommand(structType any) BACnetConstructedDataLightingCommand {
	if casted, ok := structType.(BACnetConstructedDataLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingCommand) GetTypeName() string {
	return "BACnetConstructedDataLightingCommand"
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lightingCommand)
	lengthInBits += m.LightingCommand.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLightingCommand BACnetConstructedDataLightingCommand, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightingCommand, err := ReadSimpleField[BACnetLightingCommand](ctx, "lightingCommand", ReadComplex[BACnetLightingCommand](BACnetLightingCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingCommand' field"))
	}
	m.LightingCommand = lightingCommand

	actualValue, err := ReadVirtualField[BACnetLightingCommand](ctx, "actualValue", (*BACnetLightingCommand)(nil), lightingCommand)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingCommand")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommand](ctx, "lightingCommand", m.GetLightingCommand(), WriteComplex[BACnetLightingCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightingCommand' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingCommand")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingCommand) IsBACnetConstructedDataLightingCommand() {}

func (m *_BACnetConstructedDataLightingCommand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLightingCommand) deepCopy() *_BACnetConstructedDataLightingCommand {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLightingCommandCopy := &_BACnetConstructedDataLightingCommand{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.LightingCommand.DeepCopy().(BACnetLightingCommand),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLightingCommandCopy
}

func (m *_BACnetConstructedDataLightingCommand) String() string {
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
