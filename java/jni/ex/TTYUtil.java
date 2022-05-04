package ex;
public class TTYUtil {
    static { System.loadLibrary("ttyutil"); }
    public static native boolean isTTY();
    public static native String getTTYName();
}
