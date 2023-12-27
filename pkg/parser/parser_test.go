package parser

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Test data
var testConfig = map[string]ConfigValue{
	"stringKey":   {"value1", "string"},
	"boolKey":     {"true", "bool"},
	"intKey":      {"123", "integer"},
	"int32Key":    {"12345", "integer"},
	"int64Key":    {"1234567890", "integer"},
	"invalidKey":  {"invalid", "string"},
	"arrayKey":    {"[feature1, feature2, feature3]", "array"},
	"intArrayKey": {"[1, 2, 3]", "array[integer]"},
}

// createTestTIDE creates a TIDE object with test data.
func createTestTIDE() *TIDE {
	return &TIDE{data: testConfig}
}

// TestGetString tests the GetString method.
func TestGetString(t *testing.T) {
	tide := createTestTIDE()

	_, err := tide.GetString("nonExistentKey")
	if err == nil {
		t.Errorf("GetString should return an error for non-existing key")
	}
}

// TestGetBool tests the GetBool method.
func TestGetBool(t *testing.T) {
	tide := createTestTIDE()

	val, err := tide.GetBool("boolKey")
	if err != nil {
		t.Errorf("GetBool returned an error: %v", err)
	}
	if val != true {
		t.Errorf("GetBool returned unexpected value: got %v want %v", val, true)
	}

	_, err = tide.GetBool("invalidKey")
	if err == nil {
		t.Errorf("GetBool should return an error for non-boolean type")
	}
}

// TestGetInt tests the GetInt method.
func TestGetInt(t *testing.T) {
	tide := createTestTIDE()

	val, err := tide.GetInt("intKey")
	if err != nil {
		t.Errorf("GetInt returned an error: %v", err)
	}
	if val != 123 {
		t.Errorf("GetInt returned unexpected value: got %v want %v", val, 123)
	}

	_, err = tide.GetInt("invalidKey")
	if err == nil {
		t.Errorf("GetInt should return an error for non-integer type")
	}
}

// TestGetInt32 tests the GetInt32 method.
func TestGetInt32(t *testing.T) {
	tide := createTestTIDE()

	val, err := tide.GetInt32("int32Key")
	if err != nil {
		t.Errorf("GetInt32 returned an error: %v", err)
	}
	if val != 12345 {
		t.Errorf("GetInt32 returned unexpected value: got %v want %v", val, 12345)
	}

	_, err = tide.GetInt32("invalidKey")
	if err == nil {
		t.Errorf("GetInt32 should return an error for non-integer type")
	}
}

// TestGetInt64 tests the GetInt64 method.
func TestGetInt64(t *testing.T) {
	tide := createTestTIDE()

	val, err := tide.GetInt64("int64Key")
	if err != nil {
		t.Errorf("GetInt64 returned an error: %v", err)
	}
	if val != 1234567890 {
		t.Errorf("GetInt64 returned unexpected value: got %v want %v", val, 1234567890)
	}

	_, err = tide.GetInt64("invalidKey")
	if err == nil {
		t.Errorf("GetInt64 should return an error for non-integer type")
	}
}

// TestGetArray tests the GetArray method.
func TestArray(t *testing.T) {
	tide := createTestTIDE()

	expectedStringArray := []string{"feature1", "feature2", "feature3"}
	array, err := tide.GetArray("arrayKey")
	if err != nil {
		t.Errorf("GetArray returned an error: %v", err)
	}
	if !reflect.DeepEqual(array, expectedStringArray) {
		t.Errorf("GetArray returned unexpected value: got %v want %v", array, expectedStringArray)
	}

	_, err = tide.GetArray("nonExistentKey")
	if err == nil {
		t.Errorf("GetArray should return an error for non-existing key")
	}
}

// TestGetConfigValueWithEnv tests the getConfigValue method with environment variables.
func TestGetConfigValueWithEnv(t *testing.T) {
	// Set environment variables
	os.Setenv("STRINGKEY", "envValue1")
	os.Setenv("INTKEY", "456")
	os.Setenv("ARRAYKEY", "envFeature1,envFeature2,envFeature3")

	tide := createTestTIDE()

	// Test string value overridden by environment variable
	stringVal, err := tide.GetString("stringKey")
	if err != nil || stringVal != "envValue1" {
		t.Errorf("Expected envValue1, got %v, error: %v", stringVal, err)
	}

	// Test int value overridden by environment variable
	intVal, err := tide.GetInt("intKey")
	if err != nil || intVal != 456 {
		t.Errorf("Expected 456, got %v, error: %v", intVal, err)
	}

	// Test array value overridden by environment variable
	expectedArray := []string{"\"envFeature1\"", "\"envFeature2\"", "\"envFeature3\""}

	arrayVal, err := tide.GetArray("arrayKey")
	if err != nil || !reflect.DeepEqual(arrayVal, expectedArray) {
		t.Errorf("Expected %v, got %v, error: %v", expectedArray, arrayVal, err)
	}

	// Unset environment variables
	os.Unsetenv("STRINGKEY")
	os.Unsetenv("INTKEY")
	os.Unsetenv("ARRAYKEY")
}

func TestImportConfig(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	configFile := filepath.Join(dir, "../../test/main_config.tide")

	cfg, err := NewTIDE(configFile)

	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Test value from main config
	host, err := cfg.GetString("database.host")
	if err != nil || host != "localhost" {
		t.Errorf("Expected 'localhost', got %v, error: %v", host, err)
	}

	// Test value from imported config
	logLevel, err := cfg.GetString("logging.level")
	if err != nil || logLevel != "info" {
		t.Errorf("Expected 'info', got %v, error: %v", logLevel, err)
	}
}
