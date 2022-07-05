# gokit-account
这是youtube上tensor programming的gokit-account microservice 例子

>go run .
```
level=info service=account time:=2022-07-05T07:57:14.33025Z caller=main.go:33 msg="service started"
listening on port :8081

```
1.post http://localhost:8081/user
```
body: raw: json: {"name":"testname", "email":"123@gmail.com"}

response:
{
  "ok": "Success"
}

server:
service=account time:=2022-07-05T07:59:39.074834Z caller=repo.go:28 repo=CreateUser repo=CreateUser Success=(MISSING)
service=account time:=2022-07-05T07:59:39.074871Z caller=logic.go:34 logic=CreateUser CreatedUser=f21813da-127d-4c77-9dee-5e745807b6b2

```
2.get http://localhost:8081/user/f21813da-127d-4c77-9dee-5e745807b6b2
```
response:
{
  "email": "123@gmail.com"
}

server:
service=account time:=2022-07-05T08:01:12.219087Z caller=repo.go:41 repo=GetUser repo=GetUser Success=(MISSING)
service=account time:=2022-07-05T08:01:12.219154Z caller=logic.go:47 logic=GetUser GetUser=f21813da-127d-4c77-9dee-5e745807b6b2

```

...
