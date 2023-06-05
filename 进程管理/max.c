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
	int x = atoi(argv[0]);
	int y = atoi(argv[1]);
	printf("Max(%d, %d)，较大的数为：%d \n", x, y, Max(x, y));
	return 0;
}