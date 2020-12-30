REM Equivalent to the "ridk enable" command, but without the exit
set TARGET_ARCH=x64
SET PATH=%PATH%;%GOPATH%/bin
if "%TARGET_ARCH%" == "x64" (
    @echo IN x64 BRANCH
    @for /f "delims=" %%x in ('"ruby" --disable-gems -x '%RIDK%' enable') do set "%%x"
)

if "%TARGET_ARCH%" == "x86" (
    @echo IN x86 BRANCH
    REM Use 64-bit toolchain to build gems
    Powershell -C "ridk enable; cd omnibus; bundle install"
)

pip install -r requirements.txt || exit /b 4
inv -e deps --verbose --dep-vendor-only --no-checks