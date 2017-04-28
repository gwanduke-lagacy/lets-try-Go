// 설정을 위한 구조체 제공
// 미리 정의된 값을 설정

package config

// Config .
type Config struct {
	DB *DBConfig
}

// DBConfig 는 상위 Config 구조체에 임베딩
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig 미리정의된 설정값을 가지는 Config 인스턴스를 반환하는 함수
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "gwanduke",
			Charset:  "utf8",
		},
	}
}
