Step 1
Install go seperti biasa

Step 2
Inisialisasi db menggunakan query yang sudah disediakan dan juga connect pada DB sesuai dengan user masing masing pada config.go

Step 3
Import beberapa Package yang digunakan
go get -u github.com/gorilla/mux
go get -u	github.com/joho/godotenv v1.5.1
go get -u	github.com/lib/pq v1.10.9
go get -u github.com/google/uuid

Step 4
Inisialiasi port pada main.go

Step 5
Run pada localhost sesuai main.go
go run main.go

Step 6
Test menggunakan postman, dll sesuai dengan router.go
