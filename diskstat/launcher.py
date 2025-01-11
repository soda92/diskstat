from diskstat.gen import gen
import diskstat.autostart as autostart
import argparse


def main():
    gen()
    parser = argparse.ArgumentParser()

    parser.add_argument(
        "-e", "--enable", action="store_true", default=False, help="enable auto start"
    )
    parser.add_argument(
        "-d", "--disable", action="store_true", default=False, help="disable auto start"
    )
    parser.add_argument(
        "-o", "--open", action="store_true", default=False, help="open startup folder"
    )

    parser.add_argument(
        "-x", "--exec", action="store_true", default=True, help="execute now"
    )

    args = parser.parse_args()

    if args.enable:
        autostart.enable()
    if args.disable:
        autostart.disable()

    if args.open:
        autostart.open_start_folder()

    if args.exec:
        autostart.exec()


if __name__ == "__main__":
    main()
