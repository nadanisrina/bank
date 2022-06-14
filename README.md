# bank
1. git checkout -b feature/create_user
2. go mod tidy 
3. run docker : make postgres
4. make createdb
5. make migrateup
6. go run main.go
