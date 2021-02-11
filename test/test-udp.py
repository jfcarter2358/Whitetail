import socket
import datetime
import json
import sys

UDP_IP = "127.0.0.1"
UDP_PORT = 9003

timestamp = datetime.datetime.now().strftime('%Y-%m-%dT%H:%M:%S.%f')
message = sys.argv[2]
service = sys.argv[1]
logger_name = 'test'
stack_trace = ''

data = {
    "@timestamp":"2021-02-01T16:30:54.43112Z",
    "fields": {
        "severity":"info",
        "application":"engine",
        "hostname":"engine-1",
        "application_name":"<undefined>",
        "module":"external_file_asset",
        "function":"fetch",
        "line":12,
        "pid":"<0.247.0>",
        "node":"nonode@nohost"
    },
    "message":"hello world"
}

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.sendto(json.dumps(data).encode('utf-8'), (UDP_IP, UDP_PORT))