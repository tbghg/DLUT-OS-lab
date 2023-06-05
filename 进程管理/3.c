#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <string.h>

int main()
{
	char name[20];
	char data[20][20];
	int pid, statue, p, i;
	while(strcmp(name,"quit")!=0) {
		printf("请输入要进行的方法(max/min/average/quit):");
		scanf("%s",name);
		if(strcmp(name,"min")==0) {
			p = fork();
			if(p == 0) {
				printf("请输入两个数来比较大小:");
				for(i=0; i<2; i++) 
					scanf("%s",data[i]);
				execl("./min",data[0],data[1]);
			}
			pid = wait(&statue);
		} else if(strcmp(name,"max")==0) {
			p = fork();
			if(p == 0)
			{
				printf("请输入两个数来比较大小:");
				for(i=0; i<2; i++) 
					scanf("%s",data[i]);
				execl("./max",data[0],data[1]);
			}
			pid = wait(&statue);
		} else if(strcmp(name,"average")==0) {
			p = fork();
			if(p == 0) {
				printf("请输入三个数来求平均值:");
				for(i=0; i<3; i++) 
					scanf("%s",data[i]);
				execl("./average",data[0],data[1],data[2]);
			}
			pid = wait(&statue);
		} else if(strcmp(name,"quit")==0) {
			printf("程序成功退出\n");
			break;
		} else {
			printf("输入错误\n");
			break;	
		}
	}
	return 0;
}
