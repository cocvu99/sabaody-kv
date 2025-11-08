# cpu_logger.py
import psutil
import time
from datetime import datetime

log_file = "/tmp/cpu_usage.log"

print(f"[{datetime.now()}] Bắt đầu ghi log CPU... (ấn Ctrl+C để dừng)")

while True:
    cpu = psutil.cpu_percent(interval=1)
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    with open(log_file, "a") as f:
        f.write(f"{timestamp} - CPU usage: {cpu}%\n")

    time.sleep(4)  # mỗi 5 giây (1s từ psutil + 4s sleep)
