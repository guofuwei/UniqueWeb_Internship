#include <stdio.h>
#include <queue>
#include <stdlib.h>
using namespace std;

queue<char> q;
#define MAX_NUM 100
bool visited[MAX_NUM]={0};

//�����ڽӾ���
typedef struct MGraph
{
	char vexs[MAX_NUM];            
	int arcs[MAX_NUM][MAX_NUM];     
	int vexnum,arcnum;          //��ǰ�Ķ������ͱ���
}MGraph;

void visit(char x){
    printf("%c\n",x);
}

int LocateVex(MGraph &G, char v)
{
	int i;
	for (i = 0; i < G.vexnum; i++){
        if (G.vexs[i] == v){
            return i;
        }
    }
    return -1;
}

void InitGraph(MGraph &G){
    int i,j,k;
    int m,n;
    char v1,v2;
    printf("�������ܽ�������ܱ���:");
    scanf("%d %d",&G.vexnum,&G.arcnum);
    printf("��������������Ϣ\n");
    getchar();
    for(i=0;i<G.vexnum;i++)
    {
        printf("��%d����:",i+1);
        scanf("%c",&G.vexs[i]);
        getchar();
    }
    for(i=0;i<G.vexnum;i++){
        for(j=0;j++;j<G.vexnum){
            G.arcs[i][j]=0;
        }
    }
    printf("����������ߵ�������Ϣ\n");
    for(i=0;i<G.arcnum;i++){
        // getchar();//��ȡ/n
        fflush(stdin);
        printf("��%d����:",i+1);
        scanf("%c %c",&v1,&v2);
        m=LocateVex(G,v1);
        n=LocateVex(G,v2);
        G.arcs[m][n]=G.arcs[n][m]=1;
    }
}


//���������BFS��ʼ����㣨��������
void BFS_Search(MGraph &G,char v0){
    int u,i,v,w;
    v=LocateVex(G,v0);
    visit(v0);
    visited[v]=1;
    q.push(v0);
    while(!q.empty()){
        u=q.front();
        v=LocateVex(G,u);
        // visited[v]=1;
        // visit(u);
        q.pop();
        for(i=0;i<G.vexnum;i++){
            w=G.vexs[i];
            if(G.arcs[v][i]&&!visited[i]){
                visit(w);
                q.push(w);
                visited[i]=1;
            }
        }
    }
}

//�ݹ��㷨
//���������DFS��ʼ����㣨��������
void DFS_Search(MGraph &G, int v0)
{
	int v,i;
	visit(v0);
    v=LocateVex(G,v0);
	visited[v] = 1;
	for (i=0;i<G.vexnum;i++){
        if (G.arcs[v][i]&&!visited[i]){
            DFS_Search(G,G.vexs[i]);
        }
    }
		
}

int main(){
    MGraph G;
    InitGraph(G);
    // printf("���ǹ����������\n");
    // BFS_Search(G,'A');
    printf("��ʱ������������㷨\n");
    DFS_Search(G,'A');
    system("pause");
    return 0;
}

