#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

SHOW_COMMAND := "{resources}\diskstat-api.exe -show"
SERVER := "diskstat-server.exe --hidden"
STOP_COMMAND := "{resources}\diskstat-api.exe -stop"

; shortcut behavior
#!O::
{
    Run SHOW_COMMAND, , "Hide"
}

; for thinkpad keyboard
#HotIf GetKeyState("Alt")
PrintScreen & O::
{
    Run SHOW_COMMAND, , "Hide"
}

; launch behavior
Run SERVER, , "Hide"

; exit behavior
Persistent

ExitPreviousSession(ExitReason, ExitCode) {
    id := 0
    Run STOP_COMMAND, , "Hide", &id
    ProcessWaitClose(id)
}

OnExit ExitPreviousSession