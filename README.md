# Request
## Deskripsi
Library sederhana untuk http request. Saat ini library hanya support untuk request GET, POST, DELETE dan PATCH.

### Contoh
#### Request Post dengan data
```
import "github.com/muhfaris/request"

url := "https://vendors.paddle.com/api/2.0/product/generate_pay_link"

payload := []byte(`{
  "vendor_id": "xxxxx",
  "vendor_auth_code": "xxxx",
  "title": "buy managix plan",
  "webhook_url": "https://app.managix.id/webhook",
  "prices": [
    "USD:60"
  ],
  "customer_email": "akun@gmail.com",
  "customer_country": "ID",
  "customer_postcode": "4000",
  "passthrough": "111",
  "recurring_prices": [
    "USD:60"
  ],
  "quantity": 1,
  "quantity_variabel": 1,
  "discountable": 0
}`)

	app, _ := request.New(url, "application/json", "", payload, nil)
	resp, _ := app.POST()
	log.Println(string(resp.Body))
```

#### Konvert body data
```
body := request.BodyByte(exampleStruct)
```

#### Request Get dengan Query string
```
import "github.com/muhfaris/request"
url := "https://jsonplaceholder.typicode.com/posts"

pq := request.ParamQuery{
    "userId":1
}
app, _ := request.New(url, "application/json", "", "", pq )
resp, _ := app.GET()
log.Println(string(resp)
```
