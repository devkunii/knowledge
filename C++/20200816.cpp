#include <iostream>
using namespace std;

int add(int a, int b);     // 普通の関数プロトタイプ
int sub(int a, int b);     // 普通の関数プロトタイプ
int (*calc)(int a, int b); // 関数へのポインタの宣言

int add(int a, int b)
{
  return a + b;
}

int sub(int a, int b)
{
  return a - b;
}

int main()
{
  calc = add;
  cout << calc(30, 20) << endl;
  calc = sub;
  cout << calc(30, 20) << endl;
  return 0;
}
