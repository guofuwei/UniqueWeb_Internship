//哈希表的构建和查找
//哈希冲突采用的是开放寻址法，可以用链式扩展优化

#include <stdio.h>
#include <stdlib.h>
#define HASHSIZE 100
#define DEFAULT 65535

typedef struct{
    int key;
    int value;
}data;

typedef struct{
    data* elem;//用数组作为哈希表 
    int count;
}HashMap;


//初始化哈希表
void InitHashMap(HashMap* hashmap){
    int i;
    hashmap->elem=(data *)malloc(sizeof(data)*HASHSIZE);
    hashmap->count=0;
    for(i=0;i<HASHSIZE;i++){
        hashmap->elem[i].key=DEFAULT;
        hashmap->elem[i].value=DEFAULT;
        }
}
//哈希散列函数
int Hash(int key){
    return key%HASHSIZE;
}

//插入
void HashInsert(HashMap* hashmap,int key,int value){
    //求哈希地址
    if(hashmap->count<HASHSIZE){
        int hashaddr=Hash(key);
        while(hashmap->elem[hashaddr].key!=DEFAULT){
            hashaddr=(hashaddr+1)%HASHSIZE;
        }
        hashmap->elem[hashaddr].key=key;
        hashmap->elem[hashaddr].value=value;
        hashmap->count++;
    }   
}

//查询
//返回哈希地址
int HashSearch(HashMap* hashmap,int key){
    int hashaddr=Hash(key);
    while(hashmap->elem[hashaddr].key!=key){
        hashaddr=(hashaddr+1)%HASHSIZE;
        if(hashmap->elem[hashaddr].key==DEFAULT||hashaddr==Hash(key)){
            printf("查找错误");
            return -1;
        }
    }
    return hashaddr;
}

//修改value值
void HashChange(HashMap* hashmap,int key,int newvalue){
    int hashaddr=HashSearch(hashmap,key);
    if(hashaddr==-1){
        printf("修改错误");
    }
    else{
        //更新value值
        hashmap->elem[hashaddr].value=newvalue;
    }
}

//删除value值
void HashDelete(HashMap* hashmap,int key){
    int hashaddr=HashSearch(hashmap,key);
    if(hashaddr==-1){
        printf("删除错误");
    }
    else{
        hashmap->elem[hashaddr].key=DEFAULT;
        hashmap->elem[hashaddr].value=DEFAULT;
    }
}


int main(){
    int i;
    int find_addr,find_value;
    HashMap hashmap;
    InitHashMap(&hashmap);
    int key[HASHSIZE]={13,29,27,28,26,30,38};
    int value[HASHSIZE]={1,2,3,4,5,6,7};
    //利用插入函数构造哈希表
    for (i=0;i<HASHSIZE;i++){
        HashInsert(&hashmap,key[i],value[i]);
    }
    //调用查找算法
    for(i=0;i<7;i++){
        find_addr=HashSearch(&hashmap,key[i]);
        find_value=hashmap.elem[find_addr].value;
        printf("%d-%d\n",key[i],find_value);
    }
    system("pause");
}
    