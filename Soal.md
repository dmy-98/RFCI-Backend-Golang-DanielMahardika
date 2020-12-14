# Sorting

Buatlah sebuah program untuk menampilkan dan menghitung jumlah swap yang diperlukan agar sebuah deret angka menjadi berurutan.

Contoh input
```
4 9 7 5 8 9 3
```

Contoh output. Jika program dijalankan, output harus sama persis seperti dibawah ini
```
1. [7,9] -> 4 7 9 5 8 9 3
2. [5,9] -> 4 7 5 9 8 9 3
3. [5,7] -> 4 5 7 9 8 9 3
4. [8,9] -> 4 5 7 8 9 9 3
5. [3,9] -> 4 5 7 8 9 3 9
6. [3,9] -> 4 5 7 8 3 9 9
7. [3,8] -> 4 5 7 3 8 9 9
8. [3,7] -> 4 5 3 7 8 9 9 
9. [3,5] -> 4 3 5 7 8 9 9
10. [3,4] -> 3 4 5 7 8 9 9 

Jumlah swap: 10
```

# Client Server

### Specification
Untuk pembuatan server, dapat memilih salah satu dari fromework berikut : 
1. Gin "https://github.com/gin-gonic/gin"
2. Mux "https://github.com/gorilla/mux"


### Instruction
Buatlah aplikasi client server dengan kriteria sebagai berikut :

## Client

Client akan mengirim data json seperti berikut menggunakan POST method :

```json
{
    "author": "nama",
    "title": "judul",
    "comments": [
        {
            "message": "message1"
        },
        {
            "message": "message2"
        }
    ]
}
```

## Server

Setiap Server menerima data dari client, payload dari body akan disimpan ke file `.log`, lalu memberikan response dengan HTTP Status Code 201.

Contoh isi file `author.log`
```
[2020-07-28T16:23:40+07:00] Success: POST http://127.0.0.0:8080/blogs {"author" : "name"}
```

Contoh isi file `title.log`
```
[2020-07-28T16:23:40+07:00] Success: POST http://127.0.0.0:8080/blogs {"title" : "title"}
```

Contoh isi file `message.log`
```
[2020-07-28T16:23:40+07:00] Success: POST http://127.0.0.0:8080/blogs {"comments" : [{"message": "messag1"}, {"message": "message2"}]}
```

## Fix the broken thing

Terdapat soal seperti berikut :

```
Input: [3, 5, 2, -4, 8, 11] dan penjumlahan = 7
Output: [[11, -4], [2, 5]]

note : 
11, -4 dan 2, 5 adalah hasil dari penjumlahan 7
```

Jalankan code golang dibawah dan pastikan input dan output seperti yang diinginkan

```go
package main

import (
	"fmt"
)

func main() {
	collection := []int{3, 5, 2, -4, 8, 11}
	fmt.Println(foo(7, collection))
}

func foo(input int, numbers []int) (int, []int) {
	var result []int
	var nextIndex int

	for idx, number := range numbers {
		nextIndex = idx + 1

		if nextIndex >= len(numbers) {
			break;
		}
		nextNumber := numbers[nextIndex]
		if number + nextNumber == input {
			result = append(result, number)
			result = append(result, nextNumber)
		}
	}
	return input, result
}


```
