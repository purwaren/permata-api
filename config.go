package permata

// Config is config for access Permata API
type Config struct {
	URL          string
	APIKey       string
	StaticKey    string
	ClientID     string
	ClientSecret string

	LogPath string
}
