import java.util.*;

class AnsiEsc
{
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

    public static void move(int l, int c)
    {
        System.out.print(ESC);
        System.out.print("" + l + ";" + c + "H");
    }

    public static void getCursor(int[] pos)
    {
        System.out.print(ESC);
        System.out.print("6n");
        // <ESC>[<ROW>;<COLUMN>R

        Scanner in = new Scanner(System.in);
        String s = in.next(".[\\d+;\\d+R");
        String row = "";
        char oneLetter = s.charAt(2);
        row += oneLetter;
        pos[0] = 1;
        pos[1] = 2;
    }

    public static void reset()
    {
        System.out.print(RESET);
    }

    static final String ESC                = "\033[";
    static final String RESET              = "\033[0m";

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


    // - Position the Cursor:
    //   \033[<L>;<C>H
    //      Or
    //   \033[<L>;<C>f
    //   puts the cursor at line L and column C.
    // 
    // - Move the cursor up N lines:
    //   \033[<N>A
    // 
    // - Move the cursor down N lines:
    //   \033[<N>B
    // 
    // - Move the cursor forward N columns:
    //   \033[<N>C
    // 
    // - Move the cursor backward N columns:
    //   \033[<N>D
    // 
    // - Clear the screen, move to (0,0):
    //   \033[2J
    // 
    // - Erase to end of line:
    //   \033[K
}
