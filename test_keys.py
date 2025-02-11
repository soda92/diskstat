import win32api  # noqa: F401
import win32gui
import win32con

def handle_hotkey():
    print("Hotkey pressed!")

# Register the hotkey Ctrl+Shift+H
hotkey_id = 1
win32gui.RegisterHotKey(None, hotkey_id, win32con.MOD_CONTROL | win32con.MOD_ALT, ord('w'))

# Keep the script running to listen for hotkey events
try:
    while True:
        msg = win32gui.GetMessage(None, 0, 0)
        if msg[0] == win32con.WM_HOTKEY:
            if msg[1] == hotkey_id:
                handle_hotkey()
finally:
    # Unregister the hotkey when the script exits
    win32gui.UnregisterHotKey(None, hotkey_id)