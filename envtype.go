package midtrans

type EnvironmentType int8

const (
    _ EnvironmentType = iota
    Sandbox
    Production
)

var typeString = map[EnvironmentType]string {
    Sandbox: "https://api.sandbox.veritrans.co.id/v2",
    Production: "https://api.veritrans.co.id/v2",
}

// implement stringer
func (e EnvironmentType) String() string {
    for k, v := range typeString {
        if k == e {
            return v
        }
    }
    return "undefined"
}