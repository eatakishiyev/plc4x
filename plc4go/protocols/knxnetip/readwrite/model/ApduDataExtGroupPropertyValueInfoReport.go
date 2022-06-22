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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtGroupPropertyValueInfoReport is the corresponding interface of ApduDataExtGroupPropertyValueInfoReport
type ApduDataExtGroupPropertyValueInfoReport interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtGroupPropertyValueInfoReportExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtGroupPropertyValueInfoReport.
// This is useful for switch cases.
type ApduDataExtGroupPropertyValueInfoReportExactly interface {
	isApduDataExtGroupPropertyValueInfoReport() bool
}

// _ApduDataExtGroupPropertyValueInfoReport is the data-structure of this message
type _ApduDataExtGroupPropertyValueInfoReport struct {
	*_ApduDataExt

	// Arguments.
	Length uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetExtApciType() uint8 {
	return 0x2B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueInfoReport) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtGroupPropertyValueInfoReport factory function for _ApduDataExtGroupPropertyValueInfoReport
func NewApduDataExtGroupPropertyValueInfoReport(length uint8) *_ApduDataExtGroupPropertyValueInfoReport {
	_result := &_ApduDataExtGroupPropertyValueInfoReport{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueInfoReport(structType interface{}) ApduDataExtGroupPropertyValueInfoReport {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueInfoReport); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueInfoReport); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueInfoReport"
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtGroupPropertyValueInfoReportParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtGroupPropertyValueInfoReport, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueInfoReport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueInfoReport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueInfoReport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueInfoReport")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtGroupPropertyValueInfoReport{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueInfoReport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueInfoReport")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueInfoReport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueInfoReport")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) isApduDataExtGroupPropertyValueInfoReport() bool {
	return true
}

func (m *_ApduDataExtGroupPropertyValueInfoReport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
