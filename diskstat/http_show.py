import requests


def main():
    try:
        r = requests.get("http://localhost:12346/show", timeout=0.5)
    except Exception as e:
        print("timeout", e)
    else:
        print(r.content)


if __name__ == "__main__":
    main()
