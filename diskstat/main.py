from PySide6 import QtWidgets, QtGui
import win32api
import shutil
from diskstat.disks import get_all_disks, main_console
import sys
import diskstat.res as _a  # noqa: F401
import argparse
import diskstat.autostart as autostart


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

    def closeEvent(self, _event):
        self.hide()


class SystemTrayApp(QtWidgets.QApplication):
    def __init__(self, argv):
        super().__init__(argv)

        self.tray_icon = QtWidgets.QSystemTrayIcon(self)
        self.tray_icon.setIcon(
            QtGui.QIcon(":/program.ico")
        )  # Replace with your icon path
        self.tray_icon.setToolTip("Disk Usage")

        # Create the menu
        menu = QtWidgets.QMenu()
        action_exit = QtGui.QAction("Exit", self)
        action_exit.triggered.connect(self.quit)
        menu.addAction(action_exit)

        self.window = MainWindow()

        # action_show = QtGui.QAction("show", self)
        # action_show.triggered.connect(self.window.show)

        # menu.addAction(action_show)
        self.tray_icon.activated.connect(self.window.refresh)

        self.tray_icon.setContextMenu(menu)
        self.tray_icon.show()

        self.window.show()
        self.setQuitOnLastWindowClosed(False)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-e", "--enable", action="store_true", default=False, help="enable auto start"
    )
    parser.add_argument(
        "-d", "--disable", action="store_true", default=False, help="disable auto start"
    )
    parser.add_argument(
        "-c", "--console", action="store_true", default=False, help="console mode"
    )
    parser.add_argument(
        "-o", "--open", action="store_true", default=False, help="open startup folder"
    )

    args = parser.parse_args()
    if args.enable:
        autostart.enable()
    if args.disable:
        autostart.disable()

    if args.open:
        autostart.open_start_folder()

    if args.console:
        main_console()
    else:
        app = SystemTrayApp(sys.argv)
        app.exec()


if __name__ == "__main__":
    main()
