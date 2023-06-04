#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

int Max(int a, int b)
{
	return (a > b) ? a:b;
}

int main(int argc, char* argv [])
{
	int x;
	int y;
	x = atoi(argv[1]);
	y = atoi(argv[2]);
	printf("The larger number in %d and %d is %d.\n\n", x, y, Max(x, y));
	return 0;
}