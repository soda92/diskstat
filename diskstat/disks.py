import shutil
from pathlib import Path
import colorama

from colorama import init

init(autoreset=True)


def get_all_disks():
    disk = "A"
    disks = []
    for i in range(26):
        disk = chr(ord("A") + i)
        path = disk + ":/"
        if Path(path).exists():
            disks.append(path)
    return disks


def main_console():
    disks = get_all_disks()
    for d in disks:
        usage = shutil.disk_usage(d)
        print(d, end=" ")
        percent = usage.used / usage.total * 10
        if percent > 9:
            print(colorama.Fore.RED)
        percent = int(percent)
        print("[" + "=" * percent + " " * (10 - percent) + "]", end=" ")

        free_gigabytes = usage.free / 1024 / 1024 / 1024
        percent = usage.free / usage.total * 100
        print(colorama.Fore.GREEN + f"free: {free_gigabytes:.1f} G ({percent:.1f}%)")


if __name__ == "__main__":
    main_console()
