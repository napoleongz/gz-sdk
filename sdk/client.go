package sdk

import (
	"fmt"
)

type SDK struct {
	APIKey string
}

func NewSDK(apiKey string) *SDK {
	return &SDK{
		APIKey: apiKey,
	}
}

func (sdk *SDK) CallAPI(endpoint string) {
	fmt.Printf("Calling API with endpoint %s and API key %s\n", endpoint, sdk.APIKey)
	// Add your API call logic here
}
