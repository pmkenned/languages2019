import ex.TTYUtil;
public class Test {
    public static void main(String[] args) {
        if (TTYUtil.isTTY()) {
            System.out.println("TTY: "+TTYUtil.getTTYName());
        } else {
            System.out.println("Not a TTY");
        }
    }
}
