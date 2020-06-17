### Reference: https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
### --> https://github.com/afdolriski/golang-docker

### Makefile Reference:  https://github.com/Vungle/docker-golang/blob/master/Makefile  

## Steps - Building Using Makefile
Review and update Makefile  
Run command 'make build'  


## Steps - Building Manually
Use one of the dockerfiles:  
- better use Dockerfile.builder - it produces small image from 'scratch'
- The other one Dockerfile.noBuilder - uses the golang image itself and image becomes very big

```
go mod init webdocker1
go get github.com/gin-gonic/gin
go mod tidy
docker build . -t go-dock
OR
docker build -f <the dockerfile name you want> -t go-dock
docker run --rm -p 3000:3000 go-dock
curl http://localhost:3000/ping
--> you should receive a response {"message":"pong"}
```

