#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

; shortcut behavior
#!O::
{
    Run "diskstat-show.exe", , "Hide"
}

; launch behavior
Run "diskstat-server.exe --hidden", , "Hide"

; exit behavior
Persistent

ExitPreviousSession(ExitReason, ExitCode) {
    id := 0
    Run "diskstat-stop.exe", , "Hide", &id
    ProcessWaitClose(id)
}

OnExit ExitPreviousSession