```bash
firebase --help

firebase database:set --help
firebase database:get --help
firebase database:update --help
firebase database:push --help
firebase database:remove --help
```

```bash
firebase database:set /users/jack/name -d '{ "first": "Jack", "last": "Sparrow" }'
```

```bash
firebase database:get /users
firebase database:get /users/jack
firebase database:get /users/jack/name
```

```bash
firebase database:update -y /users/jack/name -d '{"last":"Jones"}'
firebase database:get --pretty /users
 
firebase database:update -y /users/jack/name -d '{"name":"Jack Jones"}'
firebase database:get --pretty /users
 
firebase database:update -y /users/jack -d '{"name":"Jack Jones"}'
firebase database:get --pretty /users
```

```bash
firebase database:push /message_list -d '{"user_id" : "jack", "text" : "Ahoy!"}'
firebase database:push /message_list -d '{"user_id" : "john", "text" : "Aye!"}'
firebase database:push /message_list -d '{"user_id" : "jeff", "text" : "Cool!"}'
  
firebase database:get --pretty /message_list
```

```bash
firebase database:set /data/list "${LOCAL_DATA_DIR}/list.json"
 
firebase database:get --pretty /data/list
firebase database:get --shallow /data/list

firebase database:get --pretty /data/list/0
firebase database:get --pretty /data/list/0/i

firebase database:get --pretty /data/list --order-by-key --limit-to-first 1 
firebase database:get --pretty /data/list --order-by-key --limit-to-last 1
```

```bash
firebase database:get --shallow /
```

```bash
firebase database:remove -y /jack/name/last
firebase database:remove -y /jack/name/first
  
firebase database:remove -y /data/list
  
firebase database:remove -y /
```


