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

// BACnetConstructedDataActionText is the corresponding interface of BACnetConstructedDataActionText
type BACnetConstructedDataActionText interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetActionText returns ActionText (property field)
	GetActionText() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataActionText is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataActionText()
	// CreateBuilder creates a BACnetConstructedDataActionTextBuilder
	CreateBACnetConstructedDataActionTextBuilder() BACnetConstructedDataActionTextBuilder
}

// _BACnetConstructedDataActionText is the data-structure of this message
type _BACnetConstructedDataActionText struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ActionText           []BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataActionText = (*_BACnetConstructedDataActionText)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataActionText)(nil)

// NewBACnetConstructedDataActionText factory function for _BACnetConstructedDataActionText
func NewBACnetConstructedDataActionText(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, actionText []BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActionText {
	_result := &_BACnetConstructedDataActionText{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ActionText:                    actionText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataActionTextBuilder is a builder for BACnetConstructedDataActionText
type BACnetConstructedDataActionTextBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(actionText []BACnetApplicationTagCharacterString) BACnetConstructedDataActionTextBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataActionTextBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataActionTextBuilder
	// WithActionText adds ActionText (property field)
	WithActionText(...BACnetApplicationTagCharacterString) BACnetConstructedDataActionTextBuilder
	// Build builds the BACnetConstructedDataActionText or returns an error if something is wrong
	Build() (BACnetConstructedDataActionText, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataActionText
}

// NewBACnetConstructedDataActionTextBuilder() creates a BACnetConstructedDataActionTextBuilder
func NewBACnetConstructedDataActionTextBuilder() BACnetConstructedDataActionTextBuilder {
	return &_BACnetConstructedDataActionTextBuilder{_BACnetConstructedDataActionText: new(_BACnetConstructedDataActionText)}
}

type _BACnetConstructedDataActionTextBuilder struct {
	*_BACnetConstructedDataActionText

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataActionTextBuilder) = (*_BACnetConstructedDataActionTextBuilder)(nil)

func (b *_BACnetConstructedDataActionTextBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataActionTextBuilder) WithMandatoryFields(actionText []BACnetApplicationTagCharacterString) BACnetConstructedDataActionTextBuilder {
	return b.WithActionText(actionText...)
}

func (b *_BACnetConstructedDataActionTextBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataActionTextBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataActionTextBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataActionTextBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataActionTextBuilder) WithActionText(actionText ...BACnetApplicationTagCharacterString) BACnetConstructedDataActionTextBuilder {
	b.ActionText = actionText
	return b
}

func (b *_BACnetConstructedDataActionTextBuilder) Build() (BACnetConstructedDataActionText, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataActionText.deepCopy(), nil
}

func (b *_BACnetConstructedDataActionTextBuilder) MustBuild() BACnetConstructedDataActionText {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataActionTextBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataActionTextBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataActionTextBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataActionTextBuilder().(*_BACnetConstructedDataActionTextBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataActionTextBuilder creates a BACnetConstructedDataActionTextBuilder
func (b *_BACnetConstructedDataActionText) CreateBACnetConstructedDataActionTextBuilder() BACnetConstructedDataActionTextBuilder {
	if b == nil {
		return NewBACnetConstructedDataActionTextBuilder()
	}
	return &_BACnetConstructedDataActionTextBuilder{_BACnetConstructedDataActionText: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActionText) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActionText) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTION_TEXT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActionText) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActionText) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataActionText) GetActionText() []BACnetApplicationTagCharacterString {
	return m.ActionText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataActionText) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActionText(structType any) BACnetConstructedDataActionText {
	if casted, ok := structType.(BACnetConstructedDataActionText); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActionText); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActionText) GetTypeName() string {
	return "BACnetConstructedDataActionText"
}

func (m *_BACnetConstructedDataActionText) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ActionText) > 0 {
		for _, element := range m.ActionText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActionText) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActionText) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataActionText BACnetConstructedDataActionText, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActionText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActionText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	actionText, err := ReadTerminatedArrayField[BACnetApplicationTagCharacterString](ctx, "actionText", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actionText' field"))
	}
	m.ActionText = actionText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActionText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActionText")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActionText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActionText) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActionText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActionText")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "actionText", m.GetActionText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'actionText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActionText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActionText")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActionText) IsBACnetConstructedDataActionText() {}

func (m *_BACnetConstructedDataActionText) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataActionText) deepCopy() *_BACnetConstructedDataActionText {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataActionTextCopy := &_BACnetConstructedDataActionText{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagCharacterString, BACnetApplicationTagCharacterString](m.ActionText),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataActionTextCopy
}

func (m *_BACnetConstructedDataActionText) String() string {
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
