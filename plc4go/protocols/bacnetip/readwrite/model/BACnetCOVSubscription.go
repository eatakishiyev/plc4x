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

// BACnetCOVSubscription is the corresponding interface of BACnetCOVSubscription
type BACnetCOVSubscription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientProcessEnclosed
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
	// IsBACnetCOVSubscription is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetCOVSubscription()
	// CreateBuilder creates a BACnetCOVSubscriptionBuilder
	CreateBACnetCOVSubscriptionBuilder() BACnetCOVSubscriptionBuilder
}

// _BACnetCOVSubscription is the data-structure of this message
type _BACnetCOVSubscription struct {
	Recipient                   BACnetRecipientProcessEnclosed
	MonitoredPropertyReference  BACnetObjectPropertyReferenceEnclosed
	IssueConfirmedNotifications BACnetContextTagBoolean
	TimeRemaining               BACnetContextTagUnsignedInteger
	CovIncrement                BACnetContextTagReal
}

var _ BACnetCOVSubscription = (*_BACnetCOVSubscription)(nil)

// NewBACnetCOVSubscription factory function for _BACnetCOVSubscription
func NewBACnetCOVSubscription(recipient BACnetRecipientProcessEnclosed, monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger, covIncrement BACnetContextTagReal) *_BACnetCOVSubscription {
	if recipient == nil {
		panic("recipient of type BACnetRecipientProcessEnclosed for BACnetCOVSubscription must not be nil")
	}
	if monitoredPropertyReference == nil {
		panic("monitoredPropertyReference of type BACnetObjectPropertyReferenceEnclosed for BACnetCOVSubscription must not be nil")
	}
	if issueConfirmedNotifications == nil {
		panic("issueConfirmedNotifications of type BACnetContextTagBoolean for BACnetCOVSubscription must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetCOVSubscription must not be nil")
	}
	return &_BACnetCOVSubscription{Recipient: recipient, MonitoredPropertyReference: monitoredPropertyReference, IssueConfirmedNotifications: issueConfirmedNotifications, TimeRemaining: timeRemaining, CovIncrement: covIncrement}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetCOVSubscriptionBuilder is a builder for BACnetCOVSubscription
type BACnetCOVSubscriptionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(recipient BACnetRecipientProcessEnclosed, monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger) BACnetCOVSubscriptionBuilder
	// WithRecipient adds Recipient (property field)
	WithRecipient(BACnetRecipientProcessEnclosed) BACnetCOVSubscriptionBuilder
	// WithRecipientBuilder adds Recipient (property field) which is build by the builder
	WithRecipientBuilder(func(BACnetRecipientProcessEnclosedBuilder) BACnetRecipientProcessEnclosedBuilder) BACnetCOVSubscriptionBuilder
	// WithMonitoredPropertyReference adds MonitoredPropertyReference (property field)
	WithMonitoredPropertyReference(BACnetObjectPropertyReferenceEnclosed) BACnetCOVSubscriptionBuilder
	// WithMonitoredPropertyReferenceBuilder adds MonitoredPropertyReference (property field) which is build by the builder
	WithMonitoredPropertyReferenceBuilder(func(BACnetObjectPropertyReferenceEnclosedBuilder) BACnetObjectPropertyReferenceEnclosedBuilder) BACnetCOVSubscriptionBuilder
	// WithIssueConfirmedNotifications adds IssueConfirmedNotifications (property field)
	WithIssueConfirmedNotifications(BACnetContextTagBoolean) BACnetCOVSubscriptionBuilder
	// WithIssueConfirmedNotificationsBuilder adds IssueConfirmedNotifications (property field) which is build by the builder
	WithIssueConfirmedNotificationsBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetCOVSubscriptionBuilder
	// WithTimeRemaining adds TimeRemaining (property field)
	WithTimeRemaining(BACnetContextTagUnsignedInteger) BACnetCOVSubscriptionBuilder
	// WithTimeRemainingBuilder adds TimeRemaining (property field) which is build by the builder
	WithTimeRemainingBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVSubscriptionBuilder
	// WithCovIncrement adds CovIncrement (property field)
	WithOptionalCovIncrement(BACnetContextTagReal) BACnetCOVSubscriptionBuilder
	// WithOptionalCovIncrementBuilder adds CovIncrement (property field) which is build by the builder
	WithOptionalCovIncrementBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetCOVSubscriptionBuilder
	// Build builds the BACnetCOVSubscription or returns an error if something is wrong
	Build() (BACnetCOVSubscription, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetCOVSubscription
}

// NewBACnetCOVSubscriptionBuilder() creates a BACnetCOVSubscriptionBuilder
func NewBACnetCOVSubscriptionBuilder() BACnetCOVSubscriptionBuilder {
	return &_BACnetCOVSubscriptionBuilder{_BACnetCOVSubscription: new(_BACnetCOVSubscription)}
}

type _BACnetCOVSubscriptionBuilder struct {
	*_BACnetCOVSubscription

	err *utils.MultiError
}

var _ (BACnetCOVSubscriptionBuilder) = (*_BACnetCOVSubscriptionBuilder)(nil)

func (b *_BACnetCOVSubscriptionBuilder) WithMandatoryFields(recipient BACnetRecipientProcessEnclosed, monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed, issueConfirmedNotifications BACnetContextTagBoolean, timeRemaining BACnetContextTagUnsignedInteger) BACnetCOVSubscriptionBuilder {
	return b.WithRecipient(recipient).WithMonitoredPropertyReference(monitoredPropertyReference).WithIssueConfirmedNotifications(issueConfirmedNotifications).WithTimeRemaining(timeRemaining)
}

func (b *_BACnetCOVSubscriptionBuilder) WithRecipient(recipient BACnetRecipientProcessEnclosed) BACnetCOVSubscriptionBuilder {
	b.Recipient = recipient
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithRecipientBuilder(builderSupplier func(BACnetRecipientProcessEnclosedBuilder) BACnetRecipientProcessEnclosedBuilder) BACnetCOVSubscriptionBuilder {
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

func (b *_BACnetCOVSubscriptionBuilder) WithMonitoredPropertyReference(monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed) BACnetCOVSubscriptionBuilder {
	b.MonitoredPropertyReference = monitoredPropertyReference
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithMonitoredPropertyReferenceBuilder(builderSupplier func(BACnetObjectPropertyReferenceEnclosedBuilder) BACnetObjectPropertyReferenceEnclosedBuilder) BACnetCOVSubscriptionBuilder {
	builder := builderSupplier(b.MonitoredPropertyReference.CreateBACnetObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.MonitoredPropertyReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithIssueConfirmedNotifications(issueConfirmedNotifications BACnetContextTagBoolean) BACnetCOVSubscriptionBuilder {
	b.IssueConfirmedNotifications = issueConfirmedNotifications
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithIssueConfirmedNotificationsBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetCOVSubscriptionBuilder {
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

func (b *_BACnetCOVSubscriptionBuilder) WithTimeRemaining(timeRemaining BACnetContextTagUnsignedInteger) BACnetCOVSubscriptionBuilder {
	b.TimeRemaining = timeRemaining
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithTimeRemainingBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetCOVSubscriptionBuilder {
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

func (b *_BACnetCOVSubscriptionBuilder) WithOptionalCovIncrement(covIncrement BACnetContextTagReal) BACnetCOVSubscriptionBuilder {
	b.CovIncrement = covIncrement
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) WithOptionalCovIncrementBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetCOVSubscriptionBuilder {
	builder := builderSupplier(b.CovIncrement.CreateBACnetContextTagRealBuilder())
	var err error
	b.CovIncrement, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetCOVSubscriptionBuilder) Build() (BACnetCOVSubscription, error) {
	if b.Recipient == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'recipient' not set"))
	}
	if b.MonitoredPropertyReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'monitoredPropertyReference' not set"))
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
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetCOVSubscription.deepCopy(), nil
}

func (b *_BACnetCOVSubscriptionBuilder) MustBuild() BACnetCOVSubscription {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetCOVSubscriptionBuilder) DeepCopy() any {
	_copy := b.CreateBACnetCOVSubscriptionBuilder().(*_BACnetCOVSubscriptionBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetCOVSubscriptionBuilder creates a BACnetCOVSubscriptionBuilder
func (b *_BACnetCOVSubscription) CreateBACnetCOVSubscriptionBuilder() BACnetCOVSubscriptionBuilder {
	if b == nil {
		return NewBACnetCOVSubscriptionBuilder()
	}
	return &_BACnetCOVSubscriptionBuilder{_BACnetCOVSubscription: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVSubscription) GetRecipient() BACnetRecipientProcessEnclosed {
	return m.Recipient
}

func (m *_BACnetCOVSubscription) GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_BACnetCOVSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetCOVSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetCOVSubscription) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetCOVSubscription(structType any) BACnetCOVSubscription {
	if casted, ok := structType.(BACnetCOVSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVSubscription) GetTypeName() string {
	return "BACnetCOVSubscription"
}

func (m *_BACnetCOVSubscription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits(ctx)

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetCOVSubscription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCOVSubscriptionParse(ctx context.Context, theBytes []byte) (BACnetCOVSubscription, error) {
	return BACnetCOVSubscriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetCOVSubscriptionParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
		return BACnetCOVSubscriptionParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetCOVSubscriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
	v, err := (&_BACnetCOVSubscription{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetCOVSubscription) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetCOVSubscription BACnetCOVSubscription, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recipient, err := ReadSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", ReadComplex[BACnetRecipientProcessEnclosed](BACnetRecipientProcessEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recipient' field"))
	}
	m.Recipient = recipient

	monitoredPropertyReference, err := ReadSimpleField[BACnetObjectPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", ReadComplex[BACnetObjectPropertyReferenceEnclosed](BACnetObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredPropertyReference' field"))
	}
	m.MonitoredPropertyReference = monitoredPropertyReference

	issueConfirmedNotifications, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	m.IssueConfirmedNotifications = issueConfirmedNotifications

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	var covIncrement BACnetContextTagReal
	_covIncrement, err := ReadOptionalField[BACnetContextTagReal](ctx, "covIncrement", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	if _covIncrement != nil {
		covIncrement = *_covIncrement
		m.CovIncrement = covIncrement
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVSubscription")
	}

	return m, nil
}

func (m *_BACnetCOVSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVSubscription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetCOVSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVSubscription")
	}

	if err := WriteSimpleField[BACnetRecipientProcessEnclosed](ctx, "recipient", m.GetRecipient(), WriteComplex[BACnetRecipientProcessEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'recipient' field")
	}

	if err := WriteSimpleField[BACnetObjectPropertyReferenceEnclosed](ctx, "monitoredPropertyReference", m.GetMonitoredPropertyReference(), WriteComplex[BACnetObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitoredPropertyReference' field")
	}

	if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", m.GetIssueConfirmedNotifications(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timeRemaining' field")
	}

	if err := WriteOptionalField[BACnetContextTagReal](ctx, "covIncrement", GetRef(m.GetCovIncrement()), WriteComplex[BACnetContextTagReal](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'covIncrement' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVSubscription")
	}
	return nil
}

func (m *_BACnetCOVSubscription) IsBACnetCOVSubscription() {}

func (m *_BACnetCOVSubscription) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetCOVSubscription) deepCopy() *_BACnetCOVSubscription {
	if m == nil {
		return nil
	}
	_BACnetCOVSubscriptionCopy := &_BACnetCOVSubscription{
		m.Recipient.DeepCopy().(BACnetRecipientProcessEnclosed),
		m.MonitoredPropertyReference.DeepCopy().(BACnetObjectPropertyReferenceEnclosed),
		m.IssueConfirmedNotifications.DeepCopy().(BACnetContextTagBoolean),
		m.TimeRemaining.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.CovIncrement.DeepCopy().(BACnetContextTagReal),
	}
	return _BACnetCOVSubscriptionCopy
}

func (m *_BACnetCOVSubscription) String() string {
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
