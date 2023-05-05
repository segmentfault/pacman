package utils

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"golang.org/x/net/proxy"
	"io"
	"net"
	"net/http"
)

// RequestMaxRetry is the max retry times for http request
const RequestMaxRetry = 3

// Request fetch data from specified url
func Request(method string, url string) ([]byte, error) {
	return RequestWithSocks5(method, url, "")
}

// RequestWithSocks5 fetch data from specified url with socks5 proxy
func RequestWithSocks5(method string, url string, proxyAddr string) ([]byte, error) {
	// create a retryable http client, from hashicorp go-retryablehttp library
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = RequestMaxRetry

	client := retryClient.StandardClient() // *http.Client
	httpTransport := &http.Transport{}

	if len(proxyAddr) > 0 {
		// set our socks5 as the dialer
		httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()

			default:
				// create a socks5 dialer
				dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
				if err != nil {
					return nil, err
				}

				return dialer.Dial(network, addr)
			}
		}
	}

	// mark our httpTransport as the client's transport
	client.Transport = httpTransport

	// start request
	req, _ := http.NewRequest(method, url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
