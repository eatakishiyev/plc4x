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

package object

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/basetypes"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/primitivedata"
)

type AveragingObject struct {
	Object
	objectType           string // TODO: migrateme
	properties           []Property
	_object_supports_cov bool
}

func NewAveragingObject(arg Arg) (*AveragingObject, error) {
	o := &AveragingObject{
		objectType: "averaging",
		properties: []Property{
			NewReadableProperty("minimumValue", V2P(NewReal)),
			NewOptionalProperty("minimumValueTimestamp", V2P(NewDateTime)),
			NewReadableProperty("averageValue", V2P(NewReal)),
			NewOptionalProperty("varianceValue", V2P(NewReal)),
			NewReadableProperty("maximumValue", V2P(NewReal)),
			NewOptionalProperty("maximumValueTimestamp", V2P(NewDateTime)),
			NewWritableProperty("attemptedSamples", V2P(NewUnsigned)),
			NewReadableProperty("validSamples", V2P(NewUnsigned)),
			NewReadableProperty("objectPropertyReference", V2P(NewDeviceObjectPropertyReference)),
			NewWritableProperty("windowInterval", V2P(NewUnsigned)),
			NewWritableProperty("windowSamples", V2P(NewUnsigned)),
		},
	}
	// TODO: @register_object_type
	return o, nil
}
