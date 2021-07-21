# metservice-go [![godocs.io](https://godocs.io/git.sr.ht/~kota/metservice-go?status.svg)](https://godocs.io/git.sr.ht/~kota/metservice-go)

A go library for reading weather data from [Metservice.](https://www.metservice.com/)

Discussion and patches can be found [here](https://lists.sr.ht/~kota/public-inbox).

## Example

```go
import (
	"context"
	"fmt"

	"git.sr.ht/~kota/metservice-go"
)

func main() {
	client := NewClient()
	ctx := context.Background()

	forecast, _, err := client.Forecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*forecast.LocationIPS)
}
