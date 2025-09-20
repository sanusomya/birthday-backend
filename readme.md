setting up the server -:
1. setup a local mongodb server
2. create a .env file with the keys 
    1. db_url
    2. database
    3. db_coll
    4. birthday_app_port
3. run go mod tidy to get all dependencies
4. run go run main.go and the server starts at 8002 port on local host.