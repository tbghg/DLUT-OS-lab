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
	int x; 
	int y; 
	int z; 
	x = atoi(argv[0]);
	y = atoi(argv[1]);
	z = atoi(argv[2]);
	printf("The average is %f.\n\n", Average(x, y, z));
	return 0;
}