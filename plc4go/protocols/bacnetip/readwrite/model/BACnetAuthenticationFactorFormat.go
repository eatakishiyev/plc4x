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

// BACnetAuthenticationFactorFormat is the corresponding interface of BACnetAuthenticationFactorFormat
type BACnetAuthenticationFactorFormat interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFormatType returns FormatType (property field)
	GetFormatType() BACnetAuthenticationFactorTypeTagged
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetVendorFormat returns VendorFormat (property field)
	GetVendorFormat() BACnetContextTagUnsignedInteger
}

// BACnetAuthenticationFactorFormatExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationFactorFormat.
// This is useful for switch cases.
type BACnetAuthenticationFactorFormatExactly interface {
	BACnetAuthenticationFactorFormat
	isBACnetAuthenticationFactorFormat() bool
}

// _BACnetAuthenticationFactorFormat is the data-structure of this message
type _BACnetAuthenticationFactorFormat struct {
	FormatType   BACnetAuthenticationFactorTypeTagged
	VendorId     BACnetVendorIdTagged
	VendorFormat BACnetContextTagUnsignedInteger
}

var _ BACnetAuthenticationFactorFormat = (*_BACnetAuthenticationFactorFormat)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorFormat) GetFormatType() BACnetAuthenticationFactorTypeTagged {
	return m.FormatType
}

func (m *_BACnetAuthenticationFactorFormat) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetAuthenticationFactorFormat) GetVendorFormat() BACnetContextTagUnsignedInteger {
	return m.VendorFormat
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationFactorFormat factory function for _BACnetAuthenticationFactorFormat
func NewBACnetAuthenticationFactorFormat(formatType BACnetAuthenticationFactorTypeTagged, vendorId BACnetVendorIdTagged, vendorFormat BACnetContextTagUnsignedInteger) *_BACnetAuthenticationFactorFormat {
	return &_BACnetAuthenticationFactorFormat{FormatType: formatType, VendorId: vendorId, VendorFormat: vendorFormat}
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorFormat(structType any) BACnetAuthenticationFactorFormat {
	if casted, ok := structType.(BACnetAuthenticationFactorFormat); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorFormat); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorFormat) GetTypeName() string {
	return "BACnetAuthenticationFactorFormat"
}

func (m *_BACnetAuthenticationFactorFormat) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (formatType)
	lengthInBits += m.FormatType.GetLengthInBits(ctx)

	// Optional Field (vendorId)
	if m.VendorId != nil {
		lengthInBits += m.VendorId.GetLengthInBits(ctx)
	}

	// Optional Field (vendorFormat)
	if m.VendorFormat != nil {
		lengthInBits += m.VendorFormat.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetAuthenticationFactorFormat) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthenticationFactorFormatParse(ctx context.Context, theBytes []byte) (BACnetAuthenticationFactorFormat, error) {
	return BACnetAuthenticationFactorFormatParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthenticationFactorFormatParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorFormat, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorFormat, error) {
		return BACnetAuthenticationFactorFormatParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetAuthenticationFactorFormatParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorFormat, error) {
	v, err := NewBACnetAuthenticationFactorFormat().parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetAuthenticationFactorFormat) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetAuthenticationFactorFormat BACnetAuthenticationFactorFormat, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorFormat"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorFormat")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	formatType, err := ReadSimpleField[BACnetAuthenticationFactorTypeTagged](ctx, "formatType", ReadComplex[BACnetAuthenticationFactorTypeTagged](BACnetAuthenticationFactorTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'formatType' field"))
	}
	m.FormatType = formatType

	var vendorId BACnetVendorIdTagged
	_vendorId, err := ReadOptionalField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	if _vendorId != nil {
		vendorId = *_vendorId
		m.VendorId = vendorId
	}

	var vendorFormat BACnetContextTagUnsignedInteger
	_vendorFormat, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "vendorFormat", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorFormat' field"))
	}
	if _vendorFormat != nil {
		vendorFormat = *_vendorFormat
		m.VendorFormat = vendorFormat
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorFormat"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorFormat")
	}

	return m, nil
}

func (m *_BACnetAuthenticationFactorFormat) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorFormat) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorFormat"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorFormat")
	}

	if err := WriteSimpleField[BACnetAuthenticationFactorTypeTagged](ctx, "formatType", m.GetFormatType(), WriteComplex[BACnetAuthenticationFactorTypeTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'formatType' field")
	}

	if err := WriteOptionalField[BACnetVendorIdTagged](ctx, "vendorId", GetRef(m.GetVendorId()), WriteComplex[BACnetVendorIdTagged](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'vendorId' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "vendorFormat", GetRef(m.GetVendorFormat()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'vendorFormat' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorFormat"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorFormat")
	}
	return nil
}

func (m *_BACnetAuthenticationFactorFormat) isBACnetAuthenticationFactorFormat() bool {
	return true
}

func (m *_BACnetAuthenticationFactorFormat) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
