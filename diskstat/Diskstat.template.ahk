#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

#!O::
{
    Run "cmd /C diskstat-launcher.exe", , "Hide"
}
