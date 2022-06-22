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

// AdsInvalidRequest is the corresponding interface of AdsInvalidRequest
type AdsInvalidRequest interface {
	utils.LengthAware
	utils.Serializable
	AdsData
}

// AdsInvalidRequestExactly can be used when we want exactly this type and not a type which fulfills AdsInvalidRequest.
// This is useful for switch cases.
type AdsInvalidRequestExactly interface {
	isAdsInvalidRequest() bool
}

// _AdsInvalidRequest is the data-structure of this message
type _AdsInvalidRequest struct {
	*_AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidRequest) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidRequest) InitializeParent(parent AdsData) {}

func (m *_AdsInvalidRequest) GetParent() AdsData {
	return m._AdsData
}

// NewAdsInvalidRequest factory function for _AdsInvalidRequest
func NewAdsInvalidRequest() *_AdsInvalidRequest {
	_result := &_AdsInvalidRequest{
		_AdsData: NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidRequest(structType interface{}) AdsInvalidRequest {
	if casted, ok := structType.(AdsInvalidRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidRequest) GetTypeName() string {
	return "AdsInvalidRequest"
}

func (m *_AdsInvalidRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsInvalidRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_AdsInvalidRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsInvalidRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsInvalidRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsInvalidRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsInvalidRequest{
		_AdsData: &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsInvalidRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidRequest")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsInvalidRequest) isAdsInvalidRequest() bool {
	return true
}

func (m *_AdsInvalidRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
