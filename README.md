## A simple backend for practicing Postgresql with Golang + Gorm and implementation of Email Auth via link using google's smtp.

### Get started :

1. Run Postgresql server by using

```bash
brew services start postgresql
```

2. Add env.json in the root directory of the project

{
    "USER":"your email which will send auth links",
    "PASSWORD":"app password set up in your google account"
}

3. Compile and run :

```bash
make run
```

4. Use the server

5. Press [control + c] to stop the server.

6. Stop Postgresql

```bash
brew services stop postgresql
```