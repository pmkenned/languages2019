import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class AnsiEsc
{
    // Clear the screen, move to (1, 1)
    public static void clear()
    {
        System.out.print(ESC);
        System.out.print(CLEAR);
    }

    public static void reset()
    {
        System.out.print(ESC);
        System.out.print(RESET);
    }

    // Erase to end of line
    public static void eraseToEnd()
    {
        System.out.print(ESC);
        System.out.print(ERASE_TO_END);
    }

    // Move the cursor up N lines:
    public static void up(int n)
    {
        System.out.print(ESC);
        System.out.print(n);
        System.out.print("A");
    }

    // Move the cursor down N lines:
    public static void down(int n)
    {
        System.out.print(ESC);
        System.out.print(n);
        System.out.print("B");
    }

    // Move the cursor forward N columns:
    public static void forward(int n)
    {
        System.out.print(ESC);
        System.out.print(n);
        System.out.print("C");
    }

    // Move the cursor backward N columns:
    public static void backward(int n)
    {
        System.out.print(ESC);
        System.out.print(n);
        System.out.print("D");
    }

    // Puts the cursor at line L and column C
    public static void move(int l, int c)
    {
        System.out.print(ESC);
        System.out.print("" + l + ";" + c + "H"); // or "f"?
    }

    public static int[] getCursor()
    {
        System.out.print(ESC);
        System.out.print("6n"); // <ESC>[<ROW>;<COLUMN>R
        Scanner in = new Scanner(System.in);
        String s = in.next(".\\[\\d+;\\d+R");
        Pattern pattern = Pattern.compile(".\\[(\\d+);(\\d+)R");
        Matcher matcher = pattern.matcher(s);
        matcher.find();
        String row = matcher.group(1);
        String col = matcher.group(2);
        int[] rv = new int[2];
        rv[0] = Integer.parseInt(row);
        rv[1] = Integer.parseInt(col);
        return rv;
    }

    public static void hideCursor()
    {
        System.out.print(ESC);
        System.out.print(HIDE);
    }

    public static void showCursor()
    {
        System.out.print(ESC);
        System.out.print(SHOW);
    }

    public static void set(String... codes)
    {
        System.out.print(ESC);
        boolean first = true;
        for (String code : codes) {
            if (!first) {
                System.out.print(";");
            }
            System.out.print(code);
            first = false;
        }
        System.out.print("m");
    }

    static final String ESC                = "\033[";

    static final String CLEAR              = "2J";
    static final String RESET              = "0m";
    static final String ERASE_TO_END       = "K";
    static final String HIDE               = "?25l";
    static final String SHOW               = "?25h";

    static final String UNDERLINE          = "4";
    static final String UNDERLINE_OFF      = "24";
    static final String BOLD               = "1";
    static final String BOLD_OFF           = "21";

    static final String FG_BLACK           = "30";
    static final String FG_RED             = "31";
    static final String FG_GREEN           = "32";
    static final String FG_YELLOW          = "33";
    static final String FG_BLUE            = "34";
    static final String FG_MAGENTA         = "35";
    static final String FG_CYAN            = "36";
    static final String FG_WHITE           = "37";
    static final String FG_BRIGHT_BLACK    = "90";
    static final String FG_BRIGHT_RED      = "91";
    static final String FG_BRIGHT_GREEN    = "92";
    static final String FG_BRIGHT_YELLOW   = "93";
    static final String FG_BRIGHT_BLUE     = "94";
    static final String FG_BRIGHT_MAGENTA  = "95";
    static final String FG_BRIGHT_CYAN     = "96";
    static final String FG_BRIGHT_WHITE    = "97";

    static final String BG_BLACK           = "40";
    static final String BG_RED             = "41";
    static final String BG_GREEN           = "42";
    static final String BG_YELLOW          = "43";
    static final String BG_BLUE            = "44";
    static final String BG_MAGENTA         = "45";
    static final String BG_CYAN            = "46";
    static final String BG_WHITE           = "47";
    static final String BG_BRIGHT_BLACK    = "100";
    static final String BG_BRIGHT_RED      = "101";
    static final String BG_BRIGHT_GREEN    = "102";
    static final String BG_BRIGHT_YELLOW   = "103";
    static final String BG_BRIGHT_BLUE     = "104";
    static final String BG_BRIGHT_MAGENTA  = "105";
    static final String BG_BRIGHT_CYAN     = "106";
    static final String BG_BRIGHT_WHITE    = "107";
}
