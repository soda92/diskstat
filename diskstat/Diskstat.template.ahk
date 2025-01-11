#Requires AutoHotKey v2.0
#SingleInstance Force

TraySetIcon "{resources}\program.ico"

#!O::
{
    Run "diskstat-launcher.exe", , "Hide"
}
