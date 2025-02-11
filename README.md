# diskstat

<a href="https://pypi.org/project/diskstat/">
    <img alt="PyPI - Version" src="https://img.shields.io/pypi/v/diskstat">
</a>

Visualize your disk free space without Windows Explorer.

## install
```bash
pip install -U diskstat
```

Then install [autohotkey](https://www.autohotkey.com/).

## Options
```bash
> diskstat.exe --help
usage: diskstat [-h] [-d] [-o]

options:
  -h, --help     show this help message and exit
  -d, --disable  disable auto start
  -o, --open     open startup folder
```

for example, if you want to start this with Windows:
```pwsh
diskstat.exe
```

to disable:
```pwsh
diskstat.exe -d
```

to launch via command line:
```pwsh
diskstat-console.exe
```

## API for testing

show: <http://127.0.0.1:12347/show>

stop: <http://127.0.0.1:12347/stop>, only useful when debugging

## Screenshots
<img src="https://raw.githubusercontent.com/soda92/diskstat/refs/heads/main/image.png" alt="demo" style="width:400px;"/>

## Features

Use 'Win+Alt+O' for shortcut key.


## Note on running from source

need to install [golang](https://go.dev/); then run `hatch_build.py`.
