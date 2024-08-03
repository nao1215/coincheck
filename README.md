## coincheck public &amp; private API client
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/coincheck.svg)](https://pkg.go.dev/github.com/nao1215/coincheck)
![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/coincheck/coverage.svg)
[![MultiPlatformUnitTest](https://github.com/nao1215/coincheck/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/coincheck/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/coincheck/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/coincheck/actions/workflows/reviewdog.yml)
[![gitleaks](https://github.com/nao1215/coincheck/actions/workflows/gitleak.yml/badge.svg)](https://github.com/nao1215/coincheck/actions/workflows/gitleak.yml)


>[!IMPORTANT]
> This library is under development and is not yet ready for production use.

The coincheck package is a client for the API provided by Coincheck, Inc., which operates the cryptocurrency exchange (Coincheck). The coincheck package offers two types of APIs:

- Public API: Can be executed without authentication
- Private API: Requires authentication using the API Key and API Secret issued by the Coincheck service.

The coincheck package supports both Public and Private APIs.
- [Coincheck official API documentation](https://coincheck.com/documents/exchange/api)
- [Coincheck official API client](https://github.com/coincheckjp/coincheck-go)

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
| GET /api/ticker | [GetTicker()](https://pkg.go.dev/github.com/nao1215/coincheck#Client.GetTicker) | Check latest ticker information. |
| GET /api/trades | [GetTrades()](https://pkg.go.dev/github.com/nao1215/coincheck#Client.GetTrades) | You can get current order transactions. |
| GET /api/order_books | [GetOrderBooks()](https://pkg.go.dev/github.com/nao1215/coincheck#Client.GetOrderBooks) | Fetch order book information. |

### Private API

| API | Method Name |Description |
| :--- | :--- | :--- |
| GET /api/bank_accounts | [GetBankAccounts()](https://pkg.go.dev/github.com/nao1215/coincheck#Client.GetBankAccounts) | Display list of bank account you registered (withdrawal).|

## License

[MIT License](./LICENSE)


## Contribution
First off, thanks for taking the time to contribute! See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information. Contributions are not only related to development. For example, GitHub Star motivates me to develop! Please feel free to contribute to this project.


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://debimate.jp/"><img src="https://avatars.githubusercontent.com/u/22737008?v=4?s=70" width="70px;" alt="CHIKAMATSU Naohiro"/><br /><sub><b>CHIKAMATSU Naohiro</b></sub></a><br /><a href="https://github.com/nao1215/coincheck/commits?author=nao1215" title="Code">💻</a></td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td align="center" size="13px" colspan="7">
        <img src="https://raw.githubusercontent.com/all-contributors/all-contributors-cli/1b8533af435da9854653492b1327a23a4dbd0a10/assets/logo-small.svg">
          <a href="https://all-contributors.js.org/docs/en/bot/usage">Add your contributions</a>
        </img>
      </td>
    </tr>
  </tfoot>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
