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

// NodeIdTypeDefinition is the corresponding interface of NodeIdTypeDefinition
type NodeIdTypeDefinition interface {
	NodeIdTypeDefinitionContract
	NodeIdTypeDefinitionRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// NodeIdTypeDefinitionContract provides a set of functions which can be overwritten by a sub struct
type NodeIdTypeDefinitionContract interface {
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
}

// NodeIdTypeDefinitionRequirements provides a set of functions which need to be implemented by a sub struct
type NodeIdTypeDefinitionRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetNodeType returns NodeType (discriminator field)
	GetNodeType() NodeIdType
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
}

// NodeIdTypeDefinitionExactly can be used when we want exactly this type and not a type which fulfills NodeIdTypeDefinition.
// This is useful for switch cases.
type NodeIdTypeDefinitionExactly interface {
	NodeIdTypeDefinition
	isNodeIdTypeDefinition() bool
}

// _NodeIdTypeDefinition is the data-structure of this message
type _NodeIdTypeDefinition struct {
	_SubType NodeIdTypeDefinition
}

var _ NodeIdTypeDefinitionContract = (*_NodeIdTypeDefinition)(nil)

type NodeIdTypeDefinitionChild interface {
	utils.Serializable

	GetParent() *NodeIdTypeDefinition

	GetTypeName() string
	NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for abstract fields.
///////////////////////

func (m *_NodeIdTypeDefinition) GetIdentifier() string {
	return m._SubType.GetIdentifier()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdTypeDefinition factory function for _NodeIdTypeDefinition
func NewNodeIdTypeDefinition() *_NodeIdTypeDefinition {
	return &_NodeIdTypeDefinition{}
}

// Deprecated: use the interface for direct cast
func CastNodeIdTypeDefinition(structType any) NodeIdTypeDefinition {
	if casted, ok := structType.(NodeIdTypeDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdTypeDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdTypeDefinition) GetTypeName() string {
	return "NodeIdTypeDefinition"
}

func (m *_NodeIdTypeDefinition) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (nodeType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_NodeIdTypeDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func NodeIdTypeDefinitionParse[T NodeIdTypeDefinition](ctx context.Context, theBytes []byte) (T, error) {
	return NodeIdTypeDefinitionParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdTypeDefinitionParseWithBufferProducer[T NodeIdTypeDefinition]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := NodeIdTypeDefinitionParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func NodeIdTypeDefinitionParseWithBuffer[T NodeIdTypeDefinition](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := NewNodeIdTypeDefinition().parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_NodeIdTypeDefinition) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__nodeIdTypeDefinition NodeIdTypeDefinition, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdTypeDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdTypeDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeType, err := ReadDiscriminatorEnumField[NodeIdType](ctx, "nodeType", "NodeIdType", ReadEnum(NodeIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child NodeIdTypeDefinition
	switch {
	case nodeType == NodeIdType_nodeIdTypeTwoByte: // NodeIdTwoByte
		if _child, err = (&_NodeIdTwoByte{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdTwoByte for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeFourByte: // NodeIdFourByte
		if _child, err = (&_NodeIdFourByte{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdFourByte for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeNumeric: // NodeIdNumeric
		if _child, err = (&_NodeIdNumeric{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdNumeric for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeString: // NodeIdString
		if _child, err = (&_NodeIdString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdString for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeGuid: // NodeIdGuid
		if _child, err = (&_NodeIdGuid{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdGuid for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeByteString: // NodeIdByteString
		if _child, err = (&_NodeIdByteString{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdByteString for type-switch of NodeIdTypeDefinition")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [nodeType=%v]", nodeType)
	}

	if closeErr := readBuffer.CloseContext("NodeIdTypeDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdTypeDefinition")
	}

	return _child, nil
}

func (pm *_NodeIdTypeDefinition) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NodeIdTypeDefinition, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NodeIdTypeDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NodeIdTypeDefinition")
	}

	if err := WriteDiscriminatorEnumField(ctx, "nodeType", "NodeIdType", m.GetNodeType(), WriteEnum[NodeIdType, uint8](NodeIdType.GetValue, NodeIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 6))); err != nil {
		return errors.Wrap(err, "Error serializing 'nodeType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NodeIdTypeDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NodeIdTypeDefinition")
	}
	return nil
}

func (m *_NodeIdTypeDefinition) isNodeIdTypeDefinition() bool {
	return true
}
