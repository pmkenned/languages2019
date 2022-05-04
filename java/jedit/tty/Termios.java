package tty;

public class Termios {

    public Termios() {
        IFlag = 0;
        OFlag = 0;
        CFlag = 0;
        LFlag = 0;
        CC = new int[NCCS];
    }

    static public final int NCCS = 32;
    static public final int ECHO = 8;

    static public final int TCSAFLUSH = 2;

    public int   IFlag;  /* input modes */
    public int   OFlag;  /* output modes */
    public int   CFlag;  /* control modes */
    public int   LFlag;  /* local modes */
    public int[] CC;     /* special characters */
}
