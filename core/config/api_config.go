package config

type ApiConfiguration struct {
	Debug                bool   `envconfig:"DEBUG" defalt:"false"`
	AppEnv               string `envconfig:"APP_ENV" default:"prod"`
	HttpAddress          string `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:6969"`
	AwsAccountKey        string `envconfig:"AWS_ACCESS_KEY_ID"`
	AwsAccountSecret     string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	AwsBucketName        string `envconfig:"AWS_BUCKET_NAME"`
	DiscordCrashWebhook  string `envconfig:"CRASHES_WEBHOOK"`
	NewRelicLicenseKey   string `envconfig:"NEW_RELIC_LICENSE_KEY"`
	DatabaseUri          string `envconfig:"DATABASE_URI"`
	DatabaseName         string `envconfig:"DATABASE_NAME"`
	PoogieEmail          string `envconfig:"POOGIE_EMAIL"`
	PoogiePassword       string `envconfig:"POOGIE_PASSWORD"`
	PatreonWebhookSecret string `envconfig:"PATREON_WEBHOOK_SECRET"`
	HashSalt             string `envconfig:"HASH_SALT"`
	CryptoSalt           string `envconfig:"AES_CRYPTO_SALT"`
	CryptoKey            string `envconfig:"AES_CRYPTO_KEY"`
	JwtKey               string `envconfig:"JWT_SIGNING_KEY"`
}
