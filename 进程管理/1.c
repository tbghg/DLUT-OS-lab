#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>

/*
每个进程都执行自己独立的程序，打印自己的pid，每个父进程打印其子进程的pid
    父 -> 子1 -> 子2
*/

int main(int argc, char *argv[])
{
	int p1 = fork();  // 创建子进程p1
	int statue; // 收集进程退出时的一些状态
	int ppid;
	if(p1 == 0)
	{
		int p2 = fork(); // 创建子进程p1
		int statue;  // 收集进程退出时的一些状态
		int ppid;
		char* Argv[] = {"10", "20", NULL};
		char* envp[] = {NULL};
		if(p2 == 0)
		{
			char* Argv[] = {"40", "30", NULL};
			char* envp[] = {NULL};
			printf("子进程p2 pid：%d\n", getpid());
			// p2执行自己独立的程序
			execve("./min", Argv, envp);
			return 0;
		}
		ppid = wait(&statue);  // 阻塞等待子进程结束
		printf("子进程p2 pid：%d 运行结束\n", ppid);
		printf("子进程p1 pid：%d 子进程p1的子进程pid：%d\n", getpid(), p2);
		// p1执行自己独立的程序
		execve("./max", Argv, envp);
		return 0;
	}
	ppid = wait(&statue);  // 阻塞等待子进程结束
	printf("子进程p1 pid：%d 运行结束\n", ppid);
	printf("父进程pid：%d 父进程的子进程pid：%d\n", getpid(), p1);
	return 0;
}
