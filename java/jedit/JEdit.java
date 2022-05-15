import java.io.IOException;
import tty.Termios;
import tty.TTYUtil;

// TODO: Consider using Maven
// TODO: package jedit;

class Key
{
    Key() { }
    static final int K_A        = 0;
    static final int K_B        = 1;
    static final int K_F1       = 2;
    static final int K_F2       = 3;
    static final int K_F3       = 4;
    static final int K_F4       = 5;
    static final int K_UP       = 6;
    static final int K_DOWN     = 7;
    static final int K_LEFT     = 8;
    static final int K_RIGHT    = 9;
    static final int K_CTRL_Q   = 10;
}

public class JEdit {

    public static int ctrl_key(int k) { return k & 0x1f; }

    public static int getKey() throws IOException
    {
        int ci = System.in.read();
        char c = (char) ci;

        if (ci == -1) {
        } else if (ci == 033) {
            ci = System.in.read();
            c = (char) ci;
            if (c == '[') {
                ci = System.in.read();
                c = (char) ci;
                if (c == 'A') {
                    return Key.K_UP;
                } else if (c == 'B') {
                    return Key.K_DOWN;
                } else if (c == 'C') {
                    return Key.K_RIGHT;
                } else if (c == 'D') {
                    return Key.K_LEFT;
                } else {
                    //System.out.print(c);
                }
            } else if (ci == 79) {
                ci = System.in.read();
                c = (char) ci;
                if (c == 'P') {
                    return Key.K_F1;
                } else if (c == 'Q') {
                    return Key.K_F2;
                } else if (c == 'R') {
                    return Key.K_F3;
                } else if (c == 'S') {
                    return Key.K_F4;
                }
            } else {
                System.out.print("" + ci + " ");
            }
        } else if (ci == ctrl_key('q')) {
            return Key.K_CTRL_Q;
        } else if (ci <= 32 || ci >= 127) {
            System.out.print("" + ci + " ");
        } else {
            System.out.print("'" + c + "' ");
        }

        return -1;
    }

    public static void main(String[] args) throws IOException {
        if (!TTYUtil.isTTY()) {
            System.err.println("JEdit: cannot run in non-interactive mode");
            return;
        }

        int mode = 0;
        int row = 1;
        int col = 1;

        try {
            TTYUtil.enableRawMode();

            AnsiEsc.clear();
            AnsiEsc.move(999, 999);
            int[] pos = AnsiEsc.getCursor();
            int height = pos[0];
            int width = pos[1];
            AnsiEsc.move(1, 1);
            System.out.print("Ctrl-Q to quit");
            AnsiEsc.move(2, 1);

            boolean done = false;
            while (!done) {

                //AnsiEsc.move(row, col);

                int k = getKey();

                switch (k) {
                    case Key.K_UP:
                        if (row > 1) {
                            row--;
                        }
                        break;
                    case Key.K_DOWN:
                        if (row < height) {
                            row++;
                        }
                        break;
                    case Key.K_LEFT:
                        if (col > 0) {
                            col--;
                        }
                        break;
                    case Key.K_RIGHT:
                        if (col < width) {
                            col++;
                        }
                        break;
                    case Key.K_CTRL_Q:
                        done = true;
                        break;
                    default:
                        break;
                }
            }
        } finally {
            AnsiEsc.reset();
            TTYUtil.disableRawMode();
            AnsiEsc.clear();
            AnsiEsc.move(1, 1);
        }
    }

}
