# API LIST

## 本機docker-compose up 即可 api port 8888
## mysql phpmyadmin prot 8090 (root/password)

#### 新增
```
[POST] /quiz/v1/insert
參數 {"name":"bbbb"}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 99999
}
```

#### 查詢
```
[GET] /quiz/v1/?id=1
參數 id
```

#### Response
```bash
{
    "data": [
        {
            "id": 3,
            "name": "cccc",
            "created_at": "2023-05-11T16:58:16Z",
            "updated_at": "2023-05-11T16:58:16Z"
        }
    ]
}


{
     "error_code": 99999
}
```

#### 刪除
```
[DELETE] /quiz/v1/:id
參數 id
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 99999
}
```

#### 更新
```
[PUT] /quiz/v1/update
參數 {"id": 1, "name":"cccc"}
```

#### Response
```bash
{
    "result": "ok"
}

{
     "error_code": 99999
}
```
