# quasar
Challenge Technical

## Go Version
  - go1.17

## Install Dependencies
  ``` go mod tidy ```

## Eviroment on Production
  - Linux or MAC OS
  ``` 
      export ENV=production
      export SERVER_PORT=80 
   ```
## Enviroment on Development
  ### All Systems
  - Create file .env
  - Add var SERVER_PORT=300

## Run Test
  ``` go test ./test/services/ && go test ./test/repository/ ```
  
## Run Proyect
  ``` go run ./cmd/main.go ```
  
