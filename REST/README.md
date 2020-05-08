We store [global variables](../END.md) locally to use them all over the place.

```bash
curl -X PUT -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{ "first": "Jack", "last": "Sparrow" }' ${FIREBASE_URL}/users/jack/name.json
```

```bash
curl -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users.json
curl -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users/jack.json
curl -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users/jack/name.json
```

```bash
curl -X PATCH -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"last":"Jones"}' ${FIREBASE_URL}/users/jack/name.json
 
curl -X GET -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users.json?print=pretty'
 
 
curl -X PATCH -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"name":"Jack Jones"}' ${FIREBASE_URL}/users/jack/name.json
 
curl -X GET -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users.json?print=pretty'
 
 
curl -X PATCH -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"name":"Jack Jones"}' ${FIREBASE_URL}/users/jack.json
 
curl -X GET -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users.json?print=pretty
```

```bash
curl -X POST -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"user_id" : "jack", "text" : "Ahoy!"}' \
  ${FIREBASE_URL}/message_list.json
  
curl -X POST -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"user_id" : "john", "text" : "Aye!"}' \
  ${FIREBASE_URL}/message_list.json
  
curl -X POST -H "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"user_id" : "jeff", "text" : "Cool!"}' \
  ${FIREBASE_URL}/message_list.json
  
curl -X GET -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/message_list.json?print=pretty
```

```bash
curl -s -X PUT -T data/list.json -H "Authorization: Bearer ${ACCESS_TOKEN}" \
  ${FIREBASE_URL}/data/list.json?print=silent
  

curl -s -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/data/list.json
curl -s -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/data/list.json?print=pretty
curl -s -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/data/list/0.json?print=pretty
```

```bash
curl -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/.json
curl -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/.json?shallow=true
```

```bash
curl -X DELETE -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users/jack/name/last.json
curl -X DELETE -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/users/jack/name/first.json
  
curl -X DELETE -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/data/list.json
  
curl -X DELETE -H "Authorization: Bearer ${ACCESS_TOKEN}" ${FIREBASE_URL}/.json
```


