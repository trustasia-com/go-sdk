// Package session provides ...
package session

// api host
const (
	// test host
	testHost = "https://api-test.wekey.com"
	// prod host
	prodHost = "https://api.wekey.com"
)

// Options session options
type Options struct {
	IsProduction bool // default false
	AccessKey    string
	SecretKey    string
	// you can custom fido server if you are privacy
	Host string // eg. https://fido.example.com
}

// Option some set config
type Option func(opt *Options)

// WithCredential credential with accessKey, secretKey
func WithCredential(ak, sk string) Option {
	return func(opt *Options) {
		opt.AccessKey = ak
		opt.SecretKey = sk
	}
}

// WithProduction set run production mode
func WithProduction() Option {
	return func(opt *Options) { opt.IsProduction = true }
}

// WithHost api host
func WithHost(host string) Option {
	return func(opt *Options) { opt.Host = host }
}
