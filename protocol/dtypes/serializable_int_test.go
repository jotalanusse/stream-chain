package dtypes_test

import (
	"math/big"
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	big_testutil "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/big"
	"github.com/stretchr/testify/require"
)

var TestCases = map[string]struct {
	b []byte
}{
	"0":                            {b: []byte{0x02}},
	"-0":                           {b: []byte{0x02}},
	"1":                            {b: []byte{0x02, 0x01}},
	"-1":                           {b: []byte{0x03, 0x01}},
	"255":                          {b: []byte{0x02, 0xFF}},
	"-255":                         {b: []byte{0x03, 0xFF}},
	"256":                          {b: []byte{0x02, 0x01, 0x00}},
	"-256":                         {b: []byte{0x03, 0x01, 0x00}},
	"123456789":                    {b: []byte{0x02, 0x07, 0x5b, 0xcd, 0x15}},
	"-123456789":                   {b: []byte{0x03, 0x07, 0x5b, 0xcd, 0x15}},
	"123456789123456789":           {b: []byte{0x02, 0x01, 0xb6, 0x9b, 0x4b, 0xac, 0xd0, 0x5f, 0x15}},
	"-123456789123456789":          {b: []byte{0x03, 0x01, 0xb6, 0x9b, 0x4b, 0xac, 0xd0, 0x5f, 0x15}},
	"123456789123456789123456789":  {b: []byte{0x02, 0x66, 0x1e, 0xfd, 0xf2, 0xe3, 0xb1, 0x9f, 0x7c, 0x04, 0x5f, 0x15}},
	"-123456789123456789123456789": {b: []byte{0x03, 0x66, 0x1e, 0xfd, 0xf2, 0xe3, 0xb1, 0x9f, 0x7c, 0x04, 0x5f, 0x15}},
}

var stringTestCases = map[string]dtypes.SerializableInt{
	"0":                            dtypes.NewIntFromBigInt(big.NewInt(0)),
	"1":                            dtypes.NewIntFromBigInt(big.NewInt(1)),
	"255":                          dtypes.NewIntFromBigInt(big.NewInt(255)),
	"256":                          dtypes.NewIntFromBigInt(big.NewInt(256)),
	"123456789":                    dtypes.NewIntFromBigInt(big.NewInt(123456789)),
	"-123456789":                   dtypes.NewIntFromBigInt(big.NewInt(-123456789)),
	"123456789123456789":           dtypes.NewIntFromBigInt(big.NewInt(123456789123456789)),
	"123456789123456789123456789":  dtypes.NewIntFromBigInt(big_testutil.MustFirst(new(big.Int).SetString("123456789123456789123456789", 10))),
	"-123456789123456789123456789": dtypes.NewIntFromBigInt(big_testutil.MustFirst(new(big.Int).SetString("-123456789123456789123456789", 10))),
	"1_000_000":                    dtypes.NewIntFromBigInt(big.NewInt(1000000)),
	"-1_000_000":                   dtypes.NewIntFromBigInt(big.NewInt(-1000000)),
	"_1__":                         dtypes.NewIntFromBigInt(big.NewInt(1)),
	"__-___1___":                   dtypes.NewIntFromBigInt(big.NewInt(-1)),
}

func TestSerializableInt_ZeroInt(t *testing.T) {
	require.Equal(t, dtypes.ZeroInt().BigInt(), big.NewInt(0))
}

func TestSerializableInt_MaxUint256SerializableInt(t *testing.T) {
	require.Equal(
		t,
		dtypes.MaxUint256SerializableInt(),
		dtypes.NewIntFromBigInt(big_testutil.MustFirst(new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10))),
	)
}

func TestSerializableInt_NewIntFromString(t *testing.T) {
	for input, expected := range stringTestCases {
		actual := dtypes.NewIntFromString(input)
		require.Equal(
			t,
			expected,
			actual,
		)
	}
}

func TestSerializableInt_Marshal(t *testing.T) {
	for name, tc := range TestCases {
		si := dtypes.NewIntFromBigInt(bigIntFromString(name))
		b, err := si.Marshal()
		if err != nil {
			panic(err)
		}
		require.Equal(t, tc.b, b)
	}
}

func TestSerializableInt_MarshalTo(t *testing.T) {
	for name, tc := range TestCases {
		si := dtypes.NewIntFromBigInt(bigIntFromString(name))
		size := si.Size()
		allocBytes := make([]byte, size)
		numBytes, err := si.MarshalTo(allocBytes)
		if err != nil {
			panic(err)
		}
		require.Equal(t, size, numBytes)
		require.Equal(t, tc.b, allocBytes)
	}
}

func TestSerializableInt_Unmarshal(t *testing.T) {
	for name, tc := range TestCases {
		si := new(dtypes.SerializableInt)
		err := si.Unmarshal(tc.b)
		if err != nil {
			panic(err)
		}
		require.Equal(t, bigIntFromString(name), si.BigInt())
	}
}

func TestSerializableInt_Size(t *testing.T) {
	for name := range TestCases {
		si := dtypes.NewIntFromBigInt(bigIntFromString(name))
		b, err := si.Marshal()
		if err != nil {
			panic(err)
		}
		require.Equal(t, len(b), si.Size())
	}
}

func TestSerializableInt_JSON(t *testing.T) {
	for name := range TestCases {
		si := dtypes.NewIntFromBigInt(bigIntFromString(name))
		bi1 := si.BigInt()

		b, err := si.MarshalJSON()
		if err != nil {
			panic(err)
		}

		err = si.UnmarshalJSON(b)
		if err != nil {
			panic(err)
		}
		bi2 := si.BigInt()

		require.Equal(t, bi1, bi2)
	}
}

func bigIntFromString(s string) *big.Int {
	bi, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Cannot create big.Int from string")
	}
	return bi
}
