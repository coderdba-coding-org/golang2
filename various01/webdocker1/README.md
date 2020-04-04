### Reference: https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
### --> https://github.com/afdolriski/golang-docker

## Steps
Use one of the dockerfiles:  
- better use Dockerfile.builder - it produces small image from 'scratch'
- The other one Dockerfile.noBuilder - uses the golang image itself and image becomes very big

```
go mod init webdocker1
go get github.com/gin-gonic/gin
go mod tidy
docker build . -t go-dock
curl http://localhost:3000/ping
--> you should receive a response {"message":"pong"}
```

