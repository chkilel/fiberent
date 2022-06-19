# FiberEnt | *Clean Architecture in Go* 🎉

<p align='center'>
  <img src='https://res.cloudinary.com/chkilel/image/upload/v1655654392/fiberent/fiberent-preview_lp0p4b.png' alt='FiberEnt' width='60%'/>
</p>


 FiberEnt is a clean architecture implementation in Go with the following frameworks:
- [Fiber](https://github.com/gofiber/fiber) 🚀 is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go.
- [Ent](https://github.com/ent/ent) 🎉 is an entity framework for Go,
Simple, yet powerful ORM for modeling and querying data.

<br/>

## Start development
> Docker must be installed.

Start docker container
```bash
  make docker-dev # or docker-compose up
```
then migrate database

```bash
  make migrate
```
<br />

# Steps to create a new entity

Install ent entity framework, check [https://entgo.io/docs/getting-started#installation](https://entgo.io/docs/getting-started#installation) for more information.

> **In the following example, we will create a new entity called `User`.**

1. Create an entity schema

   ```bash
   go run entgo.io/ent/cmd/ent init User # User is the name of the entity
   ```

2. Open up `<project>/ent/schema/user.go`

   - add your fields to the User schema, check **[Ent Field creation](https://entgo.io/docs/schema-fields)** for more information.
   - add your edges to the User schema, check **[Ent Edges creation](https://entgo.io/docs/schema-edges)** for more information.

3. Run go generate from the root directory of the project.

     ```bash presenter to the `presenter` folder, `user.go` file.
     go generate ./ent
     ```

4. Create `user entity` in the `<project>/entity` directory, `user.go` file

5. Define the `user` repository (Reader and Writer) Interface and the usecase (service) Interface in the `<project>/usecase/user` folder

6. Create the User **service** in the `<project>/usecase/user` folder, `service.go` file that implements the `Usecase` interface.

7. Create the User **repository** implementation in the `<project>/infrastructure/ent/repository/user_ent.go` that implements the `Repository` interface.

8. Add the handlers to the `<project>/api/handler` folder, `user.go` file, and the presenter to the `<project>/api/presenter` folder, `user.go` file.

## API requests

### Add user

```
curl -X "POST" "http://localhost:3030/api/v1/users" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
          "email": "adil@mail.com",
          "first_name": "Adil",
          "last_name": "Chehabi",
          "password": "password"
          }'
```
### Update user

```
curl -X "POST" "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
          "email": "adil@mail.com",
          "first_name": "Adil",
          "last_name": "Chkilel",
          "password": "password"
          }'
```

### Get a user

```
curl "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### Delete a user

```
curl -X "DELETE" "http://localhost:3030/api/v1/users/[USER_ID]" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

### List all users

```
curl "http://localhost:3030/api/v1/users" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```
