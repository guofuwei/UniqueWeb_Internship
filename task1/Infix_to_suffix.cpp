#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void Infix_to_suffix_detail(char* ,char );

//新建两个栈分别储存表达式和运算符
char num[1000]="0";
char opera[1000]="0";
//新建两个栈分别储存表达式和运算符
int i=0;int j=0;int op=0;

void Opera_NUm(char* str){
    num[j]=opera[op-1];
    j++;
    op--;
}

char* Infix_to_suffix(char *str){    
    while(str[i]!='\0'){
    Infix_to_suffix_detail(str,str[i]);
    }
    while(op!=0){
        Opera_NUm(str);
    }
    num[j]='\0';
    i=0;
    while(num[i]!='\0'){
        str[i]=num[i];
        i++;
        str[i]='\0';
    }
    return str;
}

void PushOpera(char* stack1,char* stack2,int* index1 ,int* index2){
    stack1[*index1]=stack2[*index2];
    (*index1)++;
    (*index2)++;
}




void Infix_to_suffix_detail(char* str,char x){
    if('0'<=x&&x<='9'){
            num[j++]=str[i++];
        }
    else{
        //如果是操作数
        if(op==0){
                PushOpera(opera,str,&op,&i);
            }
        
        else if(x=='+'||x=='-'){
            if(opera[j-1]=='+'||opera[j-1]=='-'||opera[j-1]=='('){
                PushOpera(opera,str,&op,&i);
            }
            if(opera[j-1]=='*'||opera[j-1]=='/'){
                Opera_NUm(str);
                Infix_to_suffix_detail(str,x);
            }
        }

        else if(x=='*'||x=='/'||x=='('){
            PushOpera(opera,str,&op,&i);
        }

        else if(x==')'){
            while(opera[op-1]!='('){
                num[j]=opera[op-1];
                j++;
                op--;
            }
            i++;
            op--;//保证覆盖'('
        }
    }
}

int main(){
    char str[1000]={'\0'};
    printf("请输入一个字符串:");
    gets(str);
    char *The_str;
    The_str=Infix_to_suffix(str);
    printf("%s\n",The_str);
    system("pause");
    return 0;
}