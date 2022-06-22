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

// BACnetConfirmedServiceRequestAtomicReadFile is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFile
type BACnetConfirmedServiceRequestAtomicReadFile interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetFileIdentifier returns FileIdentifier (property field)
	GetFileIdentifier() BACnetApplicationTagObjectIdentifier
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
}

// BACnetConfirmedServiceRequestAtomicReadFileExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestAtomicReadFile.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestAtomicReadFileExactly interface {
	isBACnetConfirmedServiceRequestAtomicReadFile() bool
}

// _BACnetConfirmedServiceRequestAtomicReadFile is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFile struct {
	*_BACnetConfirmedServiceRequest
	FileIdentifier BACnetApplicationTagObjectIdentifier
	AccessMethod   BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetFileIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.FileIdentifier
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetAccessMethod() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFile factory function for _BACnetConfirmedServiceRequestAtomicReadFile
func NewBACnetConfirmedServiceRequestAtomicReadFile(fileIdentifier BACnetApplicationTagObjectIdentifier, accessMethod BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestAtomicReadFile {
	_result := &_BACnetConfirmedServiceRequestAtomicReadFile{
		FileIdentifier:                 fileIdentifier,
		AccessMethod:                   accessMethod,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFile(structType interface{}) BACnetConfirmedServiceRequestAtomicReadFile {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFile); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFile"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileIdentifier)
	lengthInBits += m.FileIdentifier.GetLengthInBits()

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestAtomicReadFile, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileIdentifier)
	if pullErr := readBuffer.PullContext("fileIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileIdentifier")
	}
	_fileIdentifier, _fileIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _fileIdentifierErr != nil {
		return nil, errors.Wrap(_fileIdentifierErr, "Error parsing 'fileIdentifier' field")
	}
	fileIdentifier := _fileIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("fileIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileIdentifier")
	}

	// Simple Field (accessMethod)
	if pullErr := readBuffer.PullContext("accessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessMethod")
	}
	_accessMethod, _accessMethodErr := BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordParse(readBuffer)
	if _accessMethodErr != nil {
		return nil, errors.Wrap(_accessMethodErr, "Error parsing 'accessMethod' field")
	}
	accessMethod := _accessMethod.(BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord)
	if closeErr := readBuffer.CloseContext("accessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessMethod")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFile")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestAtomicReadFile{
		FileIdentifier:                 fileIdentifier,
		AccessMethod:                   accessMethod,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFile")
		}

		// Simple Field (fileIdentifier)
		if pushErr := writeBuffer.PushContext("fileIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileIdentifier")
		}
		_fileIdentifierErr := writeBuffer.WriteSerializable(m.GetFileIdentifier())
		if popErr := writeBuffer.PopContext("fileIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileIdentifier")
		}
		if _fileIdentifierErr != nil {
			return errors.Wrap(_fileIdentifierErr, "Error serializing 'fileIdentifier' field")
		}

		// Simple Field (accessMethod)
		if pushErr := writeBuffer.PushContext("accessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessMethod")
		}
		_accessMethodErr := writeBuffer.WriteSerializable(m.GetAccessMethod())
		if popErr := writeBuffer.PopContext("accessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessMethod")
		}
		if _accessMethodErr != nil {
			return errors.Wrap(_accessMethodErr, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFile")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) isBACnetConfirmedServiceRequestAtomicReadFile() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
