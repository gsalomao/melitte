// Copyright 2023 Gustavo Salomao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConnectTestSuite struct {
	suite.Suite
}

func (s *ConnectTestSuite) TestType() {
	var p Connect
	s.Require().Equal(TypeConnect, p.Type())
}

func (s *ConnectTestSuite) TestSize() {
	testCases := []struct {
		name   string
		packet Connect
		size   int
	}{
		{"V3.1, client ID", Connect{Version: MQTT31, ClientID: []byte("a")}, 17},
		{"V3.1.1, client ID and keep alive", Connect{Version: MQTT311, KeepAlive: 30, ClientID: []byte("ab")}, 16},
		{"V5.0", Connect{Version: MQTT50, KeepAlive: 500, ClientID: []byte("abc")}, 18},
		{
			"V5.0, empty properties",
			Connect{
				Version:        MQTT50,
				KeepAlive:      500,
				ClientID:       []byte("abc"),
				Properties:     &ConnectProperties{},
				WillProperties: &WillProperties{},
			},
			18,
		},
		{
			"V5.0, with Will, and empty properties",
			Connect{
				Version:        MQTT50,
				ClientID:       []byte("abc"),
				Flags:          connectFlagWillFlag,
				Properties:     &ConnectProperties{},
				WillTopic:      []byte("a"),
				WillPayload:    []byte("b"),
				WillProperties: &WillProperties{},
			},
			25,
		},
		{
			"V5.0, with Will, username, password, and properties",
			Connect{
				Version:  MQTT50,
				ClientID: []byte("abc"),
				Flags:    ConnectFlags(connectFlagWillFlag | connectFlagUsernameFlag | connectFlagPasswordFlag),
				Properties: &ConnectProperties{
					Flags:                 PropertyFlags(0).Set(PropertySessionExpiryInterval),
					SessionExpiryInterval: 30,
				},
				WillTopic:   []byte("a"),
				WillPayload: []byte("b"),
				WillProperties: &WillProperties{
					Flags:             PropertyFlags(0).Set(PropertyWillDelayInterval),
					WillDelayInterval: 60,
				},
				Username: []byte("user"),
				Password: []byte("pass"),
			},
			47,
		},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			size := test.packet.Size()
			s.Require().Equal(test.size, size)
		})
	}
}

func (s *ConnectTestSuite) TestDecodeSuccess() {
	testCases := []struct {
		name   string
		data   []byte
		packet Connect
	}{
		{
			"V3.1",
			[]byte{
				0, 6, 'M', 'Q', 'I', 's', 'd', 'p', // Protocol name
				3,     // Protocol version
				0,     // Packet flags
				0, 30, // Keep alive
				0, 1, 'a', // Client ID
			},
			Connect{Version: MQTT31, KeepAlive: 30, ClientID: []byte("a")},
		},
		{
			"V3.1.1",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,      // Protocol version
				0,      // Packet flags
				0, 255, // Keep alive
				0, 2, 'a', 'b', // Client ID
			},
			Connect{Version: MQTT311, KeepAlive: 255, ClientID: []byte("ab")},
		},
		{
			"V3.1.1, clean session",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,      // Protocol version
				2,      // Packet flags (Clean Session)
				0, 255, // Keep alive
				0, 2, 'a', 'b', // Client ID
			},
			Connect{Version: MQTT311, KeepAlive: 255, Flags: connectFlagCleanSession, ClientID: []byte("ab")},
		},
		{
			"V3.1.1, clean session + no client ID",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,      // Protocol version
				2,      // Packet flags (Clean Session)
				0, 255, // Keep alive
				0, 0, // Client ID
			},
			Connect{
				Version: MQTT311, KeepAlive: 255, Flags: connectFlagCleanSession, ClientID: []byte{},
			},
		},
		{
			"V3.1.1, Will flags",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,    // Protocol version
				0x2c, // Packet flags (Will Retain + Will QoS + Will Flag)
				1, 0, // Keep alive
				0, 2, 'a', 'b', // Client ID
				0, 1, 'a', // Will Topic
				0, 1, 'b', // Will Payload
			},
			Connect{
				Version:   MQTT311,
				KeepAlive: 256,
				Flags: ConnectFlags(
					connectFlagWillRetain | (QoS1 << connectFlagShiftWillQoS) | connectFlagWillFlag,
				),
				ClientID:    []byte("ab"),
				WillTopic:   []byte("a"),
				WillPayload: []byte("b"),
			},
		},
		{
			"V3.1.1, Will flags + no will payload",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,    // Protocol version
				0x2c, // Packet flags (Will Retain + Will QoS + Will Flag)
				1, 0, // Keep alive
				0, 2, 'a', 'b', // Client ID
				0, 1, 'a', // Will Topic
				0, 0, // Will Payload
			},
			Connect{
				Version:   MQTT311,
				KeepAlive: 256,
				Flags: ConnectFlags(
					connectFlagWillRetain | (QoS1 << connectFlagShiftWillQoS) | connectFlagWillFlag,
				),
				ClientID:    []byte("ab"),
				WillTopic:   []byte("a"),
				WillPayload: []byte{},
			},
		},
		{
			"V3.1.1, username/password",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				4,    // Protocol version
				0xc0, // Packet flags (Username + Password)
				0, 1, // Keep alive
				0, 2, 'a', 'b', // Client ID
				0, 1, 'a', // Username
				0, 1, 'b', // Password
			},
			Connect{
				Version:   MQTT311,
				KeepAlive: 1,
				ClientID:  []byte("ab"),
				Flags:     ConnectFlags(connectFlagUsernameFlag | connectFlagPasswordFlag),
				Username:  []byte("a"),
				Password:  []byte("b"),
			},
		},
		{
			"V5.0, no properties",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				5,        // Protocol version
				0,        // Packet flags
				255, 255, // Keep alive
				0,                   // Properties Length
				0, 3, 'a', 'b', 'c', // Client ID
			},
			Connect{Version: MQTT50, KeepAlive: 65535, ClientID: []byte("abc")},
		},
		{
			"V5.0, no properties, password",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				5,        // Protocol version
				0x40,     // Packet flags
				255, 255, // Keep alive
				0,                   // Properties Length
				0, 3, 'a', 'b', 'c', // Client ID
				0, 1, 'd', // Password
			},
			Connect{
				Version: MQTT50, KeepAlive: 65535, ClientID: []byte("abc"), Flags: connectFlagPasswordFlag,
				Password: []byte("d"),
			},
		},
		{
			"V5.0, properties",
			[]byte{
				0, 4, 'M', 'Q', 'T', 'T', // Protocol name
				5,        // Protocol version
				0x04,     // Packet flags
				255, 255, // Keep alive
				5,               // Property length
				17, 0, 0, 0, 10, // Session expiry interval
				0, 3, 'a', 'b', 'c', // Client ID
				5,               // Will Property length
				24, 0, 0, 0, 15, // Will delay interval
				0, 1, 'a', // Will Topic
				0, 1, 'b', // Will Payload
			},
			Connect{
				Version:   MQTT50,
				KeepAlive: 65535,
				Flags:     connectFlagWillFlag,
				Properties: &ConnectProperties{
					Flags:                 PropertyFlags(0).Set(PropertySessionExpiryInterval),
					SessionExpiryInterval: 10,
				},
				ClientID: []byte("abc"),
				WillProperties: &WillProperties{
					Flags:             PropertyFlags(0).Set(PropertyWillDelayInterval),
					WillDelayInterval: 15,
				},
				WillTopic:   []byte("a"),
				WillPayload: []byte("b"),
			},
		},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			var packet Connect
			header := FixedHeader{PacketType: TypeConnect, RemainingLength: len(test.data)}

			n, err := packet.Decode(test.data, header)
			s.Require().NoError(err)
			s.Assert().Equal(test.packet, packet)
			s.Assert().Equal(len(test.data), n)
		})
	}
}

func (s *ConnectTestSuite) TestDecodeError() {
	testCases := []struct {
		name string
		data []byte
		err  error
	}{
		{"Missing protocol name", []byte{0, 4}, ErrMalformedProtocolName},
		{"Invalid protocol name", []byte{0, 4, 'M', 'Q', 'T', 'T', 3}, ErrMalformedProtocolName},
		{"Missing protocol version", []byte{0, 4, 'M', 'Q', 'T', 'T'}, ErrMalformedProtocolVersion},
		{"Invalid protocol version", []byte{0, 4, 'M', 'Q', 'T', 'T', 0}, ErrMalformedProtocolVersion},
		{"Missing CONNECT flags", []byte{0, 4, 'M', 'Q', 'T', 'T', 4}, ErrMalformedConnectFlags},
		{"Invalid CONNECT flags", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 1}, ErrMalformedConnectFlags},
		{
			"Invalid WillFlag and WillQoS combination",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0x10},
			ErrMalformedConnectFlags,
		},
		{"Invalid WillQoS", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0x1c}, ErrMalformedConnectFlags},
		{
			"Invalid WillFlag and WillRetain combination",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0x20},
			ErrMalformedConnectFlags,
		},
		{
			"Invalid Username and Password flags combination",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0x40},
			ErrMalformedConnectFlags,
		},
		{"Missing keep alive", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0}, ErrMalformedKeepAlive},
		{"Missing property length", []byte{0, 4, 'M', 'Q', 'T', 'T', 5, 0, 0, 10}, ErrMalformedPropertyLength},
		{"Invalid property", []byte{0, 4, 'M', 'Q', 'T', 'T', 5, 0, 0, 10, 1, 255}, ErrMalformedPropertyInvalid},
		{"Missing client ID", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0, 0, 10}, ErrMalformedClientID},
		{
			"Zero-byte Client ID with Clean Session flag (V3.1)",
			[]byte{0, 6, 'M', 'Q', 'I', 's', 'd', 'p', 3, 2, 0, 10, 0, 0},
			ErrV3ClientIDRejected,
		},
		{
			"Missing Will property length",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 5, 4, 0, 10, 0, 0, 0},
			ErrMalformedPropertyLength,
		},
		{
			"Invalid Will property",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 5, 4, 0, 10, 0, 0, 0, 1, 255},
			ErrMalformedPropertyInvalid,
		},
		{"Missing Will topic", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 6, 0, 10, 0, 0}, ErrMalformedWillTopic},
		{"Empty Will topic", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 6, 0, 10, 0, 0, 0, 0}, ErrMalformedWillTopic},
		{"Invalid Will Topic", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 6, 0, 10, 0, 0, 0, 1, '#'}, ErrMalformedWillTopic},
		{
			"Missing Will Payload",
			[]byte{0, 4, 'M', 'Q', 'T', 'T', 4, 6, 0, 10, 0, 0, 0, 1, 'a'},
			ErrMalformedWillPayload,
		},
		{"Missing username", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0x82, 0, 10, 0, 0}, ErrMalformedUsername},
		{"Missing password", []byte{0, 4, 'M', 'Q', 'T', 'T', 5, 0x40, 0, 10, 0, 0, 0}, ErrMalformedPassword},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			var packet Connect
			header := FixedHeader{PacketType: TypeConnect, RemainingLength: len(test.data)}

			_, err := packet.Decode(test.data, header)
			s.Require().ErrorIs(err, test.err)
			s.Assert().NotEmpty(err.Error())
		})
	}
}

func (s *ConnectTestSuite) TestDecodeErrorInvalidHeader() {
	testCases := []struct {
		name   string
		header FixedHeader
		err    error
	}{
		{"Invalid packet type", FixedHeader{PacketType: TypeReserved}, ErrMalformedPacketType},
		{"Invalid flags", FixedHeader{PacketType: TypeConnect, Flags: 1}, ErrMalformedFlags},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			var packet Connect

			_, err := packet.Decode(nil, test.header)
			s.Require().ErrorIs(err, test.err)
		})
	}
}

func TestConnectTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectTestSuite))
}

func BenchmarkConnectSize(b *testing.B) {
	testCases := []struct {
		name   string
		packet Connect
	}{
		{"V3", Connect{Version: MQTT311, KeepAlive: 30, ClientID: []byte("ab")}},
		{"V5", Connect{Version: MQTT50, KeepAlive: 500, ClientID: []byte("abc")}},
	}

	for _, test := range testCases {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.packet.Size()
			}
		})
	}
}

func BenchmarkConnectDecode(b *testing.B) {
	testCases := []struct {
		name string
		data []byte
	}{
		{"V3", []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0, 0, 255, 0, 2, 'a', 'b'}},
		{"V5", []byte{0, 4, 'M', 'Q', 'T', 'T', 5, 0, 255, 255, 0, 0, 3, 'a', 'b', 'c'}},
	}

	for _, test := range testCases {
		b.Run(test.name, func(b *testing.B) {
			header := FixedHeader{PacketType: TypeConnect, RemainingLength: len(test.data)}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				packet := Connect{}

				_, err := packet.Decode(test.data, header)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

type ConnectPropertiesTestSuite struct {
	suite.Suite
}

func (s *ConnectPropertiesTestSuite) TestHas() {
	testCases := []struct {
		props  *ConnectProperties
		id     PropertyID
		result bool
	}{
		{&ConnectProperties{}, PropertyUserProperty, true},
		{&ConnectProperties{}, PropertyAuthenticationMethod, true},
		{&ConnectProperties{}, PropertyAuthenticationData, true},
		{&ConnectProperties{}, PropertySessionExpiryInterval, true},
		{&ConnectProperties{}, PropertyMaximumPacketSize, true},
		{&ConnectProperties{}, PropertyReceiveMaximum, true},
		{&ConnectProperties{}, PropertyTopicAliasMaximum, true},
		{&ConnectProperties{}, PropertyRequestResponseInfo, true},
		{&ConnectProperties{}, PropertyRequestProblemInfo, true},
		{nil, 0, false},
	}

	for _, test := range testCases {
		s.Run(fmt.Sprint(test.id), func() {
			test.props.Set(test.id)
			s.Require().Equal(test.result, test.props.Has(test.id))
		})
	}
}

func (s *ConnectPropertiesTestSuite) TestSize() {
	var flags PropertyFlags

	flags = flags.Set(PropertyUserProperty)
	flags = flags.Set(PropertyAuthenticationMethod)
	flags = flags.Set(PropertyAuthenticationData)
	flags = flags.Set(PropertySessionExpiryInterval)
	flags = flags.Set(PropertyMaximumPacketSize)
	flags = flags.Set(PropertyReceiveMaximum)
	flags = flags.Set(PropertyTopicAliasMaximum)
	flags = flags.Set(PropertyRequestResponseInfo)
	flags = flags.Set(PropertyRequestProblemInfo)

	props := ConnectProperties{
		Flags: flags,
		UserProperties: []UserProperty{
			{[]byte("a"), []byte("b")},
			{[]byte("c"), []byte("d")},
		},
		AuthenticationMethod: []byte("auth"),
		AuthenticationData:   []byte("data"),
	}

	size := props.size()
	s.Assert().Equal(48, size)
}

func (s *ConnectPropertiesTestSuite) TestSizeOnNil() {
	var props *ConnectProperties

	size := props.size()
	s.Assert().Equal(0, size)
}

func (s *ConnectPropertiesTestSuite) TestDecodeSuccess() {
	data := []byte{
		0,               // Property Length
		17, 0, 0, 0, 10, // Session Expiry Interval
		33, 0, 50, // Receive Maximum
		39, 0, 0, 0, 200, // Maximum Packet Size
		34, 0, 50, // Topic Alias Maximum
		25, 1, // Request Response Info
		23, 0, // Request Problem Info
		38, 0, 1, 'a', 0, 1, 'b', // User Property
		38, 0, 1, 'c', 0, 1, 'd', // User Property
		21, 0, 2, 'e', 'f', // Authentication Method
		22, 0, 1, 10, // Authentication Data
	}
	data[0] = byte(len(data) - 1)

	props, n, err := decodeProperties[ConnectProperties](data)
	s.Require().NoError(err)
	s.Assert().Equal(len(data), n)
	s.Assert().True(props.Has(PropertySessionExpiryInterval))
	s.Assert().True(props.Has(PropertyReceiveMaximum))
	s.Assert().True(props.Has(PropertyMaximumPacketSize))
	s.Assert().True(props.Has(PropertyTopicAliasMaximum))
	s.Assert().True(props.Has(PropertyRequestResponseInfo))
	s.Assert().True(props.Has(PropertyRequestProblemInfo))
	s.Assert().True(props.Has(PropertyAuthenticationMethod))
	s.Assert().True(props.Has(PropertyAuthenticationData))
	s.Assert().True(props.Has(PropertyUserProperty))
	s.Assert().Equal(10, int(props.SessionExpiryInterval))
	s.Assert().Equal(50, int(props.ReceiveMaximum))
	s.Assert().Equal(200, int(props.MaximumPacketSize))
	s.Assert().Equal(50, int(props.TopicAliasMaximum))
	s.Assert().True(props.RequestResponseInfo)
	s.Assert().False(props.RequestProblemInfo)
	s.Assert().Equal([]byte("a"), props.UserProperties[0].Key)
	s.Assert().Equal([]byte("b"), props.UserProperties[0].Value)
	s.Assert().Equal([]byte("c"), props.UserProperties[1].Key)
	s.Assert().Equal([]byte("d"), props.UserProperties[1].Value)
	s.Assert().Equal([]byte("ef"), props.AuthenticationMethod)
	s.Assert().Equal([]byte{10}, props.AuthenticationData)
}

func (s *ConnectPropertiesTestSuite) TestDecodeError() {
	testCases := []struct {
		name string
		data []byte
		err  error
	}{
		{"No property", []byte{1}, ErrMalformedPropertyConnect},
		{"Missing Session Expiry Interval", []byte{1, 17}, ErrMalformedPropertySessionExpiryInterval},
		{"Session Expiry Interval - uint", []byte{2, 17, 0}, ErrMalformedPropertySessionExpiryInterval},
		{
			"Duplicated Session Expiry Interval",
			[]byte{10, 17, 0, 0, 0, 10, 17, 0, 0, 0, 11},
			ErrMalformedPropertySessionExpiryInterval,
		},
		{"Missing Receive Max", []byte{1, 33}, ErrMalformedPropertyReceiveMaximum},
		{"Invalid Receive Max", []byte{3, 33, 0, 0}, ErrMalformedPropertyReceiveMaximum},
		{"Duplicated Receive Max", []byte{6, 33, 0, 50, 33, 0, 51}, ErrMalformedPropertyReceiveMaximum},
		{"Missing Maximum Packet Size", []byte{1, 39}, ErrMalformedPropertyMaxPacketSize},
		{"Invalid Maximum Packet Size", []byte{5, 39, 0, 0, 0, 0}, ErrMalformedPropertyMaxPacketSize},
		{
			"Duplicated Maximum Packet Size",
			[]byte{10, 39, 0, 0, 0, 200, 39, 0, 0, 0, 201},
			ErrMalformedPropertyMaxPacketSize,
		},
		{"Missing Topic Alias Max", []byte{1, 34}, ErrMalformedPropertyTopicAliasMaximum},
		{"Duplicated Topic Alias Max", []byte{6, 34, 0, 50, 34, 0, 51}, ErrMalformedPropertyTopicAliasMaximum},
		{"Missing Request Response Info", []byte{1, 25}, ErrMalformedPropertyRequestResponseInfo},
		{"Invalid Request Response Info", []byte{2, 25, 2}, ErrMalformedPropertyRequestResponseInfo},
		{"Duplicated Request Response Info", []byte{4, 25, 0, 25, 1}, ErrMalformedPropertyRequestResponseInfo},
		{"Missing Request Problem Info", []byte{1, 23}, ErrMalformedPropertyRequestProblemInfo},
		{"Invalid Request Problem Info", []byte{2, 23, 2}, ErrMalformedPropertyRequestProblemInfo},
		{"Duplicated Request Problem Info", []byte{4, 23, 0, 23, 1}, ErrMalformedPropertyRequestProblemInfo},
		{"Missing User Property", []byte{1, 38}, ErrMalformedPropertyUserProperty},
		{"Missing User Property Value", []byte{4, 38, 0, 1, 'a'}, ErrMalformedPropertyUserProperty},
		{"User Prop Value - Incomplete str", []byte{4, 38, 0, 5, 'a'}, ErrMalformedPropertyUserProperty},
		{"Missing Authentication Method", []byte{1, 21}, ErrMalformedPropertyAuthenticationMethod},
		{
			"Duplicated Auth Method",
			[]byte{10, 21, 0, 2, 'a', 'b', 21, 0, 2, 'c', 'd'},
			ErrMalformedPropertyAuthenticationMethod,
		},
		{"Missing Auth Data", []byte{1, 22}, ErrMalformedPropertyAuthenticationData},
		{"Duplicated Auth Data", []byte{8, 22, 0, 1, 10, 22, 0, 1, 11}, ErrMalformedPropertyAuthenticationData},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			_, _, err := decodeProperties[ConnectProperties](test.data)
			s.Require().ErrorIs(err, test.err)
		})
	}
}

func TestConnectPropertiesTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectPropertiesTestSuite))
}

type WillPropertiesTestSuite struct {
	suite.Suite
}

func (s *WillPropertiesTestSuite) TestHas() {
	testCases := []struct {
		props  *WillProperties
		id     PropertyID
		result bool
	}{
		{&WillProperties{}, PropertyUserProperty, true},
		{&WillProperties{}, PropertyCorrelationData, true},
		{&WillProperties{}, PropertyContentType, true},
		{&WillProperties{}, PropertyResponseTopic, true},
		{&WillProperties{}, PropertyWillDelayInterval, true},
		{&WillProperties{}, PropertyMessageExpiryInterval, true},
		{&WillProperties{}, PropertyPayloadFormatIndicator, true},
		{nil, 0, false},
	}

	for _, test := range testCases {
		s.Run(fmt.Sprint(test.id), func() {
			test.props.Set(test.id)
			s.Require().Equal(test.result, test.props.Has(test.id))
		})
	}
}

func (s *WillPropertiesTestSuite) TestSize() {
	var flags PropertyFlags
	flags = flags.Set(PropertyUserProperty)
	flags = flags.Set(PropertyCorrelationData)
	flags = flags.Set(PropertyContentType)
	flags = flags.Set(PropertyResponseTopic)
	flags = flags.Set(PropertyWillDelayInterval)
	flags = flags.Set(PropertyMessageExpiryInterval)
	flags = flags.Set(PropertyPayloadFormatIndicator)

	props := WillProperties{
		Flags:                  flags,
		UserProperties:         []UserProperty{{[]byte("a"), []byte("b")}},
		CorrelationData:        []byte{20, 1},
		ContentType:            []byte("json"),
		ResponseTopic:          []byte("b"),
		WillDelayInterval:      10,
		MessageExpiryInterval:  100,
		PayloadFormatIndicator: true,
	}

	size := props.size()
	s.Assert().Equal(35, size)
}

func (s *WillPropertiesTestSuite) TestSizeOnNil() {
	var props *WillProperties

	size := props.size()
	s.Assert().Equal(0, size)
}

func (s *WillPropertiesTestSuite) TestDecodeSuccess() {
	data := []byte{
		0,               // Property Length
		24, 0, 0, 0, 15, // Will Delay Interval
		1, 1, // Payload Format Indicator
		2, 0, 0, 0, 10, // Message Expiry Interval
		3, 0, 4, 'j', 's', 'o', 'n', // Content Type
		8, 0, 1, 'b', // Response Topic
		9, 0, 2, 20, 1, // Correlation Data
		38, 0, 1, 'a', 0, 1, 'b', // User Property
	}
	data[0] = byte(len(data) - 1)

	props, n, err := decodeProperties[WillProperties](data)
	s.Require().NoError(err)
	s.Assert().Equal(len(data), n)
	s.Assert().True(props.Has(PropertyWillDelayInterval))
	s.Assert().True(props.Has(PropertyPayloadFormatIndicator))
	s.Assert().True(props.Has(PropertyMessageExpiryInterval))
	s.Assert().True(props.Has(PropertyContentType))
	s.Assert().True(props.Has(PropertyResponseTopic))
	s.Assert().True(props.Has(PropertyCorrelationData))
	s.Assert().True(props.Has(PropertyUserProperty))
	s.Assert().Equal(15, int(props.WillDelayInterval))
	s.Assert().True(props.PayloadFormatIndicator)
	s.Assert().Equal(10, int(props.MessageExpiryInterval))
	s.Assert().Equal([]byte("json"), props.ContentType)
	s.Assert().Equal([]byte("b"), props.ResponseTopic)
	s.Assert().Equal([]byte{20, 1}, props.CorrelationData)
	s.Assert().Equal([]byte("a"), props.UserProperties[0].Key)
	s.Assert().Equal([]byte("b"), props.UserProperties[0].Value)
}

func (s *WillPropertiesTestSuite) TestDecodeError() {
	testCases := []struct {
		name string
		data []byte
		err  error
	}{
		{"No property", []byte{1}, ErrMalformedPropertyWill},
		{"Missing Will Delay Interval", []byte{1, 24}, ErrMalformedPropertyWillDelayInterval},
		{
			"Duplicated Will Delay Interval",
			[]byte{10, 24, 0, 0, 0, 15, 24, 0, 0, 0, 16},
			ErrMalformedPropertyWillDelayInterval,
		},
		{"Missing Payload Format Indicator", []byte{1, 1}, ErrMalformedPropertyPayloadFormatIndicator},
		{"Invalid Payload Format Indicator", []byte{2, 1, 2}, ErrMalformedPropertyPayloadFormatIndicator},
		{"Duplicated Payload Format Indicator", []byte{4, 1, 0, 1, 1}, ErrMalformedPropertyPayloadFormatIndicator},
		{"Missing Message Expiry Interval", []byte{1, 2}, ErrMalformedPropertyMessageExpiryInterval},
		{
			"Duplicated Message Expiry Interval",
			[]byte{10, 2, 0, 0, 0, 10, 2, 0, 0, 0, 11},
			ErrMalformedPropertyMessageExpiryInterval,
		},
		{"Missing Content Type", []byte{1, 3}, ErrMalformedPropertyContentType},
		{"Content Type - Missing string", []byte{3, 3, 0, 4}, ErrMalformedPropertyContentType},
		{
			"Duplicated Content Type",
			[]byte{13, 3, 0, 4, 'j', 's', 'o', 'n', 3, 0, 3, 'x', 'm', 'l'},
			ErrMalformedPropertyContentType,
		},
		{"Missing Response Topic", []byte{1, 8}, ErrMalformedPropertyResponseTopic},
		{"Response Topic - Missing string", []byte{3, 8, 0, 0}, ErrMalformedPropertyResponseTopic},
		{"Response Topic - Incomplete string", []byte{3, 8, 0, 1}, ErrMalformedPropertyResponseTopic},
		{"Invalid Response Topic", []byte{4, 8, 0, 1, '#'}, ErrMalformedPropertyResponseTopic},
		{"Duplicated Response Topic", []byte{8, 8, 0, 1, 'b', 8, 0, 1, 'c'}, ErrMalformedPropertyResponseTopic},
		{"Missing Correlation Data", []byte{1, 9}, ErrMalformedPropertyCorrelationData},
		{"Correlation Data - Missing data", []byte{3, 9, 0, 2}, ErrMalformedPropertyCorrelationData},
		{
			"Duplicated Correlation Data",
			[]byte{10, 9, 0, 2, 20, 1, 9, 0, 2, 20, 2},
			ErrMalformedPropertyCorrelationData,
		},
	}

	for _, test := range testCases {
		s.Run(test.name, func() {
			_, _, err := decodeProperties[WillProperties](test.data)
			s.Require().ErrorIs(err, test.err)
		})
	}
}

func TestWillPropertiesTestSuite(t *testing.T) {
	suite.Run(t, new(WillPropertiesTestSuite))
}
