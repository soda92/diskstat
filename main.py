from PySide6 import QtWidgets
import win32api
import shutil
from disks import get_all_disks


def get_disk_name(path):
    return win32api.GetVolumeInformation(f"{path[0]}:\\")[0]


def get_usage(path):
    stat = shutil.disk_usage(path)
    return stat.free / 1024**3, stat.used / 1024**3, stat.total / 1024**3


class Disk(QtWidgets.QWidget):
    def __init__(self, path):
        super().__init__()
        self._layout = QtWidgets.QVBoxLayout()
        label = get_disk_name(path)
        self.label = QtWidgets.QLabel(f"{label} ({path[0]}:)")

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

        if used_percent>90:
            self.bar.setStyleSheet(
                self.bar.styleSheet().replace("#05B8CC", "#930006")
            )

        label2 = f"{free:.1f} GB free of {int(total)} GB"
        self.info = QtWidgets.QLabel(label2)

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
        for d in get_all_disks():
            self.disks.append(Disk(d))

        for d in self.disks:
            self._layout.addWidget(d)

        self.widget.setLayout(self._layout)
        self.setCentralWidget(self.widget)
        self.resize(585, 444)

    def resizeEvent(self, event):
        print(event.size())
        super().resizeEvent(event)


if __name__ == "__main__":
    app = QtWidgets.QApplication()
    window = MainWindow()
    window.show()
    app.exec()
