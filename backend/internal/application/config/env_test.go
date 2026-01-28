package config_test

import (
	"os"
	"sub-watch-backend/internal/application/config"
	"testing"
)

func setupEnvVars() {
	os.Setenv("API_PORT", "8080")
	os.Setenv("SERVICE_NAME", "sub-watch-backend")
	os.Setenv("JWT_SECRET_KEY", "secret")
}

func clearEnvVars() {
	os.Unsetenv("API_PORT")
	os.Unsetenv("SERVICE_NAME")
	os.Unsetenv("JWT_SECRET_KEY")
}

func TestLoadEnvs(t *testing.T) {
	setupEnvVars()
	defer clearEnvVars()

	config := config.LoadEnvs()

	if config.ApiPort != 8080 {
		t.Errorf("Expected PORT to be 8080, got %d", config.ApiPort)
	}
}

func TestGetEnvString(t *testing.T) {
	tests := []struct {
		name      string
		envVar    string
		value     string
		shouldPanic bool
	}{
		{
			name: "Success - existing env",
			envVar: "TEST_STRING_ENV",
			value: "value",
			shouldPanic: false,
		},
		{
			name: "Panic - missing env",
			envVar: "MISSING_STRING_ENV",
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected GetEnvString to panic")
					}
				}()
			}

			if tt.value != "" {
				os.Setenv(tt.envVar, tt.value)
				defer os.Unsetenv(tt.envVar)
			} else {
				os.Unsetenv(tt.envVar)
			}

			got := config.GetEnvString(tt.envVar)
			if !tt.shouldPanic && got != tt.value {
				t.Errorf("GetEnvString() = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestGetEnvNumber(t *testing.T) {
	tests := []struct {
		name      string
		envVar    string
		value     string
		want      int
		shouldPanic bool
	}{
		{
			name: "Success - existing env",
			envVar: "TEST_NUMBER_ENV",
			value: "1234",
			want: 1234,
			shouldPanic: false,
		},
		{
			name: "Panic - missing env",
			envVar: "MISSING_NUMBER_ENV",
			shouldPanic: true,
		},
		{
			name: "Panic - invalid value",
			envVar: "INVALID_NUMBER_ENV",
			value: "invalid",
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected GetEnvNumber to panic")
					}
				}()
			}

			if tt.value != "" {
				os.Setenv(tt.envVar, tt.value)
				defer os.Unsetenv(tt.envVar)
			} else {
				os.Unsetenv(tt.envVar)
			}

			got := config.GetEnvNumber(tt.envVar)
			if !tt.shouldPanic && got != tt.want {
				t.Errorf("GetEnvNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadEnvs_ConfigAlreadyInitialized(t *testing.T) {
	setupEnvVars()
	defer clearEnvVars()

	firstConfig := config.LoadEnvs()
	secondConfig := config.LoadEnvs()

	if firstConfig != secondConfig {
		t.Errorf("Expected LoadEnvs to return the same instance, but got different instances")
	}
}
