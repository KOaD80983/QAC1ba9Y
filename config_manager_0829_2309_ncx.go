// 代码生成时间: 2025-08-29 23:09:27
package main

import (
    "buffalo"
    "fmt"
    "os"
    "path/filepath"
)

// ConfigManager is a struct to handle configuration files.
type ConfigManager struct {
    // Directory where the configuration files are stored.
    dir string
    // File extension of the configuration files.
    ext string
}

// NewConfigManager creates a new ConfigManager instance with a specified directory and extension.
func NewConfigManager(directory, extension string) *ConfigManager {
    return &ConfigManager{
        dir: directory,
        ext: extension,
    }
}

// LoadConfig loads a configuration file by its name.
func (m *ConfigManager) LoadConfig(filename string) (map[string]interface{}, error) {
    // Construct the full path to the configuration file.
    path := filepath.Join(m.dir, fmt.Sprintf("%s.%s", filename, m.ext))

    // Check if the file exists.
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return nil, fmt.Errorf("configuration file '%s' not found", filename)
    }

    // Read and parse the configuration file.
    // Assuming the configuration file is in JSON format for simplicity.
    var config map[string]interface{}
    if err := readAndParseConfigFile(path, &config); err != nil {
        return nil, err
    }

    return config, nil
}

// saveConfig saves a configuration file by its name.
func (m *ConfigManager) SaveConfig(filename string, config map[string]interface{}) error {
    // Construct the full path to the configuration file.
    path := filepath.Join(m.dir, fmt.Sprintf("%s.%s", filename, m.ext))

    // Convert the configuration map to a JSON string.
    jsonConfig, err := convertMapToJSON(config)
    if err != nil {
        return err
    }

    // Write the JSON string to the file.
    if err := os.WriteFile(path, []byte(jsonConfig), 0644); err != nil {
        return err
    }

    return nil
}

// readAndParseConfigFile reads and parses a configuration file into a map.
func readAndParseConfigFile(path string, config *map[string]interface{}) error {
    // Open the file.
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    // Decode the JSON file into the map.
    decoder := json.NewDecoder(file)
    return decoder.Decode(config)
}

// convertMapToJSON converts a map to a JSON string.
func convertMapToJSON(config map[string]interface{}) (string, error) {
    // Marshal the map into a JSON byte slice.
    jsonBytes, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return "", err
    }
    // Convert the byte slice to a string.
    return string(jsonBytes), nil
}
