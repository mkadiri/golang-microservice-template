#### Setup
```    
# Build the application
bash build.sh

# Running the application
docker-compose up -d

# View logging 
docker logs -f {container-name}
```

#### Using the application

to access service to go 

`http://localhost:8181`
    
#### API endpoints


#### DEV mode
Setting `MODE=DEV` will enable dev fixtures to be run (seed data)

### Run the app locally
`docker-compose up mysql`

in goland 
- go to main.go
- press play button
- if not working as expected edit `go build main.go` config
- add env variables 
    - MYSQL_HOST=localhost
    - MYSQL_PORT=3307
    - the rest should be the same as docker-compose.yml