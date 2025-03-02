package passphrase

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	state = `
    {"salt": "v1:fozI5u6B030=:v1:F+6ZduKKd8G0/V7L:PGMFeIzwobWRKmEAzUdaQHqC5mMRIQ=="}
`
	brokenState = `
    {"salt": "fozI5u6B030=:v1:F+6ZduL:PGMFeIzwobWRKmEAzUdaQHqC5mMRIQ=="}
`
)

func setIncorrectPassphraseTestEnvVars() func() {
	clearCachedSecretsManagers()

	oldPassphrase := os.Getenv("PULUMI_CONFIG_PASSPHRASE")
	oldPassphraseFile := os.Getenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	os.Setenv("PULUMI_CONFIG_PASSPHRASE", "password123")
	os.Unsetenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	return func() {
		os.Setenv("PULUMI_CONFIG_PASSPHRASE", oldPassphrase)
		os.Setenv("PULUMI_CONFIG_PASSPHRASE_FILE", oldPassphraseFile)
	}
}

//nolint:paralleltest // mutates environment variables
func TestPassphraseManagerIncorrectPassphraseReturnsErrorCrypter(t *testing.T) {
	setupEnv := setIncorrectPassphraseTestEnvVars()
	defer setupEnv()

	manager, err := NewPromptingPassphaseSecretsManagerFromState([]byte(state))
	assert.NoError(t, err) // even if we pass the wrong provider, we should get a lockedPassphraseProvider

	assert.Equal(t, manager, &localSecretsManager{
		state:   localSecretsManagerState{Salt: "v1:fozI5u6B030=:v1:F+6ZduKKd8G0/V7L:PGMFeIzwobWRKmEAzUdaQHqC5mMRIQ=="},
		crypter: &errorCrypter{},
	})
}

func setCorrectPassphraseTestEnvVars() func() {
	clearCachedSecretsManagers()

	oldPassphrase := os.Getenv("PULUMI_CONFIG_PASSPHRASE")
	oldPassphraseFile := os.Getenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	os.Setenv("PULUMI_CONFIG_PASSPHRASE", "password")
	os.Unsetenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	return func() {
		os.Setenv("PULUMI_CONFIG_PASSPHRASE", oldPassphrase)
		os.Setenv("PULUMI_CONFIG_PASSPHRASE_FILE", oldPassphraseFile)
	}
}

//nolint:paralleltest // mutates environment variables
func TestPassphraseManagerIncorrectStateReturnsError(t *testing.T) {
	setupEnv := setCorrectPassphraseTestEnvVars()
	defer setupEnv()

	_, err := NewPromptingPassphaseSecretsManagerFromState([]byte(brokenState))
	assert.Error(t, err)
}

//nolint:paralleltest // mutates environment variables
func TestPassphraseManagerCorrectPassphraseReturnsSecretsManager(t *testing.T) {
	setupEnv := setCorrectPassphraseTestEnvVars()
	defer setupEnv()

	sm, _ := NewPromptingPassphaseSecretsManagerFromState([]byte(state))
	assert.NotNil(t, sm)
}

func unsetAllPassphraseEnvVars() func() {
	clearCachedSecretsManagers()

	oldPassphrase := os.Getenv("PULUMI_CONFIG_PASSPHRASE")
	oldPassphraseFile := os.Getenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	os.Unsetenv("PULUMI_CONFIG_PASSPHRASE")
	os.Unsetenv("PULUMI_CONFIG_PASSPHRASE_FILE")
	return func() {
		os.Setenv("PULUMI_CONFIG_PASSPHRASE", oldPassphrase)
		os.Setenv("PULUMI_CONFIG_PASSPHRASE_FILE", oldPassphraseFile)
	}
}

//nolint:paralleltest // mutates environment variables
func TestPassphraseManagerNoEnvironmentVariablesReturnsError(t *testing.T) {
	setupEnv := unsetAllPassphraseEnvVars()
	defer setupEnv()

	_, err := NewPromptingPassphaseSecretsManagerFromState([]byte(state))
	assert.NotNil(t, err, strings.Contains(err.Error(), "unable to find either `PULUMI_CONFIG_PASSPHRASE` nor "+
		"`PULUMI_CONFIG_PASSPHRASE_FILE`"))
}
