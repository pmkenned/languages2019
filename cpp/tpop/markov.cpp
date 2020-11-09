#include <iostream>
#include <string>
#include <deque>
#include <map>
#include <vector>

using namespace std;

typedef deque<string> Prefix;
map<Prefix, vector<string> > statetab;

enum {
    NPREF   = 2,        /* number of prefix words */
    MAXGEN  = 10000,    /* maximum words generated */
};

char NONWORD[] = "\n";

void add(Prefix & prefix, const string & s)
{
    if (prefix.size() == NPREF) {
        statetab[prefix].push_back(s);
        prefix.pop_front();
    }
    prefix.push_back(s);
}

void build(Prefix & prefix, istream & in)
{
    string buf;

    while (in >> buf)
        add(prefix, buf);
}

void generate(int nwords)
{
    Prefix prefix;
    int i;

    for (i=0; i < NPREF; i++)
        add(prefix, NONWORD);

    for (i=0; i < nwords; i++) {
        vector<string>& suf = statetab[prefix];
        const string & w = suf[rand() % suf.size()];
        if (w == NONWORD)
            break;
        cout << w << "\n";
        prefix.pop_front();
        prefix.push_back(w);
    }
}

int main(void)
{
    int nwords = MAXGEN;
    Prefix prefix;

    for (int i = 0; i < NPREF; i++)
        add(prefix, NONWORD);
    build(prefix, cin);
    add(prefix, NONWORD);
    generate(nwords);
    return 0;
}
