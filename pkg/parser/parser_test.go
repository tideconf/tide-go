package parser

import (
	"reflect"
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
