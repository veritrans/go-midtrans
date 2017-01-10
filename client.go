package midtrans

import (
    "log"
    "os"
    "net/http"
    "time"
    "io"
    "strings"
    "io/ioutil"
    "encoding/json"
)

type Client struct {
    ApiEnvType EnvironmentType
    ClientKey string
    ServerKey string

    LogLevel int
    Logger *log.Logger
}

// this function will always be called when the library is in use
func NewClient() Client {
    return Client{
        ApiEnvType: Sandbox,

        // LogLevel is the logging level used by the Midtrans library
        // 0: No logging
        // 1: Errors only
        // 2: Errors + informational (default)
        // 3: Errors + informational + debug
        LogLevel: 2,
        Logger: log.New(os.Stderr, "", log.LstdFlags),
    }
}

// ===================== HTTP CLIENT ================================================
var defHttpTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHttpTimeout}

func (c *Client) NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
    logLevel := c.LogLevel
    logger := c.Logger

    if !strings.HasPrefix(path, "/") {
        path = "/" + path
    }

    path = c.ApiEnvType.String() + path

    req, err := http.NewRequest(method, path, body)
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

func (c *Client) ExecuteRequest(req *http.Request, v interface{}) error {
    logLevel := c.LogLevel
    logger := c.Logger

    if logLevel > 1 {
        logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
    }

    start := time.Now()

    res, err := httpClient.Do(req)
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
        logger.Println("Midtrans response: ", resBody)
    }

    if v != nil {
        return json.Unmarshal(resBody, v)
    }

    return nil
}

// Call the Midtrans API at specific `path` using the specified HTTP `method`. The result will be
// given to `v` if there is no error. If any error occured, the return of this function is the error
// itself, otherwise nil.
func (c *Client) Call(method, path string, body io.Reader, v interface{}) error {
    req, err := c.NewRequest(method, path, body)

    if err != nil {
        return err
    }

    if err := c.ExecuteRequest(req, v); err != nil {
        return err
    }

    return nil
}
// ===================== END HTTP CLIENT ================================================