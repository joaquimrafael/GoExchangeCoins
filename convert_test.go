package main

import "testing"

func TestConvert(t *testing.T) {
	cases := []struct {
		name           string
		value          float64
		origin, target string
		want           float64
		wantErr        bool
	}{
		{name: "USDtoBRL", value: 1, origin: "USD", target: "BRL", want: 5.10},
		{name: "USDtoBRLwithValue", value: 2, origin: "USD", target: "BRL", want: 10.2},
		{name: "USDtoEUR", value: 1, origin: "USD", target: "EUR", want: 0.87},
		{name: "BRLtoUSD", value: 5.10, origin: "BRL", target: "USD", want: 1},
		{name: "USDtoUSD", value: 1, origin: "USD", target: "USD", want: 1},
		{name: "EURtoEUR", value: 1, origin: "EUR", target: "EUR", want: 1},
		{name: "ZeroValue", value: 0, origin: "USD", target: "BRL", want: 0},
		{name: "UnknownOrigin", value: 1, origin: "XYZ", target: "BRL", wantErr: true},
		{name: "UnknownTarget", value: 1, origin: "USD", target: "XYZ", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := convert(tc.value, tc.origin, tc.target)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("convert(%v, %q, %q): expected an error, got nil", tc.value, tc.origin, tc.target)
				}
				return
			}

			if err != nil {
				t.Fatalf("convert(%v, %q, %q): unexpected an error: %v", tc.value, tc.origin, tc.target, err)
			}
			if got != tc.want {
				t.Fatalf("convert(%v, %q, %q) = %v, want %v", tc.value, tc.origin, tc.target, got, tc.want)
			}
		})
	}

}
