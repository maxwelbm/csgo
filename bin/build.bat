go version

REM configure csgo.exe dist is 386
set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
set PATH=%PATH%;C:\MinGW\bin

REM build .\csboost.exe
set GOARCH=amd64

for /f %%a in ('powershell -Command "git rev-parse --short HEAD"') do set VERSION=%%a
for /f %%a in ('powershell -Command "Get-Date -format yyyyMMdd.HHmmss"') do set DATE=%%a

go build -v -ldflags "-X github.com/maxwelbm/csboost.Version=%VERSION% -X github.com/maxwelbm/csboost.Date=%DATE%" cmd/csboost/main.go

if %errorlevel% neq 0 exit /b %errorlevel%
