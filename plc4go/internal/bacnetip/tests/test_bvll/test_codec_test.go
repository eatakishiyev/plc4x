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

package test_bvll

import (
	"testing"

	"github.com/apache/plc4x/plc4go/internal/bacnetip"
	"github.com/apache/plc4x/plc4go/internal/bacnetip/tests"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/testutils"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

func Result(i uint16) *bacnetip.Result {
	result, err := bacnetip.NewResult(bacnetip.WithResultBvlciResultCode(readWriteModel.BVLCResultCode(i)))
	if err != nil {
		panic(err)
	}
	return result
}

func WriteBroadcastDistributionTable(bdt ...*bacnetip.Address) *bacnetip.WriteBroadcastDistributionTable {
	writeBroadcastDistributionTable, err := bacnetip.NewWriteBroadcastDistributionTable(bacnetip.WithWriteBroadcastDistributionTable(bdt...))
	if err != nil {
		panic(err)
	}
	return writeBroadcastDistributionTable
}

type TestAnnexJCodecSuite struct {
	suite.Suite

	client *tests.TrappedClient
	codec  *bacnetip.AnnexJCodec
	server *tests.TrappedServer

	log zerolog.Logger
}

func (suite *TestAnnexJCodecSuite) SetupSuite() {
	suite.log = testutils.ProduceTestingLogger(suite.T())
}

func (suite *TestAnnexJCodecSuite) SetupTest() {
	// minature trapped stack
	var err error
	suite.codec, err = bacnetip.NewAnnexJCodec(suite.log)
	suite.Require().NoError(err)
	suite.client, err = tests.NewTrappedClient(suite.log)
	suite.Require().NoError(err)
	suite.server, err = tests.NewTrappedServer(suite.log)
	suite.Require().NoError(err)
	err = bacnetip.Bind(suite.log, suite.client, suite.codec, suite.server)
	suite.Require().NoError(err)
}

// Pass the PDU to the client to send down the stack.
func (suite *TestAnnexJCodecSuite) Request(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Request")

	return suite.client.Request(args, kwargs)
}

// Check what the server received.
func (suite *TestAnnexJCodecSuite) Indication(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Indication")

	var pduType any
	if len(args) > 0 {
		pduType = args[0].(any)
	}
	pduAttrs := kwargs
	suite.Assert().True(tests.MatchPdu(suite.log, suite.server.GetIndicationReceived(), pduType, pduAttrs))
	return nil
}

// Check what the server received.
func (suite *TestAnnexJCodecSuite) Response(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Response")

	return suite.server.Response(args, kwargs)
}

// Check what the server received.
func (suite *TestAnnexJCodecSuite) Confirmation(args bacnetip.Args, kwargs bacnetip.KWArgs) error {
	suite.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Confirmation")

	pduType := args[0].(any)
	pduAttrs := kwargs
	suite.Assert().True(tests.MatchPdu(suite.log, suite.client.GetConfirmationReceived(), pduType, pduAttrs))
	return nil
}

func (suite *TestAnnexJCodecSuite) TestResult() {
	// Request successful
	pduBytes, err := bacnetip.Xtob("81.00.0006.0000")
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.BVLCParse(testutils.TestContext(suite.T()), pduBytes)
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(Result(0)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs((*bacnetip.Result)(nil)), bacnetip.NewKWArgs(bacnetip.KWBvlciResultCode, readWriteModel.BVLCResultCode(0)))

	// Request error condition
	pduBytes, err = bacnetip.Xtob("81.00.0006.0010") // TODO: check if this is right or if it should be 01 as there is no code for 1
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.BVLCParse(testutils.TestContext(suite.T()), pduBytes)
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(Result(0x0010)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs((*bacnetip.Result)(nil)), bacnetip.NewKWArgs(bacnetip.KWBvlciResultCode, readWriteModel.BVLCResultCode(0x0010)))
}

func (suite *TestAnnexJCodecSuite) TestWriteBroadcastDistributionTable() {
	// write an empty table
	pduBytes, err := bacnetip.Xtob("81.01.0004")
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.BVLCParse(testutils.TestContext(suite.T()), pduBytes)
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(WriteBroadcastDistributionTable()), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs((*bacnetip.WriteBroadcastDistributionTable)(nil)), bacnetip.NewKWArgs(bacnetip.KWBvlciBDT, nil))

	// write table with an element
	addr, _ := bacnetip.NewAddress(zerolog.Nop(), "192.168.0.254/24")
	pduBytes, err = bacnetip.Xtob("81.01.000e" +
		"c0.a8.00.fe.ba.c0 ff.ff.ff.00") // address and mask
	suite.Require().NoError(err)
	{ // Parse with plc4x parser to validate
		parse, err := readWriteModel.BVLCParse(testutils.TestContext(suite.T()), pduBytes)
		suite.Assert().NoError(err)
		if parse != nil {
			suite.T().Log("\n" + parse.String())
		}
	}

	err = suite.Request(bacnetip.NewArgs(WriteBroadcastDistributionTable(addr)), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Indication(bacnetip.NoArgs, bacnetip.NewKWArgs(bacnetip.KWPDUData, pduBytes))
	suite.Assert().NoError(err)

	err = suite.Response(bacnetip.NewArgs(bacnetip.NewPDU(&bacnetip.MessageBridge{Bytes: pduBytes})), bacnetip.NoKWArgs)
	suite.Assert().NoError(err)
	err = suite.Confirmation(bacnetip.NewArgs((*bacnetip.WriteBroadcastDistributionTable)(nil)), bacnetip.NewKWArgs(bacnetip.KWBvlciBDT, []*bacnetip.Address{addr}))

}

func TestAnnexJCodec(t *testing.T) {
	suite.Run(t, new(TestAnnexJCodecSuite))
}
