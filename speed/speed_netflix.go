package speed

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ddo/go-fast"
)

func speedNetflix() (*TestResult, error) {
	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		return nil, fmt.Errorf("fastCom.Init err: %w", err)
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, fmt.Errorf("fastCom.GetUrls err: %w", err)
	}

	urls = urls[:1]

	chKbps := make(chan float64)

	var (
		downMbps []float64
		averMbps float64
	)

	const (
		divKtoM = 1000
	)

	go func() {
		for Kbps := range chKbps {
			// fmt.Printf("%.2f Kbps %.2f Mbps\n", Kbps, Kbps/1000)
			downMbps = append(downMbps, Kbps/divKtoM)
		}
	}()

	err = fastCom.Measure(urls, chKbps)
	if err != nil {
		return nil, fmt.Errorf("fastCom.Measure err: %w", err)
	}

	for _, v := range downMbps {
		averMbps += v
	}

	averMbps /= float64(len(downMbps))

	u, err := url.Parse(strings.Split(urls[0], "?")[0])
	if err != nil {
		return nil, fmt.Errorf("url.Parse err: %w", err)
	}

	out := new(TestResult)
	// out.Ip = user.IP
	out.Host = strings.Join([]string{u.Hostname(), u.Port()}, ":")
	out.DownloadMbps = averMbps
	// out.UploadMbps = srv.ULSpeed
	// out.LatencyNs = srv.Latency.Nanoseconds()

	return out, nil
}
