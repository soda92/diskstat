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
> diskstat.exe --help
usage: diskstat [-h] [-e] [-d] [-o] [-x]

options:
  -h, --help     show this help message and exit
  -e, --enable   enable auto start
  -d, --disable  disable auto start
  -o, --open     open startup folder
  -x, --exec     execute now
```

for example, if you want to start this with Windows:
```bash
diskstat.exe -e
```

## Screenshots
<img src="https://raw.githubusercontent.com/soda92/diskstat/refs/heads/main/image.png" alt="demo" style="width:400px;"/>

## Features

Red bar when free space is lower than 10% (Same as Windows Explorer).

Use 'Win+Alt+O' for shortcut key.
