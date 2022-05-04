import java.io.IOException;
import java.util.*;

import com.sun.jna.Platform;

class Hello
{
    public static void main(String args[]) throws IOException
    {
        AnsiEsc.set(AnsiEsc.FG_RED);
        System.out.println("red");

        AnsiEsc.set(AnsiEsc.FG_GREEN, AnsiEsc.BOLD);
        System.out.println("green bold");

        //AnsiEsc.move(3, 4);
        //System.out.println("HELLO");

        AnsiEsc.reset();
        System.out.println("Hello, world");

        //int[] pos = new int[2];
        //AnsiEsc.getCursor(pos);

        //Scanner in = new Scanner(System.in);
        //String s = in.next("\\d+;\\d+R");
        //System.out.println(s);
        int ch = System.in.read();
        char c = (char) ch;
        System.out.println(c);

        //System.out.println(pos[0]);
        //System.out.println(pos[1]);

    }
}
