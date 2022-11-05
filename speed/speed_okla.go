package speed

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func speedOkla() (*TestResult, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		return nil, fmt.Errorf("speedtest.FetchUserInfo err: %w", err)
	}

	servers, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, fmt.Errorf("speedtest.FetchServers err: %w", err)
	}

	srv := servers[0]

	err = srv.PingTest()
	if err != nil {
		return nil, fmt.Errorf("srv.PingTest err: %w", err)
	}

	err = srv.DownloadTest(true)
	if err != nil {
		return nil, fmt.Errorf("srv.DownloadTest err: %w", err)
	}

	err = srv.UploadTest(true)
	if err != nil {
		return nil, fmt.Errorf("srv.UploadTest err: %w", err)
	}

	out := new(TestResult)
	out.Ip = user.IP
	out.Host = srv.Host
	out.DownloadMbps = srv.DLSpeed
	out.UploadMbps = srv.ULSpeed
	out.LatencyNs = srv.Latency.Nanoseconds()

	return out, nil
}
