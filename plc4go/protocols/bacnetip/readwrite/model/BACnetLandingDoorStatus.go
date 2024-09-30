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

// BACnetLandingDoorStatus is the corresponding interface of BACnetLandingDoorStatus
type BACnetLandingDoorStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetLandingDoors returns LandingDoors (property field)
	GetLandingDoors() BACnetLandingDoorStatusLandingDoorsList
	// IsBACnetLandingDoorStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLandingDoorStatus()
	// CreateBuilder creates a BACnetLandingDoorStatusBuilder
	CreateBACnetLandingDoorStatusBuilder() BACnetLandingDoorStatusBuilder
}

// _BACnetLandingDoorStatus is the data-structure of this message
type _BACnetLandingDoorStatus struct {
	LandingDoors BACnetLandingDoorStatusLandingDoorsList
}

var _ BACnetLandingDoorStatus = (*_BACnetLandingDoorStatus)(nil)

// NewBACnetLandingDoorStatus factory function for _BACnetLandingDoorStatus
func NewBACnetLandingDoorStatus(landingDoors BACnetLandingDoorStatusLandingDoorsList) *_BACnetLandingDoorStatus {
	if landingDoors == nil {
		panic("landingDoors of type BACnetLandingDoorStatusLandingDoorsList for BACnetLandingDoorStatus must not be nil")
	}
	return &_BACnetLandingDoorStatus{LandingDoors: landingDoors}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLandingDoorStatusBuilder is a builder for BACnetLandingDoorStatus
type BACnetLandingDoorStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(landingDoors BACnetLandingDoorStatusLandingDoorsList) BACnetLandingDoorStatusBuilder
	// WithLandingDoors adds LandingDoors (property field)
	WithLandingDoors(BACnetLandingDoorStatusLandingDoorsList) BACnetLandingDoorStatusBuilder
	// WithLandingDoorsBuilder adds LandingDoors (property field) which is build by the builder
	WithLandingDoorsBuilder(func(BACnetLandingDoorStatusLandingDoorsListBuilder) BACnetLandingDoorStatusLandingDoorsListBuilder) BACnetLandingDoorStatusBuilder
	// Build builds the BACnetLandingDoorStatus or returns an error if something is wrong
	Build() (BACnetLandingDoorStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLandingDoorStatus
}

// NewBACnetLandingDoorStatusBuilder() creates a BACnetLandingDoorStatusBuilder
func NewBACnetLandingDoorStatusBuilder() BACnetLandingDoorStatusBuilder {
	return &_BACnetLandingDoorStatusBuilder{_BACnetLandingDoorStatus: new(_BACnetLandingDoorStatus)}
}

type _BACnetLandingDoorStatusBuilder struct {
	*_BACnetLandingDoorStatus

	err *utils.MultiError
}

var _ (BACnetLandingDoorStatusBuilder) = (*_BACnetLandingDoorStatusBuilder)(nil)

func (b *_BACnetLandingDoorStatusBuilder) WithMandatoryFields(landingDoors BACnetLandingDoorStatusLandingDoorsList) BACnetLandingDoorStatusBuilder {
	return b.WithLandingDoors(landingDoors)
}

func (b *_BACnetLandingDoorStatusBuilder) WithLandingDoors(landingDoors BACnetLandingDoorStatusLandingDoorsList) BACnetLandingDoorStatusBuilder {
	b.LandingDoors = landingDoors
	return b
}

func (b *_BACnetLandingDoorStatusBuilder) WithLandingDoorsBuilder(builderSupplier func(BACnetLandingDoorStatusLandingDoorsListBuilder) BACnetLandingDoorStatusLandingDoorsListBuilder) BACnetLandingDoorStatusBuilder {
	builder := builderSupplier(b.LandingDoors.CreateBACnetLandingDoorStatusLandingDoorsListBuilder())
	var err error
	b.LandingDoors, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLandingDoorStatusLandingDoorsListBuilder failed"))
	}
	return b
}

func (b *_BACnetLandingDoorStatusBuilder) Build() (BACnetLandingDoorStatus, error) {
	if b.LandingDoors == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'landingDoors' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetLandingDoorStatus.deepCopy(), nil
}

func (b *_BACnetLandingDoorStatusBuilder) MustBuild() BACnetLandingDoorStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetLandingDoorStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetLandingDoorStatusBuilder().(*_BACnetLandingDoorStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetLandingDoorStatusBuilder creates a BACnetLandingDoorStatusBuilder
func (b *_BACnetLandingDoorStatus) CreateBACnetLandingDoorStatusBuilder() BACnetLandingDoorStatusBuilder {
	if b == nil {
		return NewBACnetLandingDoorStatusBuilder()
	}
	return &_BACnetLandingDoorStatusBuilder{_BACnetLandingDoorStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingDoorStatus) GetLandingDoors() BACnetLandingDoorStatusLandingDoorsList {
	return m.LandingDoors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLandingDoorStatus(structType any) BACnetLandingDoorStatus {
	if casted, ok := structType.(BACnetLandingDoorStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingDoorStatus) GetTypeName() string {
	return "BACnetLandingDoorStatus"
}

func (m *_BACnetLandingDoorStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (landingDoors)
	lengthInBits += m.LandingDoors.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLandingDoorStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLandingDoorStatusParse(ctx context.Context, theBytes []byte) (BACnetLandingDoorStatus, error) {
	return BACnetLandingDoorStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLandingDoorStatusParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingDoorStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingDoorStatus, error) {
		return BACnetLandingDoorStatusParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLandingDoorStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingDoorStatus, error) {
	v, err := (&_BACnetLandingDoorStatus{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLandingDoorStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetLandingDoorStatus BACnetLandingDoorStatus, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	landingDoors, err := ReadSimpleField[BACnetLandingDoorStatusLandingDoorsList](ctx, "landingDoors", ReadComplex[BACnetLandingDoorStatusLandingDoorsList](BACnetLandingDoorStatusLandingDoorsListParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'landingDoors' field"))
	}
	m.LandingDoors = landingDoors

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingDoorStatus")
	}

	return m, nil
}

func (m *_BACnetLandingDoorStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLandingDoorStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingDoorStatus")
	}

	if err := WriteSimpleField[BACnetLandingDoorStatusLandingDoorsList](ctx, "landingDoors", m.GetLandingDoors(), WriteComplex[BACnetLandingDoorStatusLandingDoorsList](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'landingDoors' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingDoorStatus")
	}
	return nil
}

func (m *_BACnetLandingDoorStatus) IsBACnetLandingDoorStatus() {}

func (m *_BACnetLandingDoorStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLandingDoorStatus) deepCopy() *_BACnetLandingDoorStatus {
	if m == nil {
		return nil
	}
	_BACnetLandingDoorStatusCopy := &_BACnetLandingDoorStatus{
		m.LandingDoors.DeepCopy().(BACnetLandingDoorStatusLandingDoorsList),
	}
	return _BACnetLandingDoorStatusCopy
}

func (m *_BACnetLandingDoorStatus) String() string {
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
