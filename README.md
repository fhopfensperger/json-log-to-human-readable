# json-log-to-human-readable

Convert JSON log messages from Quarkus JSON Logging (https://quarkus.io/guides/logging#json-logging) or from Spring Boot to human readable output.

##Usage Examples:
##### **`test.json`**
```json 
{ "level": "INFO", "timestamp": "2020-07-14T09:38:14.977Z", "message": "sample output", "loggerName": "org.acme.MyClass" }
```
```bash
cat test.json | go run main.go
```
##### **`Output`**
```
INFO 2020-07-14T09:38:14.977Z    org.acme.MyClass       sample output
```

###This works also for Pods running in Kubernetes: 
```bash
kubectl logs -f pod1 | json-log-to-human-readable
```
### Alternative JSON Logging format could be transform with `-a`
##### **`test-alternative.json`**
```json 
{"@timestamp":"2020-07-15T19:09:39.983Z","@version":"1","message":"My log message","logger_name":"org.acme.MyClass","thread_name":"pool-1-thread-1","level":"INFO","level_value":20000}
```
```bash
cat test.json | go run main.go -a
```
##### **`Output`**
```
INFO 2020-07-15T19:09:39.983Z    org.acme.MyClass       My log message
```