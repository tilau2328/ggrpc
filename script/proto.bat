@echo off
setlocal
cd proto
if "%1" equ "gen" (
    if "%2" equ "" (
        for /d %%d in (*) do call :gen %%d
    ) else call :gen %2
)
exit /b %ERRORLEVEL%
endlocal

:gen
set CLIENT_OUT=../web/service
set SERVICE_OUT=../pkg/service
set GO_OUT=--go_out=%SERVICE_OUT%
set GO_OPT=--go_opt=paths=source_relative
set GRPC_OUT=--go-grpc_out=%SERVICE_OUT%
set GRPC_OPT=--go-grpc_opt=paths=source_relative
set JS_OUT=--js_out=import_style=commonjs:%CLIENT_OUT%
set WEB_OUT=--grpc-web_out=import_style=typescript,mode=grpcweb:%CLIENT_OUT%
set FILES=%1/src/domain/*.proto %1/src/command/*.proto %1/src/query/*.proto %1/src/event/*.proto
protoc -I=. %JS_OUT% %WEB_OUT% %GO_OUT% %GO_OPT% %GRPC_OUT% %GRPC_OPT% %FILES%
exit /b 0
