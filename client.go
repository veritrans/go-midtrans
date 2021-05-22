package midtrans

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	defaultHTTPTimeout = 80 * time.Second
	defaultHTTPClient  = &http.Client{Timeout: defaultHTTPTimeout}
)

// Client struct
type Client struct {
	APIEnvType EnvironmentType
	ClientKey  string
	ServerKey  string

	LogLevel int
	Logger   *log.Logger

	HTTPClient *http.Client
}

type ClientOption func(*Client)

func WithEnvType(envType EnvironmentType) ClientOption {
	return func(c *Client) {
		c.APIEnvType = envType
	}
}

func WithClientKey(key string) ClientOption {
	return func(c *Client) {
		c.ClientKey = key
	}
}

func WithServerKey(key string) ClientOption {
	return func(c *Client) {
		c.ServerKey = key
	}
}

func WithLogLevel(level int) ClientOption {
	return func(c *Client) {
		c.LogLevel = level
	}
}

func WithLogger(logger *log.Logger) ClientOption {
	return func(c *Client) {
		c.Logger = logger
	}
}

func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// NewClient : this function will always be called when the library is in use
func NewClient(opts ...ClientOption) Client {
	client := Client{
		APIEnvType: Sandbox,

		// LogLevel is the logging level used by the Midtrans library
		// 0: No logging
		// 1: Errors only
		// 2: Errors + informational (default)
		// 3: Errors + informational + debug
		LogLevel:   2,
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
		HTTPClient: defaultHTTPClient,
	}

	for _, apply := range opts {
		apply(&client)
	}

	return client
}

// NewRequest : send new request
func (c *Client) NewRequest(method string, fullPath string, body io.Reader) (*http.Request, error) {
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Request creation failed: ", err)
		}
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.ServerKey, "")

	return req, nil
}

// ExecuteRequest : execute request
func (c *Client) ExecuteRequest(req *http.Request, v interface{}) error {
	logLevel := c.LogLevel
	logger := c.Logger

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot send request: ", err)
		}
		return err
	}
	defer res.Body.Close()

	if logLevel > 2 {
		logger.Println("Completed in ", time.Since(start))
	}

	if err != nil {
		if logLevel > 0 {
			logger.Println("Request failed: ", err)
		}
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		if logLevel > 0 {
			logger.Println("Cannot read response body: ", err)
		}
		return err
	}

	if logLevel > 2 {
		logger.Println("Midtrans response: ", string(resBody))
	}

	if v != nil {
		if err = json.Unmarshal(resBody, v); err != nil {
			return err
		}

		// when return unexpected error, midtrans not return `status_message` but `message`, so this to catch it
		error := make(map[string]string)
		if res.StatusCode >= 500 {
			err := json.Unmarshal(resBody, &error)
			if err != nil {
				return err
			}
		}

		// we're safe to reflect status_code if response not return status code
		if reflect.ValueOf(v).Elem().Kind() == reflect.Struct {
			if reflect.ValueOf(v).Elem().FieldByName("StatusCode").Len() == 0 {
				reflect.ValueOf(v).Elem().FieldByName("StatusCode").SetString(strconv.Itoa(res.StatusCode))
				// response of snap transaction not return StatusMessage
				if req.URL.Path != "/snap/v1/transactions" {
					reflect.ValueOf(v).Elem().FieldByName("StatusMessage").SetString(error["message"])
				}
			}
		}
	}

	return nil
}

// Call the Midtrans API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occurred, the return of this function is the error
// itself, otherwise nil.
func (c *Client) Call(method, path string, body io.Reader, v interface{}) error {
	req, err := c.NewRequest(method, path, body)

	if err != nil {
		return err
	}

	return c.ExecuteRequest(req, v)
}

// ===================== END HTTP CLIENT ================================================
