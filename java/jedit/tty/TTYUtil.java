package tty;

// TODO: finish tcgetattr and tcsetattr so that they can be used this way:
//  Termios raw = new Termios();
//  TTYUtil.tcgetattr(STDIN_FILENO, raw);
//  raw.LFlag &= ~(Termios.ECHO);
//  TTYUtil.tcsetattr(STDIN_FILENO, Termios.TCSAFLUSH, raw);

public class TTYUtil {
    static { System.loadLibrary("ttyutil"); }
    public static native boolean isTTY();
    public static native String getTTYName();
    public static native int tcgetattr(int fd, Termios termios);
    public static native int tcsetattr(int fd, int optional_actions, Termios termios);
    public static native void disableRawMode();
    public static native void enableRawMode();
}
