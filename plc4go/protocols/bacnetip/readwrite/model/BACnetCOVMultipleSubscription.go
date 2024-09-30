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

// BACnetCOVMultipleSubscription is the corresponding interface of BACnetCOVMultipleSubscription
type BACnetCOVMultipleSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientProcessEnclosed
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetMaxNotificationDelay returns MaxNotificationDelay (property field)
	GetMaxNotificationDelay() BACnetContextTagUnsignedInteger
	// GetListOfCovSubscriptionSpecification returns ListOfCovSubscriptionSpecification (property field)
	GetListOfCovSubscriptionSpecification() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
	// IsBACnetCOVMultipleSubscription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCOVMultipleSubscription()
	// CreateBuilder creates a BACnetCOVMultipleSubscriptionBuilder
	CreateBACnetCOVMultipleSubscriptionBuilder() BACnetCOVMultipleSubscriptionBuilder
}

// _BACnetCOVMultipleSubscription is the data-structure of this message
type _BACnetCOVMultipleSubscription struct {
	Recipient                          BACnetRecipientProcessEnclosed
	IssueConfirmedNotifications        BACnetContextTagBoolean
	TimeRemaining                      BACnetContextTagUnsignedInteger
	MaxNotificationDelay               BACnetContextTagUnsignedInteger
	ListOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
}

var _ BACnetCOVMultipleSubscription = (*_BACnetCOVMultipleSubscription)(nil)

// NewBACnetCOVMultipleSubscription factory function for _BACnetCOVMultipleSubscription
func NewBACnetCOVMultipleSubscription(recipient BACnetRecipientProcessEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) *_BACnetCOVMultipleSubscription {
	if recipient == nil {
		panic("recipient of type BACnetRecipientProcessEnclosed for BACnetCOVMultipleSubscription must not be nil")
	}
	if issueConfirmedNotifications == nil {
		panic("issueConfirmedNotifications of type BACnetContextTagBoolean for BACnetCOVMultipleSubscription must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetCOVMultipleSubscription must not be nil")
	}
	if maxNotificationDelay == nil {
		panic("maxNotificationDelay of type BACnetContextTagUnsignedInteger for BACnetCOVMultipleSubscription must not be nil")
	}
	if listOfCovSubscriptionSpecification == nil {
		panic("listOfCovSubscriptionSpecification of type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification for BACnetCOVMultipleSubscription must not be nil")
	}
	return &_BACnetCOVMultipleSubscription{Recipient: recipient, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining, MaxNotificationDelay: maxNotificationDelay, ListOfCovSubscriptionSpecification: listOfCovSubscriptionSpecification}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCOVMultipleSubscriptionBuilder is a builder for BACnetCOVMultipleSubscription
type BACnetCOVMultipleSubscriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recipient BACnetRecipientProcessEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) BACnetCOVMultipleSubscriptionBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipientProcessEnclosed) BACnetCOVMultipleSubscriptionBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientProcessEnclosedBuilder) BACnetRecipientProcessEnclosedBuilder) BACnetCOVMultipleSubscriptionBuilder
	// WithIssueConfirmedNotifications adds IssueConfirmedNotifications (property field)
	WithIssueConfirmedNotifications(BACnetContextTagBoolean) BACnetCOVMultipleSubscriptionBuilder
	// WithIssueConfirmedNotificationsBuilder adds IssueConfirmedNotifications (property field) which is build by the builder
	WithIssueConfirmedNotificationsBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetCOVMultipleSubscriptionBuilder
	// WithTimeRemaining adds TimeRemaining (property field)
	WithTimeRemaining(BACnetContextTagUnsignedInteger) BACnetCOVMultipleSubscriptionBuilder
	// WithTimeRemainingBuilder adds TimeRemaining (property field) which is build by the builder
	WithTimeRemainingBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVMultipleSubscriptionBuilder
	// WithMaxNotificationDelay adds MaxNotificationDelay (property field)
	WithMaxNotificationDelay(BACnetContextTagUnsignedInteger) BACnetCOVMultipleSubscriptionBuilder
	// WithMaxNotificationDelayBuilder adds MaxNotificationDelay (property field) which is build by the builder
	WithMaxNotificationDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVMultipleSubscriptionBuilder
	// WithListOfCovSubscriptionSpecification adds ListOfCovSubscriptionSpecification (property field)
	WithListOfCovSubscriptionSpecification(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) BACnetCOVMultipleSubscriptionBuilder
	// WithListOfCovSubscriptionSpecificationBuilder adds ListOfCovSubscriptionSpecification (property field) which is build by the builder
	WithListOfCovSubscriptionSpecificationBuilder(func(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) BACnetCOVMultipleSubscriptionBuilder
	// Build builds the BACnetCOVMultipleSubscription or returns an error if something is wrong
	Build() (BACnetCOVMultipleSubscription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCOVMultipleSubscription
}

// NewBACnetCOVMultipleSubscriptionBuilder() creates a BACnetCOVMultipleSubscriptionBuilder
func NewBACnetCOVMultipleSubscriptionBuilder() BACnetCOVMultipleSubscriptionBuilder {
	return &_BACnetCOVMultipleSubscriptionBuilder{_BACnetCOVMultipleSubscription: new(_BACnetCOVMultipleSubscription)}
}

type _BACnetCOVMultipleSubscriptionBuilder struct {
	*_BACnetCOVMultipleSubscription

	err *utils.MultiError
}

var _ (BACnetCOVMultipleSubscriptionBuilder) = (*_BACnetCOVMultipleSubscriptionBuilder)(nil)

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithMandatoryFields(recipient BACnetRecipientProcessEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) BACnetCOVMultipleSubscriptionBuilder {
	return b.WithRecipient(recipient).WithIssueConfirmedNotifications(issueConfirmedNotifications).WithTimeRemaining(timeRemaining).WithMaxNotificationDelay(maxNotificationDelay).WithListOfCovSubscriptionSpecification(listOfCovSubscriptionSpecification)
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithRecipient(recipient BACnetRecipientProcessEnclosed) BACnetCOVMultipleSubscriptionBuilder {
	b.Recipient = recipient
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientProcessEnclosedBuilder) BACnetRecipientProcessEnclosedBuilder) BACnetCOVMultipleSubscriptionBuilder {
	builder := builderSupplier(b.Recipient.CreateBACnetRecipientProcessEnclosedBuilder())
	var err error
	b.Recipient, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetRecipientProcessEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithIssueConfirmedNotifications(issueConfirmedNotifications BACnetContextTagBoolean) BACnetCOVMultipleSubscriptionBuilder {
	b.IssueConfirmedNotifications = issueConfirmedNotifications
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithIssueConfirmedNotificationsBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetCOVMultipleSubscriptionBuilder {
	builder := builderSupplier(b.IssueConfirmedNotifications.CreateBACnetContextTagBooleanBuilder())
	var err error
	b.IssueConfirmedNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithTimeRemaining(timeRemaining BACnetContextTagUnsignedInteger) BACnetCOVMultipleSubscriptionBuilder {
	b.TimeRemaining = timeRemaining
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithTimeRemainingBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVMultipleSubscriptionBuilder {
	builder := builderSupplier(b.TimeRemaining.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeRemaining, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithMaxNotificationDelay(maxNotificationDelay BACnetContextTagUnsignedInteger) BACnetCOVMultipleSubscriptionBuilder {
	b.MaxNotificationDelay = maxNotificationDelay
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithMaxNotificationDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVMultipleSubscriptionBuilder {
	builder := builderSupplier(b.MaxNotificationDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.MaxNotificationDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithListOfCovSubscriptionSpecification(listOfCovSubscriptionSpecification BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) BACnetCOVMultipleSubscriptionBuilder {
	b.ListOfCovSubscriptionSpecification = listOfCovSubscriptionSpecification
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) WithListOfCovSubscriptionSpecificationBuilder(builderSupplier func(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder) BACnetCOVMultipleSubscriptionBuilder {
	builder := builderSupplier(b.ListOfCovSubscriptionSpecification.CreateBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder())
	var err error
	b.ListOfCovSubscriptionSpecification, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) Build() (BACnetCOVMultipleSubscription, error) {
	if b.Recipient == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'recipient' not set"))
	}
	if b.IssueConfirmedNotifications == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'issueConfirmedNotifications' not set"))
	}
	if b.TimeRemaining == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeRemaining' not set"))
	}
	if b.MaxNotificationDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxNotificationDelay' not set"))
	}
	if b.ListOfCovSubscriptionSpecification == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfCovSubscriptionSpecification' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetCOVMultipleSubscription.deepCopy(), nil
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) MustBuild() BACnetCOVMultipleSubscription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetCOVMultipleSubscriptionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetCOVMultipleSubscriptionBuilder().(*_BACnetCOVMultipleSubscriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetCOVMultipleSubscriptionBuilder creates a BACnetCOVMultipleSubscriptionBuilder
func (b *_BACnetCOVMultipleSubscription) CreateBACnetCOVMultipleSubscriptionBuilder() BACnetCOVMultipleSubscriptionBuilder {
	if b == nil {
		return NewBACnetCOVMultipleSubscriptionBuilder()
	}
	return &_BACnetCOVMultipleSubscriptionBuilder{_BACnetCOVMultipleSubscription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVMultipleSubscription) GetRecipient() BACnetRecipientProcessEnclosed {
	return m.Recipient
}

func (m *_BACnetCOVMultipleSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetCOVMultipleSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetCOVMultipleSubscription) GetMaxNotificationDelay() BACnetContextTagUnsignedInteger {
	return m.MaxNotificationDelay
}

func (m *_BACnetCOVMultipleSubscription) GetListOfCovSubscriptionSpecification() BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	return m.ListOfCovSubscriptionSpecification
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCOVMultipleSubscription(structType any) BACnetCOVMultipleSubscription {
	if casted, ok := structType.(BACnetCOVMultipleSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscription) GetTypeName() string {
	return "BACnetCOVMultipleSubscription"
}

func (m *_BACnetCOVMultipleSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Simple field (maxNotificationDelay)
	lengthInBits += m.MaxNotificationDelay.GetLengthInBits(ctx)

	// Simple field (listOfCovSubscriptionSpecification)
	lengthInBits += m.ListOfCovSubscriptionSpecification.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCOVMultipleSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVMultipleSubscriptionParse(ctx context.Context, theBytes []byte) (BACnetCOVMultipleSubscription, error) {
	return BACnetCOVMultipleSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVMultipleSubscriptionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
		return BACnetCOVMultipleSubscriptionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVMultipleSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVMultipleSubscription, error) {
	v, err := (&_BACnetCOVMultipleSubscription{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetCOVMultipleSubscription) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetCOVMultipleSubscription BACnetCOVMultipleSubscription, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVMultipleSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientProcessEnclosed](BACnetRecipientProcessEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	issueConfirmedNotifications, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	m.IssueConfirmedNotifications = issueConfirmedNotifications

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	maxNotificationDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationDelay' field"))
	}
	m.MaxNotificationDelay = maxNotificationDelay

	listOfCovSubscriptionSpecification, err := ReadSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](ctx, "listOfCovSubscriptionSpecification", ReadComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecification' field"))
	}
	m.ListOfCovSubscriptionSpecification = listOfCovSubscriptionSpecification

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVMultipleSubscription")
	}

	return m, nil
}

func (m *_BACnetCOVMultipleSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVMultipleSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVMultipleSubscription")
	}

	if err := WriteSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipientProcessEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", m.GetIssueConfirmedNotifications(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRemaining' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", m.GetMaxNotificationDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'maxNotificationDelay' field")
	}

	if err := WriteSimpleField[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](ctx, "listOfCovSubscriptionSpecification", m.GetListOfCovSubscriptionSpecification(), WriteComplex[BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfCovSubscriptionSpecification' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVMultipleSubscription")
	}
	return nil
}

func (m *_BACnetCOVMultipleSubscription) IsBACnetCOVMultipleSubscription() {}

func (m *_BACnetCOVMultipleSubscription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCOVMultipleSubscription) deepCopy() *_BACnetCOVMultipleSubscription {
	if m == nil {
		return nil
	}
	_BACnetCOVMultipleSubscriptionCopy := &_BACnetCOVMultipleSubscription{
		m.Recipient.DeepCopy().(BACnetRecipientProcessEnclosed),
		m.IssueConfirmedNotifications.DeepCopy().(BACnetContextTagBoolean),
		m.TimeRemaining.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.MaxNotificationDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ListOfCovSubscriptionSpecification.DeepCopy().(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification),
	}
	return _BACnetCOVMultipleSubscriptionCopy
}

func (m *_BACnetCOVMultipleSubscription) String() string {
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
