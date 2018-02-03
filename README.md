# Firebase REST Client

Install using go packager of choice. 

```
github.com/aj0strow/firebase
```

[Godoc](https://godoc.org/github.com/aj0strow/firebase)

## Usage

```go
package main

import (
    "github.com/aj0strow/firebase"
    "os"
    "log"
)

func main() {
    // If you want legacy admin access. 
    secret := os.Getenv("FIREBASE_SECRET")

    fbase := firebase.App{
        DatabaseURL: os.Getenv("FIREBASE_URL"),
        Secret: secret,
    }

    // If you sign a token.
    client := fbase.Auth(token)
    
    // If you prefer admin access.
    admin := fbase.Admin()
    
    // Set up a new path.
    users := firebase.Reference{"users"}
    
    // Reference child path.
    me := users.Child("aj0strow")
    
    // Create a new path.
    newbie := users.Push()
    
    // Write data to a path.
    type User struct {
        Name string `json:"name"`
    }
    err = admin.Write(newbie, &User{Name: "Newbie"})
    
    // Query and parse data.
    var userIds map[string]interface{}
    err = admin.Query(users, &firebase.Params{Shallow: true}, &userIds)
    
    // Remove data at paths.
    err = admin.Remove(me)
    
    // Update multiple paths in a single transaction.
    batch := firebase.NewBatch()
    for uid := range userIds {
        batch.Set(users.Child(uid, "private"), true)
    }
    err = admin.UpdateByMerge(batch)
}
```

## Custom Tokens

Signing custom JWTs is out of scope. If you want to [create custom tokens](https://firebase.google.com/docs/auth/admin/create-custom-tokens) check out **[knq/jwt](https://github.com/knq/jwt)** and use the following snippet.

```go
import (
    "github.com/knq/jwt"
    "github.com/knq/jwt/gserviceaccount"
    "time"
)

gSvcAct := &gserviceaccount.GServiceAccount{
    // private fields
}

signer, err := gSvcAct.Signer()
if err != nil {
    // handle error
}

now := time.Now().Unix()
token, err := jwt.Encode(jwt.RS256, signer, map[string]interface{}{
    "iss": gSvcAct.ClientEmail,
    "sub": gSvcAct.ClientEmail,
    "aud": "https://identitytoolkit.googleapis.com/google.identity.identitytoolkit.v1.IdentityToolkit",
    "iat": now,
    "exp": now + 60*60,
    "uid": "custom uid here",
    "claims": map[string]interface{}{
        "custom": "stuff",
    },
}
if err != nil {
    // handle error
}

// Ready to go.
client := fbase.Auth(token)
```
