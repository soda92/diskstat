import subprocess
from pathlib import Path

CURRENT = Path(__file__).resolve().parent


def main():
    subprocess.run([CURRENT.joinpath("diskstat.exe"), "-hide"], check=True)


if __name__ == "__main__":
    main()
