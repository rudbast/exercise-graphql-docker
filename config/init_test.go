package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	// Negative test.
	configFilename = ""
	assert.Error(t, Init())

	content := []byte(`
	[Database]
	[InvalidEntry]
	`)

	// Create temporary config file.
	func() {
		tmpFile, err := ioutil.TempFile(".", "sampleconfig")
		require.NoError(t, err)
		defer tmpFile.Close()

		_, err = tmpFile.Write(content)
		require.NoError(t, err)

		configFilename = tmpFile.Name()
	}()

	// Temporary file cleanup.
	defer os.Remove(configFilename)

	err := Init()
	assert.Error(t, err)
	assert.NotEqual(t, err, ErrMissingConfigFile)

	// Normal test.
	content = []byte(`
		[Database]
			Host = "localhost"
			User = "localuser"
			Pass = "localpass"
			Name = "localdb"
	`)

	err = ioutil.WriteFile(configFilename, content, 0644)
	require.NoError(t, err)

	assert.NoError(t, Init())
}

func TestGet(t *testing.T) {
	assert.Equal(t, globalCfg, Get())
}
