@echo off
go-bindata assets/...
go install -ldflags "-H windowsgui"
exit