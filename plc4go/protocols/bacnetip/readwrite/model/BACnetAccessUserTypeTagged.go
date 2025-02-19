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

// BACnetAccessUserTypeTagged is the corresponding interface of BACnetAccessUserTypeTagged
type BACnetAccessUserTypeTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessUserType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// IsBACnetAccessUserTypeTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccessUserTypeTagged()
	// CreateBuilder creates a BACnetAccessUserTypeTaggedBuilder
	CreateBACnetAccessUserTypeTaggedBuilder() BACnetAccessUserTypeTaggedBuilder
}

// _BACnetAccessUserTypeTagged is the data-structure of this message
type _BACnetAccessUserTypeTagged struct {
	Header           BACnetTagHeader
	Value            BACnetAccessUserType
	ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccessUserTypeTagged = (*_BACnetAccessUserTypeTagged)(nil)

// NewBACnetAccessUserTypeTagged factory function for _BACnetAccessUserTypeTagged
func NewBACnetAccessUserTypeTagged(header BACnetTagHeader, value BACnetAccessUserType, proprietaryValue uint32, tagNumber uint8, tagClass TagClass) *_BACnetAccessUserTypeTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccessUserTypeTagged must not be nil")
	}
	return &_BACnetAccessUserTypeTagged{Header: header, Value: value, ProprietaryValue: proprietaryValue, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccessUserTypeTaggedBuilder is a builder for BACnetAccessUserTypeTagged
type BACnetAccessUserTypeTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccessUserType, proprietaryValue uint32) BACnetAccessUserTypeTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccessUserTypeTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessUserTypeTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccessUserType) BACnetAccessUserTypeTaggedBuilder
	// WithProprietaryValue adds ProprietaryValue (property field)
	WithProprietaryValue(uint32) BACnetAccessUserTypeTaggedBuilder
	// Build builds the BACnetAccessUserTypeTagged or returns an error if something is wrong
	Build() (BACnetAccessUserTypeTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccessUserTypeTagged
}

// NewBACnetAccessUserTypeTaggedBuilder() creates a BACnetAccessUserTypeTaggedBuilder
func NewBACnetAccessUserTypeTaggedBuilder() BACnetAccessUserTypeTaggedBuilder {
	return &_BACnetAccessUserTypeTaggedBuilder{_BACnetAccessUserTypeTagged: new(_BACnetAccessUserTypeTagged)}
}

type _BACnetAccessUserTypeTaggedBuilder struct {
	*_BACnetAccessUserTypeTagged

	err *utils.MultiError
}

var _ (BACnetAccessUserTypeTaggedBuilder) = (*_BACnetAccessUserTypeTaggedBuilder)(nil)

func (b *_BACnetAccessUserTypeTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccessUserType, proprietaryValue uint32) BACnetAccessUserTypeTaggedBuilder {
	return b.WithHeader(header).WithValue(value).WithProprietaryValue(proprietaryValue)
}

func (b *_BACnetAccessUserTypeTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccessUserTypeTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccessUserTypeTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccessUserTypeTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetAccessUserTypeTaggedBuilder) WithValue(value BACnetAccessUserType) BACnetAccessUserTypeTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccessUserTypeTaggedBuilder) WithProprietaryValue(proprietaryValue uint32) BACnetAccessUserTypeTaggedBuilder {
	b.ProprietaryValue = proprietaryValue
	return b
}

func (b *_BACnetAccessUserTypeTaggedBuilder) Build() (BACnetAccessUserTypeTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccessUserTypeTagged.deepCopy(), nil
}

func (b *_BACnetAccessUserTypeTaggedBuilder) MustBuild() BACnetAccessUserTypeTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccessUserTypeTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccessUserTypeTaggedBuilder().(*_BACnetAccessUserTypeTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccessUserTypeTaggedBuilder creates a BACnetAccessUserTypeTaggedBuilder
func (b *_BACnetAccessUserTypeTagged) CreateBACnetAccessUserTypeTaggedBuilder() BACnetAccessUserTypeTaggedBuilder {
	if b == nil {
		return NewBACnetAccessUserTypeTaggedBuilder()
	}
	return &_BACnetAccessUserTypeTaggedBuilder{_BACnetAccessUserTypeTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessUserTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessUserTypeTagged) GetValue() BACnetAccessUserType {
	return m.Value
}

func (m *_BACnetAccessUserTypeTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetAccessUserTypeTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccessUserTypeTagged(structType any) BACnetAccessUserTypeTagged {
	if casted, ok := structType.(BACnetAccessUserTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessUserTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessUserTypeTagged) GetTypeName() string {
	return "BACnetAccessUserTypeTagged"
}

func (m *_BACnetAccessUserTypeTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetAccessUserTypeTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessUserTypeTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessUserTypeTagged, error) {
	return BACnetAccessUserTypeTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessUserTypeTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessUserTypeTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessUserTypeTagged, error) {
		return BACnetAccessUserTypeTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccessUserTypeTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessUserTypeTagged, error) {
	v, err := (&_BACnetAccessUserTypeTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccessUserTypeTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccessUserTypeTagged BACnetAccessUserTypeTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccessUserTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessUserTypeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetAccessUserType](ctx, "value", readBuffer, EnsureType[BACnetAccessUserType](ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetAccessUserType_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetAccessUserTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessUserTypeTagged")
	}

	return m, nil
}

func (m *_BACnetAccessUserTypeTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessUserTypeTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessUserTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessUserTypeTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccessUserType](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessUserTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessUserTypeTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccessUserTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessUserTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccessUserTypeTagged) IsBACnetAccessUserTypeTagged() {}

func (m *_BACnetAccessUserTypeTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccessUserTypeTagged) deepCopy() *_BACnetAccessUserTypeTagged {
	if m == nil {
		return nil
	}
	_BACnetAccessUserTypeTaggedCopy := &_BACnetAccessUserTypeTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Value,
		m.ProprietaryValue,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccessUserTypeTaggedCopy
}

func (m *_BACnetAccessUserTypeTagged) String() string {
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
