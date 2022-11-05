//nolint:all
package dult_test

import (
	"fmt"
	"testing"

	"github.com/ac-i/dult"
	"github.com/ac-i/dult/proto/speedpb"
	"github.com/ac-i/dult/speed"
)

func TestQuickDULT(t *testing.T) {
	type args struct {
		opt speed.Options
	}
	tests := []struct {
		name    string
		args    args
		want    *speed.TestResult
		wantErr bool
	}{
		{
			name:    "empty PROVIDER",
			wantErr: false,
		},
		{
			name:    "PROVIDER_UNSPECIFIED",
			args:    args{opt: speed.Options{Provider: speedpb.Provider_PROVIDER_UNSPECIFIED}},
			wantErr: false,
		},
		{
			name:    "PROVIDER_SPEEDTEST_NET",
			args:    args{opt: speed.Options{Provider: speedpb.Provider_PROVIDER_SPEEDTEST_NET}},
			wantErr: false,
		},
		{
			name:    "PROVIDER_FAST_COM",
			args:    args{opt: speed.Options{Provider: speedpb.Provider_PROVIDER_FAST_COM}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := speed.QuickDULT(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuickDULT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("QuickDULT() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func BenchmarkQuickDULT(b *testing.B) {
	type args struct {
		opt speed.Options
	}
	tests := []struct {
		name    string
		args    args
		want    *speed.TestResult
		wantErr bool
	}{
		{
			name:    "PROVIDER_SPEEDTEST_NET",
			args:    args{opt: speed.Options{Provider: speedpb.Provider_PROVIDER_SPEEDTEST_NET}},
			wantErr: false,
		},
		{
			name:    "PROVIDER_FAST_COM",
			args:    args{opt: speed.Options{Provider: speedpb.Provider_PROVIDER_FAST_COM}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				speed.QuickDULT(tt.args.opt)
			}
		})
	}
}

func ExampleQuickDULT() {

	tst, err := dult.QuickDULT(speed.Options{Provider: speedpb.Provider_PROVIDER_SPEEDTEST_NET})
	_ = fmt.Sprintf("result: %+v\nerror: %v\n\n", tst, err)
	if err != nil {
		fmt.Printf("SPEEDTEST_NET error: %v", err)
	} else {
		fmt.Println("SPEEDTEST_NET OK")
	}

	tst, err = dult.QuickDULT(speed.Options{Provider: speedpb.Provider_PROVIDER_FAST_COM})
	_ = fmt.Sprintf("result: %+v\nerror: %v\n\n", tst, err)
	if err != nil {
		fmt.Printf("FAST_COM error: %v", err)
	} else {
		fmt.Println("FAST_COM OK")
	}

	// Output:
	//
	// SPEEDTEST_NET OK
	// FAST_COM OK
}
