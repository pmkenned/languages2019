import java.io.IOException;
import tty.Termios;
import tty.TTYUtil;

public class JEdit {

    public static int ctrl_key(int k) { return k & 0x1f; }

    public static void main(String[] args) throws IOException {
        if (!TTYUtil.isTTY()) {
            System.err.println("JEdit: cannot run in non-interactive mode");
            return;
        }

        //Termios raw = new Termios();
        //TTYUtil.tcgetattr(STDIN_FILENO, raw);
        //System.out.println(raw.IFlag);
        //System.out.println(raw.OFlag);
        //System.out.println(raw.CFlag);
        //System.out.println(raw.LFlag);
        //raw.LFlag &= ~(Termios.ECHO);
        //TTYUtil.tcsetattr(STDIN_FILENO, Termios.TCSAFLUSH, raw);

        TTYUtil.enableRawMode();
        while (true) {
            int ci = System.in.read();
            char c = (char) ci;
            if (ci == -1) {
            //} else if (c == 'q') {
            //    break;
            } else {
                System.out.print(ci);
            }
            if (ci == ctrl_key('q')) {
                break;
            }
        }

        TTYUtil.disableRawMode();
    }

    static final int STDIN_FILENO = 0;
    static final int STDOUT_FILENO = 1;
    static final int STDERR_FILENO = 2;
}
