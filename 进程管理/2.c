#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>


int main(int argc, char *argv[])
{
	int p1;
	int p2;
	int statu1; // 收集进程退出时的一些状态
	int statu2;
	int ppid1;
	int ppid2;
	
	p1 = fork();  // 创建子进程
	if(p1 == 0)
	{
		char* Argv[] = {"./test", "200", "300", NULL};
		char* envp[] = {NULL};
		
		printf("I am son1 process. My pid is %d.\n", getpid());
		// 执行自己独立的程序
		execve("./max", Argv, envp);
		return 0;
	}
	else
	{
		p2 = fork();
		if(p2 == 0)
		{
			char* Argv[] = {"./min", "100", "200", NULL};
			char* envp[] = {NULL};
		
			printf("I am son2 process. My pid is %d. \n", getpid());
			// 执行自己独立的程序
			execve("./min", Argv, envp);
			return 0;
		}
	}
	ppid1 = wait(&statu1);
	ppid2 = wait(&statu2);
	printf("The wait is over, child pid1 = %d finished, child pid2 = %d finished\n", ppid1, ppid2);
	printf("I am parent process. My pid is %d. My child is %d and %d.\n", getpid(), p1, p2);
	return 0;	
}
