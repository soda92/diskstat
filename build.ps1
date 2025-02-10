Push-Location .\diskstat
go build diskstat-api.go
pyside6-rcc res.qrc -o res.py
Pop-Location