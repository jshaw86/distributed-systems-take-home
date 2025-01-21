package client

import (
    "bytes"
    "context"
    "fmt"
    "math/rand"
    "math"
    "net/http"
    "time"
)


type Client struct {
    servers []string
    baseDelay time.Duration
    maxRetries int
}

type IClient interface {
    SendData(ctx context.Context, user UserInfo)
}

func AddServers(servers []string) Client {
    baseDelay := 1 * time.Second
    maxRetries := 3
    return Client{servers, baseDelay, maxRetries}
}

func (c Client) makeRequest(ctx context.Context, method, uri string, headers map[string]string, body string) (int, error) {
	client := &http.Client{}
    ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
    defer cancel()
    for attempt := 0; attempt < c.maxRetries; attempt++ {
        url := fmt.Sprintf("%s%s", c.servers[rand.Intn(len(c.servers))], uri)
        req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer([]byte(body)))
        if err != nil {
            return -1, err
        }
        for key, value := range headers {
            req.Header.Set(key, value)
        }
        resp, err := client.Do(req)
        if err != nil {
            return -1, err
        }
        defer resp.Body.Close()
        // retry on 429
        if resp.StatusCode == http.StatusTooManyRequests {
            retryAfter := resp.Header.Get("Retry-After")
            var waitTime time.Duration

            if retryAfter != "" {
                retryAfterSec, err := time.ParseDuration(retryAfter + "s")
                if err == nil {
                    waitTime = retryAfterSec
                }
            } else {
                // Use exponential backoff with jitter
                waitTime = c.baseDelay * time.Duration(math.Pow(2, float64(attempt))) +
                    time.Duration(rand.Intn(1000))*time.Millisecond
            }

            time.Sleep(waitTime)
            continue
        }

        return resp.StatusCode, nil
    }
    
    return -1, fmt.Errorf("Failed to send data after %d retries", c.maxRetries)
}

func (c Client) SendData(ctx context.Context, user UserInfo) error {
    _,  err := c.makeRequest(ctx, "POST", "/send", map[string]string{
        "X-Session-ID": user.ID,
        "X-IP-Address": user.IPAddr,
        "User-Agent": user.UserAgent,
	}, StringWithCharset(rand.Intn(10)))

    return err
}
