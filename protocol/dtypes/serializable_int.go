package dtypes

import (
	"encoding"
	"encoding/json"
	"math/big"
	"strings"
)

// SerializableInt is basically copied from cosmos-sdk/types/Int but:
// - doesn’t have a bit-length restriction
// - uses GobEncode/GobDecode instead of serializing to an ascii string
// - removes superfluous functions to do `big.Int` math on the underlying value
type SerializableInt struct {
	i *big.Int
}

// BigInt converts Int to big.Int
func (i SerializableInt) BigInt() *big.Int {
	if i.IsNil() {
		return nil
	}
	return new(big.Int).Set(i.i)
}

// IsNil returns true if Int is uninitialized
func (i SerializableInt) IsNil() bool {
	return i.i == nil
}

// NewInt constructs Int from int64
func NewInt(n int64) SerializableInt {
	return SerializableInt{big.NewInt(n)}
}

// NewIntFromUint64 constructs an Int from a uint64.
func NewIntFromUint64(n uint64) SerializableInt {
	b := big.NewInt(0)
	b.SetUint64(n)
	return SerializableInt{b}
}

// NewIntFromBigInt constructs Int from big.Int. If the provided big.Int is nil,
func NewIntFromBigInt(i *big.Int) SerializableInt {
	if i == nil {
		return SerializableInt{}
	}

	return SerializableInt{i}
}

// NewIntFromString constructs an Int from a String. The string is expected to
// to represent a number in base 10. A string can contain underscores. Underscores
// are interpreted as empty strings.
func NewIntFromString(i string) SerializableInt {

	i = strings.ReplaceAll(i, "_", "")

	b, ok := new(big.Int).SetString(i, 10)
	if !ok {
		return SerializableInt{}
	}

	return NewIntFromBigInt(b)
}

// ZeroInt returns Int value with zero
func ZeroInt() SerializableInt {
	return SerializableInt{big.NewInt(0)}
}

// Creates the maximum serializable int (2^256 - 1)
func MaxUint256SerializableInt() SerializableInt {
	max := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	return NewIntFromBigInt(max)
}

func (i SerializableInt) String() string {
	if i.IsNil() {
		return "nil"
	}
	return i.BigInt().String()
}

// Add sets z to the sum of x and y and returns z.
// This is similar to big.Int.Add.
func (z *SerializableInt) Add(x SerializableInt, y SerializableInt) SerializableInt {
	if x.IsNil() {
		panic("SerializableInt: cannot perform addition with nil-value for x")
	}

	if y.IsNil() {
		panic("SerializableInt: cannot perform addition with nil-value for y")
	}

	bigIntAdd := big.NewInt(0).Add(
		x.BigInt(),
		y.BigInt(),
	)

	z.i = bigIntAdd

	// TODO: [SINT-1]
	return NewIntFromBigInt(bigIntAdd)
}

// Sub sets z to the difference of x and y and returns z.
// This is similar to big.Int.Sub.
func (z *SerializableInt) Sub(x SerializableInt, y SerializableInt) SerializableInt {
	if x.IsNil() {
		panic("SerializableInt: cannot perform subtraction with nil-value for x")
	}

	if y.IsNil() {
		panic("SerializableInt: cannot perform subtraction with nil-value for y")
	}

	bigIntSub := big.NewInt(0).Sub(
		x.BigInt(),
		y.BigInt(),
	)

	z.i = bigIntSub

	// TODO: [SINT-1]
	return NewIntFromBigInt(bigIntSub)
}

// Mul sets z to the product of x and y and returns z.
// This is similar to big.Int.Mul.
func (z *SerializableInt) Mul(x SerializableInt, y SerializableInt) SerializableInt {
	if x.IsNil() {
		panic("SerializableInt: cannot perform multiplication with nil-value for x")
	}

	if y.IsNil() {
		panic("SerializableInt: cannot perform multiplication with nil-value for y")
	}

	bigIntMul := big.NewInt(0).Mul(
		x.BigInt(),
		y.BigInt(),
	)

	z.i = bigIntMul

	// TODO: [SINT-1]
	return NewIntFromBigInt(bigIntMul)
}

// Div sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements a Euclidean division.
// This is similar to big.Int.Div.
func (z *SerializableInt) Div(x SerializableInt, y SerializableInt) SerializableInt {
	if x.IsNil() {
		panic("SerializableInt: cannot perform division with nil-value for x")
	}

	if y.IsNil() {
		panic("SerializableInt: cannot perform division with nil-value for y")
	}

	bigIntDiv := big.NewInt(0).Div(
		x.BigInt(),
		y.BigInt(),
	)

	z.i = bigIntDiv

	// TODO: [SINT-1]
	return NewIntFromBigInt(bigIntDiv)
}

// Cmp compares x and y and returns:
//
//	-1 if (x <  y) OR (x is nil and y is not nil)
//	 0 if (x == y) OR (x is nil and y is nil)
//	+1 if (x >  y) OR (x is not nil and y is nil
//
// This is similar to big.Int.Cmp where nil values sort first.
func (i SerializableInt) Cmp(j SerializableInt) int {
	if i.IsNil() {
		if j.IsNil() {
			return 0
		}
		return -1
	}
	if j.IsNil() {
		return 1
	}
	return i.BigInt().Cmp(j.BigInt())
}

// Marshal implements the gogo proto custom type interface.
func (i SerializableInt) Marshal() ([]byte, error) {
	i.ensureNonNil()
	return i.i.GobEncode()
}

// MarshalTo implements the gogo proto custom type interface.
func (i *SerializableInt) MarshalTo(data []byte) (n int, err error) {
	bz, err := i.Marshal()
	if err != nil {
		return 0, err
	}

	n = copy(data, bz)
	return n, nil
}

// Unmarshal implements the gogo proto custom type interface.
func (i *SerializableInt) Unmarshal(data []byte) error {
	i.ensureNonNil()

	if err := i.i.GobDecode(data); err != nil {
		return err
	}

	return nil
}

// Size implements the gogo proto custom type interface.
func (i *SerializableInt) Size() int {
	i.ensureNonNil()
	n := i.i.BitLen()
	return 1 + ((n + 7) / 8)
}

// MarshalJSON defines custom encoding scheme
func (i SerializableInt) MarshalJSON() ([]byte, error) {
	i.ensureNonNil()
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *SerializableInt) UnmarshalJSON(bz []byte) error {
	i.ensureNonNil()
	return unmarshalJSON(i.i, bz)
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i encoding.TextMarshaler) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}

	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	if err := json.Unmarshal(bz, &text); err != nil {
		return err
	}

	return unmarshalText(i, text)
}

func unmarshalText(i *big.Int, text string) error {
	if err := i.UnmarshalText([]byte(text)); err != nil {
		return err
	}

	return nil
}

func (i *SerializableInt) ensureNonNil() {
	if i.i == nil {
		i.i = new(big.Int)
	}
}
