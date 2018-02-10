@echo off
color 0a
title Compile EasyDoc

set goName=easydoc.go
set execName=easydoc
set version=2.0.5

::Cross compile is not support CGO
set CGO_ENABLED=0

@echo Compile windows 32 bit...
set GOOS=windows
set GOARCH=386
set GOEXE=.exe
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile windows 64 bit...
set GOOS=windows
set GOARCH=amd64
set GOEXE=.exe
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Linux 32 bit...
set GOOS=linux
set GOARCH=386
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Linux 64 bit...
set GOOS=linux
set GOARCH=amd64
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Mac 32 bit...
set GOOS=darwin
set GOARCH=386
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.

@echo Compile Mac 64 bit...
set GOOS=darwin
set GOARCH=amd64
set GOEXE=
mkdir %execName%-%version%-%GOOS%-%GOARCH%
go build -ldflags="-s -w" -o %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE% %goName%
upx %execName%-%version%-%GOOS%-%GOARCH%\%execName%%GOEXE%
::rar a %execName%-%version%-%GOOS%-%GOARCH%.zip -r %execName%-%version%-%GOOS%-%GOARCH%
@echo.
