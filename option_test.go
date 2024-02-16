package option

import (
	"errors"
	"net/url"
	"testing"
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
		requireNoError(t, err)

		got, err := Apply[Client](Client{},
			OptionAPIKey(apiKey),
			OptionBaseURL(baseURL.String()),
		)
		requireNoError(t, err)

		want := Client{APIKey: apiKey, BaseURL: baseURL}
		if got.APIKey != want.APIKey || got.BaseURL.String() != want.BaseURL.String() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("fails on error", func(t *testing.T) {
		_, got := Apply[Client](Client{},
			Func[Client](func(Client) (Client, error) { return Client{}, errors.New("foo error") }),
		)

		want := "failed to apply option 0: foo error"
		if got.Error() != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func requireNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error when none expected: %v", err)
	}
}
