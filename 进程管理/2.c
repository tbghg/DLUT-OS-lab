#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>

/*
每个进程都执行自己独立的程序，打印自己的pid，父进程打印其子进程的pid;
父进程创建两个子进程：
        -> 子1
    父 -｜
        -> 子2
*/

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
		char* Argv[] = {"10", "20", NULL};
		char* envp[] = {NULL};
		printf("子进程p1 pid：%d\n", getpid());
		// p1执行自己独立的程序
		execve("./max", Argv, envp);
		return 0;
	} else {
		p2 = fork();
		if(p2 == 0)
		{
			char* Argv[] = {"30", "40", NULL};
			char* envp[] = {NULL};
			printf("子进程p2 pid：%d\n", getpid());
			// p2执行自己独立的程序
			execve("./min", Argv, envp);
			return 0;
		}
	}
	ppid1 = wait(&statu1);
	ppid2 = wait(&statu2);
	printf("子进程p1 pid：%d 运行结束，子进程p2 pid：%d 运行结束\n", ppid1, ppid2);
	printf("父进程pid：%d 父进程的子进程p1 pid：%d，父进程的子进程p2 pid：%d\n", getpid(), p1, p2);
	return 0;
}
