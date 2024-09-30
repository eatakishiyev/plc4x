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

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the corresponding interface of SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsSubscribeCOVPropertyMultipleErrorFirstFailedSubscription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSubscribeCOVPropertyMultipleErrorFirstFailedSubscription()
	// CreateBuilder creates a SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	CreateSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder() SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
}

// _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription is the data-structure of this message
type _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription struct {
	OpeningTag                 BACnetOpeningTag
	MonitoredObjectIdentifier  BACnetContextTagObjectIdentifier
	MonitoredPropertyReference BACnetPropertyReferenceEnclosed
	ErrorType                  ErrorEnclosed
	ClosingTag                 BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ SubscribeCOVPropertyMultipleErrorFirstFailedSubscription = (*_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription)(nil)

// NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription factory function for _SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
func NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(openingTag BACnetOpeningTag, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, monitoredPropertyReference BACnetPropertyReferenceEnclosed, errorType ErrorEnclosed, closingTag BACnetClosingTag, tagNumber uint8) *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription must not be nil")
	}
	if monitoredObjectIdentifier == nil {
		panic("monitoredObjectIdentifier of type BACnetContextTagObjectIdentifier for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription must not be nil")
	}
	if monitoredPropertyReference == nil {
		panic("monitoredPropertyReference of type BACnetPropertyReferenceEnclosed for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription must not be nil")
	}
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription must not be nil")
	}
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{OpeningTag: openingTag, MonitoredObjectIdentifier: monitoredObjectIdentifier, MonitoredPropertyReference: monitoredPropertyReference, ErrorType: errorType, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder is a builder for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
type SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, monitoredPropertyReference BACnetPropertyReferenceEnclosed, errorType ErrorEnclosed, closingTag BACnetClosingTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithMonitoredObjectIdentifier adds MonitoredObjectIdentifier (property field)
	WithMonitoredObjectIdentifier(BACnetContextTagObjectIdentifier) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithMonitoredObjectIdentifierBuilder adds MonitoredObjectIdentifier (property field) which is build by the builder
	WithMonitoredObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithMonitoredPropertyReference adds MonitoredPropertyReference (property field)
	WithMonitoredPropertyReference(BACnetPropertyReferenceEnclosed) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithMonitoredPropertyReferenceBuilder adds MonitoredPropertyReference (property field) which is build by the builder
	WithMonitoredPropertyReferenceBuilder(func(BACnetPropertyReferenceEnclosedBuilder) BACnetPropertyReferenceEnclosedBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
	// Build builds the SubscribeCOVPropertyMultipleErrorFirstFailedSubscription or returns an error if something is wrong
	Build() (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

// NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder() creates a SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
func NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder() SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder{_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription: new(_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription)}
}

type _SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder struct {
	*_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription

	err *utils.MultiError
}

var _ (SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) = (*_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder)(nil)

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, monitoredObjectIdentifier BACnetContextTagObjectIdentifier, monitoredPropertyReference BACnetPropertyReferenceEnclosed, errorType ErrorEnclosed, closingTag BACnetClosingTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	return b.WithOpeningTag(openingTag).WithMonitoredObjectIdentifier(monitoredObjectIdentifier).WithMonitoredPropertyReference(monitoredPropertyReference).WithErrorType(errorType).WithClosingTag(closingTag)
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithOpeningTag(openingTag BACnetOpeningTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithMonitoredObjectIdentifier(monitoredObjectIdentifier BACnetContextTagObjectIdentifier) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	b.MonitoredObjectIdentifier = monitoredObjectIdentifier
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithMonitoredObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	builder := builderSupplier(b.MonitoredObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.MonitoredObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithMonitoredPropertyReference(monitoredPropertyReference BACnetPropertyReferenceEnclosed) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	b.MonitoredPropertyReference = monitoredPropertyReference
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithMonitoredPropertyReferenceBuilder(builderSupplier func(BACnetPropertyReferenceEnclosedBuilder) BACnetPropertyReferenceEnclosedBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	builder := builderSupplier(b.MonitoredPropertyReference.CreateBACnetPropertyReferenceEnclosedBuilder())
	var err error
	b.MonitoredPropertyReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithErrorType(errorType ErrorEnclosed) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	builder := builderSupplier(b.ErrorType.CreateErrorEnclosedBuilder())
	var err error
	b.ErrorType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithClosingTag(closingTag BACnetClosingTag) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) Build() (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.MonitoredObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'monitoredObjectIdentifier' not set"))
	}
	if b.MonitoredPropertyReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'monitoredPropertyReference' not set"))
	}
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SubscribeCOVPropertyMultipleErrorFirstFailedSubscription.deepCopy(), nil
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) MustBuild() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder) DeepCopy() any {
	_copy := b.CreateSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder().(*_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder creates a SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder
func (b *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) CreateSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder() SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder {
	if b == nil {
		return NewSubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder()
	}
	return &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionBuilder{_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetMonitoredPropertyReference() BACnetPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleErrorFirstFailedSubscription(structType any) SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleErrorFirstFailedSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits(ctx)

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(ctx context.Context, theBytes []byte, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	return SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
		return SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, error) {
	v, err := (&_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__subscribeCOVPropertyMultipleErrorFirstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	monitoredObjectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredObjectIdentifier' field"))
	}
	m.MonitoredObjectIdentifier = monitoredObjectIdentifier

	monitoredPropertyReference, err := ReadSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", ReadComplex[BACnetPropertyReferenceEnclosed](BACnetPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredPropertyReference' field"))
	}
	m.MonitoredPropertyReference = monitoredPropertyReference

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	return m, nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "monitoredObjectIdentifier", m.GetMonitoredObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredObjectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", m.GetMonitoredPropertyReference(), WriteComplex[BACnetPropertyReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredPropertyReference' field")
	}

	if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'errorType' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleErrorFirstFailedSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleErrorFirstFailedSubscription")
	}
	return nil
}

////
// Arguments Getter

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) IsSubscribeCOVPropertyMultipleErrorFirstFailedSubscription() {
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) deepCopy() *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	if m == nil {
		return nil
	}
	_SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionCopy := &_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.MonitoredObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.MonitoredPropertyReference.DeepCopy().(BACnetPropertyReferenceEnclosed),
		m.ErrorType.DeepCopy().(ErrorEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionCopy
}

func (m *_SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) String() string {
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
