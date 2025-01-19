from diskstat.program_window import App
import requests
import multiprocessing
import sys
import argparse


def run_app_impl(hidden, queue_app: multiprocessing.Queue):
    app = App(sys.argv, queue_app)
    if not hidden:
        app.window.show()
    app.exec()


def run_app_manager(hidden, queue_signals: multiprocessing.Queue):
    queue_app = multiprocessing.Queue()
    p = multiprocessing.Process(target=run_app_impl, args=(hidden, queue_app))
    p.start()

    while True:
        signal = queue_signals.get()
        if signal == "quit":
            queue_app.put("quit")
            break
        elif signal == "show":
            queue_app.put("show")
        elif signal == "hide":
            queue_app.put("hide")


def run_server_impl(
    queue_signals: multiprocessing.Queue, queue_server: multiprocessing.Queue
):
    from fastapi import FastAPI

    app = FastAPI()

    @app.get("/show")
    def show():
        queue_signals.put("show")
        return "success"

    @app.get("/hide")
    def hide():
        queue_signals.put("hide")
        return "success"

    @app.get("/quit")
    def quit():
        queue_signals.put("quit")
        queue_server.put("quit")
        return "success"

    import uvicorn

    uvicorn.run(app=app, port=12346)


def run_server_manager(
    queue_signals: multiprocessing.Queue, queue_server: multiprocessing.Queue
):
    p = multiprocessing.Process(
        target=run_server_impl, args=[queue_signals, queue_server]
    )
    p.start()

    while True:
        server_signal = queue_server.get()
        if server_signal == "quit":
            p.terminate()
            break


def run_app_and_server(
    hidden, queue_signals: multiprocessing.Queue, queue_server: multiprocessing.Queue
):
    p = multiprocessing.Process(target=run_app_manager, args=(hidden, queue_signals))
    p.start()

    p2 = multiprocessing.Process(
        target=run_server_manager, args=(queue_signals, queue_server)
    )
    p2.start()

    p2.join()


def try_run():
    parser = argparse.ArgumentParser()
    parser.add_argument("--hidden", action="store_true", default=False)
    args = parser.parse_args()
    hidden = args.hidden
    r = ""  # noqa: F841
    try:
        r = requests.get("http://127.0.0.1:12346/show", timeout=7)  # noqa: F841
    except Exception as e:
        print(e)
        queue_signals = multiprocessing.Queue()
        queue_server = multiprocessing.Queue()
        run_app_and_server(
            hidden=hidden, queue_signals=queue_signals, queue_server=queue_server
        )


if __name__ == "__main__":
    try_run()
