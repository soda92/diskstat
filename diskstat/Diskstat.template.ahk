#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

SHOW_COMMAND := "{resources}\diskstat-api.exe -show -port 12347"
SERVER := "diskstat-server.exe"
STOP_COMMAND := "{resources}\diskstat-api.exe -stop -port 12347"

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