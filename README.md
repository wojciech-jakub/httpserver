# httpserver


## to build 
``docker build -t hello_go_http . `` 
or 
``docker compose build``

## to run 
``docker run -p 8080:8080 -t hello_go_http``
or
``docker compose run``

## to execute endpoint from terminal
``curl "http://0.0.0.0:8080/random/mean?requests=1&length=10"
``

## what needs to be updated or refactored 
1. use .env for the http server configuration and external API address
2. fix error handling from the go routines in external data fetching 
3. create directories to for the clean architecture approach
4. change random org api defintion to use API key - current solution can be blocked for huge amount of the requests - cloudflare 