#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <string.h>

int main()
{
	char name[20];
	char data[20][20];
	int pid;
	int statu;
	int p;
	int i;
	while(strcmp(name,"exit")!=0)
	{
		printf("Please input operation name:");
		scanf("%s",name);
		
		if(strcmp(name,"min")==0) 
		{
			p = fork();
			if(p == 0)
			{
				printf("please input two numbers:");
				for(i=0; i<2; i++) 
					scanf("%s",data[i]);
				execl("./min3",data[0],data[1],NULL);
			}
			pid = wait(&statu);
		}
		else if(strcmp(name,"max")==0) 
		{
			p = fork();
			if(p == 0)
			{
				printf("please input two numbers:");
				for(i=0; i<2; i++) 
					scanf("%s",data[i]);
				execl("./max3",data[0],data[1],NULL);
			}
			pid = wait(&statu);
		}
		else if(strcmp(name,"average")==0) 
		{
			p = fork();
			if(p == 0)
			{
				printf("please input three numbers:");
				for(i=0; i<3; i++) 
					scanf("%s",data[i]);
				execl("./average3",data[0],data[1],data[2],NULL);
			}
			pid = wait(&statu);
		}
		else if(strcmp(name,"exit")==0)
		{
			printf("Exit\n");
			break;
		}
		else 
		{
			printf("Input Error\n");
			break;	
		}
	}
	return 0;
}
