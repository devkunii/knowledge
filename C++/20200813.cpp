// #include <iostream>
// using namespace std;
//
// namespace spc1 {
//   int dt;
//   void disp() { cout << "spc1_dt=" << dt << endl; }
// }
//
// namespace spc2 {
//   int dt;
//   void disp() { cout << "spc2_dt=" << dt << endl; }
// }
//
// int dt;
// void disp()
// {
//   cout << "global_dt=" << dt << endl;
// }
//
// int main()
// {
//   spc1::dt = 100;
//   spc1::disp();
//   spc2::dt = 200;
//   spc2::disp();
//   ::dt = 200;
//   ::disp();
//   dt = 400;
//   disp();
//   return 0;
// }

// #include <iostream>
// using namespace std;
//
// namespace outside {
//   void disp1() { cout << "outside-disp1\n"; }
//   namespace inside {
//     void disp2() { cout << "inside-disp2\n"; }
//   }
// }
//
// using namespace outside;
// using namespace outside::inside;
//
// int main()
// {
//   outside::disp1();
//   outside::inside::disp2();
//   disp1();
//   disp2();
//   return 0;
// }
