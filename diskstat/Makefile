all: build build-api


build:
	fyne package
	mv "Display Disk Usage Status.exe" diskstat.exe

build-api:
	cd ../diskstat_api && make

clean:
	pwsh -nop -c "rm *.exe"
	pwsh -nop -c "rm *.exe~"
	pwsh -nop -c "rm *.syso"
	pwsh -nop -c "rm -r __pycache__"