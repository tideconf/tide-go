package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"path/filepath"
)

type ConfigValue struct {
	Value string
	Type  string
}

type TIDE struct {
	data map[string]ConfigValue
}

type TypeConverter interface {
	ToString(string) (string, error)
	ToBool(string) (bool, error)
	ToInt(string) (int, error)
	ToInt32(string) (int32, error)
	ToInt64(string) (int64, error)
	ToArray(string) ([]string, error)
	ToIntArray(string) ([]int, error)
}

type ConversionHelper struct{}

func parseImportStatement(line string) (string, bool) {
	if strings.HasPrefix(line, "import") {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			return strings.Trim(parts[1], "\""), true
		}
	}
	return "", false
}

func loadConfigFile(basePath, importFile string, importedFiles map[string]bool) (*TIDE, error) {
	// Append '.tide' extension if not present
	if !strings.HasSuffix(importFile, ".tide") {
		importFile = strings.ReplaceAll(importFile, ".", "/") + ".tide"
	}

	// Resolve the full path of the importFile relative to basePath
	dir := filepath.Dir(basePath)
	fullPath := filepath.Join(dir, importFile)

	// Check for circular imports
	if importedFiles[fullPath] {
		return nil, fmt.Errorf("circular import detected: %s", fullPath)
	}
	importedFiles[fullPath] = true

	// Load the configuration file
	tide, err := NewTIDE(fullPath)
	if err != nil {
		return nil, err
	}

	return tide, nil
}

func (cv ConfigValue) Validate() error {
	switch {
	case cv.Type == "string":
		// No validation needed for strings
		return nil
	case cv.Type == "integer":
		_, err := strconv.Atoi(cv.Value)
		if err != nil {
			return fmt.Errorf("invalid integer value: %s", err)
		}
	case strings.HasPrefix(cv.Type, "array"):
		elementType := strings.TrimPrefix(cv.Type, "array[")
		elementType = strings.TrimSuffix(elementType, "]")

		switch elementType {
		case "string":
			arrayElements, err := ConversionHelper{}.ToArray(cv.Value)
			if err != nil {
				return fmt.Errorf("invalid array format: %s", err)
			}
			for _, element := range arrayElements {
				if _, err := strconv.Atoi(element); err == nil {
					return fmt.Errorf("invalid array element type: expected string, got integer")
				}
			}
		case "integer":
			_, err := ConversionHelper{}.ToIntArray(cv.Value)
			if err != nil {
				return fmt.Errorf("invalid array format: %s", err)
			}
		default:
			return fmt.Errorf("unsupported array element type: %s", elementType)
		}
		return nil
	default:
		return fmt.Errorf("unsupported type: %s", cv.Type)
	}
	return nil
}

func (ConversionHelper) ToString(value string) (string, error) {
	return value, nil
}

func (ConversionHelper) ToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

func (ConversionHelper) ToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func (ConversionHelper) ToInt32(value string) (int32, error) {
	val, err := strconv.ParseInt(value, 10, 32)
	return int32(val), err
}

func (ConversionHelper) ToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func (ConversionHelper) ToArray(value string) ([]string, error) {
	value = strings.Trim(value, "[]")
	items := strings.Split(value, ",")
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items, nil
}

func (ConversionHelper) ToIntArray(value string) ([]int, error) {
	value = strings.Trim(value, "[]")
	items := strings.Split(value, ",")
	var intArray []int
	for _, item := range items {
		trimmedItem := strings.TrimSpace(item)
		intVal, err := strconv.Atoi(trimmedItem)
		if err != nil {
			return nil, fmt.Errorf("array element is not an integer: %v", err)
		}
		intArray = append(intArray, intVal)
	}
	return intArray, nil
}

func (c *TIDE) getConfigValue(key string) (ConfigValue, error) {
	// Convert hierarchical key to environment variable format
	envKey := strings.ReplaceAll(key, ".", "_")
	envKey = strings.ToUpper(envKey)

	// Check for an environment variable
	envVal, exists := os.LookupEnv(envKey)
	if exists {
		// Determine if the key is expected to be an array
		if strings.Contains(key, "array") {
			// Split the environment variable string by comma to create an array
			arrayElements := strings.Split(envVal, ",")
			// Join elements with the required format (e.g., ["elem1", "elem2"])
			joinedElements := fmt.Sprintf("[\"%s\"]", strings.Join(arrayElements, "\", \""))
			return ConfigValue{Value: joinedElements, Type: "array[string]"}, nil
		}
		// If it's not an array, assume it's a string for simplicity
		return ConfigValue{Value: envVal, Type: "string"}, nil
	}

	// If not found in environment, check the file configuration
	configVal, ok := c.data[key]
	if !ok {
		return ConfigValue{}, fmt.Errorf("key not found")
	}

	return configVal, nil
}

func (c *TIDE) GetString(key string) (string, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return "", fmt.Errorf("type mismatch: expected string, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToString(configVal.Value)
}

func (c *TIDE) GetBool(key string) (bool, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return false, fmt.Errorf("type mismatch: expected bool, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToBool(configVal.Value)
}

func (c *TIDE) GetInt(key string) (int, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return 0, fmt.Errorf("type mismatch: expected integer, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToInt(configVal.Value)
}

func (c *TIDE) GetInt32(key string) (int32, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return 0, fmt.Errorf("type mismatch: expected integer, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToInt32(configVal.Value)
}

func (c *TIDE) GetInt64(key string) (int64, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return 0, fmt.Errorf("type mismatch: expected integer, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToInt64(configVal.Value)
}

func (c *TIDE) GetArray(key string) ([]string, error) {
	configVal, err := c.getConfigValue(key)
	if err != nil {
		return nil, fmt.Errorf("type mismatch: expected array, got %s", configVal.Type)
	}

	return ConversionHelper{}.ToArray(configVal.Value)
}
func NewTIDE(filepath string) (*TIDE, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &TIDE{data: make(map[string]ConfigValue)}
	scanner := bufio.NewScanner(file)

	var currentContext []string
	importedFiles := make(map[string]bool) // Track imported files to prevent circular imports

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if importFile, ok := parseImportStatement(line); ok {
			// Pass the current file path as basePath and the importFile to loadConfigFile
			importedConfig, err := loadConfigFile(filepath, importFile, importedFiles)
			if err != nil {
				return nil, err
			}
			// Merge configurations: importedConfig takes precedence
			for k, v := range importedConfig.data {
				cfg.data[k] = v
			}
			continue
		}

		if strings.HasSuffix(line, "{") {
			contextKey := strings.TrimSuffix(line, " {")
			currentContext = append(currentContext, contextKey)
			continue
		} else if line == "}" {
			if len(currentContext) > 0 {
				currentContext = currentContext[:len(currentContext)-1]
			}
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		keyParts := strings.SplitN(strings.TrimSpace(parts[0]), ":", 2)
		if len(keyParts) != 2 {
			continue
		}

		key := buildKey(currentContext, keyParts[0])
		cfgValue := ConfigValue{
			Value: strings.TrimSpace(parts[1]),
			Type:  strings.TrimSpace(keyParts[1]),
		}

		// Trim quotation marks for string values
		if cfgValue.Type == "string" {
			cfgValue.Value = strings.Trim(cfgValue.Value, "\"")
		}

		if err := cfgValue.Validate(); err != nil {
			return nil, fmt.Errorf("validation error for key %s: %v", key, err)
		}
		cfg.data[key] = cfgValue
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func buildKey(context []string, key string) string {
	fullKey := strings.Join(context, ".") + "." + key
	return strings.Trim(fullKey, ".")
}
