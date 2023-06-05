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
	int x = atoi(argv[0]);
	int y = atoi(argv[1]);
	printf("Min(%d, %d)，较小的数为：%d \n", x, y, Min(x, y));
	return 0;
}