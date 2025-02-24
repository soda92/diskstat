all: build build-api


build:
	go-winres make
	go build .

build-api:
	cd ../diskstat_api && make

clean:
	pwsh -nop -c "rm *.exe"
	pwsh -nop -c "rm *.exe~"
	pwsh -nop -c "rm *.syso"
	pwsh -nop -c "rm -r __pycache__"