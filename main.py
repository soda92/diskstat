from PySide6 import QtWidgets
import win32api


def get_disk_name(path):
    return win32api.GetVolumeInformation(f"{path}:\\")[0]


class Disk(QtWidgets.QWidget):
    def __init__(self, path):
        super().__init__()
        self._layout = QtWidgets.QVBoxLayout()
        label = get_disk_name(path)
        self.label = QtWidgets.QLabel(label)

        self.bar = QtWidgets.QProgressBar()
        self.bar.setValue(50)

        self._layout.addWidget(self.label)
        self._layout.addWidget(self.bar)

        self.setLayout(self._layout)


class MainWindow(QtWidgets.QMainWindow):
    def __init__(self):
        super().__init__()

        self.widget = QtWidgets.QWidget()
        self._layout = QtWidgets.QVBoxLayout()
        self.disks = []
        self.disks.append(Disk("E"))

        for d in self.disks:
            self._layout.addWidget(d)
        
        self.widget.setLayout(self._layout)
        self.setCentralWidget(self.widget)


if __name__ == "__main__":
    app = QtWidgets.QApplication()
    window = MainWindow()
    window.show()
    app.exec()
