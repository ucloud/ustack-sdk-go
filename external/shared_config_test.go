package external

import (
	"github.com/ucloud/ustack-sdk-go/common"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ucloud/ustack-sdk-go/common/auth"
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
