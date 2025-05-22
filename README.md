# API CRUD Golang
## Instruction Setup
1. Install MySQL this project using mysql `https://dev.mysql.com/downloads/installer/`
2. Install all package dependencies using
    ```
   go mod download
   ```
3. Check dependency injection using `wire` on terminal
   ```
   wire
   ```
4. Change configuration on MySQL in `wire.go` funtcion `InitializeDB`
    ```
   db.ConnectDB("change username:change password@tcp(127.0.0.1:3306)/change database?charset=utf8mb4&parseTime=True&loc=Local")
    ```
5. Now test migration

   For migration up
   ```
   goose -dir ./migrations mysql "change username:change password@tcp(localhost:3306)/change database" up
   ```

   For migration down/rollback
   ```
   goose -dir ./migrations mysql "change username:change password@tcp(localhost:3306)/change database" down
   ```
6. Run project install air golang
   ```
   go install github.com/air-verse/air@latest
   ```
   on terminal
   ```
   air
   ```
7. Run project local
   ```
   http://localhost:8080/api/v1/
   ```

## Daily Task
### 27 April 2025
- Setup Project and GitHub
- Add dockerfile

### 28 April
- Add wire dependency injection
- Adding migration using goose

### 30 April
- Create feature Signin
- Create feature Signup
- Create feature RBAC for crud
- add drop migration RBAC

### 1 May
- landing page api
- Create feature read collection profile
- Create feature read document profile
- Create feature read insert profile
- Create feature read update profile
- Create feature read delete profile

## Endpoints API
**1. Signin [GET]**
   ```
   localhost:8080/api/v1/auth/signin
   
   body:
      - email
      - password
   ```
   ![img_2.png](img_2.png)

**2. Signup [GET]**
   ```
   http://localhost:8080/api/v1/auth/signup

   body:
      - username
      - email
      - password
   ```
   ![img_1.png](img_1.png)

**3. Collection Profile [GET]**
   ```
   http://localhost:8080/api/v1/profile/
   
   Header:
      - Authorization: (from signin token)
   ```
   ![img_3.png](img_3.png)

**4. Document Profile [GET]**
   ```
   http://localhost:8080/api/v1/profile/:id_profile
   
   Header:
      - Authorization: (from signin token)
   ```
   ![img_4.png](img_4.png)

**5. Create Profile [POST]**
   ```
   http://localhost:8080/api/v1/profile/management
   
   Header:
      - Authorization: (from signin token)
   Body:
      - Bio
      - AvatarUrl
   ```
   ![img_5.png](img_5.png)

**6. Update Profile [PUT]**
   ```
   http://localhost:8080/api/v1/profile/management
   
   Header:
      - Authorization: (from signin token)
   Body:
      - Bio
      - AvatarUrl
   ```
   ![img_7.png](img_7.png)
   ![img_6.png](img_6.png)

**7. Delete Profile [DELETE]**
   ```
   http://localhost:8080/api/v1/profile/management
   
   Header:
      - Authorization: (from signin token)
   ```
   ![img_8.png](img_8.png)
   
## Note
for RBAC add manualy on table user_role, example:
1 - admin
2 - member 
3 - guest
```
INSERT INTO user_role (user_id, role_id) VALUES (1, 1)
```"# api_pegadaian" 
