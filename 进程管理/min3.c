#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

int Min(int a, int b)
{
	return (a < b) ? a:b;
}

int main(int argc, char* argv [])
{
	int x;
	int y;
	x = atoi(argv[0]);
	y = atoi(argv[1]);
	printf("The smaller number in %d and %d is %d.\n\n", x, y, Min(x, y));
	return 0;
}