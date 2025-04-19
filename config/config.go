package config

// Config holds all configuration for the application
type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Extract   ExtractConfig   `mapstructure:"extract"`
	Transform TransformConfig `mapstructure:"transform"`
	Load      LoadConfig      `mapstructure:"load"`
}

// AppConfig holds general app configuration
type AppConfig struct {
	Name     string `mapstructure:"name"`
	Version  string `mapstructure:"version"`
	LogLevel string `mapstructure:"log_level"`
}

// ExtractConfig holds configuration for data extraction
type ExtractConfig struct {
	API      APIConfig      `mapstructure:"api"`
	Database DatabaseConfig `mapstructure:"database"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
}

// APIConfig holds API source configuration
type APIConfig struct {
	URL           string `mapstructure:"url"`
	Method        string `mapstructure:"method"`
	AuthToken     string `mapstructure:"auth_token"`
	Timeoutsecs   int    `mapstructure:"timeout_seconds"`
	RetryAttempts int    `mapstructure:"retry_attempts"`
}

// DatabaseConfig holds database source configuration
type DatabaseConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Query     string `mapstructure:"query"`
	BatchSize int    `mapstructure:"batch_size"`
}

// KafkaConfig holds kafka source configuration
type KafkaConfig struct {
	Brokers         []string `mapstructure:"brokers"`
	Topic           string   `mapstructure:"topic"`
	GroupID         string   `mapstructure:"group_id"`
	AutoOffsetReset string   `mapstructure:"auto_offset_reset"`
}

// TransformConfig holds configuration for data transformation
type TransformConfig struct {
	WorkerPoolSize int              `mapstructure:"woker_pool_size"`
	Validation     ValidationConfig `mapstructure:"validation"`
}

// ValidationConfig holds validation configuration
type ValidationConfig struct {
	RequiredFields []string `mapstructure:"required_fields"`
}

// LoadConfig holds configuration for data loading
type LoadConfig struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	BigQuery BigQueryConfig `mapstructure:"bigquery"`
	S3       S3Config       `mapstructure:"s3"`
}

// PostgresConfig holds Postgres target configuration
type PostgresConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Table     string `mapstructure:"table"`
	BatchSize int    `mapstructure:"batch_size"`
}

// BigQueryConfig holds BigQuery target configuration
type BigQueryConfig struct {
	ProjectID string `mapstructure:"project_id"`
	Dataset   string `mapstructure:"dataset"`
	Table     string `mapstructure:"table"`
}

// S3Config holds S3 target configuration
type S3Config struct {
	Bucket     string `mapstructure:"bucket"`
	Region     string `mapstructure:"region"`
	Path       string `mapstructure:"path"`
	FileFormat string `mapstructure:"file_format"`
}


