# diskstat

<a href="https://pypi.org/project/diskstat/">
    <img alt="PyPI - Version" src="https://img.shields.io/pypi/v/diskstat">
</a>

Visualize your disk free space without Windows Explorer.

## install
```bash
pip install -U diskstat
```

## Options
```bash
py -m diskstat --help
usage: __main__.py [-h] [-e] [-d] [-c] [-o]

options:
  -h, --help     show this help message and exit
  -e, --enable   enable auto start
  -d, --disable  disable auto start
  -c, --console  console mode
  -o, --open     open startup folder
```

for example, if you want to start this with Windows:
```bash
py -m diskstat -ceo
```
Then click the shortcut in the popped folder to run it.

## Screenshots
<img src="https://raw.githubusercontent.com/soda92/diskstat/refs/heads/main/image.png" alt="demo" style="width:400px;"/>

## Features

Red bar when free space is lower than 10% (Same as Windows Explorer).

System tray for easy reopen.

Use `Ctrl + R` to refresh (reopen will also refresh automatically).
