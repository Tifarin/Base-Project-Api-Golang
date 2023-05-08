- menginstall goose untuk membuat migration

$ go install github.com/pressly/goose/v3/cmd/goose@latest

- membuat migration table baru

$ goose create add_users_table sql

- membuat migration seeder baru

$ goose create add_users_seeder sql

- menjalankan migration ke database

$ goose postgres "host=localhost port=5432 user=postgres password=password dbname=BaseApiGolang sslmode=disable" up

- mengecek status

$ goose postgres "user=postgres password=password dbname=BaseApiGolang sslmode=disable" status

- melakukan reset database

$ goose postgres "user=postgres password=password dbname=BaseApiGolang sslmode=disable" reset
