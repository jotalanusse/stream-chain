package address

import (
	"fmt"
	"testing"
)

func TestAddressConversion(t *testing.T) {
	newAddr, err := ConvertAddressPrefix("klyra1n88uc38xhjgxzw9nwre4ep2c8ga4fjxcttpsmz", "klyravalcons")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(newAddr)
}
