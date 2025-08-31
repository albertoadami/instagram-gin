package configuration

import (
	"testing"

	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigSuccessfully(t *testing.T) {
	viper.AddConfigPath("../..") // project root for test
	viper.AddConfigPath(".")     // current dir for local/dev
	viper.AddConfigPath("/app")  // docker WORKDIR

	configuration, err := LoadConfig()
	assert.NoError(t, err)

	dbConfig := configuration.Database

	assert.Equal(t, dbConfig.Host, "localhost")
	assert.Equal(t, dbConfig.Port, 5432)
	assert.Equal(t, dbConfig.User, "postgres")
	assert.Equal(t, dbConfig.Password, "password")
	assert.Equal(t, dbConfig.Name, "instagram_gin_db")
}
