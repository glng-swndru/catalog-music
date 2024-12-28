package configs

// Config adalah struktur utama yang merepresentasikan konfigurasi aplikasi.
type Config struct {
	Service  Service  `mapstructure:"service"`  // Konfigurasi untuk bagian service.
	Database Database `mapstructure:"database"` // Konfigurasi untuk bagian database
}

// Service adalah struktur untuk menyimpan konfigurasi terkait layanan (service).
type Service struct {
	Port      string `mapstructure:"port"`
	SecretJWT string `mapstructure:"secretJWT"`
}

// Database adalah struktur utama yang menyimpan konfigurasi terkait database.
type Database struct {
	Host string `mapstructure:"host"` // Host database
}
