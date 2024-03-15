package config

type CryptoConfig struct {
	AESLength        string
	RSAPublicKeyPath string
}

type ServerConfig struct {
	Address string
}

type CollectorConfig struct {
	SearchPath      string
	FilePathPattern string
	ParseTemplate   string
	ReportInterval  int
}
