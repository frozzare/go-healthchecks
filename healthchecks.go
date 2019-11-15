package healthchecks

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrBadStatusCode is the error returned when a bad status code occurs.
	ErrBadStatusCode = errors.New("Bad status code from healthchecks")

	// DefaultURL is the default url to healthchecks.
	DefaultURL = "https://hc-ping.com/"
)

// A Client manages communication with Healthchecks.
type Client struct {
	http *http.Client
	url  string
}

// NewClient returns a new Healthchecks client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, args ...string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		http: httpClient,
	}

	if len(args) > 0 {
		c.url = args[0]
	}

	return c
}

// Success sends a success request to Healthchecks, a error is returned
// if a error has occurred.
func (c *Client) Success(ctx context.Context, id string) error {
	return c.get(ctx, c.createURL(id))
}

// Fail sends a fail request to Healthchecks, a error is returned
// if a error has occurred.
func (c *Client) Fail(ctx context.Context, id string) error {
	return c.get(ctx, fmt.Sprintf("%s/fail", c.createURL(id)))
}

// Start sends a start request to Healthchecks.io to indicate that the job has started.
func (c *Client) Start(ctx context.Context, id string) error {
	return c.get(ctx, fmt.Sprintf("%s/start", c.createURL(id)))
}

func (c *Client) createURL(id string) string {
	url := c.url

	if url == "" {
		url = DefaultURL
	}

	if url[len(url)-1] != '/' {
		url += "/"
	}

	return fmt.Sprintf("%s%s", url, id)
}

func (c *Client) get(ctx context.Context, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	res, err := c.http.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		return nil
	}

	return ErrBadStatusCode
}
