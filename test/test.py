import socket
import datetime
import sys
import json

'''
Timestamp  string `json:"@timestamp"`
Message    string `json:"message"`
Service    string `json:"appName"`
LoggerName string `json:"logger_name"`
Level      string `json:"level"`
StackTrace string `json:"stack_trace"`
'''
message = sys.argv[3]
service = sys.argv[1]
logger_name = 'test'
stack_trace = ''

host = "localhost"
port = 9002
for i in range(0, 30):
    level = 'WARN'
    timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
    data = {
        "@timestamp": timestamp,
        "message": message + " {}".format(i),
        "appName": service,
        "logger_name": logger_name,
        "level": level,
        "stack_trace": stack_trace
    }
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.sendall(json.dumps(data).encode('utf-8'))
    data = s.recv(1024)
    s.close()
    print('Received', repr(data))
for i in range(0, 30):
    level = 'ERROR'
    timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
    data = {
        "@timestamp": timestamp,
        "message": message + " {}".format(i),
        "appName": service,
        "logger_name": logger_name,
        "level": level,
        "stack_trace": stack_trace
    }
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.sendall(json.dumps(data).encode('utf-8'))
    data = s.recv(1024)
    s.close()
    print('Received', repr(data))
for i in range(0, 30):
    level = 'TRACE'
    timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
    data = {
        "@timestamp": timestamp,
        "message": message + " {}".format(i),
        "appName": service,
        "logger_name": logger_name,
        "level": level,
        "stack_trace": stack_trace
    }
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.sendall(json.dumps(data).encode('utf-8'))
    data = s.recv(1024)
    s.close()
    print('Received', repr(data))
for i in range(0, 30):
    level = 'INFO'
    timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
    data = {
        "@timestamp": timestamp,
        "message": message + " {}".format(i),
        "appName": service,
        "logger_name": logger_name,
        "level": level,
        "stack_trace": stack_trace
    }
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.sendall(json.dumps(data).encode('utf-8'))
    data = s.recv(1024)
    s.close()
    print('Received', repr(data))
for i in range(0, 30):
    level = 'DEBUG'
    timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
    data = {
        "@timestamp": timestamp,
        "message": message + " {}".format(i),
        "appName": service,
        "logger_name": logger_name,
        "level": level,
        "stack_trace": stack_trace
    }
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((host, port))
    s.sendall(json.dumps(data).encode('utf-8'))
    data = s.recv(1024)
    s.close()
    print('Received', repr(data))
