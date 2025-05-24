# consul-blocking-query-demo

## Setup
Run below command in terminal to start the demo.
```bash
docker-compose up -d
```

## Scenario 1
Say the service is not existing in consul.

### Q: How do I start the blocking query without knowing the start index?
Always start the blocking query with index 0, it will return the latest index and then you can use that index for next blocking query.

```
2025/05/24 13:57:21 Starting blocking query watcher...
2025/05/24 13:57:21 [Query] Waiting for service updates, index 0...
2025/05/24 13:57:21 [Query] Waiting for service updates, index 17...
[Query] Got 0 instances at index 17
```

## Scenario 2
### Q: A different service is registered e.g. my-service-2.
There is no impact on the blocking query for my-service-1 since update index maintained for each resource separately.

## Scenario 3
### The service `my-service-1` gets registered in consul.
This case as-well, start with index 0, it will return the latest index and then you can use that index for next blocking query.
```
2025/05/24 14:40:23 Starting blocking query watcher...
2025/05/24 14:40:23 [Query] Waiting for service updates, index 0...
2025/05/24 14:40:23 [Query] Waiting for service updates, index 20...
[Query] Got 1 instances at index 20
[Query] Got 1 instances at index 21
2025/05/24 14:40:27 [Query] Waiting for service updates, index 21...
```


## Scenario 4
### Q: When service is aLready existing in consul, How do I start the blocking query without knowing the start index?
Always start the blocking query with index 0, it will return the latest index and then you can use that index for next blocking query.

```
2025/05/24 14:23:31 Starting blocking query watcher...
2025/05/24 14:23:31 [Query] Waiting for service updates, index 0...
2025/05/24 14:23:31 [Query] Waiting for service updates, index 59...
[Query] Got 1 instances at index 59
[Query] Got 1 instances at index 116
2025/05/24 14:30:13 [Query] Waiting for service updates, index 116...
[Query] Got 1 instances at index 122
2025/05/24 14:31:33 [Query] Waiting for service updates, index 122...
[Query] Got 1 instances at index 135
2025/05/24 14:34:17 [Query] Waiting for service updates, index 135...
[Query] Got 1 instances at index 137
2025/05/24 14:34:27 [Query] Waiting for service updates, index 137...
2025/05/24 14:34:37 [Query] Waiting for service updates, index 138...
[Query] Got 1 instances at index 138
```

## Scenario 5
If you start at index bigger than the latest index, it will be blocked until the requested resource reaches the index.

```
2025/05/24 16:31:30 Starting blocking query watcher...
2025/05/24 16:31:34 [Query] Waiting for service updates, index 200...
```

## Cleanup
Run below command to stop the demo and remove all containers and volumes.
```bash
docker compose down -v
```