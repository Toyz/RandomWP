@echo off
bash -c 'git rev-parse HEAD > assets/version.txt'
go-bindata assets/...
go install -ldflags "-H windowsgui"
exit