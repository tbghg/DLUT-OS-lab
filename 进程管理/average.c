#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

float Average(int a, int b, int c)
{
	float ave;
	ave = (a+b+c)/3.0;
	return ave;
}

int main(int argc, char* argv [])
{
	int x = atoi(argv[0]);
	int y = atoi(argv[1]);
	int z = atoi(argv[2]);
	printf("Average(%d, %d, %d)，平均数为：%f \n",x ,y ,z ,Average(x, y, z));
	return 0;
}