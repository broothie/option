package option

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Client struct {
	APIKey  string
	BaseURL *url.URL
}

func OptionAPIKey(apiKey string) Func[Client] {
	return func(client Client) (Client, error) {
		client.APIKey = apiKey
		return client, nil
	}
}

func OptionBaseURL(baseURL string) Func[Client] {
	return func(client Client) (Client, error) {
		parsed, err := url.Parse(baseURL)
		if err != nil {
			return Client{}, err
		}

		client.BaseURL = parsed
		return client, nil
	}
}

func TestApply(t *testing.T) {
	t.Run("applies options", func(t *testing.T) {
		apiKey := "foo-api-key"
		baseURL, err := url.Parse("https://example.com")
		require.NoError(t, err)

		client, err := Apply[Client](Client{},
			OptionAPIKey(apiKey),
			OptionBaseURL(baseURL.String()),
		)
		require.NoError(t, err)

		assert.Equal(t, Client{APIKey: apiKey, BaseURL: baseURL}, client)
	})

	t.Run("collects errors", func(t *testing.T) {
		_, err := Apply[Client](Client{},
			OptionBaseURL("%"),
			Func[Client](func(Client) (Client, error) { return Client{}, errors.New("foo error") }),
		)

		assert.EqualError(t, err, "failed to apply option 0: parse \"%\": invalid URL escape \"%\"\nfailed to apply option 1: foo error")
	})
}
