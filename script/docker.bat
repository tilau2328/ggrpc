@echo off
setlocal
cd dkr/local
if "%1" equ "build" (
    if "%2" equ "" (
        for /d %%d in (*) do call docker-compose build
    ) else call docker-compose build -d %2
)
if "%1" equ "run" (
    if "%2" equ "" (
        for /d %%d in (*) do call docker-compose up -d
    ) else call docker-compose up -d %2
)
exit /b %ERRORLEVEL%
endlocal
