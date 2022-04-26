// #include <iostream>
// using namespace std;
//
// struct STname1 {
//   int x;
//   int y;
// };
//
// struct STname2 {
//   int x;
//   int y;
//   void disp() { cout << x << " " << y << endl; };
// };
//
// int main()
// {
//   STname1 d1;
//   STname2 d2;
//
//   d1.x = 10;
//   d1.y = 20;
//   cout << d1.x << " " << d1.y << endl;
//   d2.x = 30;
//   d2.y = 40;
//   d2.disp();
//   return 0;
// }

// #include <iostream>
// using namespace std;
//
// // 共用体
// union Utype {
//   // int型のメンバ
//   int idt;
//
//   // double型のメンバ
//   double ddt;
// };
//
// int main()
// {
//   Utype a;
//   a.idt = 1234;
//   cout << a.idt << endl;
//   a.ddt = 567.89;
//   cout << a.ddt << endl;
//   cout << &a.idt << " " << &a.ddt << endl;
//   return 0;
// }

// #include <iostream>
// using namespace std;
// void test()
// {
//   union {
//     int       INT_DT;
//     short int SHRT_DT;
//   };
//
//   INT_DT = 0x41424344;
//   cout << hex;
//   cout << INT_DT << endl;
//   cout << SHRT_DT << endl;
// }

#include <iostream>
using namespace std;

// グローバル変数
int g = 0;

void func()
{
  // 自動変数
  int a = 0;

  // ローカルな静的変数
  static int s = 0;
  ++g; ++a; ++a;
  cout << "g=" << g << " a=" << a << " s=" << s << endl;
}

int main()
{
  func(); // g=1 a=1 s=1
  func(); // g=2 a=1 s=2
  func(); // g=3 a=1 s=3
  return 0;
}
