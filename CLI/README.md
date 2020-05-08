```bash
firebase --help
firebase database:get --help
```

```bash
firebase database:get /
```



```bash
firebase database:push -d '{ "first": "Jack", "last": "Sparrow" }' \
  'https://hybrid-cloud-22365.firebaseio.com/users/jack/name.json'
```

```bash
curl 'https://hybrid-cloud-22365.firebaseio.com/users.json'
curl 'https://hybrid-cloud-22365.firebaseio.com/users/jack.json'
curl 'https://hybrid-cloud-22365.firebaseio.com/users/jack/name.json'
```

```bash
curl -X PATCH -d '{"last":"Jones"}' \
 'https://hybrid-cloud-22365.firebaseio.com/users/jack/name.json'
 
curl -X GET 'https://hybrid-cloud-22365.firebaseio.com/users.json?print=pretty'
 
 
curl -X PATCH -d '{"name":"Jack Jones"}' \
 'https://hybrid-cloud-22365.firebaseio.com/users/jack/name.json'
 
curl -X GET 'https://hybrid-cloud-22365.firebaseio.com/users.json?print=pretty'
 
 
curl -X PATCH -d '{"name":"Jack Jones"}' \
 'https://hybrid-cloud-22365.firebaseio.com/users/jack.json'
 
curl -X GET 'https://hybrid-cloud-22365.firebaseio.com/users.json?print=pretty'
```

```bash
curl -X POST -d '{"user_id" : "jack", "text" : "Ahoy!"}' \
  'https://hybrid-cloud-22365.firebaseio.com/message_list.json'
  
curl -X POST -d '{"user_id" : "john", "text" : "Aye!"}' \
  'https://hybrid-cloud-22365.firebaseio.com/message_list.json'
  
curl -X POST -d '{"user_id" : "jeff", "text" : "Cool!"}' \
  'https://hybrid-cloud-22365.firebaseio.com/message_list.json'
  
curl -X GET 'https://hybrid-cloud-22365.firebaseio.com/message_list.json?print=pretty'
```

```bash
curl -s -X PUT -T data/list.json \
  'https://hybrid-cloud-22365.firebaseio.com/data/list.json?print=silent'
  

curl -s 'https://hybrid-cloud-22365.firebaseio.com/data/list.json'
curl -s 'https://hybrid-cloud-22365.firebaseio.com/data/list.json?print=pretty'
curl -s 'https://hybrid-cloud-22365.firebaseio.com/data/list/0.json?print=pretty'
```

```bash
curl 'https://hybrid-cloud-22365.firebaseio.com/.json'
curl 'https://hybrid-cloud-22365.firebaseio.com/.json?shallow=true'
```

```bash
curl -X DELETE 'https://hybrid-cloud-22365.firebaseio.com/users/jack/name/last.json'
curl -X DELETE 'https://hybrid-cloud-22365.firebaseio.com/users/jack/name/first.json'
  
curl -X DELETE 'https://hybrid-cloud-22365.firebaseio.com/data/list.json'
  
curl -X DELETE 'https://hybrid-cloud-22365.firebaseio.com/.json'
```


