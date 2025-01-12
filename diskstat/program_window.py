from PySide6 import QtWidgets, QtGui, QtCore
import win32api
import shutil
from diskstat.disks import get_all_disks
import diskstat.res as _a  # noqa: F401


def get_disk_name(path):
    return win32api.GetVolumeInformation(f"{path[0]}:\\")[0]


def get_usage(path):
    stat = shutil.disk_usage(path)
    return stat.free / 1024**3, stat.used / 1024**3, stat.total / 1024**3


class Disk(QtWidgets.QWidget):
    def __init__(self, path):
        super().__init__()
        self._layout = QtWidgets.QVBoxLayout()

        label_style = """
font: "Noto Sans SC" 12pt;
"""
        label = get_disk_name(path)
        self.label = QtWidgets.QLabel(f"{label} ({path[0]}:)")
        self.label.setStyleSheet(label_style)

        self.bar = QtWidgets.QProgressBar()
        self.bar.setTextVisible(False)
        self.bar.setStyleSheet("""
QProgressBar {
    /* Styles for the entire progress bar */
    border: 2px solid grey;
    border-radius: 5px;
    text-align: center;
}

QProgressBar::chunk {
    /* Styles for the filled part of the progress bar */
    background-color: #05B8CC;
    width: 3px; /* Fixed width for chunks */
    margin: 0px; /* Space between chunks */
}
        """)

        free, used, total = get_usage(path)
        used_percent = used / total * 100
        self.bar.setValue(int(used_percent))

        if used_percent > 90:
            self.bar.setStyleSheet(self.bar.styleSheet().replace("#05B8CC", "#930006"))

        label2 = f"{free:.1f} GB free of {int(total)} GB"
        self.info = QtWidgets.QLabel(label2)
        self.info.setStyleSheet(label_style)

        self._layout.addWidget(self.label)
        self._layout.addWidget(self.bar)
        self._layout.addWidget(self.info)

        self.setLayout(self._layout)


class MainWindow(QtWidgets.QMainWindow):
    def __init__(self):
        super().__init__()
        self.setWindowTitle("Disk Usage")

        self.widget = QtWidgets.QWidget()
        self._layout = QtWidgets.QVBoxLayout()
        self.disks = []

        self.setWindowIcon(QtGui.QIcon(":/program.ico"))
        for d in get_all_disks():
            self.disks.append(Disk(d))

        for d in self.disks:
            self._layout.addWidget(d)

        self.widget.setLayout(self._layout)
        self.setCentralWidget(self.widget)
        self.resize(585, 444)

        shortcut = QtGui.QShortcut(QtGui.QKeySequence("Ctrl+R"), self)
        shortcut.activated.connect(self.refresh)
        shortcut2 = QtGui.QShortcut(QtGui.QKeySequence("Esc"), self)
        shortcut2.activated.connect(self.hide)

    def refresh(self):
        # self.label.setText("Shortcut activated!")

        for d in self.disks:
            self._layout.removeWidget(d)
            d.setParent(None)
            del d

        self.disks = []
        for d in get_all_disks():
            self.disks.append(Disk(d))

        for d in self.disks:
            self._layout.addWidget(d)
        self.resize(585, 444)
        self.show()

    def resizeEvent(self, event):
        # print(event.size())
        super().resizeEvent(event)


class QueueHandler(QtCore.QThread):
    signal_show: QtCore.Signal = QtCore.Signal()
    signal_hide: QtCore.Signal = QtCore.Signal()
    signal_quit: QtCore.Signal = QtCore.Signal()

    def __init__(self, queue_app):
        super().__init__()
        self.queue_app = queue_app

    def run(self):
        while True:
            signal = self.queue_app.get()
            if signal == "show":
                self.signal_show.emit()
            elif signal == "hide":
                self.signal_hide.emit()
            elif signal == "quit":
                self.signal_quit.emit()
                break


class App(QtWidgets.QApplication):
    def __init__(self, argv, queue_app):
        super().__init__(argv)

        self.window = MainWindow()
        self.setQuitOnLastWindowClosed(False)

        self.queue_handler = QueueHandler(queue_app=queue_app)

        def show():
            self.window.refresh()
            self.window.show()
        self.queue_handler.signal_show.connect(show)
        self.queue_handler.signal_hide.connect(lambda: self.window.hide())
        self.queue_handler.signal_quit.connect(lambda: self.quit())
        self.queue_handler.start()
