import subprocess
from win32com.client import Dispatch
from pathlib import Path
import sys

home_folder = Path.home()
python_path = Path(sys.executable).resolve().parent
start_folder = home_folder.joinpath(
    r"AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup"
)
lnk_file = start_folder.joinpath("diskstat.lnk")


def enable():
    shell = Dispatch("WScript.Shell")
    shortcut = shell.CreateShortCut(str(lnk_file))
    shortcut.Targetpath = str(python_path.joinpath("pythonw.exe"))
    shortcut.Arguments = "-m diskstat"
    shortcut.WorkingDirectory = str(python_path)
    shortcut.save()


def disable():
    lnk_file.unlink(missing_ok=True)


def open_start_folder():
    subprocess.Popen(f"explorer {str(start_folder)}")


if __name__ == "__main__":
    enable()

    subprocess.Popen(f"explorer /select,{str(lnk_file)}")
    import time

    time.sleep(5)
    disable()
