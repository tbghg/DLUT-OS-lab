#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>

/*
		(1)在父进程中，fork返回新创建子进程的进程ID
		(2)在子进程中，fork返回0
		(3)如果出现错误，fork返回负值
*/
int main(int argc, char *argv[])
{
	int p1 = fork();  // 创建子进程
	int statu; // 收集进程退出时的一些状态
	int ppid;
	if(p1 == 0)
	{
		int p2 = fork();
		int statu; // 收集进程退出时的一些状态
		int ppid;
		char* Argv[] = {"./max", "400", "500", NULL};
		char* envp[] = {NULL};
		if(p2 == 0)
		{
			char* Argv[] = {"./min", "100", "200", NULL};
			char* envp[] = {NULL};
			printf("I am son2 process. My pid is %d. I have not child.\n", getpid());
			// 执行自己独立的程序
			execve("./min", Argv, envp);
			return 0;
		}
		ppid = wait(&statu);  // 阻塞等待子进程结束
		printf("The wait is over, child pid %d finished.\n", ppid);
		printf("I am son1 process. My pid is %d. My child pid is %d.\n", getpid(), p2);
		// 执行自己独立的程序
		execve("./max", Argv, envp);
		return 0;
	}
	ppid = wait(&statu);  // 阻塞等待子进程结束
	printf("The wait is over, child pid %d finished.\n", ppid);
	printf("I am parent process. My pid is %d. My child pid is %d.\n", getpid(), p1);
	return 0;	
}
