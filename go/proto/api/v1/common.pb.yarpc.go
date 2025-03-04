// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/common.proto

package apiv1

var yarpcFileDescriptorClosure0ff151d4a308b356 = [][]byte{
	// uber/cadence/api/v1/common.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
		0x14, 0x9d, 0xe2, 0xd8, 0x69, 0xaf, 0xdd, 0xd4, 0x63, 0xd6, 0xd4, 0xc9, 0xbe, 0x3c, 0x03, 0x43,
		0xb3, 0x01, 0x93, 0x11, 0xf7, 0xa5, 0x58, 0x51, 0x0c, 0x49, 0xec, 0xac, 0x6a, 0xb7, 0xc4, 0x90,
		0x8d, 0x06, 0xdb, 0x80, 0x09, 0xb4, 0x78, 0xe5, 0x72, 0x96, 0x48, 0x81, 0xa2, 0x9c, 0xf8, 0x6d,
		0xbf, 0x64, 0x0f, 0xfb, 0x4b, 0xfb, 0x43, 0x03, 0x25, 0x3a, 0x76, 0x3a, 0x0f, 0x7d, 0x19, 0xf6,
		0x46, 0xde, 0x73, 0xee, 0xb9, 0xe7, 0x12, 0x97, 0x24, 0xb4, 0xf3, 0x09, 0xaa, 0x6e, 0x48, 0x19,
		0x8a, 0x10, 0xbb, 0x34, 0xe5, 0xdd, 0xf9, 0x71, 0x37, 0x94, 0x49, 0x22, 0x85, 0x9b, 0x2a, 0xa9,
		0x25, 0xd9, 0x33, 0x0c, 0xd7, 0x32, 0x5c, 0x9a, 0x72, 0x77, 0x7e, 0x7c, 0xf8, 0xd9, 0x54, 0xca,
		0x69, 0x8c, 0xdd, 0x82, 0x32, 0xc9, 0xa3, 0x2e, 0xcb, 0x15, 0xd5, 0x7c, 0x99, 0xd4, 0x79, 0x0d,
		0x1f, 0x5e, 0x49, 0x35, 0x8b, 0x62, 0x79, 0x3d, 0xb8, 0xc1, 0x30, 0x37, 0x10, 0xf9, 0x1c, 0xea,
		0xd7, 0x36, 0x18, 0x70, 0xd6, 0x72, 0xda, 0xce, 0xd1, 0x7d, 0x1f, 0x96, 0x21, 0x8f, 0x91, 0x47,
		0x50, 0x53, 0xb9, 0x30, 0xd8, 0x56, 0x81, 0x55, 0x55, 0x2e, 0x3c, 0xd6, 0xe9, 0x40, 0x63, 0x29,
		0x36, 0x5e, 0xa4, 0x48, 0x08, 0x6c, 0x0b, 0x9a, 0xa0, 0x15, 0x28, 0xd6, 0x86, 0x73, 0x12, 0x6a,
		0x3e, 0xe7, 0x7a, 0xf1, 0xaf, 0x9c, 0x4f, 0x61, 0x67, 0x48, 0x17, 0xb1, 0xa4, 0xcc, 0xc0, 0x8c,
		0x6a, 0x5a, 0xc0, 0x0d, 0xbf, 0x58, 0x77, 0x9e, 0xc3, 0xce, 0x39, 0xe5, 0x71, 0xae, 0x90, 0xec,
		0x43, 0x4d, 0x21, 0xcd, 0xa4, 0xb0, 0xf9, 0x76, 0x47, 0x5a, 0xb0, 0xc3, 0x50, 0x53, 0x1e, 0x67,
		0x85, 0xc3, 0x86, 0xbf, 0xdc, 0x76, 0xfe, 0x70, 0x60, 0xfb, 0x47, 0x4c, 0x24, 0x79, 0x01, 0xb5,
		0x88, 0x63, 0xcc, 0xb2, 0x96, 0xd3, 0xae, 0x1c, 0xd5, 0x7b, 0x5f, 0xba, 0x1b, 0xce, 0xcf, 0x35,
		0x54, 0xf7, 0xbc, 0xe0, 0x0d, 0x84, 0x56, 0x0b, 0xdf, 0x26, 0x1d, 0x5e, 0x41, 0x7d, 0x2d, 0x4c,
		0x9a, 0x50, 0x99, 0xe1, 0xc2, 0xba, 0x30, 0x4b, 0xd2, 0x83, 0xea, 0x9c, 0xc6, 0x39, 0x16, 0x06,
		0xea, 0xbd, 0x4f, 0x36, 0xca, 0xdb, 0x36, 0xfd, 0x92, 0xfa, 0xed, 0xd6, 0x33, 0xa7, 0xf3, 0xa7,
		0x03, 0xb5, 0x97, 0x48, 0x19, 0x2a, 0xf2, 0xdd, 0x3b, 0x16, 0x9f, 0x6c, 0xd4, 0x28, 0xc9, 0xff,
		0xaf, 0xc9, 0xbf, 0x1c, 0x68, 0x8e, 0x90, 0xaa, 0xf0, 0xed, 0x89, 0xd6, 0x8a, 0x4f, 0x72, 0x8d,
		0x19, 0x09, 0x60, 0x97, 0x0b, 0x86, 0x37, 0xc8, 0x82, 0x3b, 0xb6, 0x9f, 0x6d, 0x54, 0x7d, 0x37,
		0xdd, 0xf5, 0xca, 0xdc, 0xf5, 0x3e, 0x1e, 0xf0, 0xf5, 0xd8, 0xe1, 0xaf, 0x40, 0xfe, 0x49, 0xfa,
		0x0f, 0xbb, 0x8a, 0xe0, 0x5e, 0x9f, 0x6a, 0x7a, 0x1a, 0xcb, 0x09, 0x39, 0x87, 0x07, 0x28, 0x42,
		0xc9, 0xb8, 0x98, 0x06, 0x7a, 0x91, 0x96, 0x03, 0xba, 0xdb, 0xfb, 0x62, 0xa3, 0xd6, 0xc0, 0x32,
		0xcd, 0x44, 0xfb, 0x0d, 0x5c, 0xdb, 0xdd, 0x0e, 0xf0, 0xd6, 0xda, 0x00, 0x0f, 0xcb, 0x4b, 0x87,
		0xea, 0x0d, 0xaa, 0x8c, 0x4b, 0xe1, 0x89, 0x48, 0x1a, 0x22, 0x4f, 0xd2, 0x78, 0x79, 0x11, 0xcc,
		0x9a, 0x3c, 0x81, 0x87, 0x11, 0x52, 0x9d, 0x2b, 0x0c, 0xe6, 0x25, 0xd5, 0x5e, 0xb8, 0x5d, 0x1b,
		0xb6, 0x02, 0x9d, 0xd7, 0xf0, 0x78, 0x94, 0xa7, 0xa9, 0x54, 0x1a, 0xd9, 0x59, 0xcc, 0x51, 0x68,
		0x8b, 0x64, 0xe6, 0xae, 0x4e, 0x65, 0x90, 0xb1, 0x99, 0x55, 0xae, 0x4e, 0xe5, 0x88, 0xcd, 0xc8,
		0x01, 0xdc, 0xfb, 0x8d, 0xce, 0x69, 0x01, 0x94, 0x9a, 0x3b, 0x66, 0x3f, 0x62, 0xb3, 0xce, 0xef,
		0x15, 0xa8, 0xfb, 0xa8, 0xd5, 0x62, 0x28, 0x63, 0x1e, 0x2e, 0x48, 0x1f, 0x9a, 0x5c, 0x70, 0xcd,
		0x69, 0x1c, 0x70, 0xa1, 0x51, 0xcd, 0x69, 0xe9, 0xb2, 0xde, 0x3b, 0x70, 0xcb, 0xe7, 0xc5, 0x5d,
		0x3e, 0x2f, 0x6e, 0xdf, 0x3e, 0x2f, 0xfe, 0x43, 0x9b, 0xe2, 0xd9, 0x0c, 0xd2, 0x85, 0xbd, 0x09,
		0x0d, 0x67, 0x32, 0x8a, 0x82, 0x50, 0x62, 0x14, 0xf1, 0xd0, 0xd8, 0x2c, 0x6a, 0x3b, 0x3e, 0xb1,
		0xd0, 0xd9, 0x0a, 0x31, 0x65, 0x13, 0x7a, 0xc3, 0x93, 0x3c, 0x59, 0x95, 0xad, 0xbc, 0xb7, 0xac,
		0x4d, 0xb9, 0x2d, 0xfb, 0xd5, 0x4a, 0x85, 0x6a, 0x8d, 0x49, 0xaa, 0xb3, 0xd6, 0x76, 0xdb, 0x39,
		0xaa, 0xde, 0x52, 0x4f, 0x6c, 0x98, 0xbc, 0x80, 0x8f, 0x85, 0x14, 0x81, 0x32, 0xad, 0xd3, 0x49,
		0x8c, 0x01, 0x2a, 0x25, 0x55, 0x50, 0x3e, 0x29, 0x59, 0xab, 0xda, 0xae, 0x1c, 0xdd, 0xf7, 0x5b,
		0x42, 0x0a, 0x7f, 0xc9, 0x18, 0x18, 0x82, 0x5f, 0xe2, 0xe4, 0x15, 0xec, 0xe1, 0x4d, 0xca, 0x4b,
		0x23, 0x2b, 0xcb, 0xb5, 0xf7, 0x59, 0x26, 0xab, 0xac, 0xa5, 0xeb, 0xaf, 0xaf, 0xa1, 0xb1, 0x3e,
		0x53, 0xe4, 0x00, 0x1e, 0x0d, 0x2e, 0xce, 0x2e, 0xfb, 0xde, 0xc5, 0xf7, 0xc1, 0xf8, 0xa7, 0xe1,
		0x20, 0xf0, 0x2e, 0xde, 0x9c, 0xfc, 0xe0, 0xf5, 0x9b, 0x1f, 0x90, 0x43, 0xd8, 0xbf, 0x0b, 0x8d,
		0x5f, 0xfa, 0xde, 0xf9, 0xd8, 0xbf, 0x6a, 0x3a, 0x64, 0x1f, 0xc8, 0x5d, 0xec, 0xd5, 0xe8, 0xf2,
		0xa2, 0xb9, 0x45, 0x5a, 0xf0, 0xd1, 0xdd, 0xf8, 0xd0, 0xbf, 0x1c, 0x5f, 0x3e, 0x6d, 0x56, 0x4e,
		0x7f, 0x81, 0xc7, 0xa1, 0x4c, 0x36, 0x0d, 0xf9, 0x69, 0xfd, 0xac, 0xf8, 0x6d, 0x86, 0xa6, 0x81,
		0xa1, 0xf3, 0xf3, 0xf1, 0x94, 0xeb, 0xb7, 0xf9, 0xc4, 0x0d, 0x65, 0xd2, 0x5d, 0xff, 0x9b, 0xbe,
		0xe1, 0x2c, 0xee, 0x4e, 0x65, 0xf9, 0xe3, 0xd8, 0x8f, 0xea, 0x39, 0x4d, 0xf9, 0xfc, 0x78, 0x52,
		0x2b, 0x62, 0x4f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x13, 0xdb, 0xef, 0xb7, 0xcc, 0x06, 0x00,
		0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
}
