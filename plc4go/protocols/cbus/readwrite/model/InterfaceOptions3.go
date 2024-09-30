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

// InterfaceOptions3 is the corresponding interface of InterfaceOptions3
type InterfaceOptions3 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetExstat returns Exstat (property field)
	GetExstat() bool
	// GetPun returns Pun (property field)
	GetPun() bool
	// GetLocalSal returns LocalSal (property field)
	GetLocalSal() bool
	// GetPcn returns Pcn (property field)
	GetPcn() bool
	// IsInterfaceOptions3 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsInterfaceOptions3()
	// CreateBuilder creates a InterfaceOptions3Builder
	CreateInterfaceOptions3Builder() InterfaceOptions3Builder
}

// _InterfaceOptions3 is the data-structure of this message
type _InterfaceOptions3 struct {
	Exstat   bool
	Pun      bool
	LocalSal bool
	Pcn      bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}

var _ InterfaceOptions3 = (*_InterfaceOptions3)(nil)

// NewInterfaceOptions3 factory function for _InterfaceOptions3
func NewInterfaceOptions3(exstat bool, pun bool, localSal bool, pcn bool) *_InterfaceOptions3 {
	return &_InterfaceOptions3{Exstat: exstat, Pun: pun, LocalSal: localSal, Pcn: pcn}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// InterfaceOptions3Builder is a builder for InterfaceOptions3
type InterfaceOptions3Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(exstat bool, pun bool, localSal bool, pcn bool) InterfaceOptions3Builder
	// WithExstat adds Exstat (property field)
	WithExstat(bool) InterfaceOptions3Builder
	// WithPun adds Pun (property field)
	WithPun(bool) InterfaceOptions3Builder
	// WithLocalSal adds LocalSal (property field)
	WithLocalSal(bool) InterfaceOptions3Builder
	// WithPcn adds Pcn (property field)
	WithPcn(bool) InterfaceOptions3Builder
	// Build builds the InterfaceOptions3 or returns an error if something is wrong
	Build() (InterfaceOptions3, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() InterfaceOptions3
}

// NewInterfaceOptions3Builder() creates a InterfaceOptions3Builder
func NewInterfaceOptions3Builder() InterfaceOptions3Builder {
	return &_InterfaceOptions3Builder{_InterfaceOptions3: new(_InterfaceOptions3)}
}

type _InterfaceOptions3Builder struct {
	*_InterfaceOptions3

	err *utils.MultiError
}

var _ (InterfaceOptions3Builder) = (*_InterfaceOptions3Builder)(nil)

func (b *_InterfaceOptions3Builder) WithMandatoryFields(exstat bool, pun bool, localSal bool, pcn bool) InterfaceOptions3Builder {
	return b.WithExstat(exstat).WithPun(pun).WithLocalSal(localSal).WithPcn(pcn)
}

func (b *_InterfaceOptions3Builder) WithExstat(exstat bool) InterfaceOptions3Builder {
	b.Exstat = exstat
	return b
}

func (b *_InterfaceOptions3Builder) WithPun(pun bool) InterfaceOptions3Builder {
	b.Pun = pun
	return b
}

func (b *_InterfaceOptions3Builder) WithLocalSal(localSal bool) InterfaceOptions3Builder {
	b.LocalSal = localSal
	return b
}

func (b *_InterfaceOptions3Builder) WithPcn(pcn bool) InterfaceOptions3Builder {
	b.Pcn = pcn
	return b
}

func (b *_InterfaceOptions3Builder) Build() (InterfaceOptions3, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._InterfaceOptions3.deepCopy(), nil
}

func (b *_InterfaceOptions3Builder) MustBuild() InterfaceOptions3 {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_InterfaceOptions3Builder) DeepCopy() any {
	_copy := b.CreateInterfaceOptions3Builder().(*_InterfaceOptions3Builder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateInterfaceOptions3Builder creates a InterfaceOptions3Builder
func (b *_InterfaceOptions3) CreateInterfaceOptions3Builder() InterfaceOptions3Builder {
	if b == nil {
		return NewInterfaceOptions3Builder()
	}
	return &_InterfaceOptions3Builder{_InterfaceOptions3: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions3) GetExstat() bool {
	return m.Exstat
}

func (m *_InterfaceOptions3) GetPun() bool {
	return m.Pun
}

func (m *_InterfaceOptions3) GetLocalSal() bool {
	return m.LocalSal
}

func (m *_InterfaceOptions3) GetPcn() bool {
	return m.Pcn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastInterfaceOptions3(structType any) InterfaceOptions3 {
	if casted, ok := structType.(InterfaceOptions3); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions3); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions3) GetTypeName() string {
	return "InterfaceOptions3"
}

func (m *_InterfaceOptions3) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (exstat)
	lengthInBits += 1

	// Simple field (pun)
	lengthInBits += 1

	// Simple field (localSal)
	lengthInBits += 1

	// Simple field (pcn)
	lengthInBits += 1

	return lengthInBits
}

func (m *_InterfaceOptions3) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceOptions3Parse(ctx context.Context, theBytes []byte) (InterfaceOptions3, error) {
	return InterfaceOptions3ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions3ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions3, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions3, error) {
		return InterfaceOptions3ParseWithBuffer(ctx, readBuffer)
	}
}

func InterfaceOptions3ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions3, error) {
	v, err := (&_InterfaceOptions3{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_InterfaceOptions3) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__interfaceOptions3 InterfaceOptions3, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions3")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField2 = reservedField2

	reservedField3, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField3 = reservedField3

	exstat, err := ReadSimpleField(ctx, "exstat", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exstat' field"))
	}
	m.Exstat = exstat

	pun, err := ReadSimpleField(ctx, "pun", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pun' field"))
	}
	m.Pun = pun

	localSal, err := ReadSimpleField(ctx, "localSal", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localSal' field"))
	}
	m.LocalSal = localSal

	pcn, err := ReadSimpleField(ctx, "pcn", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pcn' field"))
	}
	m.Pcn = pcn

	if closeErr := readBuffer.CloseContext("InterfaceOptions3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions3")
	}

	return m, nil
}

func (m *_InterfaceOptions3) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions3) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("InterfaceOptions3"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions3")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 4")
	}

	if err := WriteSimpleField[bool](ctx, "exstat", m.GetExstat(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'exstat' field")
	}

	if err := WriteSimpleField[bool](ctx, "pun", m.GetPun(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'pun' field")
	}

	if err := WriteSimpleField[bool](ctx, "localSal", m.GetLocalSal(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'localSal' field")
	}

	if err := WriteSimpleField[bool](ctx, "pcn", m.GetPcn(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'pcn' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions3"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions3")
	}
	return nil
}

func (m *_InterfaceOptions3) IsInterfaceOptions3() {}

func (m *_InterfaceOptions3) DeepCopy() any {
	return m.deepCopy()
}

func (m *_InterfaceOptions3) deepCopy() *_InterfaceOptions3 {
	if m == nil {
		return nil
	}
	_InterfaceOptions3Copy := &_InterfaceOptions3{
		m.Exstat,
		m.Pun,
		m.LocalSal,
		m.Pcn,
		m.reservedField0,
		m.reservedField1,
		m.reservedField2,
		m.reservedField3,
	}
	return _InterfaceOptions3Copy
}

func (m *_InterfaceOptions3) String() string {
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
