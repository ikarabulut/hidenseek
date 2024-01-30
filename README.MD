# HideNSeek

Go web app to allow for encrypted and secure secret sharing. When a secret is added it will update the local secret map JSON running on your server. When a secret is viewed it will delete that secret from the file.

## Running the app
1. Ensure you have go installed locally and clone the repository
2. From the root directory run `go run main.go`
  - This will run on `localhost:3000`

**Add a secret**
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"plain_text":"hello-world"}' \
  http://localhost:3000/secret
```
**View a secret**
TO BE IMPLEMENTED