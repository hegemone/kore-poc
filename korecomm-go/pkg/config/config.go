// config is responsible for managing the config of the system and presenting
// that to comm in a controlled way. For the purposes of the POC, it is
// mostly hardcoded in its implementation. In practice, this pkg should handle
// things like validation, and file loading/parsing, in addition to env var
// loading for things like secret credentials.
package config

import (
	"os"
	"sync"
)

const (
	PluginDirEnvVar  = "KORECOMM_PLUGIN_DIR"
	AdapterDirEnvVar = "KORECOMM_ADAPTER_DIR"
)

type EngineConfig struct {
	BufferSize uint
}

type ExtensionConfig struct {
	Dir     string
	Enabled []string
}

type PluginConfig struct {
	ExtensionConfig
}

type AdapterConfig struct {
	ExtensionConfig
}

// Just fake out the yaml config for now
var mockConfig = map[string]interface{}{
	"engine": EngineConfig{
		BufferSize: 8,
	},
	"plugins": PluginConfig{
		ExtensionConfig: ExtensionConfig{
			Dir: os.Getenv(PluginDirEnvVar),
			Enabled: []string{
				"bacon.plugins.kore.nsk.io",
			},
		},
	},
	"adapters": AdapterConfig{
		ExtensionConfig: ExtensionConfig{
			Dir: os.Getenv(AdapterDirEnvVar),
			Enabled: []string{
				"ex-discord.adapters.kore.nsk.io",
				"ex-irc.adapters.kore.nsk.io",
			},
		},
	},
}

var _instance *map[string]interface{}
var once sync.Once

func instance() *map[string]interface{} {
	// Threadsafe lazy accessor
	once.Do(func() {
		_instance = loadConfigFile()
	})
	return _instance
}

func GetEngineConfig() EngineConfig {
	return (*instance())["engine"].(EngineConfig)
}

func GetPluginConfig() PluginConfig {
	return (*instance())["plugins"].(PluginConfig)
}

func GetAdapterConfig() AdapterConfig {
	return (*instance())["adapters"].(AdapterConfig)
}

func loadConfigFile() *map[string]interface{} {
	// Load file location from env var, or use default
	return &mockConfig
}
