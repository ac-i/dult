package speed

import (
	"fmt"

	"github.com/ac-i/dult/proto/speedpb"
)

type Options struct {
	speedpb.Provider
}

type TestResult struct {
	speedpb.Speed
}

func QuickDULT(opt Options) (*TestResult, error) {
	//nolint:nosnakecase
	switch opt.Provider {
	case speedpb.Provider_PROVIDER_UNSPECIFIED,
		speedpb.Provider_PROVIDER_SPEEDTEST_NET:
		return speedOkla()
	case speedpb.Provider_PROVIDER_FAST_COM:
		return speedNetflix()
	default:
		return nil, fmt.Errorf("unknown speedpb.Provider enum %d", opt.Provider) //nolint:goerr113
	}
}
