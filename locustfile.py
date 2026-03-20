import time
import socket
from locust import User, task, between, events

class TCPUser(User):
    wait_time = between(0.1, 1)

    @task
    def send_ping(self):
        start_time = time.time()
        try:
            # Init a TCP socket.
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

            # TODO: Edit to IP EC2/Server Public
            sock.connect(("", 3000))

            # Send PING
            sock.sendall(b"ping\n")

            # Wait the data from the server send back
            data = sock.recv(1024)
            sock.close()

            total_time = int((time.time() - start_time) * 1000)

            # Report on Locust Dashboard
            if data:
                events.request.fire(request_type="TCP", name="Ping", respone_time=total_time, request_length=len(data))
            else:
                events.request.fire(request_type="TCP", name="Ping", respone_time=total_time, exception=Exception("No data received"))

        except Exception as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request.fire(request_type="TCP", name="Ping", response_time=total_time, exception=e)