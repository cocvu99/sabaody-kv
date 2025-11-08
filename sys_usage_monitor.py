import os
import psutil
import time
import matplotlib.pyplot as plt
from datetime import datetime

# === CONFIG ===
log_file = "/tmp/sys_usage.log"
duration = 60  # Tổng thời gian ghi log (giây)
interval = 5   # Ghi mỗi 5 giây

# === STEP 1: GHI LOG ===
def log_system_usage():
    with open(log_file, "w") as f:
        f.write("timestamp,cpu,ram,disk\n")
        start = time.time()
        while time.time() - start < duration:
            cpu = psutil.cpu_percent(interval=1)
            ram = psutil.virtual_memory().percent
            disk = psutil.disk_usage("/").percent
            timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            f.write(f"{timestamp},{cpu},{ram},{disk}\n")
            time.sleep(interval - 1)

# === STEP 2: VẼ BIỂU ĐỒ ===
def plot_usage_graph():
    timestamps, cpus, rams, disks = [], [], [], []
    with open(log_file) as f:
        next(f)  # Bỏ header
        for line in f:
            ts, cpu, ram, disk = line.strip().split(",")
            timestamps.append(ts[-8:])  # Lấy giờ phút giây
            cpus.append(float(cpu))
            rams.append(float(ram))
            disks.append(float(disk))

    plt.figure(figsize=(10, 6))
    plt.plot(timestamps, cpus, label="CPU (%)")
    plt.plot(timestamps, rams, label="RAM (%)")
    plt.plot(timestamps, disks, label="Disk (%)")
    plt.xlabel("Time")
    plt.ylabel("Usage (%)")
    plt.title("System Resource Usage")
    plt.xticks(rotation=45)
    plt.legend()
    plt.tight_layout()
    plt.grid(True)
    plt.savefig("/tmp/sys_usage_chart.png")

# === STEP 3: DỌN LOG (tuỳ chọn) ===
def cleanup_log():
    if os.path.exists(log_file):
        os.remove(log_file)

# === THỰC THI ===
log_system_usage()
plot_usage_graph()
# cleanup_log()  # Bỏ comment nếu muốn dọn log ngay sau khi vẽ

"/tmp/sys_usage_chart.png"
