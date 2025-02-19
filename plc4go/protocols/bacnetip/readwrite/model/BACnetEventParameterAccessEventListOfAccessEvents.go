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

// BACnetEventParameterAccessEventListOfAccessEvents is the corresponding interface of BACnetEventParameterAccessEventListOfAccessEvents
type BACnetEventParameterAccessEventListOfAccessEvents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() []BACnetDeviceObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterAccessEventListOfAccessEvents is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterAccessEventListOfAccessEvents()
	// CreateBuilder creates a BACnetEventParameterAccessEventListOfAccessEventsBuilder
	CreateBACnetEventParameterAccessEventListOfAccessEventsBuilder() BACnetEventParameterAccessEventListOfAccessEventsBuilder
}

// _BACnetEventParameterAccessEventListOfAccessEvents is the data-structure of this message
type _BACnetEventParameterAccessEventListOfAccessEvents struct {
	OpeningTag         BACnetOpeningTag
	ListOfAccessEvents []BACnetDeviceObjectPropertyReference
	ClosingTag         BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventParameterAccessEventListOfAccessEvents = (*_BACnetEventParameterAccessEventListOfAccessEvents)(nil)

// NewBACnetEventParameterAccessEventListOfAccessEvents factory function for _BACnetEventParameterAccessEventListOfAccessEvents
func NewBACnetEventParameterAccessEventListOfAccessEvents(openingTag BACnetOpeningTag, listOfAccessEvents []BACnetDeviceObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterAccessEventListOfAccessEvents {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterAccessEventListOfAccessEvents must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterAccessEventListOfAccessEvents must not be nil")
	}
	return &_BACnetEventParameterAccessEventListOfAccessEvents{OpeningTag: openingTag, ListOfAccessEvents: listOfAccessEvents, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterAccessEventListOfAccessEventsBuilder is a builder for BACnetEventParameterAccessEventListOfAccessEvents
type BACnetEventParameterAccessEventListOfAccessEventsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, listOfAccessEvents []BACnetDeviceObjectPropertyReference, closingTag BACnetClosingTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// WithListOfAccessEvents adds ListOfAccessEvents (property field)
	WithListOfAccessEvents(...BACnetDeviceObjectPropertyReference) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder
	// Build builds the BACnetEventParameterAccessEventListOfAccessEvents or returns an error if something is wrong
	Build() (BACnetEventParameterAccessEventListOfAccessEvents, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterAccessEventListOfAccessEvents
}

// NewBACnetEventParameterAccessEventListOfAccessEventsBuilder() creates a BACnetEventParameterAccessEventListOfAccessEventsBuilder
func NewBACnetEventParameterAccessEventListOfAccessEventsBuilder() BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	return &_BACnetEventParameterAccessEventListOfAccessEventsBuilder{_BACnetEventParameterAccessEventListOfAccessEvents: new(_BACnetEventParameterAccessEventListOfAccessEvents)}
}

type _BACnetEventParameterAccessEventListOfAccessEventsBuilder struct {
	*_BACnetEventParameterAccessEventListOfAccessEvents

	err *utils.MultiError
}

var _ (BACnetEventParameterAccessEventListOfAccessEventsBuilder) = (*_BACnetEventParameterAccessEventListOfAccessEventsBuilder)(nil)

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, listOfAccessEvents []BACnetDeviceObjectPropertyReference, closingTag BACnetClosingTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	return b.WithOpeningTag(openingTag).WithListOfAccessEvents(listOfAccessEvents...).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
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

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithListOfAccessEvents(listOfAccessEvents ...BACnetDeviceObjectPropertyReference) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	b.ListOfAccessEvents = listOfAccessEvents
	return b
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterAccessEventListOfAccessEventsBuilder {
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

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) Build() (BACnetEventParameterAccessEventListOfAccessEvents, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
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
	return b._BACnetEventParameterAccessEventListOfAccessEvents.deepCopy(), nil
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) MustBuild() BACnetEventParameterAccessEventListOfAccessEvents {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterAccessEventListOfAccessEventsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterAccessEventListOfAccessEventsBuilder().(*_BACnetEventParameterAccessEventListOfAccessEventsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterAccessEventListOfAccessEventsBuilder creates a BACnetEventParameterAccessEventListOfAccessEventsBuilder
func (b *_BACnetEventParameterAccessEventListOfAccessEvents) CreateBACnetEventParameterAccessEventListOfAccessEventsBuilder() BACnetEventParameterAccessEventListOfAccessEventsBuilder {
	if b == nil {
		return NewBACnetEventParameterAccessEventListOfAccessEventsBuilder()
	}
	return &_BACnetEventParameterAccessEventListOfAccessEventsBuilder{_BACnetEventParameterAccessEventListOfAccessEvents: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetListOfAccessEvents() []BACnetDeviceObjectPropertyReference {
	return m.ListOfAccessEvents
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterAccessEventListOfAccessEvents(structType any) BACnetEventParameterAccessEventListOfAccessEvents {
	if casted, ok := structType.(BACnetEventParameterAccessEventListOfAccessEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEventListOfAccessEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetTypeName() string {
	return "BACnetEventParameterAccessEventListOfAccessEvents"
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfAccessEvents) > 0 {
		for _, element := range m.ListOfAccessEvents {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterAccessEventListOfAccessEventsParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetEventParameterAccessEventListOfAccessEvents, error) {
	return BACnetEventParameterAccessEventListOfAccessEventsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventParameterAccessEventListOfAccessEventsParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEventListOfAccessEvents, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEventListOfAccessEvents, error) {
		return BACnetEventParameterAccessEventListOfAccessEventsParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetEventParameterAccessEventListOfAccessEventsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterAccessEventListOfAccessEvents, error) {
	v, err := (&_BACnetEventParameterAccessEventListOfAccessEvents{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventParameterAccessEventListOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEventListOfAccessEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterAccessEventListOfAccessEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfAccessEvents, err := ReadTerminatedArrayField[BACnetDeviceObjectPropertyReference](ctx, "listOfAccessEvents", ReadComplex[BACnetDeviceObjectPropertyReference](BACnetDeviceObjectPropertyReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAccessEvents' field"))
	}
	m.ListOfAccessEvents = listOfAccessEvents

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEventListOfAccessEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterAccessEventListOfAccessEvents")
	}

	return m, nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEventListOfAccessEvents"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterAccessEventListOfAccessEvents")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfAccessEvents", m.GetListOfAccessEvents(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfAccessEvents' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEventListOfAccessEvents"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterAccessEventListOfAccessEvents")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) IsBACnetEventParameterAccessEventListOfAccessEvents() {
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) deepCopy() *_BACnetEventParameterAccessEventListOfAccessEvents {
	if m == nil {
		return nil
	}
	_BACnetEventParameterAccessEventListOfAccessEventsCopy := &_BACnetEventParameterAccessEventListOfAccessEvents{
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		utils.DeepCopySlice[BACnetDeviceObjectPropertyReference, BACnetDeviceObjectPropertyReference](m.ListOfAccessEvents),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
		m.TagNumber,
	}
	return _BACnetEventParameterAccessEventListOfAccessEventsCopy
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) String() string {
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
