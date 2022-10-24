package core

var (
	AppConf = &RuntimeConfig{
		App: AppConfig{
			Port: 38080,
		},
	}
)

// RuntimeConfig ...
type RuntimeConfig struct {
	App AppConfig `json:"app"`
}

type AppConfig struct {
	Port uint16 `json:"port"`
}
