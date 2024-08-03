## coincheck public &amp; private API client
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-0-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

>[!IMPORTANT]
> This library is under development and is not yet ready for production use.

The coincheck package is a client for the API provided by Coincheck, Inc., which operates the cryptocurrency exchange (Coincheck). The coincheck package offers two types of APIs:

- Public API: Can be executed without authentication
- Private API: Requires authentication using the API Key and API Secret issued by the Coincheck service.

The coincheck package supports both Public and Private APIs.
- [Official API documentation](https://coincheck.com/documents/exchange/api)
- [Official API client](https://github.com/coincheckjp/coincheck-go)

## Supported OS and go version

- OS: Linux, macOS, Windows
- Go: 1.20 or later

## Example

An example of executing the Public API is shown below.

```go
package main

import (
	"context"
	"fmt"

	"github.com/nao1215/coincheck"
)

func main() {
	client, err := coincheck.NewClient()
	if err != nil {
		panic(err)
	}

	// Get the latest ticker
	ticker, err := client.GetTicker(context.Background(), coincheck.GetTickerInput{
		Pair: coincheck.PairETCJPY,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Last: %d\n", ticker.Last)
	fmt.Printf("Bid: %d\n", ticker.Bid)
	fmt.Printf("Ask: %d\n", ticker.Ask)
	fmt.Printf("High: %d\n", ticker.High)
	fmt.Printf("Low: %d\n", ticker.Low)
	fmt.Printf("Volume: %s\n", ticker.Volume)
	fmt.Printf("Timestamp: %d\n", ticker.Timestamp)

    // Output:
    // Last: 4000.000000
    // Bid: 3980.020000
    // Ask: 4000.000000
    // High: 4220.000000
    // Low: 4000.000000
    // Volume: 339.150000
    // Timestamp: 1722661800.000000
}
```

If you want to execute the Private API, you need to create a client with the API Key and API Secret issued by the Coincheck service.

```go
	client, err := coincheck.NewClient(WithCredentials("API_KEY", "API_SECRET"))
```

## API List
### Public API

| API | Method Name |Description |
| :--- | :--- | :--- |
| GET /api/ticker | GetTicker() | Check latest ticker information |

### Private API

| API | Method Name |Description |
| :--- | :--- | :--- |
| GET /api/bank_accounts | GetBankAccounts() | Display list of bank account you registered (withdrawal).|

## License

[MIT License](./LICENSE)


## Contribution
First off, thanks for taking the time to contribute! See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information. Contributions are not only related to development. For example, GitHub Star motivates me to develop! Please feel free to contribute to this project.


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
