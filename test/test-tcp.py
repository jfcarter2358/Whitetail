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
message = sys.argv[2]
service = sys.argv[1]
logger_name = 'test.logger.TestLogger'
stack_trace = ''

count = 300

host = "localhost"
port = 9002
print('Sending WARN messages')
for i in range(0, count):
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
    s.close()
print('Done!')
print('Sending ERROR messages')
for i in range(0, count):
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
    s.close()
print('Done!')
print('Sending TRACE messages')
for i in range(0, count):
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
    s.close()
print('Done!')
print('Sending INFO messages')
for i in range(0, count):
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
    s.close()
print('Done!')
print('Sending DEBUG messages')
for i in range(0, count):
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
    s.close()
print('Done!')
