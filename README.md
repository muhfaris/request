# Request
## Deskripsi
Library yang memudahkan kita untuk membuat http request lebih mudah.

## Fitur
- Request dengan method Get, Post, Delete, Patch.
- Sistem retry, untuk melakukan request ulang jika terjadi error. Kamu bisa set total retry pada field `Retry` dan kamu juga bisa melakukan setting jeda waktu antar request yang satu dengan yang lain.
- Parse response data ke struct / map[string].

## Install
Untuk menggunakan paket request, Anda harus menginstall Go dan setup Go workspace.
- Install paket, jalankan perintah berikut
`go get github.com/muhfaris/request`

- Import ke dalam kode:
`import "github.com/muhfaris/request"`

## Get Request dan parse response data ke struct
```
import "github.com/muhfaris/request"

type Person struct{
    Name string
    Address string
}

func main(){
    var person = &Person{}
    resp, err := request.Get(
        &request.Config{
                "URL": "http://<domain_api>",
        }).Parse(&person)

    // handler response
    if resp.Error != nil {

    }


    fmt.Println(person.Name)
    fmt.Println(person.Address)
}
```

## Post request dengan retry
```
import "github.com/muhfaris/request"
func main(){
    resp, err := request.Post(
        &request.Config{
            "URL": "https://facebook.com/v1/api/profile",
            "Method": "POST",
            "Retry": 1, 
            "Delay": 10 * time.Seconds,
        },
    )
}
```

## Post application/json
```

func main(){
    body, _ := BodyByte(
            map[string]string{
                    "name": "faris",
                    "job":  "leader",
            },
    )

    // wrap response to map string  
    var data map[string]interface{}
    resp := request.Post(
        &request.Config{
            URL:  "https://reqres.in/api/users",
            Body: body,
        }).Parse(&data)

    // handle error
    if resp.Error != nil {
        // TODO Error
    }

    fmt.Println(data["id"])
}


```
