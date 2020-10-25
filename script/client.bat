@echo off
setlocal
cd pkg/client
if "%1" equ "build" (
    if "%2" equ "" (
        for /d %%d in (*) do call :build %%d
    ) else call :build %2
)
exit /b %ERRORLEVEL%
endlocal

:build
  setlocal
    cd %1
    call webpack
  endlocal
exit /b 0
