## A simple backend for practicing Postgresql with Golang + Gorm and implementation of Email Auth via link using google's smtp.

### Get started :

1. Clone the project

```bash
git clone https://github.com/Harsh-apk/goPostgressGorm.git
```

2. Run Postgresql server by using

```bash
brew services start postgresql
```

3. Add env.json in the root directory of the project

```bash
touch env.json
```

{
    "USER":"your email which will send auth links",
    "PASSWORD":"app password set up in your google account"
}

4. Compile and run :

```bash
make run
```

5. Use the server

6. Press [control + c] to stop the server.

7. Stop Postgresql

```bash
brew services stop postgresql
```