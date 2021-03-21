# json-log-to-human-readable
![Go](https://github.com/fhopfensperger/json-log-to-human-readable/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/fhopfensperger/json-log-to-human-readable)](https://goreportcard.com/report/github.com/fhopfensperger/json-log-to-human-readable)
[![Coverage Status](https://coveralls.io/repos/github/fhopfensperger/json-log-to-human-readable/badge.svg?branch=master)](https://coveralls.io/github/fhopfensperger/json-log-to-human-readable?branch=master)
[![Release](https://img.shields.io/github/release/fhopfensperger/json-log-to-human-readable?style=flat-square)](https://github.com//fhopfensperger/json-log-to-human-readable/releases/latest)


Convert JSON log messages to a human-readable format.

The following formats are supported:

- [Quarkus JSON Logging](https://quarkus.io/guides/logging#json-logging)
- [Spring Boot JSON Logging](https://www.baeldung.com/java-log-json-output)
- [Uber Zap](https://github.com/uber-go/zap)
- [.NET Core](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/logging/?view=aspnetcore-5.0)

```
Flags:
  -d, --dotnet       .NET JSON input
  -h, --help         help for json-log-to-human-readable
  -s, --springboot   Spring Boot JSON input
  -v, --version      version for json-log-to-human-readable
  -z, --zap          Uber zap JSON Input

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
### Spring Boot JSON Logging format could be transformed with `-s`
##### **`test-spring-boot.json`**
```json 
{"@timestamp":"2020-07-15T19:09:39.983Z","@version":"1","message":"My log message","logger_name":"org.acme.MyClass","thread_name":"pool-1-thread-1","level":"INFO","level_value":20000}
```
```bash
cat test-spring-boot.json |  json-log-to-human-readable -s
```
##### **`Output`**
```
INFO 2020-07-15T19:09:39.983Z    org.acme.MyClass       My log message
```

### Uber Zap JSON Logging format could be transformed with `-z`
##### **`test-uber-zap.json`**
```json 
{"level":"error","ts":1598445905.143377,"logger":"controller-runtime.controller","msg":"Reconciler error","controller":"scaledobject-controller","request":"default/azure-servicebus-queue-scaledobject","error":"error getting scaler for trigger #0: error parsing azure service bus metadata: no connection setting given","stacktrace":"github.com/go-logr/zapr.(*zapLogger).Error\n\t/Users/zroubali/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler\n\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:218\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem\n\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:192\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).worker\n\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:171\nk8s.io/apimachinery/pkg/util/wait.JitterUntil.func1\n\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:152\nk8s.io/apimachinery/pkg/util/wait.JitterUntil\n\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:153\nk8s.io/apimachinery/pkg/util/wait.Until\n\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:88"}
```
```bash
cat test-uber-zap.json |  json-log-to-human-readable -z
```
##### **`Output`**
```
error 2020-08-26 14:45:05.143377065 +0200 CEST   controller-runtime.controller  msg: Reconciler error   controller: scaledobject-controller      request: ugsvt-mercedes/azure-servicebus-queue-scaledobject
error: error getting scaler for trigger #0: error parsing azure service bus metadata: no connection setting givenstacktrace: github.com/go-logr/zapr.(*zapLogger).Error
        /Users/zroubali/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler
        /Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:218
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem
        /Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:192
sigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).worker
        /Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:171
k8s.io/apimachinery/pkg/util/wait.JitterUntil.func1
        /Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:152
k8s.io/apimachinery/pkg/util/wait.JitterUntil
        /Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:153
k8s.io/apimachinery/pkg/util/wait.Until
        /Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:88
```

## Installation

### Option 1 (script)

```bash
curl https://raw.githubusercontent.com/fhopfensperger/json-log-to-human-readable/master/get.sh | bash
```

### Option 2 (manually)

Go to [Releases](https://github.com/fhopfensperger/json-log-to-human-readable/releases) download the latest release according to your processor architecture and operating system, then unarchive and copy it to the right location

```bash
tar xvfz json-log-to-human-readable_x.x.x_darwin_amd64.tar.gz
cd json-log-to-human-readable_x.x.x_darwin_amd64
chmod +x json-log-to-human-readable
sudo mv json-log-to-human-readable /usr/local/bin/
```
