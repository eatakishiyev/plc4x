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

// IdentifyReplyCommandFirmwareSummary is the corresponding interface of IdentifyReplyCommandFirmwareSummary
type IdentifyReplyCommandFirmwareSummary interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetFirmwareVersion returns FirmwareVersion (property field)
	GetFirmwareVersion() string
	// GetUnitServiceType returns UnitServiceType (property field)
	GetUnitServiceType() byte
	// GetVersion returns Version (property field)
	GetVersion() string
}

// IdentifyReplyCommandFirmwareSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandFirmwareSummary.
// This is useful for switch cases.
type IdentifyReplyCommandFirmwareSummaryExactly interface {
	isIdentifyReplyCommandFirmwareSummary() bool
}

// _IdentifyReplyCommandFirmwareSummary is the data-structure of this message
type _IdentifyReplyCommandFirmwareSummary struct {
	*_IdentifyReplyCommand
	FirmwareVersion string
	UnitServiceType byte
	Version         string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandFirmwareSummary) GetAttribute() Attribute {
	return Attribute_Summary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandFirmwareSummary) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandFirmwareSummary) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandFirmwareSummary) GetFirmwareVersion() string {
	return m.FirmwareVersion
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetUnitServiceType() byte {
	return m.UnitServiceType
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetVersion() string {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandFirmwareSummary factory function for _IdentifyReplyCommandFirmwareSummary
func NewIdentifyReplyCommandFirmwareSummary(firmwareVersion string, unitServiceType byte, version string) *_IdentifyReplyCommandFirmwareSummary {
	_result := &_IdentifyReplyCommandFirmwareSummary{
		FirmwareVersion:       firmwareVersion,
		UnitServiceType:       unitServiceType,
		Version:               version,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandFirmwareSummary(structType interface{}) IdentifyReplyCommandFirmwareSummary {
	if casted, ok := structType.(IdentifyReplyCommandFirmwareSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandFirmwareSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetTypeName() string {
	return "IdentifyReplyCommandFirmwareSummary"
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (firmwareVersion)
	lengthInBits += 48

	// Simple field (unitServiceType)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 32

	return lengthInBits
}

func (m *_IdentifyReplyCommandFirmwareSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandFirmwareSummaryParse(readBuffer utils.ReadBuffer, attribute Attribute) (IdentifyReplyCommandFirmwareSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandFirmwareSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandFirmwareSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (firmwareVersion)
	_firmwareVersion, _firmwareVersionErr := readBuffer.ReadString("firmwareVersion", uint32(48))
	if _firmwareVersionErr != nil {
		return nil, errors.Wrap(_firmwareVersionErr, "Error parsing 'firmwareVersion' field")
	}
	firmwareVersion := _firmwareVersion

	// Simple Field (unitServiceType)
	_unitServiceType, _unitServiceTypeErr := readBuffer.ReadByte("unitServiceType")
	if _unitServiceTypeErr != nil {
		return nil, errors.Wrap(_unitServiceTypeErr, "Error parsing 'unitServiceType' field")
	}
	unitServiceType := _unitServiceType

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadString("version", uint32(32))
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandFirmwareSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandFirmwareSummary")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandFirmwareSummary{
		FirmwareVersion:       firmwareVersion,
		UnitServiceType:       unitServiceType,
		Version:               version,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandFirmwareSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandFirmwareSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandFirmwareSummary")
		}

		// Simple Field (firmwareVersion)
		firmwareVersion := string(m.GetFirmwareVersion())
		_firmwareVersionErr := writeBuffer.WriteString("firmwareVersion", uint32(48), "UTF-8", (firmwareVersion))
		if _firmwareVersionErr != nil {
			return errors.Wrap(_firmwareVersionErr, "Error serializing 'firmwareVersion' field")
		}

		// Simple Field (unitServiceType)
		unitServiceType := byte(m.GetUnitServiceType())
		_unitServiceTypeErr := writeBuffer.WriteByte("unitServiceType", (unitServiceType))
		if _unitServiceTypeErr != nil {
			return errors.Wrap(_unitServiceTypeErr, "Error serializing 'unitServiceType' field")
		}

		// Simple Field (version)
		version := string(m.GetVersion())
		_versionErr := writeBuffer.WriteString("version", uint32(32), "UTF-8", (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandFirmwareSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandFirmwareSummary")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandFirmwareSummary) isIdentifyReplyCommandFirmwareSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandFirmwareSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
