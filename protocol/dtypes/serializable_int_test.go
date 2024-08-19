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

func TestSerializableInt_Add_Success(t *testing.T) {
	testCases := map[string]struct {
		x      dtypes.SerializableInt
		y      dtypes.SerializableInt
		output dtypes.SerializableInt
	}{
		"zero plus zero": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(0),
		},
		"positive plus positive": {
			x:      dtypes.NewInt(5),
			y:      dtypes.NewInt(3),
			output: dtypes.NewInt(8),
		},
		"negative plus positive": {
			x:      dtypes.NewInt(-10),
			y:      dtypes.NewInt(7),
			output: dtypes.NewInt(-3),
		},
		"zero plus positive": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(100),
			output: dtypes.NewInt(100),
		},
		"positive plus zero": {
			x:      dtypes.NewInt(100),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(100),
		},
		"negative plus negative": {
			x:      dtypes.NewInt(-50),
			y:      dtypes.NewInt(-30),
			output: dtypes.NewInt(-80),
		},
		"positive plus its negative": {
			x:      dtypes.NewInt(1000000),
			y:      dtypes.NewInt(-1000000),
			output: dtypes.NewInt(0),
		},
		"large positive plus large positive": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("2_000_000_000_000_000_000_000_000_000"),
		},
		"large positive plus its negative": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("0"),
		},
		"large negative plus large negative": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-2_000_000_000_000_000_000_000_000_000"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			z := dtypes.NewInt(0)
			actualOutput := z.Add(
				tc.x,
				tc.y,
			)
			actualOutputEqual := tc.output.Cmp(actualOutput)
			require.Equal(t, 0, actualOutputEqual)
			zEqual := tc.output.Cmp(z)
			require.Equal(t, 0, zEqual)
		})
	}
}

func TestSerializableInt_Add_Fail(t *testing.T) {
	testCases := map[string]struct {
		x dtypes.SerializableInt
		y dtypes.SerializableInt
	}{
		"x is nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.NewInt(1),
		},
		"y is nil": {
			x: dtypes.NewInt(1),
			y: dtypes.SerializableInt{},
		},
		"x and y are nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.SerializableInt{},
		},
	}

	for _, tc := range testCases {
		z := dtypes.NewInt(0)
		require.Panics(t, func() { z.Add(tc.x, tc.y) })
	}
}

func TestSerializableInt_Sub_Success(t *testing.T) {
	testCases := map[string]struct {
		x      dtypes.SerializableInt
		y      dtypes.SerializableInt
		output dtypes.SerializableInt
	}{
		"zero minus zero": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(0),
		},
		"positive minus positive": {
			x:      dtypes.NewInt(5),
			y:      dtypes.NewInt(3),
			output: dtypes.NewInt(2),
		},
		"negative minus positive": {
			x:      dtypes.NewInt(-10),
			y:      dtypes.NewInt(7),
			output: dtypes.NewInt(-17),
		},
		"zero minus positive": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(100),
			output: dtypes.NewInt(-100),
		},
		"positive minus zero": {
			x:      dtypes.NewInt(100),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(100),
		},
		"negative minus negative": {
			x:      dtypes.NewInt(-50),
			y:      dtypes.NewInt(-30),
			output: dtypes.NewInt(-20),
		},
		"positive minus its negative": {
			x:      dtypes.NewInt(1000000),
			y:      dtypes.NewInt(-1000000),
			output: dtypes.NewInt(2000000),
		},
		"large positive minus large positive": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_001"),
			output: dtypes.NewIntFromString("-1"),
		},
		"large positive minus its negative": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("2_000_000_000_000_000_000_000_000_000"),
		},
		"large negative minus large negative": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("0"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			z := dtypes.NewInt(0)
			actualOutput := z.Sub(
				tc.x,
				tc.y,
			)
			actualOutputEqual := tc.output.Cmp(actualOutput)
			require.Equal(t, 0, actualOutputEqual)
			zEqual := tc.output.Cmp(z)
			require.Equal(t, 0, zEqual)
		})
	}
}

func TestSerializableInt_Sub_Fail(t *testing.T) {
	testCases := map[string]struct {
		x dtypes.SerializableInt
		y dtypes.SerializableInt
	}{
		"x is nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.NewInt(1),
		},
		"y is nil": {
			x: dtypes.NewInt(1),
			y: dtypes.SerializableInt{},
		},
		"x and y are nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.SerializableInt{},
		},
	}

	for _, tc := range testCases {
		z := dtypes.NewInt(0)
		require.Panics(t, func() { z.Sub(tc.x, tc.y) })
	}
}

func TestSerializableInt_Mul_Success(t *testing.T) {
	testCases := map[string]struct {
		x      dtypes.SerializableInt
		y      dtypes.SerializableInt
		output dtypes.SerializableInt
	}{
		"zero times zero": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(0),
		},
		"positive times positive": {
			x:      dtypes.NewInt(5),
			y:      dtypes.NewInt(3),
			output: dtypes.NewInt(15),
		},
		"negative times positive": {
			x:      dtypes.NewInt(-10),
			y:      dtypes.NewInt(7),
			output: dtypes.NewInt(-70),
		},
		"zero times positive": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(100),
			output: dtypes.NewInt(0),
		},
		"positive times zero": {
			x:      dtypes.NewInt(100),
			y:      dtypes.NewInt(0),
			output: dtypes.NewInt(0),
		},
		"negative times negative": {
			x:      dtypes.NewInt(-50),
			y:      dtypes.NewInt(-30),
			output: dtypes.NewInt(1500),
		},
		"positive times its negative": {
			x:      dtypes.NewInt(1000000),
			y:      dtypes.NewInt(-1000000),
			output: dtypes.NewIntFromString("-1_000_000_000_000"),
		},
		"large positive times large positive": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000"),
		},
		"large positive times its negative": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000"),
		},
		"large negative times large negative": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			z := dtypes.NewInt(0)
			actualOutput := z.Mul(
				tc.x,
				tc.y,
			)
			actualOutputEqual := tc.output.Cmp(actualOutput)
			require.Equal(t, 0, actualOutputEqual)
			zEqual := tc.output.Cmp(z)
			require.Equal(t, 0, zEqual)
		})
	}
}

func TestSerializableInt_Mul_Fail(t *testing.T) {
	testCases := map[string]struct {
		x dtypes.SerializableInt
		y dtypes.SerializableInt
	}{
		"x is nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.NewInt(1),
		},
		"y is nil": {
			x: dtypes.NewInt(1),
			y: dtypes.SerializableInt{},
		},
		"x and y are nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.SerializableInt{},
		},
	}

	for _, tc := range testCases {
		z := dtypes.NewInt(0)
		require.Panics(t, func() { z.Mul(tc.x, tc.y) })
	}
}

func TestSerializableInt_Div_Success(t *testing.T) {
	testCases := map[string]struct {
		x      dtypes.SerializableInt
		y      dtypes.SerializableInt
		output dtypes.SerializableInt
	}{
		"zero divided by positive": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(1),
			output: dtypes.NewInt(0),
		},
		"zero divided by negative": {
			x:      dtypes.NewInt(0),
			y:      dtypes.NewInt(-1),
			output: dtypes.NewInt(0),
		},
		"positive divided by positive not rounded": {
			x:      dtypes.NewInt(6),
			y:      dtypes.NewInt(3),
			output: dtypes.NewInt(2),
		},
		"positive divided by positive rounded": {
			x:      dtypes.NewInt(5),
			y:      dtypes.NewInt(3),
			output: dtypes.NewInt(1),
		},
		"positive divided by negative not rounded": {
			x:      dtypes.NewInt(1500),
			y:      dtypes.NewInt(-5),
			output: dtypes.NewInt(-300),
		},
		"positive divided by negative rounded": {
			x:      dtypes.NewInt(2400),
			y:      dtypes.NewInt(-7),
			output: dtypes.NewInt(-342),
		},
		"negative divided by positive not rounded": {
			x:      dtypes.NewInt(-20),
			y:      dtypes.NewInt(5),
			output: dtypes.NewInt(-4),
		},
		"negative divided by positive rounded": {
			x:      dtypes.NewInt(-20),
			y:      dtypes.NewInt(7),
			output: dtypes.NewInt(-3),
		},
		"negative divided by negative not rounded": {
			x:      dtypes.NewInt(-60),
			y:      dtypes.NewInt(-30),
			output: dtypes.NewInt(2),
		},
		"negative divided by negative rounded": {
			x:      dtypes.NewInt(-50),
			y:      dtypes.NewInt(-30),
			output: dtypes.NewInt(2),
		},
		"large positive divided by large positive not rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("1"),
		},
		"large positive divided by large positive rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_001"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("1"),
		},
		"large positive divided by large negative not rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-1"),
		},
		"large positive divided by large negative rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_001"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-1"),
		},
		"large negative divided by large positive not rounded": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-1"),
		},
		"large negative divided by large positive rounded": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_001"),
			y:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("-2"),
		},
		"large negative divided by large negative not rounded": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("1"),
		},
		"large negative divided by large negative rounded": {
			x:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_001"),
			y:      dtypes.NewIntFromString("-1_000_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("2"),
		},
		"large positive divided by small positive not rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_000"),
			y:      dtypes.NewIntFromString("200"),
			output: dtypes.NewIntFromString("5_000_000_000_000_000_000_000_000"),
		},
		"large positive divided by small positive rounded": {
			x:      dtypes.NewIntFromString("1_000_000_000_000_000_000_000_000_001"),
			y:      dtypes.NewIntFromString("200"),
			output: dtypes.NewIntFromString("5_000_000_000_000_000_000_000_000"),
		},
		"small positive divided by large positive rounded": {
			x:      dtypes.NewIntFromString("2"),
			y:      dtypes.NewIntFromString("5_000_000_000_000_000_000_000_000"),
			output: dtypes.NewIntFromString("0"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			z := dtypes.NewInt(0)
			actualOutput := z.Div(
				tc.x,
				tc.y,
			)
			actualOutputEqual := tc.output.Cmp(actualOutput)
			require.Equal(t, 0, actualOutputEqual)
			zEqual := tc.output.Cmp(z)
			require.Equal(t, 0, zEqual)
		})
	}
}

func TestSerializableInt_Div_Fail(t *testing.T) {
	testCases := map[string]struct {
		x dtypes.SerializableInt
		y dtypes.SerializableInt
	}{
		"y is zero": {
			x: dtypes.NewInt(1),
			y: dtypes.NewInt(0),
		},
		"x is nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.NewInt(1),
		},
		"y is nil": {
			x: dtypes.NewInt(1),
			y: dtypes.SerializableInt{},
		},
		"x and y are nil": {
			x: dtypes.SerializableInt{},
			y: dtypes.SerializableInt{},
		},
	}

	for _, tc := range testCases {
		z := dtypes.NewInt(0)
		require.Panics(t, func() { z.Div(tc.x, tc.y) })
	}
}

func bigIntFromString(s string) *big.Int {
	bi, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Cannot create big.Int from string")
	}
	return bi
}
