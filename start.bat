@echo off

echo "___________________________"
echo "Choose a mode:            |"
echo "1. Normal                 |"
echo "2. Hard                   |"
echo "3. Start with save        |"
echo "___________________________"
read "-p" "Enter your choice: " "choice"
IF "%choice%"=="1" (
  CALL :start_mode_normal
) ELSE IF "%choice%"=="2" (
  CALL :start_mode_hard
) ELSE IF "%choice%"=="3" (
  CALL :start_with_save
) ELSE (
  echo "Invalid choice"
)

EXIT /B %ERRORLEVEL%

:start_mode_normal
echo "Starting in normal mode"
go "run" "main.go" "asset\words.txt"
EXIT /B 0

:start_mode_hard
echo "Starting in hard mode"
go "run" "main.go" "--hard" "asset\words.txt"
EXIT /B 0

:start_with_save
echo "Starting with save"
go "run" "main.go" "--startWith" "save.txt"
EXIT /B 0