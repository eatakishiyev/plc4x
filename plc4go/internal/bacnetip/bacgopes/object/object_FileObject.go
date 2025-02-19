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

type FileObject struct {
	Object
	objectType           string // TODO: migrateme
	properties           []Property
	_object_supports_cov bool
}

func NewFileObject(arg Arg) (*FileObject, error) {
	o := &FileObject{
		objectType: "file",
		properties: []Property{
			NewReadableProperty("fileType", V2P(NewCharacterString)),
			NewReadableProperty("fileSize", V2P(NewUnsigned)),
			NewReadableProperty("modificationDate", V2P(NewDateTime)),
			NewWritableProperty("archive", V2P(NewBoolean)),
			NewReadableProperty("readOnly", V2P(NewBoolean)),
			NewReadableProperty("fileAccessMethod", V2P(NewFileAccessMethod)),
			NewOptionalProperty("recordCount", V2P(NewUnsigned)),
		},

		//-----
	}
	// TODO: @register_object_type
	return o, nil
}
