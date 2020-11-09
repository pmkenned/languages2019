import java.util.Vector;
import java.util.Hashtable;
import java.util.Random;
import java.io.InputStream;
import java.io.IOException;
import java.io.StreamTokenizer;
import java.io.Reader;
import java.io.BufferedReader;
import java.io.InputStreamReader;

class Prefix {
    public Vector pref;

    Prefix(Prefix p)
    {
        pref = (Vector) p.pref.clone();
    }

    Prefix(int n, String str)
    {
        pref = new Vector();
        for (int i=0; i<n; i++)
            pref.addElement(str);
    }

    static final int MULTIPLIER = 31;

    public int hashCode()
    {
        int h = 0;
        for (int i=0; i < pref.size(); i++)
            h = MULTIPLIER * h + pref.elementAt(i).hashCode();
        return h;
    }

    public boolean equals(Object o)
    {
        Prefix p = (Prefix) o;

        for (int i = 0; i < pref.size(); i++)
            if (!pref.elementAt(i).equals(p.pref.elementAt(i)))
                return false;
        return true;
    }
}

class Chain {
    static final int NPREF = 2;
    static final String NONWORD = "\n";
    Hashtable statetab = new Hashtable();
    Prefix prefix = new Prefix(NPREF, NONWORD);
    Random rand = new Random();

    void build(InputStream in) throws IOException
    {
        Reader r = new BufferedReader(new InputStreamReader(in));
        StreamTokenizer st = new StreamTokenizer(r);

        st.resetSyntax();
        st.wordChars(0, Character.MAX_VALUE);
        st.whitespaceChars(0, ' ');
        while (st.nextToken() != st.TT_EOF)
            add (st.sval);
        add(NONWORD);
    }

    void add(String word)
    {
        Vector suf = (Vector) statetab.get(prefix);
        if (suf == null) {
            suf = new Vector();
            statetab.put(new Prefix(prefix), suf);
        }
        suf.addElement(word);
        prefix.pref.removeElementAt(0);
        prefix.pref.addElement(word);
    }

    void generate(int nwords)
    {
        prefix = new Prefix(NPREF, NONWORD);
        for (int i = 0; i < nwords; i++) {
            Vector s = (Vector) statetab.get(prefix);
            int r = Math.abs(rand.nextInt()) % s.size();
            String suf = (String) s.elementAt(r);
            if (suf.equals(NONWORD))
                break;
            System.out.print(suf + ' ');
            if (i % 20 == 0)
                System.out.println();
            prefix.pref.removeElementAt(0);
            prefix.pref.addElement(suf);
        }
    }

}

class Markov {
    static final int MAXGEN = 10000;
    public static void main(String[] args)
    {
        Chain chain = new Chain();
        int nwords = MAXGEN;

        try {
            chain.build(System.in);
        } catch (IOException e) {
            ;
        }
        chain.generate(nwords);
    }
}
