#include <stdio.h>

void
hanoi_recursive(int n, int from, int to)
{
    if (n == 0)
        return;
    hanoi_recursive(n-1, from, 3-(from+to));
    printf("%d %d\n", from, to);
    hanoi_recursive(n-1, 3-(from+to), to);
}

void
hanoi_iterative(int n, int from, int to)
{
    for (int i = 1; i < (1<<n); i++) {
        int f = from, t = to;
	    for (int j = n-1; j >= 0; j--) {
	        int d = 1<<j;
			if (i % d == 0)
			    break;
			if (i & d)
			    f = 3-(f+t);
			else
			    t = 3-(f+t);
		}
		printf("%d %d\n", f, t);
    }
}

int main()
{
	hanoi_iterative(4, 0, 1);
    //hanoi_recursive(4, 0, 1);
	return 0;
}
