import subprocess
from win32com.client import Dispatch
from pathlib import Path
import sys

home_folder = Path.home()
CURRENT = Path(__file__).resolve().parent
python_path = Path(sys.executable).resolve().parent
start_folder = home_folder.joinpath(
    r"AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup"
)
lnk_file = start_folder.joinpath("diskstat - Shortcut.lnk")


def enable():
    shell = Dispatch("WScript.Shell")
    shortcut = shell.CreateShortCut(str(lnk_file))
    shortcut.Targetpath = str(CURRENT.parent.joinpath("diskstat_script").joinpath("Diskstat.ahk"))
    shortcut.Arguments = ""
    shortcut.save()


def disable():
    lnk_file.unlink(missing_ok=True)


def open_start_folder():
    subprocess.Popen(f"explorer {str(start_folder)}")

def exec():
    import os
    os.startfile(lnk_file)

if __name__ == "__main__":
    enable()

    subprocess.Popen(f"explorer /select,{str(lnk_file)}")
    import time

    time.sleep(5)
    disable()
