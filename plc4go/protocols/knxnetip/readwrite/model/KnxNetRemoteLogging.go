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

// KnxNetRemoteLogging is the corresponding interface of KnxNetRemoteLogging
type KnxNetRemoteLogging interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
	// IsKnxNetRemoteLogging is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsKnxNetRemoteLogging()
	// CreateBuilder creates a KnxNetRemoteLoggingBuilder
	CreateKnxNetRemoteLoggingBuilder() KnxNetRemoteLoggingBuilder
}

// _KnxNetRemoteLogging is the data-structure of this message
type _KnxNetRemoteLogging struct {
	ServiceIdContract
	Version uint8
}

var _ KnxNetRemoteLogging = (*_KnxNetRemoteLogging)(nil)
var _ ServiceIdRequirements = (*_KnxNetRemoteLogging)(nil)

// NewKnxNetRemoteLogging factory function for _KnxNetRemoteLogging
func NewKnxNetRemoteLogging(version uint8) *_KnxNetRemoteLogging {
	_result := &_KnxNetRemoteLogging{
		ServiceIdContract: NewServiceId(),
		Version:           version,
	}
	_result.ServiceIdContract.(*_ServiceId)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// KnxNetRemoteLoggingBuilder is a builder for KnxNetRemoteLogging
type KnxNetRemoteLoggingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(version uint8) KnxNetRemoteLoggingBuilder
	// WithVersion adds Version (property field)
	WithVersion(uint8) KnxNetRemoteLoggingBuilder
	// Build builds the KnxNetRemoteLogging or returns an error if something is wrong
	Build() (KnxNetRemoteLogging, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() KnxNetRemoteLogging
}

// NewKnxNetRemoteLoggingBuilder() creates a KnxNetRemoteLoggingBuilder
func NewKnxNetRemoteLoggingBuilder() KnxNetRemoteLoggingBuilder {
	return &_KnxNetRemoteLoggingBuilder{_KnxNetRemoteLogging: new(_KnxNetRemoteLogging)}
}

type _KnxNetRemoteLoggingBuilder struct {
	*_KnxNetRemoteLogging

	parentBuilder *_ServiceIdBuilder

	err *utils.MultiError
}

var _ (KnxNetRemoteLoggingBuilder) = (*_KnxNetRemoteLoggingBuilder)(nil)

func (b *_KnxNetRemoteLoggingBuilder) setParent(contract ServiceIdContract) {
	b.ServiceIdContract = contract
}

func (b *_KnxNetRemoteLoggingBuilder) WithMandatoryFields(version uint8) KnxNetRemoteLoggingBuilder {
	return b.WithVersion(version)
}

func (b *_KnxNetRemoteLoggingBuilder) WithVersion(version uint8) KnxNetRemoteLoggingBuilder {
	b.Version = version
	return b
}

func (b *_KnxNetRemoteLoggingBuilder) Build() (KnxNetRemoteLogging, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._KnxNetRemoteLogging.deepCopy(), nil
}

func (b *_KnxNetRemoteLoggingBuilder) MustBuild() KnxNetRemoteLogging {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_KnxNetRemoteLoggingBuilder) Done() ServiceIdBuilder {
	return b.parentBuilder
}

func (b *_KnxNetRemoteLoggingBuilder) buildForServiceId() (ServiceId, error) {
	return b.Build()
}

func (b *_KnxNetRemoteLoggingBuilder) DeepCopy() any {
	_copy := b.CreateKnxNetRemoteLoggingBuilder().(*_KnxNetRemoteLoggingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateKnxNetRemoteLoggingBuilder creates a KnxNetRemoteLoggingBuilder
func (b *_KnxNetRemoteLogging) CreateKnxNetRemoteLoggingBuilder() KnxNetRemoteLoggingBuilder {
	if b == nil {
		return NewKnxNetRemoteLoggingBuilder()
	}
	return &_KnxNetRemoteLoggingBuilder{_KnxNetRemoteLogging: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetRemoteLogging) GetServiceType() uint8 {
	return 0x06
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetRemoteLogging) GetParent() ServiceIdContract {
	return m.ServiceIdContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetRemoteLogging) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastKnxNetRemoteLogging(structType any) KnxNetRemoteLogging {
	if casted, ok := structType.(KnxNetRemoteLogging); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetRemoteLogging); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetRemoteLogging) GetTypeName() string {
	return "KnxNetRemoteLogging"
}

func (m *_KnxNetRemoteLogging) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ServiceIdContract.(*_ServiceId).getLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetRemoteLogging) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_KnxNetRemoteLogging) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ServiceId) (__knxNetRemoteLogging KnxNetRemoteLogging, err error) {
	m.ServiceIdContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetRemoteLogging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetRemoteLogging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	version, err := ReadSimpleField(ctx, "version", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'version' field"))
	}
	m.Version = version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteLogging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetRemoteLogging")
	}

	return m, nil
}

func (m *_KnxNetRemoteLogging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetRemoteLogging) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteLogging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetRemoteLogging")
		}

		if err := WriteSimpleField[uint8](ctx, "version", m.GetVersion(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteLogging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetRemoteLogging")
		}
		return nil
	}
	return m.ServiceIdContract.(*_ServiceId).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetRemoteLogging) IsKnxNetRemoteLogging() {}

func (m *_KnxNetRemoteLogging) DeepCopy() any {
	return m.deepCopy()
}

func (m *_KnxNetRemoteLogging) deepCopy() *_KnxNetRemoteLogging {
	if m == nil {
		return nil
	}
	_KnxNetRemoteLoggingCopy := &_KnxNetRemoteLogging{
		m.ServiceIdContract.(*_ServiceId).deepCopy(),
		m.Version,
	}
	m.ServiceIdContract.(*_ServiceId)._SubType = m
	return _KnxNetRemoteLoggingCopy
}

func (m *_KnxNetRemoteLogging) String() string {
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
