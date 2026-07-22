package external

import (
	"testing"
	"third_party/platform-sdk-go/common"

	"github.com/stretchr/testify/assert"

	"third_party/platform-sdk-go/common/auth"
)

func TestLoadSharedConfig(t *testing.T) {
	cfg, err := LoadUCloudConfigFile(
		TestValueEnvUCloudSharedConfigFile,
		TestValueEnvUCloudProfile,
	)
	assert.NoError(t, err)
	checkTestClientConfig(t, cfg)

	cred, err := LoadUCloudCredentialFile(
		TestValueEnvUCloudSharedCredentialFile,
		TestValueEnvUCloudProfile,
	)
	assert.NoError(t, err)
	checkTestCredential(t, cred)
}

func checkTestCredential(t *testing.T, cred *auth.Credential) {
	assert.Equal(t, TestValueFileUCloudPublicKey, cred.PublicKey)
	assert.Equal(t, TestValueFileUCloudPrivateKey, cred.PrivateKey)
}

func checkTestClientConfig(t *testing.T, cfg *common.Config) {
	assert.Equal(t, TestValueFileUCloudProjectId, cfg.ProjectId)
	assert.Equal(t, TestValueFileUCloudRegion, cfg.Region)
	assert.Equal(t, TestValueFileUCloudTimeout, cfg.Timeout)
	assert.Equal(t, TestValueFileUCloudBaseUrl, cfg.BaseUrl)
}
