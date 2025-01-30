#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

; shortcut behavior
#!O::
{
    Run "{resources}\diskstat-api.exe -show", , "Hide"
}

; for thinkpad keyboard
#HotIf GetKeyState("Alt")
PrintScreen & O::
{
    Run "{resources}\diskstat-api.exe -show", , "Hide"
}

; launch behavior
Run "diskstat-server.exe --hidden", , "Hide"

; exit behavior
Persistent

ExitPreviousSession(ExitReason, ExitCode) {
    id := 0
    Run "{resources}\diskstat-api.exe -stop", , "Hide", &id
    ProcessWaitClose(id)
}

OnExit ExitPreviousSession