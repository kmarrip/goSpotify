package authorize

import (
	"os"
	"testing"

	"github.com/chaithanyaMarripati/goSpotify/config"
	"github.com/stretchr/testify/assert"
)

func TestConstructAuthorizeReq(t *testing.T) {
	os.Setenv("baseUrl", "baseUrl")
	os.Setenv("clientId", "aClient")
	os.Setenv("redirectUrl", "aRedirectUri")
	os.Setenv("scopes", "aScope")
	config.SetConfigVar()

	req := ConstructAuthorizeReq()

	assert.Equal(t, "baseUrlclient_id=aClient&redirect_uri=aRedirectUri&response_type=code&scope=aScope", req)
}
