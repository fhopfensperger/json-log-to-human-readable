# json-log-to-human-readable

Convert JSON log messages from Quarkus JSON Logging (https://quarkus.io/guides/logging#json-logging) or from Spring Boot JSON logs to human readable output.

## Installation

### Option 1 (script)

```bash
curl https://raw.githubusercontent.com/fhopfensperger/json-log-to-human-readable/master/get.sh | bash
```

### Option 2 (manually)

Either go to [Releases](https://github.com/fhopfensperger/json-log-to-human-readable/releases) download the latest release according to your processor architecture and operating system, then unarchive and copy it to the right location
```bash
tar xvfz json-log-to-human-readable_x.x.x_darwin_amd64.tar.gz
cd json-log-to-human-readable_x.x.x_darwin_amd64
chmod +x json-log-to-human-readable
sudo mv json-log-to-human-readable /usr/local/bin/
```

## Usage Examples:

##### **`test.json`**
```json 
{ "level": "INFO", "timestamp": "2020-07-14T09:38:14.977Z", "message": "sample output", "loggerName": "org.acme.MyClass" }
```
```bash
cat test.json | json-log-to-human-readable
```
##### **`Output`**
```
INFO 2020-07-14T09:38:14.977Z    org.acme.MyClass       sample output
```

### This also works for Pods running in Kubernetes: 
```bash
kubectl logs -f pod1 | json-log-to-human-readable
```
### Alternative JSON Logging format could be transformed with `-a`
##### **`test-alternative.json`**
```json 
{"@timestamp":"2020-07-15T19:09:39.983Z","@version":"1","message":"My log message","logger_name":"org.acme.MyClass","thread_name":"pool-1-thread-1","level":"INFO","level_value":20000}
```
```bash
cat test.json |  json-log-to-human-readable -a
```
##### **`Output`**
```
INFO 2020-07-15T19:09:39.983Z    org.acme.MyClass       My log message
```