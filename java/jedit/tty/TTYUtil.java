package tty;

public class TTYUtil {
    static { System.loadLibrary("ttyutil"); }
    public static native boolean isTTY();
    public static native String getTTYName();
    public static native int tcgetattr(int fd, Termios termios);
    public static native int tcsetattr(int fd, int optional_actions, Termios termios);
    public static native void disableRawMode();
    public static native void enableRawMode();
}
