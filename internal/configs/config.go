package configs

import "github.com/spf13/viper"

// Variabel global untuk menyimpan konfigurasi aplikasi yang sudah dibaca.
var config *Config

// Struktur untuk menyimpan opsi konfigurasi.
type option struct {
	configFolders []string // Folder tempat file konfigurasi dicari.
	configFile    string   // Nama file konfigurasi (tanpa ekstensi).
	configType    string   // Tipe file konfigurasi (yaml, json, dll).
}

// Init membaca konfigurasi berdasarkan opsi yang diberikan atau menggunakan nilai default.
func Init(opts ...Option) error {
	// Inisialisasi opsi default.
	opt := &option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	// Terapkan fungsi-fungsi opsi yang diberikan ke `opt`.
	for _, optFunc := range opts {
		optFunc(opt) // Mengubah nilai `opt` sesuai opsi.
	}

	// Tambahkan folder tempat Viper akan mencari file konfigurasi.
	for _, folder := range opt.configFolders {
		viper.AddConfigPath(folder)
	}

	// Set nama file, tipe file, dan aktifkan pembacaan variabel lingkungan (env).
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	// Alokasikan memori untuk variabel global `config`.
	config = new(Config)

	// Baca file konfigurasi dan simpan hasilnya ke dalam `config`.
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config) // Decode konfigurasi ke dalam struct `Config`.
}

// Option adalah tipe fungsi untuk memodifikasi opsi konfigurasi.
type Option func(*option)

// Fungsi default: folder tempat file konfigurasi berada.
func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

// Fungsi default: nama file konfigurasi.
func getDefaultConfigFile() string {
	return "config"
}

// Fungsi default: tipe file konfigurasi.
func getDefaultConfigType() string {
	return "yaml"
}

// WithConfigFolder memungkinkan pengaturan folder konfigurasi secara kustom.
func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

// WithConfigFile memungkinkan pengaturan nama file konfigurasi secara kustom.
func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

// WithConfigType memungkinkan pengaturan tipe file konfigurasi secara kustom.
func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

// Get mengembalikan konfigurasi global.
// Jika belum diinisialisasi, membuat konfigurasi kosong sebagai fallback.
func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
