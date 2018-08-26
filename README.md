# L44

## Installing
```
# Install dependencies
go get github.com/rs/xid github.com/gin-gonic/gin github.com/auth0-community/go-auth0 gopkg.in/square/go-jose.v2

# Set env variables
export AUTH0_API_IDENTIFIER=https://l44
export AUTH0_DOMAIN=mproske.auth0.com

# Run
go run main.go
```