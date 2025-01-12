from diskstat.disks import main_console
import argparse
from diskstat.program_window import SystemTrayApp
import sys

def main():
    parser = argparse.ArgumentParser()

    parser.add_argument(
        "-c", "--console", action="store_true", default=False, help="console mode"
    )

    args = parser.parse_args()

    if args.console:
        main_console()
    else:
        app = SystemTrayApp(sys.argv)
        app.exec()


if __name__ == "__main__":
    main()
