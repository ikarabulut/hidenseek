# HideNSeek

Go web app to allow for encrypted and secure secret sharing. When a secret is added it will update the local secret map JSON file running on your server. When a secret is viewed it will delete that secret from the file. In order to decrypt a secret you must have the password and salt that was used to encrypt.

## Running the app
1. Ensure you have go installed locally and clone the repository
2. Set up ENV variables
```
PASSWORD= Password for encryption
SALT= Salt for encryption
```

3. From the root directory run `go run main.go`
  - This will run on `localhost:3000`

**Add a secret**
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"plain_text":"hello-world"}' \
  http://localhost:3000/secret
```
**View a secret**
Send a GET request to localhost:3000/secret with a query param of `hash`
```
curl "localhost:3000/secret?hash={ENTER HASH}"
```

## Testing
- Navigate to a test folder `/{package}_Test`
- Run:
```
  go test -v
```