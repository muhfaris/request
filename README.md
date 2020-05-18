# Request
## Deskripsi
Library yang memudahkan kita untuk membuat http request lebih mudah.

## Konten
- [Install](https://github.com/muhfaris/request#install)
- [Penggunaan](https://github.com/muhfaris/request#penggunaan)
  - [Request GET, POST, Delete](https://github.com/muhfaris/request#request-post)
  - [Request dengan Authorization dan custom header Header](https://github.com/muhfaris/request#request-get-dengan-query-string-dan-custom-header)
- [Mime Types](https://github.com/muhfaris/request#mime-types)

## Install
Untuk menggunakan paket request, Anda harus menginstall Go dan setup Go workspace.
- Install paket, jalankan perintah berikut
`go get github.com/muhfaris/request`
- Import ke dalam kode:
`import "github.com/muhfaris/request"`

## Penggunaan
### Request POST
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

req := request.ReqApp{
    URL:url,
    ContentType: request.MimeTypeJSON,
    Body:payload,
}
resp, _ := app.POST()
log.Println(string(resp.Body))
```
### Request GET
```
import "github.com/muhfaris/request"
url := "https://jsonplaceholder.typicode.com/posts"

pq := request.ParamQuery{
    "userId":"1",
}
req := request.ReqApp{
    URL:url,
    ContentType:request.MimeTypeJSON,
    ParamQuery:request.ParamQuery{
    "userId":"1",
    }
}
resp, _ := req.GET()
log.Println(string(resp.Body))
```

#### Request Get dengan Query string dan custom header
```
import "github.com/muhfaris/request"

url := "https://jsonplaceholder.typicode.com/posts"

app := request.ReqApp{
	URL:         url,
	ContentType: "application/json",
	Headers: request.CustomHeader{
		"x-api-key": "somekeyhere123",
	},
	Authorization: "sometoken",
	QueryString: request.ParamQuery{
		"userId": "1",
	},
}

resp, _ := app.GET()
log.Println(string(resp.Body))
```
## Mime Types
- MimeTypeJSON = "application/json"
- MimeTypeFormData = "multipart/form-data"
- MimeTypeFormUrl = "application/x-www-form-urlencoded"
