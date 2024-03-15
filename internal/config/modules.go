package config

type CryptoConfig struct {
	AESLength        int
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
