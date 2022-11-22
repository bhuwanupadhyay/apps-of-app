## Create index

```
PUT /tenant-logs?pretty
{
    "settings" : {
        "number_of_shards" : 1,
        "number_of_replicas" : 1
    },
    "mappings" : {
        "properties" : {
            "tenantId" : { "type" : "keyword" },
            "message" : { "type" : "keyword" }
        }
    }
}
```

## Tenant1: Message

```
POST tenant-logs/_doc/
{
  "message": "Tenant (abc-xyz)",
  "teantId": "abc-xyz"
}
```

## Tenant2: Message
```
POST tenant-logs/_doc/
{
  "message": "Tenant (123-456)",
  "teantId": "123-456"
}
```

## Tenant1: Role

```
PUT _plugins/_security/api/roles/user_data
{
  "cluster_permissions": [
    "*"
  ],
  "index_permissions": [{
    "index_patterns": [
      "tenant-logs*"
    ],
    "dls": "{\"term\": { \"tenantId\": \"${user.name}\"}}",
    "allowed_actions": [
      "read"
    ]
  }]
}
```
