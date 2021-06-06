package cmd

import (
	"bytes"
	"testing"
)

func TestGoZapLogMessage_transform(t *testing.T) {
	type fields struct {
		Level      string
		Timestamp  float64
		Logger     string
		Message    string
		Controller string
		Request    string
		Error      string
		Stacktrace string
	}

	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{
			name: "ok",
			fields: fields{
				Level:      "error",
				Timestamp:  1598445905.143377,
				Logger:     "controller-runtime.controller",
				Message:    "Reconciler error",
				Controller: "scaledobject-controller",
				Request:    "default/azure-servicebus-queue-scaledobject",
				Error:      "error getting scaler for trigger #0: error parsing azure service bus metadata: no connection setting given",
				Stacktrace: "github.com/go-logr/zapr.(*zapLogger).Error\\n\\t/Users/zroubali/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:218\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:192\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).worker\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:171\\nk8s.io/apimachinery/pkg/util/wait.JitterUntil.func1\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:152\\nk8s.io/apimachinery/pkg/util/wait.JitterUntil\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:153\\nk8s.io/apimachinery/pkg/util/wait.Until\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:88",
			},
			wantW: "error 2020-08-26 12:45:05.143377065 +0000 UTC\tcontroller-runtime.controller\tmsg: Reconciler error\tcontroller: scaledobject-controller\trequest: default/azure-servicebus-queue-scaledobject\nerror: error getting scaler for trigger #0: error parsing azure service bus metadata: no connection setting givenstacktrace: github.com/go-logr/zapr.(*zapLogger).Error\\n\\t/Users/zroubali/go/pkg/mod/github.com/go-logr/zapr@v0.1.1/zapr.go:128\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).reconcileHandler\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:218\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).processNextWorkItem\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:192\\nsigs.k8s.io/controller-runtime/pkg/internal/controller.(*Controller).worker\\n\\t/Users/zroubali/go/pkg/mod/sigs.k8s.io/controller-runtime@v0.2.2/pkg/internal/controller/controller.go:171\\nk8s.io/apimachinery/pkg/util/wait.JitterUntil.func1\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:152\\nk8s.io/apimachinery/pkg/util/wait.JitterUntil\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:153\\nk8s.io/apimachinery/pkg/util/wait.Until\\n\\t/Users/zroubali/go/pkg/mod/k8s.io/apimachinery@v0.0.0-20190404173353-6a84e37a896d/pkg/util/wait/wait.go:88\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			glm := &GoZapLogMessage{
				Level:      tt.fields.Level,
				Timestamp:  tt.fields.Timestamp,
				Logger:     tt.fields.Logger,
				Message:    tt.fields.Message,
				Controller: tt.fields.Controller,
				Request:    tt.fields.Request,
				Error:      tt.fields.Error,
				Stacktrace: tt.fields.Stacktrace,
			}
			w := &bytes.Buffer{}
			glm.transform(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("transform() = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}

func TestDotNetLogMessage_transform(t *testing.T) {
	type fields struct {
		Timestamp  string
		Level      string
		Message    string
		LoggerName string
	}

	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{
			name: "ok",
			fields: fields{
				Timestamp:  "2021-03-19T11:05:29.566Z",
				Level:      "Information",
				Message:    "GetUsers request received",
				LoggerName: "UserManagementSvc.Controllers.UserManagementController",
			},
			wantW: "Information 2021-03-19T11:05:29.566Z\tUserManagementSvc.Controllers.UserManagementController\tGetUsers request received\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dnlm := &DotNetLogMessage{
				Timestamp:  tt.fields.Timestamp,
				Level:      tt.fields.Level,
				Message:    tt.fields.Message,
				LoggerName: tt.fields.LoggerName,
			}
			w := &bytes.Buffer{}
			dnlm.transform(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("transform() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestSpringBootLogMessage_transform(t *testing.T) {
	type fields struct {
		Timestamp  string
		Level      string
		Message    string
		Exception  string
		LoggerName string
	}

	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{
			name: "ok",
			fields: fields{
				Timestamp: "2020-07-15T19:09:39.983Z",
				Level:     "INFO",
				Message:   "My log message",
				Exception: "java.lang.NullPointerException: null\\n\\tat com.daimler.ugsvt.mmenotificationmanager.service.TheftCaseNotificationService.getLatestUpdatedElement(TheftCaseNotificationService.java:204)\\n\\tat com.daimler.ugsvt.mmenotificationmanager.service.TheftCaseNotificationService.startNotificationProcess(TheftCaseNotificationService.java:39)\\n\\tat com.daimler.ugsvt.mmenotificationmanager.consumer.TheftCaseMessageConsumer.run(TheftCaseMessageConsumer.java:61)\\n\\tat java.base/java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:515)\\n\\tat java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)\\n\\tat java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)\\n\\tat java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)\\n\\tat java.base/java.lang.Thread.run(Thread.java:834)\\n", LoggerName: "",
			},
			wantW: "INFO 2020-07-15T19:09:39.983Z\t\tMy log message\nException: java.lang.NullPointerException: null\\n\\tat com.daimler.ugsvt.mmenotificationmanager.service.TheftCaseNotificationService.getLatestUpdatedElement(TheftCaseNotificationService.java:204)\\n\\tat com.daimler.ugsvt.mmenotificationmanager.service.TheftCaseNotificationService.startNotificationProcess(TheftCaseNotificationService.java:39)\\n\\tat com.daimler.ugsvt.mmenotificationmanager.consumer.TheftCaseMessageConsumer.run(TheftCaseMessageConsumer.java:61)\\n\\tat java.base/java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:515)\\n\\tat java.base/java.util.concurrent.FutureTask.run(FutureTask.java:264)\\n\\tat java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1128)\\n\\tat java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:628)\\n\\tat java.base/java.lang.Thread.run(Thread.java:834)\\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alm := &SpringBootLogMessage{
				Timestamp:  tt.fields.Timestamp,
				Level:      tt.fields.Level,
				Message:    tt.fields.Message,
				Exception:  tt.fields.Exception,
				LoggerName: tt.fields.LoggerName,
			}
			w := &bytes.Buffer{}
			alm.transform(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("transform() = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}

func TestException_transform(t *testing.T) {
	type fields struct {
		RefID         int
		ExceptionType string
		Message       string
		CausedBy      CausedBy
		Frames        *[]Frame
	}

	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{
			name: "ok",
			fields: fields{
				RefID:         0,
				ExceptionType: "javax.jms.JMSRuntimeException",
				Message:       "Unknown error from remote peer",
				CausedBy: CausedBy{&Exception{
					RefID:         1,
					ExceptionType: "javax.jms.RuntimeException",
					Message:       "Unknown error",
					CausedBy:      CausedBy{},
					Frames: &[]Frame{{
						Class:  "org.apache.qpid.jms.exceptions.ExceptionSupport",
						Method: "createRuntimeException",
						Line:   212,
					}},
				}},
				Frames: &[]Frame{{
					Class:  "org.apache.qpid.jms.exceptions.JmsExceptionSupport",
					Method: "createRuntimeException",
					Line:   211,
				}},
			},
			wantW: "Caused by: javax.jms.JMSRuntimeException. Unknown error from remote peer:\n\t at createRuntimeException(org.apache.qpid.jms.exceptions.JmsExceptionSupport:211)\nCaused by: javax.jms.RuntimeException. Unknown error:\n\t at createRuntimeException(org.apache.qpid.jms.exceptions.ExceptionSupport:212)\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ex := &Exception{
				RefID:         tt.fields.RefID,
				ExceptionType: tt.fields.ExceptionType,
				Message:       tt.fields.Message,
				CausedBy:      tt.fields.CausedBy,
				Frames:        tt.fields.Frames,
			}
			w := &bytes.Buffer{}
			ex.transform(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("transform() = \n%v\n, want \n%v", gotW, tt.wantW)
			}
		})
	}
}

func TestQuarkusLogMessage_transform(t *testing.T) {
	type fields struct {
		Timestamp  string
		Level      string
		Message    string
		Exception  Exception
		LoggerName string
		Tracing    Tracing
	}

	tests := []struct {
		name   string
		fields fields
		wantW  string
	}{
		{
			name: "With trace id",
			fields: fields{
				Timestamp: "2020-07-15T19:09:39.983Z",
				Level:     "INFO",
				Message:   "My log message",
				Exception: Exception{
					RefID:         1,
					ExceptionType: "javax.jms.RuntimeException",
					Message:       "Unknown error",
					CausedBy:      CausedBy{},
					Frames: &[]Frame{{
						Class:  "org.apache.qpid.jms.exceptions.ExceptionSupport",
						Method: "createRuntimeException",
						Line:   212,
					}},
				},
				LoggerName: "org.acme.MyClass",
				Tracing: Tracing{
					TraceID: "123",
					SpanID:  "12",
					Sampled: "1",
				},
			},
			wantW: "INFO 2020-07-15T19:09:39.983Z\ttraceId=123 org.acme.MyClass\tMy log message\nCaused by: javax.jms.RuntimeException. Unknown error:\n\t at createRuntimeException(org.apache.qpid.jms.exceptions.ExceptionSupport:212)\n",
		},
		{
			name: "Without trace id",
			fields: fields{
				Timestamp: "2020-07-15T19:09:39.983Z",
				Level:     "INFO",
				Message:   "My log message",
				Exception: Exception{
					RefID:         1,
					ExceptionType: "javax.jms.RuntimeException",
					Message:       "Unknown error",
					CausedBy:      CausedBy{},
					Frames: &[]Frame{{
						Class:  "org.apache.qpid.jms.exceptions.ExceptionSupport",
						Method: "createRuntimeException",
						Line:   212,
					}},
				},
				LoggerName: "org.acme.MyClass",
			},
			wantW: "INFO 2020-07-15T19:09:39.983Z\torg.acme.MyClass\tMy log message\nCaused by: javax.jms.RuntimeException. Unknown error:\n\t at createRuntimeException(org.apache.qpid.jms.exceptions.ExceptionSupport:212)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lm := &QuarkusLogMessage{
				Timestamp:  tt.fields.Timestamp,
				Level:      tt.fields.Level,
				Message:    tt.fields.Message,
				Exception:  tt.fields.Exception,
				LoggerName: tt.fields.LoggerName,
				Tracing:    tt.fields.Tracing,
			}
			w := &bytes.Buffer{}
			lm.transform(w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("transform() = \n%v\n, want \n%v", gotW, tt.wantW)
			}
		})
	}
}
