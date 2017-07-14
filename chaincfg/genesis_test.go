// Copyright (c) 2014-2016 The btcsuite developers
// Copyright (c) 2017 BitGo
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"bytes"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := MainNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := MainNetParams.GenesisBlock.BlockHash()
	if !MainNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(MainNetParams.GenesisHash))
	}

	// Check the size of the block against the header declared size.
	blockSize := uint32(MainNetParams.GenesisBlock.SerializeSize())
	headerSize := MainNetParams.GenesisBlock.Header.Size
	if blockSize != headerSize {
		t.Fatalf("TestGenesisBlock size value: %v does not match the "+
			"header size value %v", blockSize, headerSize)
	}
}

// TestRegTestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
// TODO(prova): Fix by replacing bytes of genesis block
func TestRegTestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := RegressionNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestRegTestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), regTestGenesisBlockBytes) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(regTestGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := RegressionNetParams.GenesisBlock.BlockHash()
	if !RegressionNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(RegressionNetParams.GenesisHash))
	}

	// Check the size of the block against the header declared size.
	blockSize := uint32(RegressionNetParams.GenesisBlock.SerializeSize())
	headerSize := RegressionNetParams.GenesisBlock.Header.Size
	if blockSize != headerSize {
		t.Fatalf("TestRegTestGenesisBlock size value: %v does not "+
			" match the header size value %v", blockSize,
			headerSize)
	}
}

// TestTestNetGenesisBlock tests the genesis block of the test network for
// validity by checking the encoded bytes and hashes.
func TestTestNetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := TestNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestTestNetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNetGenesisBlockBytes) {
		t.Fatalf("TestTestNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := TestNetParams.GenesisBlock.BlockHash()
	if !TestNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestTestNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(TestNetParams.GenesisHash))
	}

	// Check the size of the block against the header declared size.
	blockSize := uint32(TestNetParams.GenesisBlock.SerializeSize())
	headerSize := TestNetParams.GenesisBlock.Header.Size
	if blockSize != headerSize {
		t.Fatalf("TestNetParams size value: %v does not match the "+
			"header size value %v", blockSize, headerSize)
	}
}

// TestSimNetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
func TestSimNetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := SimNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestSimNetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), simNetGenesisBlockBytes) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simNetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash := SimNetParams.GenesisBlock.BlockHash()
	if !SimNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(SimNetParams.GenesisHash))
	}

	// Check the size of the block against the header declared size.
	blockSize := uint32(SimNetParams.GenesisBlock.SerializeSize())
	headerSize := SimNetParams.GenesisBlock.Header.Size
	if blockSize != headerSize {
		t.Fatalf("SimNetParams size value: %v does not match the "+
			"header size value %v", blockSize, headerSize)
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 60002.
var genesisBlockBytes = []byte{
	0x04, 0x00, 0x00, 0x00, /** Block Version */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Prev block hash */
	0xb1, 0x72, 0x5e, 0xa8, 0xf4, 0xc2, 0x0e, 0xd7,
	0x4f, 0xdd, 0x51, 0x1c, 0x4a, 0x65, 0xa0, 0x06,
	0x97, 0x38, 0x8b, 0x91, 0xb0, 0xe5, 0x58, 0xe0,
	0x2e, 0xd0, 0x47, 0xf1, 0x3c, 0x47, 0xee, 0x67, /** Tx merkle root */
	0x7c, 0x30, 0xdc, 0x58, 0x00, 0x00, 0x00, 0x00, /** Timestamp */
	0xff, 0xff, 0x00, 0x1d, /** Difficulty target **/
	0x00, 0x00, 0x00, 0x00, /** Height */
	0x46, 0x01, 0x00, 0x00, /** Block size */
	0x15, 0x11, 0x82, 0xab, 0x00, 0x00, 0x00, 0x00, /** Nonce */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, /** Validate pubkey */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Validate sig */
	0x01,                   /** Number of transactions */
	0x01, 0x00, 0x00, 0x00, /** Transaction version number */
	0x01, /** Number of inputs */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Outpoint tx hash */
	0xff, 0xff, 0xff, 0xff, /** Outpoint index */
	0x20, /** Input script byte count */
	0x8e, 0x5b, 0xd1, 0xf4, 0xc9, 0x0b, 0x80, 0xd8,
	0xf3, 0x12, 0x2b, 0xf4, 0x62, 0x9a, 0x88, 0x21,
	0xc3, 0x6a, 0x70, 0x6d, 0x4c, 0xb1, 0x8b, 0xa9,
	0xa7, 0xa5, 0xe9, 0xa5, 0x87, 0xee, 0x48, 0xfe, /** Script bytes */
	0xff, 0xff, 0xff, 0xff, /** Sequence number */
	0x03,                                           /** Output count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x00, 0xbb, /** Admin script output: root thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x51, 0xbb, /** Admin script output: provision thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x52, 0xbb, /** Admin script output: issue thread */
	0x00, 0x00, 0x00, 0x00, /** Lock Time */
}

// regTestGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 60002.
var regTestGenesisBlockBytes = []byte{
	0x04, 0x00, 0x00, 0x00, /** Block Version */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Prev block hash */
	0xd7, 0xb2, 0xc8, 0xb0, 0xfa, 0x10, 0xc8, 0xba,
	0x60, 0x1b, 0x0c, 0x80, 0xcc, 0x1b, 0x3a, 0x2a,
	0xf6, 0x20, 0x5f, 0x5f, 0x6f, 0xc1, 0xec, 0xaf,
	0xec, 0xae, 0x26, 0xde, 0x01, 0x3e, 0xa8, 0xba, /** Tx merkle root */
	0x7c, 0x30, 0xdc, 0x58, 0x00, 0x00, 0x00, 0x00, /** Timestamp */
	0x0f, 0x0f, 0x0f, 0x20, /** Difficulty target **/
	0x00, 0x00, 0x00, 0x00, /** Height */
	0x46, 0x01, 0x00, 0x00, /** Block size */
	0x09, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Nonce */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, /** Validate pubkey */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Validate sig */
	0x01,                   /** Number of transactions */
	0x01, 0x00, 0x00, 0x00, /** Transaction version number */
	0x01, /** Number of inputs */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Outpoint tx hash */
	0xff, 0xff, 0xff, 0xff, /** Outpoint index */
	0x20, /** Input script byte count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x35, 0x80, 0x85, 0x7b, 0xe1, 0x75, 0xc7,
	0xd5, 0x77, 0xb8, 0x5d, 0xb2, 0xe0, 0x06, 0xd5,
	0xb8, 0x91, 0x4e, 0x64, 0xab, 0xb7, 0x87, 0x1c, /** Script bytes */
	0xff, 0xff, 0xff, 0xff, /** Sequence number */
	0x03,                                           /** Output count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x00, 0xbb, /** Admin script output: root thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x51, 0xbb, /** Admin script output: provision thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x52, 0xbb, /** Admin script output: issue thread */
	0x00, 0x00, 0x00, 0x00, /** Lock Time */
}

// testNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network as of protocol version 60002.
var testNetGenesisBlockBytes = []byte{
	0x04, 0x00, 0x00, 0x00, /** Block Version */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Prev block hash */
	0xd7, 0xb2, 0xc8, 0xb0, 0xfa, 0x10, 0xc8, 0xba,
	0x60, 0x1b, 0x0c, 0x80, 0xcc, 0x1b, 0x3a, 0x2a,
	0xf6, 0x20, 0x5f, 0x5f, 0x6f, 0xc1, 0xec, 0xaf,
	0xec, 0xae, 0x26, 0xde, 0x01, 0x3e, 0xa8, 0xba, /** Tx merkle root */
	0x7c, 0x30, 0xdc, 0x58, 0x00, 0x00, 0x00, 0x00, /** Timestamp */
	0xff, 0xff, 0x07, 0x20, /** Difficulty target **/
	0x00, 0x00, 0x00, 0x00, /** Height */
	0x46, 0x01, 0x00, 0x00, /** Block size */
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Nonce */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, /** Validate pubkey */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Validate sig */
	0x01,                   /** Number of transactions */
	0x01, 0x00, 0x00, 0x00, /** Transaction version number */
	0x01, /** Number of inputs */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Outpoint tx hash */
	0xff, 0xff, 0xff, 0xff, /** Outpoint index */
	0x20, /** Input script byte count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x35, 0x80, 0x85, 0x7b, 0xe1, 0x75, 0xc7,
	0xd5, 0x77, 0xb8, 0x5d, 0xb2, 0xe0, 0x06, 0xd5,
	0xb8, 0x91, 0x4e, 0x64, 0xab, 0xb7, 0x87, 0x1c, /** Script bytes */
	0xff, 0xff, 0xff, 0xff, /** Sequence number */
	0x03,                                           /** Output count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x00, 0xbb, /** Admin script output: root thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x51, 0xbb, /** Admin script output: provision thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x52, 0xbb, /** Admin script output: issue thread */
	0x00, 0x00, 0x00, 0x00, /** Lock Time */
}

// simNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the simulation test network as of protocol version 70002.
var simNetGenesisBlockBytes = []byte{
	0x04, 0x00, 0x00, 0x00, /** Block Version */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Prev block hash */
	0xd7, 0xb2, 0xc8, 0xb0, 0xfa, 0x10, 0xc8, 0xba,
	0x60, 0x1b, 0x0c, 0x80, 0xcc, 0x1b, 0x3a, 0x2a,
	0xf6, 0x20, 0x5f, 0x5f, 0x6f, 0xc1, 0xec, 0xaf,
	0xec, 0xae, 0x26, 0xde, 0x01, 0x3e, 0xa8, 0xba, /** Tx merkle root */
	0x7c, 0x30, 0xdc, 0x58, 0x00, 0x00, 0x00, 0x00, /** Timestamp */
	0xff, 0xff, 0x7f, 0x20, /** Difficulty target **/
	0x00, 0x00, 0x00, 0x00, /** Height */
	0x46, 0x01, 0x00, 0x00, /** Block size */
	0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Nonce */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, /** Validate pubkey */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,

	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Validate sig */
	0x01,                   /** Number of transactions */
	0x01, 0x00, 0x00, 0x00, /** Transaction version number */
	0x01, /** Number of inputs */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Outpoint tx hash */
	0xff, 0xff, 0xff, 0xff, /** Outpoint index */
	0x20, /** Input script byte count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x35, 0x80, 0x85, 0x7b, 0xe1, 0x75, 0xc7,
	0xd5, 0x77, 0xb8, 0x5d, 0xb2, 0xe0, 0x06, 0xd5,
	0xb8, 0x91, 0x4e, 0x64, 0xab, 0xb7, 0x87, 0x1c, /** Script bytes */
	0xff, 0xff, 0xff, 0xff, /** Sequence number */
	0x03,                                           /** Output count */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x00, 0xbb, /** Admin script output: root thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x51, 0xbb, /** Admin script output: provision thread */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /** Value */
	0x02,       /** Script byte count */
	0x52, 0xbb, /** Admin script output: issue thread */
	0x00, 0x00, 0x00, 0x00, /** Lock Time */
}
