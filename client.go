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

type clientSetting struct {
    ApiEnvType EnvironmentType
    ClientKey string
    ServerKey string

    LogLevel int
    Logger *log.Logger
}

var ClientSetting clientSetting

// this function will always be called when the library is in use
func init() {
    ClientSetting = clientSetting{
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

func NewRequest(method string, path string, body io.Reader) (*http.Request, error) {
    logLevel := ClientSetting.LogLevel
    logger := ClientSetting.Logger

    if !strings.HasPrefix(path, "/") {
        path = "/" + path
    }

    path = ClientSetting.ApiEnvType.String() + path

    req, err := http.NewRequest(method, path, body)
    if err != nil {
        if logLevel > 0 {
            logger.Println("Request creation failed: ", err)
        }
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "application/json")
    req.SetBasicAuth("", ClientSetting.ServerKey)

    return req, nil
}

func ExecuteRequest(req *http.Request, v interface{}) error {
    logLevel := ClientSetting.LogLevel
    logger := ClientSetting.Logger

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
// ===================== END HTTP CLIENT ================================================