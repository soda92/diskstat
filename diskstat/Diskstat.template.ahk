#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

#!O::
{
    Run "diskstat-launcher.exe", , "Hide"
}

id := 0
id2 := 0

ExitPreviousSession(ExitReason, ExitCode) {
    Run "diskstat-stop.exe", , "Hide", &id
    ProcessWaitClose(id2)
}

ExitPreviousSession(0, 0)
Run "diskstat-launcher-hidden.exe", , "Hide", &id2

Persistent
OnExit ExitPreviousSession
