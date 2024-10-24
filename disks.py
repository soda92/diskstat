import shutil
from pathlib import Path

disk = 'A'
disks = []
for i in range(26):
    disk = chr(ord('A') + i)
    path = disk + ":/"
    if Path(path).exists():
        disks.append(path)

# shutil.disk_usage
print(disks)